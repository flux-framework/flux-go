# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: flux.proto
# Protobuf Python Version: 5.29.0
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder

_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC, 5, 29, 0, "", "flux.proto"
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(
    b'\n\nflux.proto\x12\x04\x66lux" \n\rSubmitRequest\x12\x0f\n\x07jobspec\x18\x01 \x01(\t"\x88\x01\n\x0eSubmitResponse\x12/\n\x06status\x18\x01 \x01(\x0e\x32\x1f.flux.SubmitResponse.ResultType"E\n\nResultType\x12\x12\n\x0eSUBMIT_SUCCESS\x10\x00\x12\x10\n\x0cSUBMIT_ERROR\x10\x01\x12\x11\n\rSUBMIT_DENIED\x10\x02"\x1c\n\x0bInfoRequest\x12\r\n\x05jobid\x18\x01 \x01(\t"~\n\x0cInfoResponse\x12-\n\x06status\x18\x01 \x01(\x0e\x32\x1d.flux.InfoResponse.ResultType"?\n\nResultType\x12\x10\n\x0cINFO_SUCCESS\x10\x00\x12\x0e\n\nINFO_ERROR\x10\x01\x12\x0f\n\x0bINFO_DENIED\x10\x02"\r\n\x0bJobsRequest"(\n\x07JobInfo\x12\r\n\x05jobid\x18\x01 \x01(\t\x12\x0e\n\x06status\x18\x02 \x01(\t"\x8c\x01\n\x0cJobsResponse\x12-\n\x06status\x18\x01 \x01(\x0e\x32\x1d.flux.JobsResponse.ResultType\x12\x0c\n\x04jobs\x18\x02 \x01(\t"?\n\nResultType\x12\x10\n\x0cJOBS_SUCCESS\x10\x00\x12\x0e\n\nJOBS_ERROR\x10\x01\x12\x0f\n\x0bJOBS_DENIED\x10\x02\x32u\n\x0b\x46luxService\x12\x35\n\x06Submit\x12\x13.flux.SubmitRequest\x1a\x14.flux.SubmitResponse"\x00\x12/\n\x04Jobs\x12\x11.flux.JobsRequest\x1a\x12.flux.JobsResponse"\x00\x42\x31Z/github.com/flux-framework/flux-go/pkg/flux-grpcb\x06proto3'
)

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, "flux_pb2", _globals)
if not _descriptor._USE_C_DESCRIPTORS:
    _globals["DESCRIPTOR"]._loaded_options = None
    _globals["DESCRIPTOR"]._serialized_options = (
        b"Z/github.com/flux-framework/flux-go/pkg/flux-grpc"
    )
    _globals["_SUBMITREQUEST"]._serialized_start = 20
    _globals["_SUBMITREQUEST"]._serialized_end = 52
    _globals["_SUBMITRESPONSE"]._serialized_start = 55
    _globals["_SUBMITRESPONSE"]._serialized_end = 191
    _globals["_SUBMITRESPONSE_RESULTTYPE"]._serialized_start = 122
    _globals["_SUBMITRESPONSE_RESULTTYPE"]._serialized_end = 191
    _globals["_INFOREQUEST"]._serialized_start = 193
    _globals["_INFOREQUEST"]._serialized_end = 221
    _globals["_INFORESPONSE"]._serialized_start = 223
    _globals["_INFORESPONSE"]._serialized_end = 349
    _globals["_INFORESPONSE_RESULTTYPE"]._serialized_start = 286
    _globals["_INFORESPONSE_RESULTTYPE"]._serialized_end = 349
    _globals["_JOBSREQUEST"]._serialized_start = 351
    _globals["_JOBSREQUEST"]._serialized_end = 364
    _globals["_JOBINFO"]._serialized_start = 366
    _globals["_JOBINFO"]._serialized_end = 406
    _globals["_JOBSRESPONSE"]._serialized_start = 409
    _globals["_JOBSRESPONSE"]._serialized_end = 549
    _globals["_JOBSRESPONSE_RESULTTYPE"]._serialized_start = 486
    _globals["_JOBSRESPONSE_RESULTTYPE"]._serialized_end = 549
    _globals["_FLUXSERVICE"]._serialized_start = 551
    _globals["_FLUXSERVICE"]._serialized_end = 668
# @@protoc_insertion_point(module_scope)
