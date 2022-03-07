from typing import List
from typing import Union

from signal import signal
from signal import SIGINT
from signal import SIGTERM

from kafka import KafkaProducer
from confluent_kafka import Consumer
from dotenv import dotenv_values
from common.tags import COMMON_TAGS

DEVICE_ID = COMMON_TAGS.get("host")

config = dotenv_values()
KAFKA_IP = config.get("kafka_ip")


def consume_confluence(topic: str):
    consumer = Consumer(
        {
            "bootstrap.servers": KAFKA_IP,
            "group.id": DEVICE_ID,
            "auto.offset.reset": "earliest",
        }
    )
    consumer.subscribe([topic])
    signal(SIGINT, lambda s, f: consumer.close())
    signal(SIGTERM, lambda s, f: consumer.close())
    while True:
        try:
            msg = consumer.poll(1.0)
            msg = None
            if msg is None:
                continue
            if msg.error():
                print("Consumer error: {}".format(msg.error()))
                continue
            yield msg
        except KeyboardInterrupt as e:
            break
        except Exception as e:
            print(e)
            break
    consumer.close()


def produce_message(topic: str, key: str, value: Union[str, List]):
    if isinstance(value, list):
        value = "\n".join(value)
    producer = KafkaProducer(
        bootstrap_servers=KAFKA_IP,
        key_serializer=str.encode,
        value_serializer=str.encode,
    )
    future = producer.send(topic, key=key, value=value)
    _ = future.get(timeout=60)
    producer.flush()
    producer.close()
