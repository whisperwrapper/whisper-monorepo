# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: authentication.proto
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
    'authentication.proto'
)
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\x14\x61uthentication.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\x8a\x01\n\x14TranscriptionHistory\x12\x15\n\rtranscription\x18\x01 \x01(\t\x12.\n\ncreated_at\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\n\n\x02id\x18\x03 \x01(\x05\x12\x10\n\x08language\x18\x04 \x01(\t\x12\r\n\x05title\x18\x05 \x01(\t\"\xc9\x01\n\x12TranslationHistory\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x15\n\rtranscription\x18\x02 \x01(\t\x12\x13\n\x0btranslation\x18\x03 \x01(\t\x12.\n\ncreated_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x1e\n\x16transcription_langauge\x18\x05 \x01(\t\x12\x1c\n\x14translation_langauge\x18\x06 \x01(\t\x12\r\n\x05title\x18\x07 \x01(\t\"\x9c\x01\n\x12\x44iarizationHistory\x12\x16\n\x0e\x64iarization_id\x18\x01 \x01(\x05\x12\x0f\n\x07speaker\x18\x03 \x03(\t\x12\x0c\n\x04line\x18\x04 \x03(\t\x12.\n\ncreated_at\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x10\n\x08language\x18\x06 \x01(\t\x12\r\n\x05title\x18\x07 \x01(\t\"1\n\x0bUserCredits\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"0\n\rLoginResponse\x12\x0b\n\x03JWT\x18\x01 \x01(\t\x12\x12\n\nsuccessful\x18\x02 \x01(\x08\"\x10\n\x02Id\x12\n\n\x02id\x18\x01 \x01(\x05\"/\n\x10NewTranscription\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0f\n\x07\x63ontent\x18\x02 \x01(\t\"~\n\x0eNewTranslation\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x15\n\rtranscription\x18\x02 \x01(\t\x12\x13\n\x0btranslation\x18\x03 \x01(\t\x12\x1a\n\x12\x65\x64it_transcription\x18\x04 \x01(\x08\x12\x18\n\x10\x65\x64it_translation\x18\x05 \x01(\x08\";\n\x0eNewDiarization\x12\n\n\x02id\x18\x01 \x01(\x05\x12\x0c\n\x04line\x18\x02 \x03(\t\x12\x0f\n\x07speaker\x18\x03 \x03(\t\"\xaf\x01\n\x10QueryParamethers\x12.\n\nstart_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\r\n\x05limit\x18\x03 \x01(\x05\x12\x10\n\x08language\x18\x04 \x01(\t\x12\x1c\n\x14translation_language\x18\x05 \x01(\t\"b\n\x08\x43ombined\x12(\n\x0b\x64iarization\x18\x01 \x01(\x0b\x32\x13.DiarizationHistory\x12,\n\rtranscription\x18\x02 \x01(\x0b\x32\x15.TranscriptionHistory2\xca\x05\n\rClientService\x12\'\n\x05Login\x12\x0c.UserCredits\x1a\x0e.LoginResponse\"\x00\x12\x32\n\x08Register\x12\x0c.UserCredits\x1a\x16.google.protobuf.Empty\"\x00\x12@\n\x10GetTranscription\x12\x11.QueryParamethers\x1a\x15.TranscriptionHistory\"\x00\x30\x01\x12@\n\x11\x45\x64itTranscription\x12\x11.NewTranscription\x1a\x16.google.protobuf.Empty\"\x00\x12\x34\n\x13\x44\x65leteTranscription\x12\x03.Id\x1a\x16.google.protobuf.Empty\"\x00\x12<\n\x0eGetTranslation\x12\x11.QueryParamethers\x1a\x13.TranslationHistory\"\x00\x30\x01\x12<\n\x0f\x45\x64itTranslation\x12\x0f.NewTranslation\x1a\x16.google.protobuf.Empty\"\x00\x12\x32\n\x11\x44\x65leteTranslation\x12\x03.Id\x1a\x16.google.protobuf.Empty\"\x00\x12<\n\x0eGetDiarization\x12\x11.QueryParamethers\x1a\x13.DiarizationHistory\"\x00\x30\x01\x12<\n\x0f\x45\x64itDiarization\x12\x0f.NewDiarization\x1a\x16.google.protobuf.Empty\"\x00\x12\x32\n\x11\x44\x65leteDiarization\x12\x03.Id\x1a\x16.google.protobuf.Empty\"\x00\x12\x42\n\x1eGetTranscriptionAndDiarization\x12\x11.QueryParamethers\x1a\t.Combined\"\x00\x30\x01\x42J\n\x16io.grpc.authenticationB\x13\x41uthenticationProtoP\x01Z\x19inzynierka/authenticationb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'authentication_pb2', _globals)
if not _descriptor._USE_C_DESCRIPTORS:
  _globals['DESCRIPTOR']._loaded_options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\026io.grpc.authenticationB\023AuthenticationProtoP\001Z\031inzynierka/authentication'
  _globals['_TRANSCRIPTIONHISTORY']._serialized_start=87
  _globals['_TRANSCRIPTIONHISTORY']._serialized_end=225
  _globals['_TRANSLATIONHISTORY']._serialized_start=228
  _globals['_TRANSLATIONHISTORY']._serialized_end=429
  _globals['_DIARIZATIONHISTORY']._serialized_start=432
  _globals['_DIARIZATIONHISTORY']._serialized_end=588
  _globals['_USERCREDITS']._serialized_start=590
  _globals['_USERCREDITS']._serialized_end=639
  _globals['_LOGINRESPONSE']._serialized_start=641
  _globals['_LOGINRESPONSE']._serialized_end=689
  _globals['_ID']._serialized_start=691
  _globals['_ID']._serialized_end=707
  _globals['_NEWTRANSCRIPTION']._serialized_start=709
  _globals['_NEWTRANSCRIPTION']._serialized_end=756
  _globals['_NEWTRANSLATION']._serialized_start=758
  _globals['_NEWTRANSLATION']._serialized_end=884
  _globals['_NEWDIARIZATION']._serialized_start=886
  _globals['_NEWDIARIZATION']._serialized_end=945
  _globals['_QUERYPARAMETHERS']._serialized_start=948
  _globals['_QUERYPARAMETHERS']._serialized_end=1123
  _globals['_COMBINED']._serialized_start=1125
  _globals['_COMBINED']._serialized_end=1223
  _globals['_CLIENTSERVICE']._serialized_start=1226
  _globals['_CLIENTSERVICE']._serialized_end=1940
# @@protoc_insertion_point(module_scope)
