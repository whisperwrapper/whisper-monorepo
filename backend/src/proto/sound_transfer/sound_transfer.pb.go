// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.26.0
// source: backend/src/proto/sound_transfer/sound_transfer.proto

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

type SoundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SoundData []byte   `protobuf:"bytes,1,opt,name=sound_data,json=soundData,proto3" json:"sound_data,omitempty"`
	Flags     []string `protobuf:"bytes,2,rep,name=flags,proto3" json:"flags,omitempty"`
}

func (x *SoundRequest) Reset() {
	*x = SoundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoundRequest) ProtoMessage() {}

func (x *SoundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SoundRequest.ProtoReflect.Descriptor instead.
func (*SoundRequest) Descriptor() ([]byte, []int) {
	return file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{0}
}

func (x *SoundRequest) GetSoundData() []byte {
	if x != nil {
		return x.SoundData
	}
	return nil
}

func (x *SoundRequest) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

type SoundResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,3,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *SoundResponse) Reset() {
	*x = SoundResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoundResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoundResponse) ProtoMessage() {}

func (x *SoundResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[1]
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
	return file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{1}
}

func (x *SoundResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type TextMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
}

func (x *TextMessage) Reset() {
	*x = TextMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TextMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TextMessage) ProtoMessage() {}

func (x *TextMessage) ProtoReflect() protoreflect.Message {
	mi := &file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[2]
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
	return file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{2}
}

func (x *TextMessage) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type SoundStreamResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text  string   `protobuf:"bytes,5,opt,name=text,proto3" json:"text,omitempty"`
	Flags []string `protobuf:"bytes,6,rep,name=flags,proto3" json:"flags,omitempty"`
}

func (x *SoundStreamResponse) Reset() {
	*x = SoundStreamResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SoundStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SoundStreamResponse) ProtoMessage() {}

func (x *SoundStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[3]
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
	return file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescGZIP(), []int{3}
}

func (x *SoundStreamResponse) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SoundStreamResponse) GetFlags() []string {
	if x != nil {
		return x.Flags
	}
	return nil
}

var File_backend_src_proto_sound_transfer_sound_transfer_proto protoreflect.FileDescriptor

var file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDesc = []byte{
	0x0a, 0x35, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x2f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x6f, 0x75, 0x6e, 0x64,
	0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x73, 0x6f, 0x75,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x22, 0x23, 0x0a, 0x0d,
	0x53, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78,
	0x74, 0x22, 0x21, 0x0a, 0x0b, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x65, 0x78, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x65, 0x78, 0x74, 0x22, 0x3f, 0x0a, 0x13, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x53, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74,
	0x65, 0x78, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x78, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x66, 0x6c, 0x61, 0x67, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x66, 0x6c, 0x61, 0x67, 0x73, 0x32, 0xac, 0x01, 0x0a, 0x0c, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2e, 0x0a, 0x0e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x6f,
	0x6e, 0x6e, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x0c, 0x2e, 0x54, 0x65, 0x78, 0x74, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x22, 0x00, 0x12, 0x30, 0x0a, 0x0d, 0x53, 0x65, 0x6e, 0x64, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3a, 0x0a, 0x0f, 0x53, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x53, 0x6f, 0x75, 0x6e, 0x64, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x0d, 0x2e, 0x53, 0x6f,
	0x75, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x53, 0x6f, 0x75,
	0x6e, 0x64, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x48, 0x0a, 0x15, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x73, 0x6f, 0x75, 0x6e, 0x64, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x42, 0x12, 0x53,
	0x6f, 0x75, 0x6e, 0x64, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x19, 0x69, 0x6e, 0x7a, 0x79, 0x6e, 0x69, 0x65, 0x72, 0x6b, 0x61, 0x2f,
	0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescOnce sync.Once
	file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescData = file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDesc
)

func file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescGZIP() []byte {
	file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescOnce.Do(func() {
		file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescData = protoimpl.X.CompressGZIP(file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescData)
	})
	return file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDescData
}

var file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_backend_src_proto_sound_transfer_sound_transfer_proto_goTypes = []interface{}{
	(*SoundRequest)(nil),        // 0: SoundRequest
	(*SoundResponse)(nil),       // 1: SoundResponse
	(*TextMessage)(nil),         // 2: TextMessage
	(*SoundStreamResponse)(nil), // 3: SoundStreamResponse
}
var file_backend_src_proto_sound_transfer_sound_transfer_proto_depIdxs = []int32{
	2, // 0: SoundService.TestConnection:input_type -> TextMessage
	0, // 1: SoundService.SendSoundFile:input_type -> SoundRequest
	0, // 2: SoundService.StreamSoundFile:input_type -> SoundRequest
	2, // 3: SoundService.TestConnection:output_type -> TextMessage
	1, // 4: SoundService.SendSoundFile:output_type -> SoundResponse
	3, // 5: SoundService.StreamSoundFile:output_type -> SoundStreamResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_backend_src_proto_sound_transfer_sound_transfer_proto_init() }
func file_backend_src_proto_sound_transfer_sound_transfer_proto_init() {
	if File_backend_src_proto_sound_transfer_sound_transfer_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SoundRequest); i {
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
		file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_backend_src_proto_sound_transfer_sound_transfer_proto_goTypes,
		DependencyIndexes: file_backend_src_proto_sound_transfer_sound_transfer_proto_depIdxs,
		MessageInfos:      file_backend_src_proto_sound_transfer_sound_transfer_proto_msgTypes,
	}.Build()
	File_backend_src_proto_sound_transfer_sound_transfer_proto = out.File
	file_backend_src_proto_sound_transfer_sound_transfer_proto_rawDesc = nil
	file_backend_src_proto_sound_transfer_sound_transfer_proto_goTypes = nil
	file_backend_src_proto_sound_transfer_sound_transfer_proto_depIdxs = nil
}
