// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/shortly.proto

package shortly

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type NewRequest struct {
	Url                  string   `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewRequest) Reset()         { *m = NewRequest{} }
func (m *NewRequest) String() string { return proto.CompactTextString(m) }
func (*NewRequest) ProtoMessage()    {}
func (*NewRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95cb5aec849ccd3, []int{0}
}

func (m *NewRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewRequest.Unmarshal(m, b)
}
func (m *NewRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewRequest.Marshal(b, m, deterministic)
}
func (m *NewRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewRequest.Merge(m, src)
}
func (m *NewRequest) XXX_Size() int {
	return xxx_messageInfo_NewRequest.Size(m)
}
func (m *NewRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_NewRequest.DiscardUnknown(m)
}

var xxx_messageInfo_NewRequest proto.InternalMessageInfo

func (m *NewRequest) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type IDRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IDRequest) Reset()         { *m = IDRequest{} }
func (m *IDRequest) String() string { return proto.CompactTextString(m) }
func (*IDRequest) ProtoMessage()    {}
func (*IDRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95cb5aec849ccd3, []int{1}
}

func (m *IDRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IDRequest.Unmarshal(m, b)
}
func (m *IDRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IDRequest.Marshal(b, m, deterministic)
}
func (m *IDRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDRequest.Merge(m, src)
}
func (m *IDRequest) XXX_Size() int {
	return xxx_messageInfo_IDRequest.Size(m)
}
func (m *IDRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_IDRequest.DiscardUnknown(m)
}

var xxx_messageInfo_IDRequest proto.InternalMessageInfo

func (m *IDRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type Link struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	Count                int64    `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Link) Reset()         { *m = Link{} }
func (m *Link) String() string { return proto.CompactTextString(m) }
func (*Link) ProtoMessage()    {}
func (*Link) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95cb5aec849ccd3, []int{2}
}

func (m *Link) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Link.Unmarshal(m, b)
}
func (m *Link) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Link.Marshal(b, m, deterministic)
}
func (m *Link) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Link.Merge(m, src)
}
func (m *Link) XXX_Size() int {
	return xxx_messageInfo_Link.Size(m)
}
func (m *Link) XXX_DiscardUnknown() {
	xxx_messageInfo_Link.DiscardUnknown(m)
}

var xxx_messageInfo_Link proto.InternalMessageInfo

func (m *Link) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Link) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *Link) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type Nothing struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Nothing) Reset()         { *m = Nothing{} }
func (m *Nothing) String() string { return proto.CompactTextString(m) }
func (*Nothing) ProtoMessage()    {}
func (*Nothing) Descriptor() ([]byte, []int) {
	return fileDescriptor_f95cb5aec849ccd3, []int{3}
}

func (m *Nothing) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Nothing.Unmarshal(m, b)
}
func (m *Nothing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Nothing.Marshal(b, m, deterministic)
}
func (m *Nothing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Nothing.Merge(m, src)
}
func (m *Nothing) XXX_Size() int {
	return xxx_messageInfo_Nothing.Size(m)
}
func (m *Nothing) XXX_DiscardUnknown() {
	xxx_messageInfo_Nothing.DiscardUnknown(m)
}

var xxx_messageInfo_Nothing proto.InternalMessageInfo

func init() {
	proto.RegisterType((*NewRequest)(nil), "NewRequest")
	proto.RegisterType((*IDRequest)(nil), "IDRequest")
	proto.RegisterType((*Link)(nil), "Link")
	proto.RegisterType((*Nothing)(nil), "Nothing")
}

func init() { proto.RegisterFile("proto/shortly.proto", fileDescriptor_f95cb5aec849ccd3) }

var fileDescriptor_f95cb5aec849ccd3 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0xce, 0xc8, 0x2f, 0x2a, 0xc9, 0xa9, 0xd4, 0x03, 0xf3, 0x94, 0xe4, 0xb8, 0xb8,
	0xfc, 0x52, 0xcb, 0x83, 0x52, 0x0b, 0x4b, 0x53, 0x8b, 0x4b, 0x84, 0x04, 0xb8, 0x98, 0x4b, 0x8b,
	0x72, 0x24, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83, 0x40, 0x4c, 0x25, 0x69, 0x2e, 0x4e, 0x4f, 0x17,
	0x98, 0x34, 0x1f, 0x17, 0x53, 0x66, 0x0a, 0x54, 0x96, 0x29, 0x33, 0x45, 0xc9, 0x8e, 0x8b, 0xc5,
	0x27, 0x33, 0x2f, 0x1b, 0x5d, 0x1c, 0x66, 0x0c, 0x13, 0xdc, 0x18, 0x21, 0x11, 0x2e, 0xd6, 0xe4,
	0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x47, 0x89, 0x93, 0x8b,
	0xdd, 0x2f, 0xbf, 0x24, 0x23, 0x33, 0x2f, 0xdd, 0x28, 0x83, 0x8b, 0x3d, 0x18, 0xe2, 0x30, 0x21,
	0x69, 0x2e, 0x66, 0xbf, 0xd4, 0x72, 0x21, 0x6e, 0x3d, 0x84, 0xc3, 0xa4, 0x58, 0xf5, 0x40, 0x16,
	0x29, 0x31, 0x08, 0x49, 0x71, 0x31, 0xbb, 0xa7, 0x96, 0x08, 0x71, 0xe9, 0xc1, 0x5d, 0x85, 0x90,
	0x53, 0xe1, 0xe2, 0x72, 0x06, 0x99, 0x1b, 0x96, 0x59, 0x9c, 0x89, 0xaa, 0x84, 0x43, 0x0f, 0x6a,
	0x8f, 0x12, 0x43, 0x12, 0x1b, 0xd8, 0xe3, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xcf, 0x68,
	0x71, 0x0a, 0x0f, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ShortlyClient is the client API for Shortly service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ShortlyClient interface {
	New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*Link, error)
	Get(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Link, error)
	CountVisit(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Nothing, error)
}

type shortlyClient struct {
	cc *grpc.ClientConn
}

func NewShortlyClient(cc *grpc.ClientConn) ShortlyClient {
	return &shortlyClient{cc}
}

func (c *shortlyClient) New(ctx context.Context, in *NewRequest, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := c.cc.Invoke(ctx, "/Shortly/New", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortlyClient) Get(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Link, error) {
	out := new(Link)
	err := c.cc.Invoke(ctx, "/Shortly/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shortlyClient) CountVisit(ctx context.Context, in *IDRequest, opts ...grpc.CallOption) (*Nothing, error) {
	out := new(Nothing)
	err := c.cc.Invoke(ctx, "/Shortly/CountVisit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ShortlyServer is the server API for Shortly service.
type ShortlyServer interface {
	New(context.Context, *NewRequest) (*Link, error)
	Get(context.Context, *IDRequest) (*Link, error)
	CountVisit(context.Context, *IDRequest) (*Nothing, error)
}

// UnimplementedShortlyServer can be embedded to have forward compatible implementations.
type UnimplementedShortlyServer struct {
}

func (*UnimplementedShortlyServer) New(ctx context.Context, req *NewRequest) (*Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method New not implemented")
}
func (*UnimplementedShortlyServer) Get(ctx context.Context, req *IDRequest) (*Link, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (*UnimplementedShortlyServer) CountVisit(ctx context.Context, req *IDRequest) (*Nothing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountVisit not implemented")
}

func RegisterShortlyServer(s *grpc.Server, srv ShortlyServer) {
	s.RegisterService(&_Shortly_serviceDesc, srv)
}

func _Shortly_New_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortlyServer).New(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortly/New",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortlyServer).New(ctx, req.(*NewRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortly_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortlyServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortly/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortlyServer).Get(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Shortly_CountVisit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IDRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ShortlyServer).CountVisit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/Shortly/CountVisit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ShortlyServer).CountVisit(ctx, req.(*IDRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Shortly_serviceDesc = grpc.ServiceDesc{
	ServiceName: "Shortly",
	HandlerType: (*ShortlyServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "New",
			Handler:    _Shortly_New_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Shortly_Get_Handler,
		},
		{
			MethodName: "CountVisit",
			Handler:    _Shortly_CountVisit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/shortly.proto",
}