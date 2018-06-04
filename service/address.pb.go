// Code generated by protoc-gen-go. DO NOT EDIT.
// source: address.proto

package service

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AddressGetOptions struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressGetOptions) Reset()         { *m = AddressGetOptions{} }
func (m *AddressGetOptions) String() string { return proto.CompactTextString(m) }
func (*AddressGetOptions) ProtoMessage()    {}
func (*AddressGetOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_20997c42f981569c, []int{0}
}
func (m *AddressGetOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressGetOptions.Unmarshal(m, b)
}
func (m *AddressGetOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressGetOptions.Marshal(b, m, deterministic)
}
func (dst *AddressGetOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressGetOptions.Merge(dst, src)
}
func (m *AddressGetOptions) XXX_Size() int {
	return xxx_messageInfo_AddressGetOptions.Size(m)
}
func (m *AddressGetOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressGetOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressGetOptions proto.InternalMessageInfo

func (m *AddressGetOptions) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type AddressGetByPathOptions struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressGetByPathOptions) Reset()         { *m = AddressGetByPathOptions{} }
func (m *AddressGetByPathOptions) String() string { return proto.CompactTextString(m) }
func (*AddressGetByPathOptions) ProtoMessage()    {}
func (*AddressGetByPathOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_20997c42f981569c, []int{1}
}
func (m *AddressGetByPathOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressGetByPathOptions.Unmarshal(m, b)
}
func (m *AddressGetByPathOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressGetByPathOptions.Marshal(b, m, deterministic)
}
func (dst *AddressGetByPathOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressGetByPathOptions.Merge(dst, src)
}
func (m *AddressGetByPathOptions) XXX_Size() int {
	return xxx_messageInfo_AddressGetByPathOptions.Size(m)
}
func (m *AddressGetByPathOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressGetByPathOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressGetByPathOptions proto.InternalMessageInfo

func (m *AddressGetByPathOptions) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

type AddressGetMultiOptions struct {
	Ids                  []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AddressGetMultiOptions) Reset()         { *m = AddressGetMultiOptions{} }
func (m *AddressGetMultiOptions) String() string { return proto.CompactTextString(m) }
func (*AddressGetMultiOptions) ProtoMessage()    {}
func (*AddressGetMultiOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_20997c42f981569c, []int{2}
}
func (m *AddressGetMultiOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AddressGetMultiOptions.Unmarshal(m, b)
}
func (m *AddressGetMultiOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AddressGetMultiOptions.Marshal(b, m, deterministic)
}
func (dst *AddressGetMultiOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AddressGetMultiOptions.Merge(dst, src)
}
func (m *AddressGetMultiOptions) XXX_Size() int {
	return xxx_messageInfo_AddressGetMultiOptions.Size(m)
}
func (m *AddressGetMultiOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_AddressGetMultiOptions.DiscardUnknown(m)
}

var xxx_messageInfo_AddressGetMultiOptions proto.InternalMessageInfo

func (m *AddressGetMultiOptions) GetIds() []string {
	if m != nil {
		return m.Ids
	}
	return nil
}

type Address struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Legacy               *Legacy  `protobuf:"bytes,3,opt,name=Legacy" json:"Legacy,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Address) Reset()         { *m = Address{} }
func (m *Address) String() string { return proto.CompactTextString(m) }
func (*Address) ProtoMessage()    {}
func (*Address) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_20997c42f981569c, []int{3}
}
func (m *Address) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Address.Unmarshal(m, b)
}
func (m *Address) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Address.Marshal(b, m, deterministic)
}
func (dst *Address) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Address.Merge(dst, src)
}
func (m *Address) XXX_Size() int {
	return xxx_messageInfo_Address.Size(m)
}
func (m *Address) XXX_DiscardUnknown() {
	xxx_messageInfo_Address.DiscardUnknown(m)
}

var xxx_messageInfo_Address proto.InternalMessageInfo

func (m *Address) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Address) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Address) GetLegacy() *Legacy {
	if m != nil {
		return m.Legacy
	}
	return nil
}

type Addresses struct {
	Address              []*Address `protobuf:"bytes,1,rep,name=Address" json:"Address,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *Addresses) Reset()         { *m = Addresses{} }
func (m *Addresses) String() string { return proto.CompactTextString(m) }
func (*Addresses) ProtoMessage()    {}
func (*Addresses) Descriptor() ([]byte, []int) {
	return fileDescriptor_address_20997c42f981569c, []int{4}
}
func (m *Addresses) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Addresses.Unmarshal(m, b)
}
func (m *Addresses) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Addresses.Marshal(b, m, deterministic)
}
func (dst *Addresses) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Addresses.Merge(dst, src)
}
func (m *Addresses) XXX_Size() int {
	return xxx_messageInfo_Addresses.Size(m)
}
func (m *Addresses) XXX_DiscardUnknown() {
	xxx_messageInfo_Addresses.DiscardUnknown(m)
}

var xxx_messageInfo_Addresses proto.InternalMessageInfo

func (m *Addresses) GetAddress() []*Address {
	if m != nil {
		return m.Address
	}
	return nil
}

func init() {
	proto.RegisterType((*AddressGetOptions)(nil), "proto.AddressGetOptions")
	proto.RegisterType((*AddressGetByPathOptions)(nil), "proto.AddressGetByPathOptions")
	proto.RegisterType((*AddressGetMultiOptions)(nil), "proto.AddressGetMultiOptions")
	proto.RegisterType((*Address)(nil), "proto.Address")
	proto.RegisterType((*Addresses)(nil), "proto.Addresses")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AddressServiceClient is the client API for AddressService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AddressServiceClient interface {
	Get(ctx context.Context, in *AddressGetOptions, opts ...grpc.CallOption) (*Address, error)
	GetByPath(ctx context.Context, in *AddressGetByPathOptions, opts ...grpc.CallOption) (*Address, error)
	GetMulti(ctx context.Context, in *AddressGetMultiOptions, opts ...grpc.CallOption) (*Addresses, error)
}

type addressServiceClient struct {
	cc *grpc.ClientConn
}

func NewAddressServiceClient(cc *grpc.ClientConn) AddressServiceClient {
	return &addressServiceClient{cc}
}

func (c *addressServiceClient) Get(ctx context.Context, in *AddressGetOptions, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/proto.AddressService/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) GetByPath(ctx context.Context, in *AddressGetByPathOptions, opts ...grpc.CallOption) (*Address, error) {
	out := new(Address)
	err := c.cc.Invoke(ctx, "/proto.AddressService/GetByPath", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *addressServiceClient) GetMulti(ctx context.Context, in *AddressGetMultiOptions, opts ...grpc.CallOption) (*Addresses, error) {
	out := new(Addresses)
	err := c.cc.Invoke(ctx, "/proto.AddressService/GetMulti", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AddressServiceServer is the server API for AddressService service.
type AddressServiceServer interface {
	Get(context.Context, *AddressGetOptions) (*Address, error)
	GetByPath(context.Context, *AddressGetByPathOptions) (*Address, error)
	GetMulti(context.Context, *AddressGetMultiOptions) (*Addresses, error)
}

func RegisterAddressServiceServer(s *grpc.Server, srv AddressServiceServer) {
	s.RegisterService(&_AddressService_serviceDesc, srv)
}

func _AddressService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressGetOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddressService/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).Get(ctx, req.(*AddressGetOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_GetByPath_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressGetByPathOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GetByPath(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddressService/GetByPath",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GetByPath(ctx, req.(*AddressGetByPathOptions))
	}
	return interceptor(ctx, in, info, handler)
}

func _AddressService_GetMulti_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddressGetMultiOptions)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AddressServiceServer).GetMulti(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.AddressService/GetMulti",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AddressServiceServer).GetMulti(ctx, req.(*AddressGetMultiOptions))
	}
	return interceptor(ctx, in, info, handler)
}

var _AddressService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.AddressService",
	HandlerType: (*AddressServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AddressService_Get_Handler,
		},
		{
			MethodName: "GetByPath",
			Handler:    _AddressService_GetByPath_Handler,
		},
		{
			MethodName: "GetMulti",
			Handler:    _AddressService_GetMulti_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "address.proto",
}

func init() { proto.RegisterFile("address.proto", fileDescriptor_address_20997c42f981569c) }

var fileDescriptor_address_20997c42f981569c = []byte{
	// 269 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xc3, 0x40,
	0x10, 0x85, 0xbb, 0x89, 0x56, 0x33, 0xb5, 0xa1, 0xce, 0x41, 0x43, 0x40, 0x09, 0x2b, 0x42, 0x10,
	0x2c, 0x18, 0xf1, 0xe8, 0x41, 0x2f, 0xbd, 0x28, 0x4a, 0xf4, 0x0f, 0xc4, 0xee, 0x60, 0x17, 0x8a,
	0x09, 0xd9, 0x55, 0xe8, 0xcf, 0xf3, 0x9f, 0x89, 0x9b, 0x4d, 0xd6, 0x34, 0x9e, 0x32, 0xbc, 0xf9,
	0xde, 0x9b, 0xec, 0x83, 0x69, 0x21, 0x44, 0x4d, 0x4a, 0xcd, 0xab, 0xba, 0xd4, 0x25, 0xee, 0x9a,
	0x4f, 0x7c, 0xb0, 0xa6, 0xf7, 0x62, 0xb9, 0x69, 0x44, 0x7e, 0x06, 0x87, 0x77, 0x0d, 0xb5, 0x20,
	0xfd, 0x54, 0x69, 0x59, 0x7e, 0x28, 0x0c, 0xc1, 0x93, 0x22, 0x62, 0x09, 0x4b, 0x83, 0xdc, 0x93,
	0x82, 0x5f, 0xc2, 0xb1, 0x83, 0xee, 0x37, 0xcf, 0x85, 0x5e, 0xb5, 0x28, 0xc2, 0x4e, 0x55, 0xe8,
	0x95, 0x85, 0xcd, 0xcc, 0x2f, 0xe0, 0xc8, 0xe1, 0x8f, 0x9f, 0x6b, 0x2d, 0x5b, 0x7a, 0x06, 0xbe,
	0x14, 0x2a, 0x62, 0x89, 0x9f, 0x06, 0xf9, 0xef, 0xc8, 0x5f, 0x61, 0xcf, 0xb2, 0xdb, 0x57, 0xbb,
	0x68, 0xcf, 0x45, 0xe3, 0x39, 0x8c, 0x1f, 0xcc, 0xef, 0x47, 0x7e, 0xc2, 0xd2, 0x49, 0x36, 0x6d,
	0x9e, 0x31, 0x6f, 0xc4, 0xdc, 0x2e, 0xf9, 0x0d, 0x04, 0x36, 0x95, 0x14, 0xa6, 0xdd, 0x09, 0x73,
	0x78, 0x92, 0x85, 0xd6, 0x64, 0xd5, 0xbc, 0x5d, 0x67, 0xdf, 0x0c, 0x42, 0x3b, 0xbf, 0x50, 0xfd,
	0x25, 0x97, 0x84, 0x57, 0xe0, 0x2f, 0x48, 0x63, 0xd4, 0xb7, 0xb8, 0xae, 0xe2, 0xad, 0x30, 0x3e,
	0xc2, 0x5b, 0x08, 0xba, 0x9a, 0xf0, 0x74, 0x60, 0xec, 0xf5, 0xf7, 0xaf, 0x7d, 0xbf, 0xad, 0x0d,
	0x4f, 0x06, 0xee, 0xbf, 0x75, 0xc6, 0xb3, 0xfe, 0x9a, 0x14, 0x1f, 0xbd, 0x8d, 0x8d, 0x74, 0xfd,
	0x13, 0x00, 0x00, 0xff, 0xff, 0x98, 0x58, 0x18, 0xea, 0xfd, 0x01, 0x00, 0x00,
}
