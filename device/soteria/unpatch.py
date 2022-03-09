from soteria.kafka_helper import produce_message, consume_confluence
from common.tags import COMMON_TAGS

DEVICE_UPDATES = "device-updates"
DEVICE_ID = "testing00001"

produce_message(DEVICE_UPDATES, DEVICE_ID, "urllib3==1.24.3")