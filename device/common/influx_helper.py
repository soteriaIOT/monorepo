from datetime import datetime

from dotenv import dotenv_values
from influxdb import InfluxDBClient


class InfluxHelper:
    def __init__(self, ip="", port=8086, username="", password="", db=""):
        self.client = InfluxDBClient(ip, port, username, password, db)
        self.payloads = []

    def add_metric(
        self,
        name="SAMPLE",
        fields={"value": 1},
        time=None,
        tags={"host": "local", "region": "us-east2"},
    ):
        if time is None:
            time = datetime.utcnow()
        payload = {"measurement": name, "tags": tags, "time": time, "fields": fields}
        self.payloads.append(payload)

    def send(self):
        self.client.write_points(self.payloads)
        self.payloads = []


config = dotenv_values()
INFLUX_USER = config.get("influx_username")
INFLUX_PW = config.get("influx_password")
EC2_IP = config.get("ec2_ip")

INFLUX_HELPER = InfluxHelper(
    ip=EC2_IP, port=8086, username=INFLUX_USER, password=INFLUX_PW, db="test"
)
