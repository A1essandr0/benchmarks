from typing import Optional

import faust
from faust.serializers import codecs

from event import Event
from config import config


class PydanticSerializer(codecs.Codec):
    def __init__(self, cls_type):
        self.cls_type = cls_type
        super(self.__class__, self).__init__()

    def _dumps(self, cls) -> bytes:
        return cls.json().encode()

    def _loads(self, s: bytes):
        cls_impl = self.cls_type.parse_raw(s)
        return cls_impl


app = faust.App(
    f"faust-sender",
    broker=config["KAFKA_BROKER"],
    web_host=config["HOST"],
    web_port=config["PORT"],
)

topic_from = app.topic(
    config["KAFKA_EVENTS_TOPIC_FROM"],
    partitions=2,
    value_serializer=PydanticSerializer(Event)
)
topic_to = app.topic(
    config["KAFKA_EVENTS_TOPIC_TO"],
    partitions=2,
    value_serializer=PydanticSerializer(Event)
)


@app.agent(topic_from)
async def on_event(stream) -> None:
    async for msg_key, event in stream.items():
        print(f"Received :: {event=}")
        await topic_to.send(key="key", value=event)
        print("...sent further...")