import requests
import speedtest
import psutil
import time
import datetime

from ratelimit import limits
from ratelimit import RateLimitException
from backoff import on_exception
from backoff import expo
from dotenv import dotenv_values

from InfluxHelper import InfluxHelper

config = dotenv_values()
API_KEY = config.get("api_key")
INFLUX_USER = config.get("influx_username")
INFLUX_PW = config.get("influx_password")
INFLUX_IP = config.get("influx_ip")

ONE_MINUTE = 60

influx_helper = InfluxHelper(ip=INFLUX_IP, port=8086, username=INFLUX_USER, password=INFLUX_PW, db="test")

@on_exception(expo, RateLimitException)
@limits(calls=30, period=ONE_MINUTE)
def get_weather(location):
    lat, lon = location
    URL = "https://api.openweathermap.org/data/2.5/weather?lat={lat}&lon={lon}&appid={api_key}"
    url = URL.format(lat=lat, lon=lon, api_key=API_KEY)
    response = requests.get(url)
    return response.json()

WATERLOO = (43.4643, -80.5204)

def get_wifi():
    s = speedtest.Speedtest()
    s.get_best_server()
    s.download()
    s.upload()
    RELEVANT_KEYS = [
        "download",
        "upload",
        "ping",
        "bytes_sent",
        "bytes_received",
        "timestamp",
    ]
    results = s.results.dict()
    return {k: results[k] for k in RELEVANT_KEYS}

def get_device_stats():
    p = psutil.Process()
    # provides speedup
    with p.oneshot():
        cpu_usage = psutil.cpu_percent()
        cpu_freq = psutil.cpu_freq().current
        try:
            cpu_temp = psutil.sensors_temperatures()["cpu_thermal"][0].current
        except Exception as e:
            cpu_temp = None
        memory = psutil.virtual_memory().percent
        running_processes = len(psutil.pids())
    
        return dict(
            cpu=cpu_usage,
            cpu_freq=cpu_freq,
            cpu_temp=cpu_temp,
            memory=memory,
            running_processes=running_processes,
        )

def send_metrics():
    weather_data = get_weather(WATERLOO)
    wifi_data = get_wifi()
    device_data = get_device_stats()
    
    weather_fields = {
        "temp": weather_data['main']['temp'] - 273.15,
        "feels_like": weather_data['main']['feels_like'] - 273.15,
        "humidity": weather_data['main']['humidity'],
        "wind_speed": weather_data['wind']['speed'],
        "wind_deg": weather_data['wind']['deg'],
    }

    if "gust" in weather_data['wind']:
        weather_fields["wind_gust"] = float(weather_data['wind']['gust'])

    weather_tags = {
        "location": weather_data['name']
    }
    wifi_fields = {
        "download_speed": wifi_data['download'],
        "upload_speed": wifi_data['upload'],
        "ping": wifi_data['ping'],
        "bytes_sent": wifi_data['bytes_sent'],
        "bytes_received": wifi_data['bytes_received']
    }
    device_fields = {
        "cpu": device_data['cpu'],
        "cpu_freq": device_data['cpu_freq'],
        "cpu_temp": device_data['cpu_temp'],
        "memory_usage": device_data['memory'],
        "running_processes": device_data['running_processes']
    }

    common_tags = {
        "device": "raspberry pi",
        "host": "raspi0001",
        "location": "University of Waterloo"
    }
     
    copy_tags = weather_tags.copy()
    copy_tags.update(common_tags)
    influx_helper.add_metric(name="weather", fields=weather_fields, tags=copy_tags)
    influx_helper.add_metric(name="wifi", fields=wifi_fields, tags=common_tags)
    influx_helper.add_metric(name="device", fields=device_fields, tags=common_tags)
    influx_helper.send()


if __name__ == "__main__":
    starttime = time.time()
    FREQUENCY = ONE_MINUTE
    while True:
        try:
            send_metrics()
            time.sleep(ONE_MINUTE - ((time.time() - starttime) % ONE_MINUTE))
        except Exception as e:
            print("Got Exception: " , e)
    
    
    