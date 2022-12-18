from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class AnalyseRequest(_message.Message):
    __slots__ = ["cargo_filled", "cargo_total", "distance", "timestamp"]
    CARGO_FILLED_FIELD_NUMBER: _ClassVar[int]
    CARGO_TOTAL_FIELD_NUMBER: _ClassVar[int]
    DISTANCE_FIELD_NUMBER: _ClassVar[int]
    TIMESTAMP_FIELD_NUMBER: _ClassVar[int]
    cargo_filled: int
    cargo_total: int
    distance: int
    timestamp: int
    def __init__(self, distance: _Optional[int] = ..., timestamp: _Optional[int] = ..., cargo_total: _Optional[int] = ..., cargo_filled: _Optional[int] = ...) -> None: ...

class AnalyseResponse(_message.Message):
    __slots__ = ["time_spent"]
    TIME_SPENT_FIELD_NUMBER: _ClassVar[int]
    time_spent: float
    def __init__(self, time_spent: _Optional[float] = ...) -> None: ...
