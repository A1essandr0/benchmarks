from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class RawEvent(_message.Message):
    __slots__ = ["source", "event_name", "event_status", "created", "payout"]
    SOURCE_FIELD_NUMBER: _ClassVar[int]
    EVENT_NAME_FIELD_NUMBER: _ClassVar[int]
    EVENT_STATUS_FIELD_NUMBER: _ClassVar[int]
    CREATED_FIELD_NUMBER: _ClassVar[int]
    PAYOUT_FIELD_NUMBER: _ClassVar[int]
    source: str
    event_name: str
    event_status: str
    created: str
    payout: str
    def __init__(self, source: _Optional[str] = ..., event_name: _Optional[str] = ..., event_status: _Optional[str] = ..., created: _Optional[str] = ..., payout: _Optional[str] = ...) -> None: ...
