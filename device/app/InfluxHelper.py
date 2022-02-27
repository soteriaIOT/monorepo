#!pip install influxdb

from influxdb import InfluxDBClient
import datetime
import random
import time

class InfluxHelper:
    
    def __init__(self, ip='', port=8086, username='', password='', db=''):
        self.client = InfluxDBClient(ip, port, username, password, db)
        self.payloads = []
        
    def add_metric(self, name="SAMPLE", fields={"value": 1}, time=None, tags={"host":"local", "region":"us-east2"}):
        if time is None:
            time = datetime.datetime.utcnow()
        payload = {
            "measurement": name,
            "tags": tags,
            "time": time,
            "fields": fields
        }
        self.payloads.append(payload)
        
    def send(self):
        self.client.write_points(self.payloads)
        #print(self.payloads)
        self.payloads = []

# influxHelper = InfluxHelper()
# while(True):
    # cpu_load = random.random()
    # influxHelper.add_metric(name="CPU_SAMPLE", fields={"value": cpu_load})
    # influxHelper.send()
    # print(cpu_load)
    # time.sleep(10)

