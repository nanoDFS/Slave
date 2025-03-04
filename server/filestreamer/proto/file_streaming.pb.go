// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v5.29.3
// source: proto/file_streaming.proto

package filestreamer

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

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data []byte `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_streaming_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_streaming_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_proto_file_streaming_proto_rawDescGZIP(), []int{0}
}

func (x *Payload) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type ReadReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ReadReq) Reset() {
	*x = ReadReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_streaming_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadReq) ProtoMessage() {}

func (x *ReadReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_streaming_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadReq.ProtoReflect.Descriptor instead.
func (*ReadReq) Descriptor() ([]byte, []int) {
	return file_proto_file_streaming_proto_rawDescGZIP(), []int{1}
}

type WriteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *WriteRes) Reset() {
	*x = WriteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_streaming_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRes) ProtoMessage() {}

func (x *WriteRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_streaming_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRes.ProtoReflect.Descriptor instead.
func (*WriteRes) Descriptor() ([]byte, []int) {
	return file_proto_file_streaming_proto_rawDescGZIP(), []int{2}
}

func (x *WriteRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type DeleteReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteReq) Reset() {
	*x = DeleteReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_streaming_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteReq) ProtoMessage() {}

func (x *DeleteReq) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_streaming_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteReq.ProtoReflect.Descriptor instead.
func (*DeleteReq) Descriptor() ([]byte, []int) {
	return file_proto_file_streaming_proto_rawDescGZIP(), []int{3}
}

type DeleteRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *DeleteRes) Reset() {
	*x = DeleteRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_file_streaming_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRes) ProtoMessage() {}

func (x *DeleteRes) ProtoReflect() protoreflect.Message {
	mi := &file_proto_file_streaming_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRes.ProtoReflect.Descriptor instead.
func (*DeleteRes) Descriptor() ([]byte, []int) {
	return file_proto_file_streaming_proto_rawDescGZIP(), []int{4}
}

func (x *DeleteRes) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

var File_proto_file_streaming_proto protoreflect.FileDescriptor

var file_proto_file_streaming_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x73, 0x74, 0x72,
	0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x22, 0x1d, 0x0a, 0x07, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x61,
	0x64, 0x52, 0x65, 0x71, 0x22, 0x22, 0x0a, 0x08, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x0b, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x22, 0x23, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52,
	0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0xc4, 0x01, 0x0a, 0x14, 0x46,
	0x69, 0x6c, 0x65, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x71, 0x1a, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65,
	0x72, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x05, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x12, 0x15, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61,
	0x6d, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x1a, 0x16, 0x2e, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x28, 0x01, 0x12, 0x3a, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12,
	0x17, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x17, 0x2e, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x52, 0x65,
	0x73, 0x42, 0x2a, 0x5a, 0x28, 0x2e, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x66, 0x69,
	0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x3b, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x65, 0x72, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_file_streaming_proto_rawDescOnce sync.Once
	file_proto_file_streaming_proto_rawDescData = file_proto_file_streaming_proto_rawDesc
)

func file_proto_file_streaming_proto_rawDescGZIP() []byte {
	file_proto_file_streaming_proto_rawDescOnce.Do(func() {
		file_proto_file_streaming_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_file_streaming_proto_rawDescData)
	})
	return file_proto_file_streaming_proto_rawDescData
}

var file_proto_file_streaming_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_file_streaming_proto_goTypes = []interface{}{
	(*Payload)(nil),   // 0: filestreamer.Payload
	(*ReadReq)(nil),   // 1: filestreamer.ReadReq
	(*WriteRes)(nil),  // 2: filestreamer.WriteRes
	(*DeleteReq)(nil), // 3: filestreamer.DeleteReq
	(*DeleteRes)(nil), // 4: filestreamer.DeleteRes
}
var file_proto_file_streaming_proto_depIdxs = []int32{
	1, // 0: filestreamer.FileStreamingService.Read:input_type -> filestreamer.ReadReq
	0, // 1: filestreamer.FileStreamingService.Write:input_type -> filestreamer.Payload
	3, // 2: filestreamer.FileStreamingService.Delete:input_type -> filestreamer.DeleteReq
	0, // 3: filestreamer.FileStreamingService.Read:output_type -> filestreamer.Payload
	2, // 4: filestreamer.FileStreamingService.Write:output_type -> filestreamer.WriteRes
	4, // 5: filestreamer.FileStreamingService.Delete:output_type -> filestreamer.DeleteRes
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_file_streaming_proto_init() }
func file_proto_file_streaming_proto_init() {
	if File_proto_file_streaming_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_file_streaming_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_proto_file_streaming_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadReq); i {
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
		file_proto_file_streaming_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRes); i {
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
		file_proto_file_streaming_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteReq); i {
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
		file_proto_file_streaming_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteRes); i {
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
			RawDescriptor: file_proto_file_streaming_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_file_streaming_proto_goTypes,
		DependencyIndexes: file_proto_file_streaming_proto_depIdxs,
		MessageInfos:      file_proto_file_streaming_proto_msgTypes,
	}.Build()
	File_proto_file_streaming_proto = out.File
	file_proto_file_streaming_proto_rawDesc = nil
	file_proto_file_streaming_proto_goTypes = nil
	file_proto_file_streaming_proto_depIdxs = nil
}
