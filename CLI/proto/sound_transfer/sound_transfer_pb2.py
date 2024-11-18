# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: sound_transfer.proto
# Protobuf Python Version: 5.28.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import runtime_version as _runtime_version
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
_runtime_version.ValidateProtobufRuntimeVersion(
    _runtime_version.Domain.PUBLIC,
    5,
    28,
    1,
    '',
    'sound_transfer.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14sound_transfer.proto\"1\n\x0cSoundRequest\x12\x12\n\nsound_data\x18\x01 \x01(\x0c\x12\r\n\x05\x66lags\x18\x02 \x03(\t\"\x1d\n\rSoundResponse\x12\x0c\n\x04text\x18\x03 \x01(\t\"\x1b\n\x0bTextMessage\x12\x0c\n\x04text\x18\x04 \x01(\t\"2\n\x13SoundStreamResponse\x12\x0c\n\x04text\x18\x05 \x01(\t\x12\r\n\x05\x66lags\x18\x06 \x03(\t\"3\n\x0eSpeakerAndLine\x12\x0c\n\x04text\x18\x07 \x01(\t\x12\x13\n\x0bspeakerName\x18\x08 \x01(\t2\xa6\x02\n\x0cSoundService\x12.\n\x0eTestConnection\x12\x0c.TextMessage\x1a\x0c.TextMessage\"\x00\x12\x30\n\rSendSoundFile\x12\r.SoundRequest\x1a\x0e.SoundResponse\"\x00\x12:\n\x0fStreamSoundFile\x12\r.SoundRequest\x1a\x14.SoundStreamResponse(\x01\x30\x01\x12\x41\n\x18SendSoundFileTranslation\x12\r.SoundRequest\x1a\x14.SoundStreamResponse0\x01\x12\x35\n\x11\x44iarizateSpeakers\x12\r.SoundRequest\x1a\x0f.SpeakerAndLine0\x01\x42H\n\x15io.grpc.soundtransferB\x12SoundTransferProtoP\x01Z\x19inzynierka/sound_transferb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'sound_transfer_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\025io.grpc.soundtransferB\022SoundTransferProtoP\001Z\031inzynierka/sound_transfer'
  _globals['_SOUNDREQUEST']._serialized_start=24
  _globals['_SOUNDREQUEST']._serialized_end=73
  _globals['_SOUNDRESPONSE']._serialized_start=75
  _globals['_SOUNDRESPONSE']._serialized_end=104
  _globals['_TEXTMESSAGE']._serialized_start=106
  _globals['_TEXTMESSAGE']._serialized_end=133
  _globals['_SOUNDSTREAMRESPONSE']._serialized_start=135
  _globals['_SOUNDSTREAMRESPONSE']._serialized_end=185
  _globals['_SPEAKERANDLINE']._serialized_start=187
  _globals['_SPEAKERANDLINE']._serialized_end=238
  _globals['_SOUNDSERVICE']._serialized_start=241
  _globals['_SOUNDSERVICE']._serialized_end=535
# @@protoc_insertion_point(module_scope)
