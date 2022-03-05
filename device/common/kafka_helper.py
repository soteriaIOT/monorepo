from random import choice
from kafka import KafkaProducer
from confluent_kafka import Consumer
from dotenv import dotenv_values
from tags import COMMON_TAGS

DEVICE_ID = COMMON_TAGS.get("host")

config = dotenv_values()
KAFKA_IP = config.get("kafka_ip")

def consume_confluence(topic: str):
    consumer = Consumer({
        'bootstrap.servers': KAFKA_IP,
        'group.id': DEVICE_ID,
        'auto.offset.reset': 'earliest'
    })
    consumer.subscribe([topic])
    while True:
        try:
            msg = consumer.poll(1.0)
            if msg is None:
                continue
            if msg.error():
                print("Consumer error: {}".format(msg.error()))
                continue
            yield msg
        except Exception as e:
            print(e)
            return


def produce_message(topic: str, key: str, value: str):
    producer = KafkaProducer(
        bootstrap_servers=KAFKA_IP,
        key_serializer=str.encode,
        value_serializer=str.encode,
    )
    future = producer.send(topic, key=key, value=value)
    result = future.get(timeout=60)
    producer.flush()
    producer.close()
    print("DONE", result)


def read_requirements(requirements_file: str):
    with open(requirements_file) as f:
        return [line.strip() for line in f.readlines()]