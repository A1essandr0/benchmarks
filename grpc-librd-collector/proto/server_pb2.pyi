from common.proto import rawevent_pb2 as _rawevent_pb2
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class Response(_message.Message):
    __slots__ = ["status", "event_name"]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    EVENT_NAME_FIELD_NUMBER: _ClassVar[int]
    status: str
    event_name: str
    def __init__(self, status: _Optional[str] = ..., event_name: _Optional[str] = ...) -> None: ...
