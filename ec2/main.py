import time
import datetime
import numpy as np
import math
import random

from dotenv import dotenv_values
from InfluxHelper import InfluxHelper

config = dotenv_values()
INFLUX_USER = config.get("influx_username")
INFLUX_PW = config.get("influx_password")
INFLUX_IP = config.get("influx_ip")

ONE_MINUTE = 10

influx_helper = InfluxHelper(
    ip=INFLUX_IP, port=8086, username=INFLUX_USER, password=INFLUX_PW, db="test"
)

# Num Active Vulnerabilities (Metric)
# Num Devices Connected (Metric)
# Memory Usage (Time Series) grouped by device
# CPU Usage (Time Series) group by device
# CPU/Memory Usage (Time Series) aggregated over all groups as Average
# Num Active Vulnerabilities (Time Series) grouped by device
# Disk Utilization % (Bar) grouped by device

devices = [
{
	"name": "raspi001",
	"location": "E7-4235",
	"type": "Raspberry Pi",
    "device_cpu": 5.3,
    "device_mem": 4.1,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,2),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi002",
	"location": "E5-1445",
	"type": "Raspberry Pi",
    "device_cpu": 35.3,
    "device_mem": 9.1,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,3),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi003",
	"location": "E7-6125",
	"type": "Raspberry Pi",
    "device_cpu": 8.6,
    "device_mem": 7.1,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,2),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi004",
	"location": "E6-4012",
	"type": "Raspberry Pi",
    "device_cpu": 23.5,
    "device_mem": 42.3,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,2),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi005",
	"location": "E2-3445",
	"type": "Raspberry Pi",
    "device_cpu": 68.2,
    "device_mem": 78.9,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,2),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi006",
	"location": "E2-2134",
	"type": "Raspberry Pi",
    "device_cpu": 43.2,
    "device_mem": 86.2,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,3),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi007",
	"location": "E5-3415",
	"type": "Raspberry Pi",
    "device_cpu": 7.3,
    "device_mem": 12.1,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,2),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "ecserv1",
	"location": "DC-1055",
	"type": "Linux Server",
    "device_cpu": 8.3,
    "device_mem": 23.1,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,2),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "ecserv2",
	"location": "DC-1056",
	"type": "Linux Server",
    "device_cpu": 23.5,
    "device_mem": 46.2,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,1),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "ecserv3",
	"location": "QNC-2315",
	"type": "QDOT Server",
    "device_cpu": 9.8,
    "device_mem": 10.5,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,2),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "FPGA-21",
	"location": "E7-4261",
	"type": "FPGA",
    "device_cpu": 56.1,
    "device_mem": 20.5,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,3),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "FPGA-13",
	"location": "E7-4268",
	"type": "FPGA",
    "device_cpu": 8.4,
    "device_mem": 28.2,
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,2),
	"device_disk" : np.random.randint(5, 89)
}
]



# Get # seconds passed in this hour
def get_current_hour_seconds():
	return time.time() % 3600

# Sin function with period of 3600s or 1 hour
def daily_sin(x):
	return 50 * math.sin(x*math.pi/1800) + 50

def random_noise(delta=10):
	return np.random.normal(0, delta)

def get_devices_connected(devices_connected):
	if np.random.randint(0, 200) == 5: 
		devices_connected = max(0, devices_connected + np.random.randint(-10, 10))
	return devices_connected
	
def get_active_vulnerabilities(active_vulnerabilities):
	if np.random.randint(0, 200) == 5: 
		active_vulnerabilities = max(0, active_vulnerabilities + np.random.randint(-2, 2))
	return active_vulnerabilities

def add_device_metrics(device):
    fields = {
        "device_cpu": float(device["device_cpu"] + 5*random.random()),
		"device_mem": float(device["device_mem"] + 10*random.random()),
		"device_power": int(device["device_power"] + 10*random.random()),
		"device_vulnerabilities": int(device["device_vulnerabilities"]),
		"device_disk": int(device["device_disk"] + 10*random.random()),
    }
    tags = {
        "device_type": device["type"],
        "host": device["name"],
        "location": device["location"]
    }
    influx_helper.add_metric(name="mock", fields=fields, tags=tags)



def send_metrics(devices_old, vulnerabilities_old):
	x = get_current_hour_seconds()
	
	fields = {
		"cpu_avg": float(max(3, min(100, daily_sin(x+300) + random_noise())))
	}
	tags = {
        "device_type": "server",
        "host": "all",
        "location": "E2-2254",
	}
	influx_helper.add_metric(name="mock", fields=fields, tags=tags)
	
	fields = {
		"mem_avg": float(max(3, min(1.5*daily_sin(x) + random_noise(),98)))
	}
	tags = {
		"device_type": "server",
		"host": "all",
		"location": "E2-2254",
	}
	influx_helper.add_metric(name="mock", fields=fields, tags=tags)
	
	devices_new = get_devices_connected(devices_old)
	fields = {
		"devices_connected": devices_new
	}
	influx_helper.add_metric(name="mock", fields=fields, tags={})

	vulnerabilities_new = get_active_vulnerabilities(vulnerabilities_old)
	fields = {
		"active_vulnerabilities": vulnerabilities_new
	}
	influx_helper.add_metric(name="mock", fields=fields, tags={})
	
	for device in devices:
		add_device_metrics(device)
	
	influx_helper.send()
	return (devices_new, vulnerabilities_new)

starttime = time.time()
FREQUENCY = ONE_MINUTE

n_vulnerabilities = 13
n_devices = 351
while True:
	try:
		n_devices, n_vulnerabilities = send_metrics(n_devices, n_vulnerabilities)
		time.sleep(ONE_MINUTE - ((time.time() - starttime) % ONE_MINUTE))
	except Exception as e:
		print("Got Exception: ", e)
