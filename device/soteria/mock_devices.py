from soteria.kafka_helper import produce_message

DEVICE_REQUIREMENTS = "device-requirements"

# This script is being used to setup mock devices by pushing fake device-ids and requirements
# to the device-requirements topic.
# Run as python3 -m soteria.mock_devices.py from devices/ folder

bad_devices = [
    f"raspi000{i}" for i in range(1, 7)
]

bad_dependencies = [
    "b2sdk==1.13.1",
    "requests==2.20.0",
    "urllib3==1.26.4",
]
for device_id in bad_devices:
    produce_message(DEVICE_REQUIREMENTS, device_id, bad_dependencies)


good_devices = [
    f"eceserv{i}" for i in range(1, 4)
] + [
    "raspi0007",
    "FPGA-21",
]

good_dependencies = [
    "urllib==1.26.5",
    "requests==2.27.1",
    "six==1.16.0",
]

for device_id in good_devices:
    produce_message(DEVICE_REQUIREMENTS, device_id, good_dependencies)