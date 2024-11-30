# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
"""Client and server classes corresponding to protobuf-defined services."""
import grpc
import warnings

import sound_transfer_pb2 as sound__transfer__pb2

GRPC_GENERATED_VERSION = '1.68.0'
GRPC_VERSION = grpc.__version__
_version_not_supported = False

try:
    from grpc._utilities import first_version_is_lower
    _version_not_supported = first_version_is_lower(GRPC_VERSION, GRPC_GENERATED_VERSION)
except ImportError:
    _version_not_supported = True

if _version_not_supported:
    raise RuntimeError(
        f'The grpc package installed is at version {GRPC_VERSION},'
        + f' but the generated code in sound_transfer_pb2_grpc.py depends on'
        + f' grpcio>={GRPC_GENERATED_VERSION}.'
        + f' Please upgrade your grpc module to grpcio>={GRPC_GENERATED_VERSION}'
        + f' or downgrade your generated code using grpcio-tools<={GRPC_VERSION}.'
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
        self.TranscribeFile = channel.unary_unary(
                '/SoundService/TranscribeFile',
                request_serializer=sound__transfer__pb2.TranscriptionRequest.SerializeToString,
                response_deserializer=sound__transfer__pb2.SoundResponse.FromString,
                _registered_method=True)
        self.TranscribeLive = channel.stream_stream(
                '/SoundService/TranscribeLive',
                request_serializer=sound__transfer__pb2.TranscirptionLiveRequest.SerializeToString,
                response_deserializer=sound__transfer__pb2.SoundStreamResponse.FromString,
                _registered_method=True)
        self.TranslateFile = channel.unary_stream(
                '/SoundService/TranslateFile',
                request_serializer=sound__transfer__pb2.TranslationRequest.SerializeToString,
                response_deserializer=sound__transfer__pb2.SoundResponse.FromString,
                _registered_method=True)
        self.DiarizateFile = channel.unary_unary(
                '/SoundService/DiarizateFile',
                request_serializer=sound__transfer__pb2.TranscriptionRequest.SerializeToString,
                response_deserializer=sound__transfer__pb2.SpeakerAndLineResponse.FromString,
                _registered_method=True)


class SoundServiceServicer(object):
    """Missing associated documentation comment in .proto file."""

    def TestConnection(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TranscribeFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TranscribeLive(self, request_iterator, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def TranslateFile(self, request, context):
        """Missing associated documentation comment in .proto file."""
        context.set_code(grpc.StatusCode.UNIMPLEMENTED)
        context.set_details('Method not implemented!')
        raise NotImplementedError('Method not implemented!')

    def DiarizateFile(self, request, context):
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
            'TranscribeFile': grpc.unary_unary_rpc_method_handler(
                    servicer.TranscribeFile,
                    request_deserializer=sound__transfer__pb2.TranscriptionRequest.FromString,
                    response_serializer=sound__transfer__pb2.SoundResponse.SerializeToString,
            ),
            'TranscribeLive': grpc.stream_stream_rpc_method_handler(
                    servicer.TranscribeLive,
                    request_deserializer=sound__transfer__pb2.TranscirptionLiveRequest.FromString,
                    response_serializer=sound__transfer__pb2.SoundStreamResponse.SerializeToString,
            ),
            'TranslateFile': grpc.unary_stream_rpc_method_handler(
                    servicer.TranslateFile,
                    request_deserializer=sound__transfer__pb2.TranslationRequest.FromString,
                    response_serializer=sound__transfer__pb2.SoundResponse.SerializeToString,
            ),
            'DiarizateFile': grpc.unary_unary_rpc_method_handler(
                    servicer.DiarizateFile,
                    request_deserializer=sound__transfer__pb2.TranscriptionRequest.FromString,
                    response_serializer=sound__transfer__pb2.SpeakerAndLineResponse.SerializeToString,
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
    def TranscribeFile(request,
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
            '/SoundService/TranscribeFile',
            sound__transfer__pb2.TranscriptionRequest.SerializeToString,
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
    def TranscribeLive(request_iterator,
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
            '/SoundService/TranscribeLive',
            sound__transfer__pb2.TranscirptionLiveRequest.SerializeToString,
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

    @staticmethod
    def TranslateFile(request,
            target,
            options=(),
            channel_credentials=None,
            call_credentials=None,
            insecure=False,
            compression=None,
            wait_for_ready=None,
            timeout=None,
            metadata=None):
        return grpc.experimental.unary_stream(
            request,
            target,
            '/SoundService/TranslateFile',
            sound__transfer__pb2.TranslationRequest.SerializeToString,
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
    def DiarizateFile(request,
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
            '/SoundService/DiarizateFile',
            sound__transfer__pb2.TranscriptionRequest.SerializeToString,
            sound__transfer__pb2.SpeakerAndLineResponse.FromString,
            options,
            channel_credentials,
            insecure,
            call_credentials,
            compression,
            wait_for_ready,
            timeout,
            metadata,
            _registered_method=True)