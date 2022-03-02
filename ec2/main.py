import time
import datetime
import numpy as np
import math

from dotenv import dotenv_values
from InfluxHelper import InfluxHelper

config = dotenv_values()
INFLUX_USER = config.get("influx_username")
INFLUX_PW = config.get("influx_password")
INFLUX_IP = config.get("influx_ip")

ONE_MINUTE = 60

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
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi002",
	"location": "E5-1445",
	"type": "Raspberry Pi",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi003",
	"location": "E7-6125",
	"type": "Raspberry Pi",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi004",
	"location": "E6-4012",
	"type": "Raspberry Pi",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi005",
	"location": "E2-3445",
	"type": "Raspberry Pi",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi006",
	"location": "E2-2134",
	"type": "Raspberry Pi",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "raspi007",
	"location": "E5-3415",
	"type": "Raspberry Pi",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "ecserv1",
	"location": "DC-1055",
	"type": "Linux Server",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "ecserv2",
	"location": "DC-1056",
	"type": "Linux Server",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "ecserv3",
	"location": "QNC-2315",
	"type": "QDOT Server",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "FPGA-21",
	"location": "E7-4261",
	"type": "FPGA",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
	"device_disk" : np.random.randint(5, 89)
},
{
	"name": "FPGA-13",
	"location": "E7-4268",
	"type": "FPGA",
	"device_power": np.random.randint(10, 78),
	"device_vulnerabilities": np.random.randint(0,6),
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
	if np.random.randint(0, 50) == 5: 
		devices_connected = max(0, devices_connected + np.random.randint(-10, 10))
	return devices_connected
	
def get_active_vulnerabilities(active_vulnerabilities):
	if np.random.randint(0, 300) == 5: 
		active_vulnerabilities = max(0, active_vulnerabilities + np.random.randint(-2, 2))
	return active_vulnerabilities

def add_device_metrics(name, location, type, device_power, device_vulnerabilities, device_disk):
	
    fields = {
        "device_cpu": float(50 + np.random.randint(-40,40)),
		"device_mem": float(5 + np.random.randint(-3,3)),
		"device_power": device_power,
		"device_vulnerabilities": device_vulnerabilities,
		"device_disk": device_disk
    }
    tags = {
        "device_type": type,
        "host": name,
        "location": location
    }
    influx_helper.add_metric(name="mock", fields=fields, tags=tags)



def send_metrics(devices_old, vulnerabilities_old):
	x = get_current_hour_seconds()
	
	fields = {
		"cpu_avg": float(max(0, min(100, daily_sin(x) + random_noise())))
	}
	tags = {
        "device_type": "server",
        "host": "all",
        "location": "E2-2254",
	}
	influx_helper.add_metric(name="mock", fields=fields, tags=tags)
	
	fields = {
		"mem_avg": float(min(1.5*daily_sin(x) + random_noise(),98))
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
		add_device_metrics(
			device["name"],
			device["location"],
			device["type"],
			device["device_power"],
			device["device_vulnerabilities"],
			device["device_disk"]
		)
	
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
