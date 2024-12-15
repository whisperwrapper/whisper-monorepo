// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: backend/ApiServer/proto/sound_transfer/sound_transfer.proto

package sound_transfer

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TextMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextMessage) Reset() {
	*x = TextMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessage) ProtoMessage() {}

func (x *TextMessage) ProtoReflect() protoreflect.Message {
	mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TextMessage.ProtoReflect.Descriptor instead.
func (*TextMessage) Descriptor() ([]byte, []int) {
	return file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *TextMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TranscriptionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoundData      []byte `protobuf:"bytes,1,opt,name=sound_data,json=soundData,proto3" json:"sound_data,omitempty"`
	SourceLanguage string `protobuf:"bytes,2,opt,name=source_language,json=sourceLanguage,proto3" json:"source_language,omitempty"`
}

func (x *TranscriptionRequest) Reset() {
	*x = TranscriptionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscriptionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscriptionRequest) ProtoMessage() {}

func (x *TranscriptionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscriptionRequest.ProtoReflect.Descriptor instead.
func (*TranscriptionRequest) Descriptor() ([]byte, []int) {
	return file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *TranscriptionRequest) GetSoundData() []byte {
	if x != nil {
		return x.SoundData
	}
	return nil
}

func (x *TranscriptionRequest) GetSourceLanguage() string {
	if x != nil {
		return x.SourceLanguage
	}
	return ""
}

type TranslationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoundData           []byte `protobuf:"bytes,1,opt,name=sound_data,json=soundData,proto3" json:"sound_data,omitempty"`
	SourceLanguage      string `protobuf:"bytes,2,opt,name=source_language,json=sourceLanguage,proto3" json:"source_language,omitempty"`
	TranslationLanguage string `protobuf:"bytes,3,opt,name=translation_language,json=translationLanguage,proto3" json:"translation_language,omitempty"`
}

func (x *TranslationRequest) Reset() {
	*x = TranslationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranslationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranslationRequest) ProtoMessage() {}

func (x *TranslationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranslationRequest.ProtoReflect.Descriptor instead.
func (*TranslationRequest) Descriptor() ([]byte, []int) {
	return file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *TranslationRequest) GetSoundData() []byte {
	if x != nil {
		return x.SoundData
	}
	return nil
}

func (x *TranslationRequest) GetSourceLanguage() string {
	if x != nil {
		return x.SourceLanguage
	}
	return ""
}

func (x *TranslationRequest) GetTranslationLanguage() string {
	if x != nil {
		return x.TranslationLanguage
	}
	return ""
}

type TranscirptionLiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoundData []byte `protobuf:"bytes,1,opt,name=sound_data,json=soundData,proto3" json:"sound_data,omitempty"`
}

func (x *TranscirptionLiveRequest) Reset() {
	*x = TranscirptionLiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TranscirptionLiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TranscirptionLiveRequest) ProtoMessage() {}

func (x *TranscirptionLiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TranscirptionLiveRequest.ProtoReflect.Descriptor instead.
func (*TranscirptionLiveRequest) Descriptor() ([]byte, []int) {
	return file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *TranscirptionLiveRequest) GetSoundData() []byte {
	if x != nil {
		return x.SoundData
	}
	return nil
}

type SoundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text             string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	DetectedLanguage string `protobuf:"bytes,2,opt,name=detected_language,json=detectedLanguage,proto3" json:"detected_language,omitempty"`
}

func (x *SoundResponse) Reset() {
	*x = SoundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoundResponse) ProtoMessage() {}

func (x *SoundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoundResponse.ProtoReflect.Descriptor instead.
func (*SoundResponse) Descriptor() ([]byte, []int) {
	return file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{4}
}

func (x *SoundResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SoundResponse) GetDetectedLanguage() string {
	if x != nil {
		return x.DetectedLanguage
	}
	return ""
}

type SoundStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text     string `protobuf:"bytes,1,opt,name=text,proto3" json:"text,omitempty"`
	NewChunk bool   `protobuf:"varint,2,opt,name=new_chunk,json=newChunk,proto3" json:"new_chunk,omitempty"`
}

func (x *SoundStreamResponse) Reset() {
	*x = SoundStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoundStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoundStreamResponse) ProtoMessage() {}

func (x *SoundStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoundStreamResponse.ProtoReflect.Descriptor instead.
func (*SoundStreamResponse) Descriptor() ([]byte, []int) {
	return file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{5}
}

func (x *SoundStreamResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SoundStreamResponse) GetNewChunk() bool {
	if x != nil {
		return x.NewChunk
	}
	return false
}

type SpeakerAndLineResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text             []string `protobuf:"bytes,1,rep,name=text,proto3" json:"text,omitempty"`
	SpeakerName      []string `protobuf:"bytes,2,rep,name=speakerName,proto3" json:"speakerName,omitempty"`
	DetectedLanguage string   `protobuf:"bytes,3,opt,name=detected_language,json=detectedLanguage,proto3" json:"detected_language,omitempty"`
}

func (x *SpeakerAndLineResponse) Reset() {
	*x = SpeakerAndLineResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpeakerAndLineResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpeakerAndLineResponse) ProtoMessage() {}

func (x *SpeakerAndLineResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpeakerAndLineResponse.ProtoReflect.Descriptor instead.
func (*SpeakerAndLineResponse) Descriptor() ([]byte, []int) {
	return file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{6}
}

func (x *SpeakerAndLineResponse) GetText() []string {
	if x != nil {
		return x.Text
	}
	return nil
}

func (x *SpeakerAndLineResponse) GetSpeakerName() []string {
	if x != nil {
		return x.SpeakerName
	}
	return nil
}

func (x *SpeakerAndLineResponse) GetDetectedLanguage() string {
	if x != nil {
		return x.DetectedLanguage
	}
	return ""
}

var File_backend_ApiServer_proto_sound_transfer_sound_transfer_proto protoreflect.FileDescriptor

var file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDesc = []byte{
	0x0a, 0x3b, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x41, 0x70, 0x69, 0x53, 0x65, 0x72,
	0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a,
	0x0b, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74,
	0x22, 0x5e, 0x0a, 0x14, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e,
	0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x6f,
	0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x22, 0x8f, 0x01, 0x0a, 0x12, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x6f, 0x75,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0e, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x12,
	0x31, 0x0a, 0x14, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x74,
	0x72, 0x61, 0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61,
	0x67, 0x65, 0x22, 0x39, 0x0a, 0x18, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x69, 0x72, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x22, 0x50, 0x0a,
	0x0d, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65,
	0x78, 0x74, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x6c,
	0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64,
	0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x22,
	0x46, 0x0a, 0x13, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65,
	0x77, 0x5f, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x6e,
	0x65, 0x77, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x22, 0x7b, 0x0a, 0x16, 0x53, 0x70, 0x65, 0x61, 0x6b,
	0x65, 0x72, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x73, 0x70, 0x65, 0x61, 0x6b, 0x65, 0x72,
	0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x70, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x65, 0x64, 0x5f, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x65, 0x64, 0x4c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x32, 0xfb, 0x02, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x39, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72,
	0x69, 0x62, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x15, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63,
	0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e,
	0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x45, 0x0a, 0x0e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x69,
	0x76, 0x65, 0x12, 0x19, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x69, 0x72, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x4c, 0x69, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e,
	0x53, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x28, 0x01, 0x30, 0x01, 0x12, 0x40, 0x0a, 0x11, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x63, 0x72, 0x69, 0x62, 0x65, 0x4c, 0x69, 0x76, 0x65, 0x57, 0x65, 0x62, 0x12, 0x15, 0x2e, 0x54,
	0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0d, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x13, 0x2e, 0x54, 0x72, 0x61,
	0x6e, 0x73, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0e, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x12, 0x3f, 0x0a, 0x0d, 0x44, 0x69, 0x61, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x65, 0x46, 0x69,
	0x6c, 0x65, 0x12, 0x15, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x53, 0x70, 0x65, 0x61,
	0x6b, 0x65, 0x72, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x48, 0x0a, 0x15, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x73, 0x6f,
	0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x12, 0x53, 0x6f, 0x75,
	0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x19, 0x69, 0x6e, 0x7a, 0x79, 0x6e, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x2f, 0x73, 0x6f,
	0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescOnce sync.Once
	file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescData = file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDesc
)

func file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescGZIP() []byte {
	file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescOnce.Do(func() {
		file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescData)
	})
	return file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDescData
}

var file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_goTypes = []interface{}{
	(*TextMessage)(nil),              // 0: TextMessage
	(*TranscriptionRequest)(nil),     // 1: TranscriptionRequest
	(*TranslationRequest)(nil),       // 2: TranslationRequest
	(*TranscirptionLiveRequest)(nil), // 3: TranscirptionLiveRequest
	(*SoundResponse)(nil),            // 4: SoundResponse
	(*SoundStreamResponse)(nil),      // 5: SoundStreamResponse
	(*SpeakerAndLineResponse)(nil),   // 6: SpeakerAndLineResponse
}
var file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_depIdxs = []int32{
	0, // 0: SoundService.TestConnection:input_type -> TextMessage
	1, // 1: SoundService.TranscribeFile:input_type -> TranscriptionRequest
	3, // 2: SoundService.TranscribeLive:input_type -> TranscirptionLiveRequest
	1, // 3: SoundService.TranscribeLiveWeb:input_type -> TranscriptionRequest
	2, // 4: SoundService.TranslateFile:input_type -> TranslationRequest
	1, // 5: SoundService.DiarizateFile:input_type -> TranscriptionRequest
	0, // 6: SoundService.TestConnection:output_type -> TextMessage
	4, // 7: SoundService.TranscribeFile:output_type -> SoundResponse
	5, // 8: SoundService.TranscribeLive:output_type -> SoundStreamResponse
	5, // 9: SoundService.TranscribeLiveWeb:output_type -> SoundStreamResponse
	4, // 10: SoundService.TranslateFile:output_type -> SoundResponse
	6, // 11: SoundService.DiarizateFile:output_type -> SpeakerAndLineResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_init() }
func file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_init() {
	if File_backend_ApiServer_proto_sound_transfer_sound_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TextMessage); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscriptionRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranslationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TranscirptionLiveRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoundResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoundStreamResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpeakerAndLineResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_goTypes,
		DependencyIndexes: file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_depIdxs,
		MessageInfos:      file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_msgTypes,
	}.Build()
	File_backend_ApiServer_proto_sound_transfer_sound_transfer_proto = out.File
	file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_rawDesc = nil
	file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_goTypes = nil
	file_backend_ApiServer_proto_sound_transfer_sound_transfer_proto_depIdxs = nil
}
