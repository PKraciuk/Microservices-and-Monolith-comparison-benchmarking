# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc

from proto import geo_pb2 as geo__pb2


class GeoStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.Nearby = channel.unary_unary(
                '/geo.Geo/Nearby',
                request_serializer=geo__pb2.Request.SerializeToString,
                response_deserializer=geo__pb2.Result.FromString,
                )


class GeoServicer(object):
    """Missing associated documentation comment in .proto file."""

    def Nearby(self, request, context):
        """Finds the hotels nearby given latitude/longitude.
        """
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_GeoServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'Nearby': grpc.unary_unary_rpc_method_handler(
                    servicer.Nearby,
                    request_deserializer=geo__pb2.Request.FromString,
                    response_serializer=geo__pb2.Result.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'geo.Geo', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


 # This class is part of an EXPERIMENTAL API.
class Geo(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def Nearby(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(request, target, '/geo.Geo/Nearby',
            geo__pb2.Request.SerializeToString,
            geo__pb2.Result.FromString,
            options, channel_credentials,
            insecure, call_credentials, compression, wait_for_ready, timeout, metadata)
