// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: api/maths/maths.proto

package maths

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

type SquaresRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SquaresRequest) Reset() {
	*x = SquaresRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_maths_maths_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquaresRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquaresRequest) ProtoMessage() {}

func (x *SquaresRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_maths_maths_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquaresRequest.ProtoReflect.Descriptor instead.
func (*SquaresRequest) Descriptor() ([]byte, []int) {
	return file_api_maths_maths_proto_rawDescGZIP(), []int{0}
}

func (x *SquaresRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type SquaresResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Square int32 `protobuf:"varint,1,opt,name=square,proto3" json:"square,omitempty"`
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *SquaresResponse) Reset() {
	*x = SquaresResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_maths_maths_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SquaresResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SquaresResponse) ProtoMessage() {}

func (x *SquaresResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_maths_maths_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SquaresResponse.ProtoReflect.Descriptor instead.
func (*SquaresResponse) Descriptor() ([]byte, []int) {
	return file_api_maths_maths_proto_rawDescGZIP(), []int{1}
}

func (x *SquaresResponse) GetSquare() int32 {
	if x != nil {
		return x.Square
	}
	return 0
}

func (x *SquaresResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CubesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number int32 `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CubesRequest) Reset() {
	*x = CubesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_maths_maths_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CubesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CubesRequest) ProtoMessage() {}

func (x *CubesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_maths_maths_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CubesRequest.ProtoReflect.Descriptor instead.
func (*CubesRequest) Descriptor() ([]byte, []int) {
	return file_api_maths_maths_proto_rawDescGZIP(), []int{2}
}

func (x *CubesRequest) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

type CubesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cube   int32 `protobuf:"varint,1,opt,name=cube,proto3" json:"cube,omitempty"`
	Number int32 `protobuf:"varint,2,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *CubesResponse) Reset() {
	*x = CubesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_maths_maths_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CubesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CubesResponse) ProtoMessage() {}

func (x *CubesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_maths_maths_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CubesResponse.ProtoReflect.Descriptor instead.
func (*CubesResponse) Descriptor() ([]byte, []int) {
	return file_api_maths_maths_proto_rawDescGZIP(), []int{3}
}

func (x *CubesResponse) GetCube() int32 {
	if x != nil {
		return x.Cube
	}
	return 0
}

func (x *CubesResponse) GetNumber() int32 {
	if x != nil {
		return x.Number
	}
	return 0
}

var File_api_maths_maths_proto protoreflect.FileDescriptor

var file_api_maths_maths_proto_rawDesc = []byte{
	0x0a, 0x15, 0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x73, 0x2f, 0x6d, 0x61, 0x74, 0x68,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d, 0x61, 0x74, 0x68, 0x73, 0x22, 0x28,
	0x0a, 0x0e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x41, 0x0a, 0x0f, 0x53, 0x71, 0x75, 0x61,
	0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x71, 0x75, 0x61, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x73, 0x71, 0x75,
	0x61, 0x72, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x26, 0x0a, 0x0c, 0x43,
	0x75, 0x62, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x22, 0x3b, 0x0a, 0x0d, 0x43, 0x75, 0x62, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x75, 0x62, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x04, 0x63, 0x75, 0x62, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x32, 0x81, 0x01, 0x0a, 0x05, 0x4d, 0x61, 0x74, 0x68, 0x73, 0x12, 0x3e, 0x0a, 0x07, 0x53, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x73, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x73, 0x2e, 0x53, 0x71,
	0x75, 0x61, 0x72, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6d,
	0x61, 0x74, 0x68, 0x73, 0x2e, 0x53, 0x71, 0x75, 0x61, 0x72, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x38, 0x0a, 0x05, 0x43, 0x75,
	0x62, 0x65, 0x73, 0x12, 0x13, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x73, 0x2e, 0x43, 0x75, 0x62, 0x65,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x6d, 0x61, 0x74, 0x68, 0x73,
	0x2e, 0x43, 0x75, 0x62, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x28, 0x01, 0x30, 0x01, 0x42, 0x33, 0x5a, 0x31, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x7a, 0x61, 0x63, 0x68, 0x6d, 0x61, 0x6e, 0x64, 0x65, 0x76, 0x69, 0x6c, 0x6c,
	0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x63, 0x68, 0x61, 0x6e, 0x2d, 0x66, 0x75, 0x6e, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_api_maths_maths_proto_rawDescOnce sync.Once
	file_api_maths_maths_proto_rawDescData = file_api_maths_maths_proto_rawDesc
)

func file_api_maths_maths_proto_rawDescGZIP() []byte {
	file_api_maths_maths_proto_rawDescOnce.Do(func() {
		file_api_maths_maths_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_maths_maths_proto_rawDescData)
	})
	return file_api_maths_maths_proto_rawDescData
}

var file_api_maths_maths_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_maths_maths_proto_goTypes = []interface{}{
	(*SquaresRequest)(nil),  // 0: maths.SquaresRequest
	(*SquaresResponse)(nil), // 1: maths.SquaresResponse
	(*CubesRequest)(nil),    // 2: maths.CubesRequest
	(*CubesResponse)(nil),   // 3: maths.CubesResponse
}
var file_api_maths_maths_proto_depIdxs = []int32{
	0, // 0: maths.Maths.Squares:input_type -> maths.SquaresRequest
	2, // 1: maths.Maths.Cubes:input_type -> maths.CubesRequest
	1, // 2: maths.Maths.Squares:output_type -> maths.SquaresResponse
	3, // 3: maths.Maths.Cubes:output_type -> maths.CubesResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_maths_maths_proto_init() }
func file_api_maths_maths_proto_init() {
	if File_api_maths_maths_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_maths_maths_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquaresRequest); i {
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
		file_api_maths_maths_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SquaresResponse); i {
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
		file_api_maths_maths_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CubesRequest); i {
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
		file_api_maths_maths_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CubesResponse); i {
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
			RawDescriptor: file_api_maths_maths_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_maths_maths_proto_goTypes,
		DependencyIndexes: file_api_maths_maths_proto_depIdxs,
		MessageInfos:      file_api_maths_maths_proto_msgTypes,
	}.Build()
	File_api_maths_maths_proto = out.File
	file_api_maths_maths_proto_rawDesc = nil
	file_api_maths_maths_proto_goTypes = nil
	file_api_maths_maths_proto_depIdxs = nil
}
