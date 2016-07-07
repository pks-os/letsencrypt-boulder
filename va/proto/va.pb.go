// Code generated by protoc-gen-go.
// source: va/proto/va.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	va/proto/va.proto

It has these top-level messages:
	IsSafeDomainRequest
	IsDomainSafe
	PerformValidationRequest
	AuthzMeta
	ValidationResult
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import core "github.com/letsencrypt/boulder/core/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type IsSafeDomainRequest struct {
	Domain           *string `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *IsSafeDomainRequest) Reset()                    { *m = IsSafeDomainRequest{} }
func (m *IsSafeDomainRequest) String() string            { return proto1.CompactTextString(m) }
func (*IsSafeDomainRequest) ProtoMessage()               {}
func (*IsSafeDomainRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *IsSafeDomainRequest) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

type IsDomainSafe struct {
	IsSafe           *bool  `protobuf:"varint,1,opt,name=isSafe" json:"isSafe,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *IsDomainSafe) Reset()                    { *m = IsDomainSafe{} }
func (m *IsDomainSafe) String() string            { return proto1.CompactTextString(m) }
func (*IsDomainSafe) ProtoMessage()               {}
func (*IsDomainSafe) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *IsDomainSafe) GetIsSafe() bool {
	if m != nil && m.IsSafe != nil {
		return *m.IsSafe
	}
	return false
}

type PerformValidationRequest struct {
	Domain           *string         `protobuf:"bytes,1,opt,name=domain" json:"domain,omitempty"`
	Challenge        *core.Challenge `protobuf:"bytes,2,opt,name=challenge" json:"challenge,omitempty"`
	Authz            *AuthzMeta      `protobuf:"bytes,3,opt,name=authz" json:"authz,omitempty"`
	XXX_unrecognized []byte          `json:"-"`
}

func (m *PerformValidationRequest) Reset()                    { *m = PerformValidationRequest{} }
func (m *PerformValidationRequest) String() string            { return proto1.CompactTextString(m) }
func (*PerformValidationRequest) ProtoMessage()               {}
func (*PerformValidationRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *PerformValidationRequest) GetDomain() string {
	if m != nil && m.Domain != nil {
		return *m.Domain
	}
	return ""
}

func (m *PerformValidationRequest) GetChallenge() *core.Challenge {
	if m != nil {
		return m.Challenge
	}
	return nil
}

func (m *PerformValidationRequest) GetAuthz() *AuthzMeta {
	if m != nil {
		return m.Authz
	}
	return nil
}

type AuthzMeta struct {
	Id               *string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	RegID            *int64  `protobuf:"varint,2,opt,name=regID" json:"regID,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *AuthzMeta) Reset()                    { *m = AuthzMeta{} }
func (m *AuthzMeta) String() string            { return proto1.CompactTextString(m) }
func (*AuthzMeta) ProtoMessage()               {}
func (*AuthzMeta) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AuthzMeta) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *AuthzMeta) GetRegID() int64 {
	if m != nil && m.RegID != nil {
		return *m.RegID
	}
	return 0
}

type ValidationResult struct {
	Records          []*core.ValidationRecord `protobuf:"bytes,1,rep,name=records" json:"records,omitempty"`
	Problems         *core.ProblemDetails     `protobuf:"bytes,2,opt,name=problems" json:"problems,omitempty"`
	XXX_unrecognized []byte                   `json:"-"`
}

func (m *ValidationResult) Reset()                    { *m = ValidationResult{} }
func (m *ValidationResult) String() string            { return proto1.CompactTextString(m) }
func (*ValidationResult) ProtoMessage()               {}
func (*ValidationResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ValidationResult) GetRecords() []*core.ValidationRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

func (m *ValidationResult) GetProblems() *core.ProblemDetails {
	if m != nil {
		return m.Problems
	}
	return nil
}

func init() {
	proto1.RegisterType((*IsSafeDomainRequest)(nil), "va.IsSafeDomainRequest")
	proto1.RegisterType((*IsDomainSafe)(nil), "va.IsDomainSafe")
	proto1.RegisterType((*PerformValidationRequest)(nil), "va.PerformValidationRequest")
	proto1.RegisterType((*AuthzMeta)(nil), "va.AuthzMeta")
	proto1.RegisterType((*ValidationResult)(nil), "va.ValidationResult")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for VA service

type VAClient interface {
	IsSafeDomain(ctx context.Context, in *IsSafeDomainRequest, opts ...grpc.CallOption) (*IsDomainSafe, error)
	PerformValidation(ctx context.Context, in *PerformValidationRequest, opts ...grpc.CallOption) (*ValidationResult, error)
}

type vAClient struct {
	cc *grpc.ClientConn
}

func NewVAClient(cc *grpc.ClientConn) VAClient {
	return &vAClient{cc}
}

func (c *vAClient) IsSafeDomain(ctx context.Context, in *IsSafeDomainRequest, opts ...grpc.CallOption) (*IsDomainSafe, error) {
	out := new(IsDomainSafe)
	err := grpc.Invoke(ctx, "/va.VA/IsSafeDomain", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vAClient) PerformValidation(ctx context.Context, in *PerformValidationRequest, opts ...grpc.CallOption) (*ValidationResult, error) {
	out := new(ValidationResult)
	err := grpc.Invoke(ctx, "/va.VA/PerformValidation", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for VA service

type VAServer interface {
	IsSafeDomain(context.Context, *IsSafeDomainRequest) (*IsDomainSafe, error)
	PerformValidation(context.Context, *PerformValidationRequest) (*ValidationResult, error)
}

func RegisterVAServer(s *grpc.Server, srv VAServer) {
	s.RegisterService(&_VA_serviceDesc, srv)
}

func _VA_IsSafeDomain_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsSafeDomainRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VAServer).IsSafeDomain(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/va.VA/IsSafeDomain",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VAServer).IsSafeDomain(ctx, req.(*IsSafeDomainRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VA_PerformValidation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PerformValidationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VAServer).PerformValidation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/va.VA/PerformValidation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VAServer).PerformValidation(ctx, req.(*PerformValidationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VA_serviceDesc = grpc.ServiceDesc{
	ServiceName: "va.VA",
	HandlerType: (*VAServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsSafeDomain",
			Handler:    _VA_IsSafeDomain_Handler,
		},
		{
			MethodName: "PerformValidation",
			Handler:    _VA_PerformValidation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto1.RegisterFile("va/proto/va.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 311 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x74, 0x91, 0x4f, 0x4f, 0xc2, 0x30,
	0x18, 0xc6, 0x61, 0x04, 0x81, 0x57, 0x51, 0xa8, 0xa8, 0x0b, 0x21, 0xc6, 0x34, 0x11, 0x3d, 0x8d,
	0x84, 0xab, 0x27, 0x74, 0x97, 0x1d, 0x4c, 0x88, 0x26, 0x1c, 0xbc, 0xd5, 0xad, 0x40, 0x93, 0x8e,
	0x6a, 0xdb, 0xed, 0xe0, 0x67, 0xf0, 0x43, 0xdb, 0x3f, 0x53, 0x88, 0xca, 0xad, 0x7b, 0x7e, 0x4f,
	0xdf, 0xf7, 0xe9, 0x33, 0xe8, 0x97, 0x64, 0xf2, 0x26, 0x85, 0x16, 0x93, 0x92, 0x44, 0xee, 0x80,
	0x82, 0x92, 0x0c, 0xcf, 0x52, 0x21, 0x69, 0x05, 0xec, 0xd1, 0x23, 0x7c, 0x0d, 0xa7, 0x89, 0x7a,
	0x26, 0x4b, 0x1a, 0x8b, 0x9c, 0xb0, 0xcd, 0x13, 0x7d, 0x2f, 0xa8, 0xd2, 0xe8, 0x18, 0x0e, 0x32,
	0x27, 0x84, 0xf5, 0xab, 0xfa, 0x6d, 0x07, 0x5f, 0xc2, 0x51, 0xa2, 0xbc, 0xc5, 0x9a, 0x2d, 0x67,
	0xee, 0x9a, 0xe3, 0x6d, 0xcc, 0x21, 0x9c, 0x53, 0xb9, 0x14, 0x32, 0x5f, 0x10, 0xce, 0x32, 0xa2,
	0x99, 0xd8, 0x37, 0x0b, 0x61, 0xe8, 0xa4, 0x6b, 0xc2, 0x39, 0xdd, 0xac, 0x68, 0x18, 0x18, 0xe9,
	0x70, 0x7a, 0x12, 0xb9, 0x48, 0x0f, 0xdf, 0x32, 0x1a, 0x41, 0x93, 0x14, 0x7a, 0xfd, 0x11, 0x36,
	0x1c, 0xef, 0x46, 0xe6, 0x2d, 0x33, 0x2b, 0x3c, 0x52, 0x4d, 0xf0, 0x18, 0x3a, 0x3f, 0x1f, 0x08,
	0x20, 0x60, 0x59, 0x35, 0xba, 0x0b, 0x4d, 0x49, 0x57, 0x49, 0xec, 0xc6, 0x36, 0x70, 0x0a, 0xbd,
	0xdd, 0x38, 0xaa, 0xe0, 0x1a, 0xdd, 0x40, 0x4b, 0x52, 0xb3, 0x2d, 0x53, 0xe6, 0x4e, 0xc3, 0xcc,
	0x3e, 0xf7, 0xbb, 0x77, 0x8d, 0x16, 0xa3, 0x31, 0xb4, 0x4d, 0x45, 0xaf, 0x9c, 0xe6, 0xaa, 0x4a,
	0x39, 0xf0, 0xce, 0xb9, 0x57, 0x63, 0xb3, 0x9c, 0x71, 0x35, 0xfd, 0xac, 0x43, 0xb0, 0x98, 0xa1,
	0x3b, 0xdb, 0xd0, 0xb6, 0x48, 0x74, 0x61, 0x23, 0xff, 0x53, 0xed, 0xb0, 0xe7, 0xc1, 0xb6, 0x4c,
	0x5c, 0x43, 0x09, 0xf4, 0xff, 0xd4, 0x87, 0x46, 0xd6, 0xb8, 0xaf, 0xd5, 0xe1, 0xc0, 0xd2, 0xdf,
	0xaf, 0xc3, 0xb5, 0xfb, 0xd6, 0x4b, 0xd3, 0xfd, 0xd9, 0xaf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7d,
	0xf6, 0x5c, 0x13, 0x08, 0x02, 0x00, 0x00,
}
