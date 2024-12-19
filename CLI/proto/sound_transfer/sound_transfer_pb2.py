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




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14sound_transfer.proto\"\x1b\n\x0bTextMessage\x12\x0c\n\x04text\x18\x01 \x01(\t\"C\n\x14TranscriptionRequest\x12\x12\n\nsound_data\x18\x01 \x01(\x0c\x12\x17\n\x0fsource_language\x18\x02 \x01(\t\"_\n\x12TranslationRequest\x12\x12\n\nsound_data\x18\x01 \x01(\x0c\x12\x17\n\x0fsource_language\x18\x02 \x01(\t\x12\x1c\n\x14translation_language\x18\x03 \x01(\t\".\n\x18TranscirptionLiveRequest\x12\x12\n\nsound_data\x18\x01 \x01(\x0c\"D\n\rSoundResponse\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x19\n\x11\x64\x65tected_language\x18\x02 \x01(\t\x12\n\n\x02id\x18\x03 \x01(\x05\"6\n\x13SoundStreamResponse\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x11\n\tnew_chunk\x18\x02 \x01(\x08\"V\n\x16SpeakerAndLineResponse\x12\x0c\n\x04text\x18\x01 \x03(\t\x12\x13\n\x0bspeakerName\x18\x02 \x03(\t\x12\x19\n\x11\x64\x65tected_language\x18\x03 \x01(\t\"h\n\tTextAndId\x12\x0c\n\x04text\x18\x01 \x01(\t\x12\x18\n\x10transcription_id\x18\x02 \x01(\x05\x12\x15\n\rtext_language\x18\x03 \x01(\t\x12\x1c\n\x14translation_language\x18\x04 \x01(\t2\xa6\x03\n\x0cSoundService\x12.\n\x0eTestConnection\x12\x0c.TextMessage\x1a\x0c.TextMessage\"\x00\x12\x39\n\x0eTranscribeFile\x12\x15.TranscriptionRequest\x1a\x0e.SoundResponse\"\x00\x12\x45\n\x0eTranscribeLive\x12\x19.TranscirptionLiveRequest\x1a\x14.SoundStreamResponse(\x01\x30\x01\x12@\n\x11TranscribeLiveWeb\x12\x15.TranscriptionRequest\x1a\x14.SoundStreamResponse\x12\x36\n\rTranslateFile\x12\x13.TranslationRequest\x1a\x0e.SoundResponse0\x01\x12?\n\rDiarizateFile\x12\x15.TranscriptionRequest\x1a\x17.SpeakerAndLineResponse\x12)\n\rTranslateText\x12\n.TextAndId\x1a\x0c.TextMessageBH\n\x15io.grpc.soundtransferB\x12SoundTransferProtoP\x01Z\x19inzynierka/sound_transferb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'sound_transfer_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\025io.grpc.soundtransferB\022SoundTransferProtoP\001Z\031inzynierka/sound_transfer'
  _globals['_TEXTMESSAGE']._serialized_start=24
  _globals['_TEXTMESSAGE']._serialized_end=51
  _globals['_TRANSCRIPTIONREQUEST']._serialized_start=53
  _globals['_TRANSCRIPTIONREQUEST']._serialized_end=120
  _globals['_TRANSLATIONREQUEST']._serialized_start=122
  _globals['_TRANSLATIONREQUEST']._serialized_end=217
  _globals['_TRANSCIRPTIONLIVEREQUEST']._serialized_start=219
  _globals['_TRANSCIRPTIONLIVEREQUEST']._serialized_end=265
  _globals['_SOUNDRESPONSE']._serialized_start=267
  _globals['_SOUNDRESPONSE']._serialized_end=335
  _globals['_SOUNDSTREAMRESPONSE']._serialized_start=337
  _globals['_SOUNDSTREAMRESPONSE']._serialized_end=391
  _globals['_SPEAKERANDLINERESPONSE']._serialized_start=393
  _globals['_SPEAKERANDLINERESPONSE']._serialized_end=479
  _globals['_TEXTANDID']._serialized_start=481
  _globals['_TEXTANDID']._serialized_end=585
  _globals['_SOUNDSERVICE']._serialized_start=588
  _globals['_SOUNDSERVICE']._serialized_end=1010
# @@protoc_insertion_point(module_scope)
