from json import loads
from time import sleep
from time import time


from backoff import on_exception
from backoff import expo
from certifi import where
from dotenv import dotenv_values
from ratelimit import limits
from ratelimit import RateLimitException

from urllib3 import PoolManager

from common.influx_helper import INFLUX_HELPER
from common.tags import COMMON_TAGS

config = dotenv_values()
API_KEY = config.get("api_key")

ONE_MINUTE = 60

@on_exception(expo, RateLimitException)
@limits(calls=30, period=ONE_MINUTE)
def get_weather(location):
    lat, lon = location
    http = PoolManager(
        cert_reqs="CERT_REQUIRED",
        ca_certs=where()
    )
    response = http.request(
        'GET',
        'https://api.openweathermap.org/data/2.5/weather',
        fields={'appid': API_KEY, 'lat': lat, 'lon': lon},
    )
    return loads(response.data.decode('utf-8'))


WATERLOO = (43.4643, -80.5204)

def send_weather():
    weather_data = get_weather(WATERLOO)
    weather_fields = {
        "temp": weather_data["main"]["temp"] - 273.15,
        "feels_like": weather_data["main"]["feels_like"] - 273.15,
        "humidity": weather_data["main"]["humidity"],
        "wind_speed": weather_data["wind"]["speed"],
        "wind_deg": weather_data["wind"]["deg"],
    }

    if "gust" in weather_data["wind"]:
        weather_fields["wind_gust"] = float(weather_data["wind"]["gust"])
    
    weather_tags = {"location": weather_data["name"]}
    tags = weather_tags.copy()
    tags.update(COMMON_TAGS)
    INFLUX_HELPER.add_metric(name="weather", fields=weather_fields, tags=tags)
    INFLUX_HELPER.send()

if __name__ == "__main__":
    starttime = time()
    FREQUENCY = ONE_MINUTE
    while True:
        try:
            send_weather()
            sleep(ONE_MINUTE - ((time() - starttime) % ONE_MINUTE))
        except Exception as e:
            print("Got Exception: ", e)
