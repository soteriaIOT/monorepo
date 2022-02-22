import requests
import speedtest
import psutil

from ratelimit import limits
from ratelimit import RateLimitException
from backoff import on_exception
from backoff import expo
from pprint import pprint as pp

from dotenv import dotenv_values

config = dotenv_values()
API_KEY = config.get("api_key")

ONE_MINUTE = 60

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
    cpu = psutil.cpu_percent()
    cpu_freq = psutil.cpu_freq().current
    try:
        cpu_temp = psutil.sensors_temperatures()["cpu_thermal"][0].current
    except Exception as e:
        cpu_temp = None
    memory = psutil.virtual_memory().percent
    return dict(
        cpu=cpu,
        cpu_freq=cpu_freq,
        cpu_temp=cpu_temp,
        memory=memory,

    )

pp(get_weather(WATERLOO))
pp(get_wifi())
pp(get_device_stats())
