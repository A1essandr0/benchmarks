# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/proto/rawevent.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x1b\x63ommon/proto/rawevent.proto\x12\x17go_grpc_collector_proto\"e\n\x08RawEvent\x12\x0e\n\x06source\x18\x01 \x01(\t\x12\x12\n\nevent_name\x18\x02 \x01(\t\x12\x14\n\x0c\x65vent_status\x18\x03 \x01(\t\x12\x0f\n\x07\x63reated\x18\x04 \x01(\t\x12\x0e\n\x06payout\x18\x05 \x01(\tB\x1bZ\x19./go_grpc_collector_protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'common.proto.rawevent_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\031./go_grpc_collector_proto'
  _globals['_RAWEVENT']._serialized_start=56
  _globals['_RAWEVENT']._serialized_end=157
# @@protoc_insertion_point(module_scope)
