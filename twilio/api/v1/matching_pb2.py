# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/v1/matching.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.type import datetime_pb2 as google_dot_type_dot_datetime__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x15\x61pi/v1/matching.proto\x12\x06\x61pi.v1\x1a\x1agoogle/type/datetime.proto\x1a\x1egoogle/protobuf/duration.proto\"\xad\x03\n\x08Matching\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x34\n\ncreated_at\x18\x02 \x01(\x0b\x32\x15.google.type.DateTimeR\tcreatedAt\x12#\n\rserial_number\x18\x03 \x01(\x03R\x0cserialNumber\x12\x1d\n\nscammer_id\x18\x04 \x03(\tR\tscammerId\x12\x17\n\x07\x63\x61ll_id\x18\x05 \x03(\tR\x06\x63\x61llId\x12\x18\n\x07matched\x18\x06 \x01(\x08R\x07matched\x12\x18\n\x07\x63hecked\x18\x07 \x01(\x08R\x07\x63hecked\x12;\n\x0bmatching_at\x18\x08 \x01(\x0b\x32\x15.google.type.DateTimeH\x00R\nmatchingAt\x88\x01\x01\x12;\n\ttalk_time\x18\t \x01(\x0b\x32\x19.google.protobuf.DurationH\x01R\x08talkTime\x88\x01\x01\x12#\n\ntranscript\x18\n \x01(\tH\x02R\ntranscript\x88\x01\x01\x42\x0e\n\x0c_matching_atB\x0c\n\n_talk_timeB\r\n\x0b_transcript\"$\n\x12GetMatchingRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\"C\n\x13GetMatchingResponse\x12,\n\x08matching\x18\x01 \x01(\x0b\x32\x10.api.v1.MatchingR\x08matching\"g\n\x13ListMatchingRequest\x12)\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x15.google.type.DateTimeR\x04\x66rom\x12%\n\x02to\x18\x02 \x01(\x0b\x32\x15.google.type.DateTimeR\x02to\"D\n\x14ListMatchingResponse\x12,\n\x08matching\x18\x01 \x03(\x0b\x32\x10.api.v1.MatchingR\x08matching\"\xc9\x02\n\x15UpdateMatchingRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x1d\n\x07matched\x18\x02 \x01(\x08H\x00R\x07matched\x88\x01\x01\x12\x1d\n\x07\x63hecked\x18\x03 \x01(\x08H\x01R\x07\x63hecked\x88\x01\x01\x12;\n\x0bmatching_at\x18\x04 \x01(\x0b\x32\x15.google.type.DateTimeH\x02R\nmatchingAt\x88\x01\x01\x12;\n\ttalk_time\x18\x05 \x01(\x0b\x32\x19.google.protobuf.DurationH\x03R\x08talkTime\x88\x01\x01\x12#\n\ntranscript\x18\x06 \x01(\tH\x04R\ntranscript\x88\x01\x01\x42\n\n\x08_matchedB\n\n\x08_checkedB\x0e\n\x0c_matching_atB\x0c\n\n_talk_timeB\r\n\x0b_transcript\"F\n\x16UpdateMatchingResponse\x12,\n\x08matching\x18\x01 \x01(\x0b\x32\x10.api.v1.MatchingR\x08matching\"\xb7\x03\n\x12PutMatchingRequest\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x34\n\ncreated_at\x18\x02 \x01(\x0b\x32\x15.google.type.DateTimeR\tcreatedAt\x12#\n\rserial_number\x18\x03 \x01(\x03R\x0cserialNumber\x12\x1d\n\nscammer_id\x18\x04 \x03(\tR\tscammerId\x12\x17\n\x07\x63\x61ll_id\x18\x05 \x03(\tR\x06\x63\x61llId\x12\x18\n\x07matched\x18\x06 \x01(\x08R\x07matched\x12\x18\n\x07\x63hecked\x18\x07 \x01(\x08R\x07\x63hecked\x12;\n\x0bmatching_at\x18\x08 \x01(\x0b\x32\x15.google.type.DateTimeH\x00R\nmatchingAt\x88\x01\x01\x12;\n\ttalk_time\x18\t \x01(\x0b\x32\x19.google.protobuf.DurationH\x01R\x08talkTime\x88\x01\x01\x12#\n\ntranscript\x18\n \x01(\tH\x02R\ntranscript\x88\x01\x01\x42\x0e\n\x0c_matching_atB\x0c\n\n_talk_timeB\r\n\x0b_transcript\"C\n\x13PutMatchingResponse\x12,\n\x08matching\x18\x01 \x01(\x0b\x32\x10.api.v1.MatchingR\x08matching\"8\n\x15\x44\x65leteMatchingRequest\x12\x1f\n\x0bmatching_id\x18\x01 \x01(\tR\nmatchingId\"\x18\n\x16\x44\x65leteMatchingResponse2\x9d\x03\n\x14MatchingStoreService\x12H\n\x0bGetMatching\x12\x1a.api.v1.GetMatchingRequest\x1a\x1b.api.v1.GetMatchingResponse\"\x00\x12K\n\x0cListMatching\x12\x1b.api.v1.ListMatchingRequest\x1a\x1c.api.v1.ListMatchingResponse\"\x00\x12H\n\x0bPutMatching\x12\x1a.api.v1.PutMatchingRequest\x1a\x1b.api.v1.PutMatchingResponse\"\x00\x12Q\n\x0eUpdateMatching\x12\x1d.api.v1.UpdateMatchingRequest\x1a\x1e.api.v1.UpdateMatchingResponse\"\x00\x12Q\n\x0e\x44\x65leteMatching\x12\x1d.api.v1.DeleteMatchingRequest\x1a\x1e.api.v1.DeleteMatchingResponse\"\x00\x42#Z!github.com/tk42/auto-redial;apiv1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.v1.matching_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z!github.com/tk42/auto-redial;apiv1'
  _MATCHING._serialized_start=94
  _MATCHING._serialized_end=523
  _GETMATCHINGREQUEST._serialized_start=525
  _GETMATCHINGREQUEST._serialized_end=561
  _GETMATCHINGRESPONSE._serialized_start=563
  _GETMATCHINGRESPONSE._serialized_end=630
  _LISTMATCHINGREQUEST._serialized_start=632
  _LISTMATCHINGREQUEST._serialized_end=735
  _LISTMATCHINGRESPONSE._serialized_start=737
  _LISTMATCHINGRESPONSE._serialized_end=805
  _UPDATEMATCHINGREQUEST._serialized_start=808
  _UPDATEMATCHINGREQUEST._serialized_end=1137
  _UPDATEMATCHINGRESPONSE._serialized_start=1139
  _UPDATEMATCHINGRESPONSE._serialized_end=1209
  _PUTMATCHINGREQUEST._serialized_start=1212
  _PUTMATCHINGREQUEST._serialized_end=1651
  _PUTMATCHINGRESPONSE._serialized_start=1653
  _PUTMATCHINGRESPONSE._serialized_end=1720
  _DELETEMATCHINGREQUEST._serialized_start=1722
  _DELETEMATCHINGREQUEST._serialized_end=1778
  _DELETEMATCHINGRESPONSE._serialized_start=1780
  _DELETEMATCHINGRESPONSE._serialized_end=1804
  _MATCHINGSTORESERVICE._serialized_start=1807
  _MATCHINGSTORESERVICE._serialized_end=2220
# @@protoc_insertion_point(module_scope)
