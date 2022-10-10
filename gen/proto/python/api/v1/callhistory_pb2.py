# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/v1/callhistory.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.type import datetime_pb2 as google_dot_type_dot_datetime__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x18\x61pi/v1/callhistory.proto\x12\x06\x61pi.v1\x1a\x1agoogle/type/datetime.proto\x1a\x1egoogle/protobuf/duration.proto\"\x87\x02\n\x0b\x43\x61llHistory\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n\nscammer_id\x18\x02 \x01(\tR\tscammerId\x12.\n\x07\x63\x61ll_at\x18\x03 \x01(\x0b\x32\x15.google.type.DateTimeR\x06\x63\x61llAt\x12\x36\n\tcall_time\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08\x63\x61llTime\x12\x16\n\x06result\x18\x05 \x01(\x08R\x06result\x12;\n\ttalk_time\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationH\x00R\x08talkTime\x88\x01\x01\x42\x0c\n\n_talk_time\"\'\n\x15GetCallHistoryRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"A\n\x16GetCallHistoryResponse\x12\'\n\x04\x63\x61ll\x18\x01 \x03(\x0b\x32\x13.api.v1.CallHistoryR\x04\x63\x61ll\"\x91\x02\n\x15PutCallHistoryRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n\nscammer_id\x18\x02 \x01(\tR\tscammerId\x12.\n\x07\x63\x61ll_at\x18\x03 \x01(\x0b\x32\x15.google.type.DateTimeR\x06\x63\x61llAt\x12\x36\n\tcall_time\x18\x04 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08\x63\x61llTime\x12\x16\n\x06result\x18\x05 \x01(\x08R\x06result\x12;\n\ttalk_time\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationH\x00R\x08talkTime\x88\x01\x01\x42\x0c\n\n_talk_time\"A\n\x16PutCallHistoryResponse\x12\'\n\x04\x63\x61ll\x18\x01 \x01(\x0b\x32\x13.api.v1.CallHistoryR\x04\x63\x61ll\"*\n\x18\x44\x65leteCallHistoryRequest\x12\x0e\n\x02id\x18\x01 \x03(\tR\x02id\"\x1b\n\x19\x44\x65leteCallHistoryResponse2\x9b\x02\n\x17\x43\x61llHistoryStoreService\x12Q\n\x0eGetCallHistory\x12\x1d.api.v1.GetCallHistoryRequest\x1a\x1e.api.v1.GetCallHistoryResponse\"\x00\x12Q\n\x0ePutCallHistory\x12\x1d.api.v1.PutCallHistoryRequest\x1a\x1e.api.v1.PutCallHistoryResponse\"\x00\x12Z\n\x11\x44\x65leteCallHistory\x12 .api.v1.DeleteCallHistoryRequest\x1a!.api.v1.DeleteCallHistoryResponse\"\x00\x42,Z*github.com/tk42/sqlc-buf-pg-template;apiv1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.v1.callhistory_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z*github.com/tk42/sqlc-buf-pg-template;apiv1'
  _CALLHISTORY._serialized_start=97
  _CALLHISTORY._serialized_end=360
  _GETCALLHISTORYREQUEST._serialized_start=362
  _GETCALLHISTORYREQUEST._serialized_end=401
  _GETCALLHISTORYRESPONSE._serialized_start=403
  _GETCALLHISTORYRESPONSE._serialized_end=468
  _PUTCALLHISTORYREQUEST._serialized_start=471
  _PUTCALLHISTORYREQUEST._serialized_end=744
  _PUTCALLHISTORYRESPONSE._serialized_start=746
  _PUTCALLHISTORYRESPONSE._serialized_end=811
  _DELETECALLHISTORYREQUEST._serialized_start=813
  _DELETECALLHISTORYREQUEST._serialized_end=855
  _DELETECALLHISTORYRESPONSE._serialized_start=857
  _DELETECALLHISTORYRESPONSE._serialized_end=884
  _CALLHISTORYSTORESERVICE._serialized_start=887
  _CALLHISTORYSTORESERVICE._serialized_end=1170
# @@protoc_insertion_point(module_scope)
