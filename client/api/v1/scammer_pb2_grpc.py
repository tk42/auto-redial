# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from api.v1 import scammer_pb2 as api_dot_v1_dot_scammer__pb2


class ScammerStoreServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.GetScammer = channel.unary_unary(
                '/api.v1.ScammerStoreService/GetScammer',
                request_serializer=api_dot_v1_dot_scammer__pb2.GetScammerRequest.SerializeToString,
                response_deserializer=api_dot_v1_dot_scammer__pb2.GetScammerResponse.FromString,
                )
        self.PutScammer = channel.unary_unary(
                '/api.v1.ScammerStoreService/PutScammer',
                request_serializer=api_dot_v1_dot_scammer__pb2.PutScammerRequest.SerializeToString,
                response_deserializer=api_dot_v1_dot_scammer__pb2.PutScammerResponse.FromString,
                )
        self.UpdateScammer = channel.unary_unary(
                '/api.v1.ScammerStoreService/UpdateScammer',
                request_serializer=api_dot_v1_dot_scammer__pb2.UpdateScammerRequest.SerializeToString,
                response_deserializer=api_dot_v1_dot_scammer__pb2.UpdateScammerResponse.FromString,
                )
        self.DeleteScammer = channel.unary_unary(
                '/api.v1.ScammerStoreService/DeleteScammer',
                request_serializer=api_dot_v1_dot_scammer__pb2.DeleteScammerRequest.SerializeToString,
                response_deserializer=api_dot_v1_dot_scammer__pb2.DeleteScammerResponse.FromString,
                )


class ScammerStoreServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def GetScammer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def PutScammer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def UpdateScammer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DeleteScammer(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_ScammerStoreServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'GetScammer': grpc.unary_unary_rpc_method_handler(
                    servicer.GetScammer,
                    request_deserializer=api_dot_v1_dot_scammer__pb2.GetScammerRequest.FromString,
                    response_serializer=api_dot_v1_dot_scammer__pb2.GetScammerResponse.SerializeToString,
            ),
            'PutScammer': grpc.unary_unary_rpc_method_handler(
                    servicer.PutScammer,
                    request_deserializer=api_dot_v1_dot_scammer__pb2.PutScammerRequest.FromString,
                    response_serializer=api_dot_v1_dot_scammer__pb2.PutScammerResponse.SerializeToString,
            ),
            'UpdateScammer': grpc.unary_unary_rpc_method_handler(
                    servicer.UpdateScammer,
                    request_deserializer=api_dot_v1_dot_scammer__pb2.UpdateScammerRequest.FromString,
                    response_serializer=api_dot_v1_dot_scammer__pb2.UpdateScammerResponse.SerializeToString,
            ),
            'DeleteScammer': grpc.unary_unary_rpc_method_handler(
                    servicer.DeleteScammer,
                    request_deserializer=api_dot_v1_dot_scammer__pb2.DeleteScammerRequest.FromString,
                    response_serializer=api_dot_v1_dot_scammer__pb2.DeleteScammerResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'api.v1.ScammerStoreService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class ScammerStoreService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def GetScammer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.v1.ScammerStoreService/GetScammer',
            api_dot_v1_dot_scammer__pb2.GetScammerRequest.SerializeToString,
            api_dot_v1_dot_scammer__pb2.GetScammerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def PutScammer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.v1.ScammerStoreService/PutScammer',
            api_dot_v1_dot_scammer__pb2.PutScammerRequest.SerializeToString,
            api_dot_v1_dot_scammer__pb2.PutScammerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def UpdateScammer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.v1.ScammerStoreService/UpdateScammer',
            api_dot_v1_dot_scammer__pb2.UpdateScammerRequest.SerializeToString,
            api_dot_v1_dot_scammer__pb2.UpdateScammerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)

    @staticmethod
    def DeleteScammer(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/api.v1.ScammerStoreService/DeleteScammer',
            api_dot_v1_dot_scammer__pb2.DeleteScammerRequest.SerializeToString,
            api_dot_v1_dot_scammer__pb2.DeleteScammerResponse.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
