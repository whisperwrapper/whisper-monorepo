# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import sound_transfer_pb2 as sound__transfer__pb2

GRPC_GENERATED_VERSION = '1.64.1'
GRPC_VERSION = grpc.__version__
EXPECTED_ERROR_RELEASE = '1.65.0'
SCHEDULED_RELEASE_DATE = 'June 25, 2024'
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    warnings.warn(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in sound_transfer_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
        + f' This warning will become an error in {EXPECTED_ERROR_RELEASE},'
        + f' scheduled for release on {SCHEDULED_RELEASE_DATE}.',
        RuntimeWarning
    )


class SoundServiceStub(object):
    """Missing associated documentation comment in .proto file."""

    def __init__(self, channel):
        """Constructor.

        Args:
            channel: A grpc.Channel.
        """
        self.TestConnection = channel.unary_unary(
                '/SoundService/TestConnection',
                request_serializer=sound__transfer__pb2.TextMessage.SerializeToString,
                response_deserializer=sound__transfer__pb2.TextMessage.FromString,
                _registered_method=True)
        self.SendSoundFile = channel.unary_unary(
                '/SoundService/SendSoundFile',
                request_serializer=sound__transfer__pb2.SoundRequest.SerializeToString,
                response_deserializer=sound__transfer__pb2.SoundResponse.FromString,
                _registered_method=True)
        self.StreamSoundFile = channel.stream_stream(
                '/SoundService/StreamSoundFile',
                request_serializer=sound__transfer__pb2.SoundRequest.SerializeToString,
                response_deserializer=sound__transfer__pb2.SoundStreamResponse.FromString,
                _registered_method=True)


class SoundServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def TestConnection(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def SendSoundFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def StreamSoundFile(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')


def add_SoundServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
            'TestConnection': grpc.unary_unary_rpc_method_handler(
                    servicer.TestConnection,
                    request_deserializer=sound__transfer__pb2.TextMessage.FromString,
                    response_serializer=sound__transfer__pb2.TextMessage.SerializeToString,
            ),
            'SendSoundFile': grpc.unary_unary_rpc_method_handler(
                    servicer.SendSoundFile,
                    request_deserializer=sound__transfer__pb2.SoundRequest.FromString,
                    response_serializer=sound__transfer__pb2.SoundResponse.SerializeToString,
            ),
            'StreamSoundFile': grpc.stream_stream_rpc_method_handler(
                    servicer.StreamSoundFile,
                    request_deserializer=sound__transfer__pb2.SoundRequest.FromString,
                    response_serializer=sound__transfer__pb2.SoundStreamResponse.SerializeToString,
            ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
            'SoundService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))
    server.add_registered_method_handlers('SoundService', rpc_method_handlers)


 # This class is part of an EXPERIMENTAL API.
class SoundService(object):
    """Missing associated documentation comment in .proto file."""

    @staticmethod
    def TestConnection(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/SoundService/TestConnection',
            sound__transfer__pb2.TextMessage.SerializeToString,
            sound__transfer__pb2.TextMessage.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def SendSoundFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_unary(
            request,
            target,
            '/SoundService/SendSoundFile',
            sound__transfer__pb2.SoundRequest.SerializeToString,
            sound__transfer__pb2.SoundResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)

    @staticmethod
    def StreamSoundFile(request_iterator,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.stream_stream(
            request_iterator,
            target,
            '/SoundService/StreamSoundFile',
            sound__transfer__pb2.SoundRequest.SerializeToString,
            sound__transfer__pb2.SoundStreamResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)
