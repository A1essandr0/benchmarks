# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: common/proto/server.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from common.proto import rawevent_pb2 as common_dot_proto_dot_rawevent__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x19\x63ommon/proto/server.proto\x12\x17go_grpc_collector_proto\x1a\x1b\x63ommon/proto/rawevent.proto\".\n\x08Response\x12\x0e\n\x06status\x18\x01 \x01(\t\x12\x12\n\nevent_name\x18\x02 \x01(\t2j\n\x10\x43ollectorService\x12V\n\x0cPostRawEvent\x12!.go_grpc_collector_proto.RawEvent\x1a!.go_grpc_collector_proto.Response\"\x00\x42\x1bZ\x19./go_grpc_collector_protob\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'common.proto.server_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\031./go_grpc_collector_proto'
  _globals['_RESPONSE']._serialized_start=83
  _globals['_RESPONSE']._serialized_end=129
  _globals['_COLLECTORSERVICE']._serialized_start=131
  _globals['_COLLECTORSERVICE']._serialized_end=237
# @@protoc_insertion_point(module_scope)