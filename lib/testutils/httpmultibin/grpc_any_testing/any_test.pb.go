// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: any_test.proto

package grpc_any_testing

import (
	context "context"
	reflect "reflect"
	sync "sync"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SumRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SumRequest) Reset() {
	*x = SumRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_any_test_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequest) ProtoMessage() {}

func (x *SumRequest) ProtoReflect() protoreflect.Message {
	mi := &file_any_test_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequest.ProtoReflect.Descriptor instead.
func (*SumRequest) Descriptor() ([]byte, []int) {
	return file_any_test_proto_rawDescGZIP(), []int{0}
}

func (x *SumRequest) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type SumRequestData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	A int64 `protobuf:"varint,1,opt,name=a,proto3" json:"a,omitempty"`
	B int64 `protobuf:"varint,2,opt,name=b,proto3" json:"b,omitempty"`
}

func (x *SumRequestData) Reset() {
	*x = SumRequestData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_any_test_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumRequestData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumRequestData) ProtoMessage() {}

func (x *SumRequestData) ProtoReflect() protoreflect.Message {
	mi := &file_any_test_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumRequestData.ProtoReflect.Descriptor instead.
func (*SumRequestData) Descriptor() ([]byte, []int) {
	return file_any_test_proto_rawDescGZIP(), []int{1}
}

func (x *SumRequestData) GetA() int64 {
	if x != nil {
		return x.A
	}
	return 0
}

func (x *SumRequestData) GetB() int64 {
	if x != nil {
		return x.B
	}
	return 0
}

type SumReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data *anypb.Any `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SumReply) Reset() {
	*x = SumReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_any_test_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumReply) ProtoMessage() {}

func (x *SumReply) ProtoReflect() protoreflect.Message {
	mi := &file_any_test_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumReply.ProtoReflect.Descriptor instead.
func (*SumReply) Descriptor() ([]byte, []int) {
	return file_any_test_proto_rawDescGZIP(), []int{2}
}

func (x *SumReply) GetData() *anypb.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

type SumReplyData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	V   int64  `protobuf:"varint,1,opt,name=v,proto3" json:"v,omitempty"`
	Err string `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *SumReplyData) Reset() {
	*x = SumReplyData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_any_test_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SumReplyData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SumReplyData) ProtoMessage() {}

func (x *SumReplyData) ProtoReflect() protoreflect.Message {
	mi := &file_any_test_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SumReplyData.ProtoReflect.Descriptor instead.
func (*SumReplyData) Descriptor() ([]byte, []int) {
	return file_any_test_proto_rawDescGZIP(), []int{3}
}

func (x *SumReplyData) GetV() int64 {
	if x != nil {
		return x.V
	}
	return 0
}

func (x *SumReplyData) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_any_test_proto protoreflect.FileDescriptor

var file_any_test_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x6e, 0x79, 0x5f, 0x74, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74, 0x69,
	0x6e, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a,
	0x0a, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x28, 0x0a, 0x04, 0x64,
	0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2c, 0x0a, 0x0e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x61, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x01, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x62, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x01, 0x62, 0x22, 0x34, 0x0a, 0x08, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x2e, 0x0a, 0x0c, 0x53, 0x75, 0x6d,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0c, 0x0a, 0x01, 0x76, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x76, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x53, 0x0a, 0x0e, 0x41, 0x6e, 0x79,
	0x54, 0x65, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x03, 0x53,
	0x75, 0x6d, 0x12, 0x1c, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x61, 0x6e, 0x79, 0x2e, 0x74, 0x65, 0x73, 0x74,
	0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x6d, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x14,
	0x5a, 0x12, 0x2e, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x61, 0x6e, 0x79, 0x5f, 0x74, 0x65, 0x73,
	0x74, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_any_test_proto_rawDescOnce sync.Once
	file_any_test_proto_rawDescData = file_any_test_proto_rawDesc
)

func file_any_test_proto_rawDescGZIP() []byte {
	file_any_test_proto_rawDescOnce.Do(func() {
		file_any_test_proto_rawDescData = protoimpl.X.CompressGZIP(file_any_test_proto_rawDescData)
	})
	return file_any_test_proto_rawDescData
}

var file_any_test_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_any_test_proto_goTypes = []interface{}{
	(*SumRequest)(nil),     // 0: grpc.any.testing.SumRequest
	(*SumRequestData)(nil), // 1: grpc.any.testing.SumRequestData
	(*SumReply)(nil),       // 2: grpc.any.testing.SumReply
	(*SumReplyData)(nil),   // 3: grpc.any.testing.SumReplyData
	(*anypb.Any)(nil),      // 4: google.protobuf.Any
}
var file_any_test_proto_depIdxs = []int32{
	4, // 0: grpc.any.testing.SumRequest.data:type_name -> google.protobuf.Any
	4, // 1: grpc.any.testing.SumReply.data:type_name -> google.protobuf.Any
	0, // 2: grpc.any.testing.AnyTestService.Sum:input_type -> grpc.any.testing.SumRequest
	2, // 3: grpc.any.testing.AnyTestService.Sum:output_type -> grpc.any.testing.SumReply
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_any_test_proto_init() }
func file_any_test_proto_init() {
	if File_any_test_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_any_test_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequest); i {
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
		file_any_test_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumRequestData); i {
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
		file_any_test_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumReply); i {
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
		file_any_test_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SumReplyData); i {
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
			RawDescriptor: file_any_test_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_any_test_proto_goTypes,
		DependencyIndexes: file_any_test_proto_depIdxs,
		MessageInfos:      file_any_test_proto_msgTypes,
	}.Build()
	File_any_test_proto = out.File
	file_any_test_proto_rawDesc = nil
	file_any_test_proto_goTypes = nil
	file_any_test_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// AnyTestServiceClient is the client API for AnyTestService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AnyTestServiceClient interface {
	Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error)
}

type anyTestServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAnyTestServiceClient(cc grpc.ClientConnInterface) AnyTestServiceClient {
	return &anyTestServiceClient{cc}
}

func (c *anyTestServiceClient) Sum(ctx context.Context, in *SumRequest, opts ...grpc.CallOption) (*SumReply, error) {
	out := new(SumReply)
	err := c.cc.Invoke(ctx, "/grpc.any.testing.AnyTestService/Sum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AnyTestServiceServer is the server API for AnyTestService service.
type AnyTestServiceServer interface {
	Sum(context.Context, *SumRequest) (*SumReply, error)
}

// UnimplementedAnyTestServiceServer can be embedded to have forward compatible implementations.
type UnimplementedAnyTestServiceServer struct {
}

func (*UnimplementedAnyTestServiceServer) Sum(context.Context, *SumRequest) (*SumReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Sum not implemented")
}

func RegisterAnyTestServiceServer(s *grpc.Server, srv AnyTestServiceServer) {
	s.RegisterService(&_AnyTestService_serviceDesc, srv)
}

func _AnyTestService_Sum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AnyTestServiceServer).Sum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.any.testing.AnyTestService/Sum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AnyTestServiceServer).Sum(ctx, req.(*SumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _AnyTestService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.any.testing.AnyTestService",
	HandlerType: (*AnyTestServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Sum",
			Handler:    _AnyTestService_Sum_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "any_test.proto",
}
