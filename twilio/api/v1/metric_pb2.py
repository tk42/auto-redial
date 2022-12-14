# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: api/v1/metric.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.type import datetime_pb2 as google_dot_type_dot_datetime__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x13\x61pi/v1/metric.proto\x12\x06\x61pi.v1\x1a\x1agoogle/type/datetime.proto\x1a\x1egoogle/protobuf/duration.proto\"\x96\x02\n\x06Metric\x12\x34\n\ncreated_at\x18\x01 \x01(\x0b\x32\x15.google.type.DateTimeR\tcreatedAt\x12\x14\n\x05\x63\x61lls\x18\x02 \x01(\x03R\x05\x63\x61lls\x12\x1a\n\x08scammers\x18\x03 \x01(\x03R\x08scammers\x12\x1c\n\tinactives\x18\x04 \x01(\x03R\tinactives\x12\x36\n\tcall_time\x18\x05 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08\x63\x61llTime\x12\x36\n\ttalk_time\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08talkTime\x12\x16\n\x06\x61mount\x18\x07 \x01(\x03R\x06\x61mount\"d\n\x10GetMetricRequest\x12)\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x15.google.type.DateTimeR\x04\x66rom\x12%\n\x02to\x18\x02 \x01(\x0b\x32\x15.google.type.DateTimeR\x02to\"=\n\x11GetMetricResponse\x12(\n\x07metrics\x18\x01 \x03(\x0b\x32\x0e.api.v1.MetricR\x07metrics\"\xa0\x02\n\x10PutMetricRequest\x12\x34\n\ncreated_at\x18\x01 \x01(\x0b\x32\x15.google.type.DateTimeR\tcreatedAt\x12\x14\n\x05\x63\x61lls\x18\x02 \x01(\x03R\x05\x63\x61lls\x12\x1a\n\x08scammers\x18\x03 \x01(\x03R\x08scammers\x12\x1c\n\tinactives\x18\x04 \x01(\x03R\tinactives\x12\x36\n\tcall_time\x18\x05 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08\x63\x61llTime\x12\x36\n\ttalk_time\x18\x06 \x01(\x0b\x32\x19.google.protobuf.DurationR\x08talkTime\x12\x16\n\x06\x61mount\x18\x07 \x01(\x03R\x06\x61mount\";\n\x11PutMetricResponse\x12&\n\x06metric\x18\x01 \x01(\x0b\x32\x0e.api.v1.MetricR\x06metric\"g\n\x13\x44\x65leteMetricRequest\x12)\n\x04\x66rom\x18\x01 \x01(\x0b\x32\x15.google.type.DateTimeR\x04\x66rom\x12%\n\x02to\x18\x02 \x01(\x0b\x32\x15.google.type.DateTimeR\x02to\"\x16\n\x14\x44\x65leteMetricResponse2\xe9\x01\n\x12MetricStoreService\x12\x42\n\tGetMetric\x12\x18.api.v1.GetMetricRequest\x1a\x19.api.v1.GetMetricResponse\"\x00\x12\x42\n\tPutMetric\x12\x18.api.v1.PutMetricRequest\x1a\x19.api.v1.PutMetricResponse\"\x00\x12K\n\x0c\x44\x65leteMetric\x12\x1b.api.v1.DeleteMetricRequest\x1a\x1c.api.v1.DeleteMetricResponse\"\x00\x42#Z!github.com/tk42/auto-redial;apiv1b\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'api.v1.metric_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z!github.com/tk42/auto-redial;apiv1'
  _METRIC._serialized_start=92
  _METRIC._serialized_end=370
  _GETMETRICREQUEST._serialized_start=372
  _GETMETRICREQUEST._serialized_end=472
  _GETMETRICRESPONSE._serialized_start=474
  _GETMETRICRESPONSE._serialized_end=535
  _PUTMETRICREQUEST._serialized_start=538
  _PUTMETRICREQUEST._serialized_end=826
  _PUTMETRICRESPONSE._serialized_start=828
  _PUTMETRICRESPONSE._serialized_end=887
  _DELETEMETRICREQUEST._serialized_start=889
  _DELETEMETRICREQUEST._serialized_end=992
  _DELETEMETRICRESPONSE._serialized_start=994
  _DELETEMETRICRESPONSE._serialized_end=1016
  _METRICSTORESERVICE._serialized_start=1019
  _METRICSTORESERVICE._serialized_end=1252
# @@protoc_insertion_point(module_scope)
