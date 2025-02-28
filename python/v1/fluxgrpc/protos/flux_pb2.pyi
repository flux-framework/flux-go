from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class SubmitRequest(_message.Message):
    __slots__ = ("jobspec",)
    JOBSPEC_FIELD_NUMBER: _ClassVar[int]
    jobspec: str
    def __init__(self, jobspec: _Optional[str] = ...) -> None: ...

class SubmitResponse(_message.Message):
    __slots__ = ("status",)
    class ResultType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        SUBMIT_SUCCESS: _ClassVar[SubmitResponse.ResultType]
        SUBMIT_ERROR: _ClassVar[SubmitResponse.ResultType]
        SUBMIT_DENIED: _ClassVar[SubmitResponse.ResultType]
    SUBMIT_SUCCESS: SubmitResponse.ResultType
    SUBMIT_ERROR: SubmitResponse.ResultType
    SUBMIT_DENIED: SubmitResponse.ResultType
    STATUS_FIELD_NUMBER: _ClassVar[int]
    status: SubmitResponse.ResultType
    def __init__(self, status: _Optional[_Union[SubmitResponse.ResultType, str]] = ...) -> None: ...

class InfoRequest(_message.Message):
    __slots__ = ("jobid",)
    JOBID_FIELD_NUMBER: _ClassVar[int]
    jobid: str
    def __init__(self, jobid: _Optional[str] = ...) -> None: ...

class InfoResponse(_message.Message):
    __slots__ = ("status",)
    class ResultType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        INFO_SUCCESS: _ClassVar[InfoResponse.ResultType]
        INFO_ERROR: _ClassVar[InfoResponse.ResultType]
        INFO_DENIED: _ClassVar[InfoResponse.ResultType]
    INFO_SUCCESS: InfoResponse.ResultType
    INFO_ERROR: InfoResponse.ResultType
    INFO_DENIED: InfoResponse.ResultType
    STATUS_FIELD_NUMBER: _ClassVar[int]
    status: InfoResponse.ResultType
    def __init__(self, status: _Optional[_Union[InfoResponse.ResultType, str]] = ...) -> None: ...

class JobsRequest(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class JobInfo(_message.Message):
    __slots__ = ("jobid", "status")
    JOBID_FIELD_NUMBER: _ClassVar[int]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    jobid: str
    status: str
    def __init__(self, jobid: _Optional[str] = ..., status: _Optional[str] = ...) -> None: ...

class JobsResponse(_message.Message):
    __slots__ = ("status", "jobs")
    class ResultType(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
        __slots__ = ()
        JOBS_SUCCESS: _ClassVar[JobsResponse.ResultType]
        JOBS_ERROR: _ClassVar[JobsResponse.ResultType]
        JOBS_DENIED: _ClassVar[JobsResponse.ResultType]
    JOBS_SUCCESS: JobsResponse.ResultType
    JOBS_ERROR: JobsResponse.ResultType
    JOBS_DENIED: JobsResponse.ResultType
    STATUS_FIELD_NUMBER: _ClassVar[int]
    JOBS_FIELD_NUMBER: _ClassVar[int]
    status: JobsResponse.ResultType
    jobs: str
    def __init__(self, status: _Optional[_Union[JobsResponse.ResultType, str]] = ..., jobs: _Optional[str] = ...) -> None: ...
