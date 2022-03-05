import speedtest
import psutil
import time

from common.influx_helper import INFLUX_HELPER
from common.tags import COMMON_TAGS

ONE_MINUTE = 60


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
    wifi_data = get_wifi()
    device_data = get_device_stats()

    wifi_fields = {
        "download_speed": wifi_data["download"],
        "upload_speed": wifi_data["upload"],
        "ping": wifi_data["ping"],
        "bytes_sent": wifi_data["bytes_sent"],
        "bytes_received": wifi_data["bytes_received"],
    }
    device_fields = {
        "cpu": device_data["cpu"],
        "cpu_freq": device_data["cpu_freq"],
        "cpu_temp": device_data["cpu_temp"],
        "memory_usage": device_data["memory"],
        "running_processes": device_data["running_processes"],
    }

    INFLUX_HELPER.add_metric(name="wifi", fields=wifi_fields, tags=COMMON_TAGS)
    INFLUX_HELPER.add_metric(name="device", fields=device_fields, tags=COMMON_TAGS)
    INFLUX_HELPER.send()


if __name__ == "__main__":
    starttime = time.time()
    FREQUENCY = ONE_MINUTE
    while True:
        try:
            send_metrics()
            time.sleep(ONE_MINUTE - ((time.time() - starttime) % ONE_MINUTE))
        except Exception as e:
            print("Got Exception: ", e)
