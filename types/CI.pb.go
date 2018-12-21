// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CI.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"

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

type Build struct {
	ID                   int64                `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	GithubRepoID         int64                `protobuf:"varint,2,opt,name=GithubRepoID,proto3" json:"GithubRepoID,omitempty"`
	IsSuccessfull        bool                 `protobuf:"varint,3,opt,name=IsSuccessfull,proto3" json:"IsSuccessfull,omitempty"`
	Date                 *timestamp.Timestamp `protobuf:"bytes,4,opt,name=Date,proto3" json:"Date,omitempty"`
	Branch               string               `protobuf:"bytes,5,opt,name=Branch,proto3" json:"Branch,omitempty"`
	Stdout               string               `protobuf:"bytes,6,opt,name=Stdout,proto3" json:"Stdout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Build) Reset()         { *m = Build{} }
func (m *Build) String() string { return proto.CompactTextString(m) }
func (*Build) ProtoMessage()    {}
func (*Build) Descriptor() ([]byte, []int) {
	return fileDescriptor_CI_fb63f34039493eb6, []int{0}
}
func (m *Build) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Build.Unmarshal(m, b)
}
func (m *Build) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Build.Marshal(b, m, deterministic)
}
func (dst *Build) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Build.Merge(dst, src)
}
func (m *Build) XXX_Size() int {
	return xxx_messageInfo_Build.Size(m)
}
func (m *Build) XXX_DiscardUnknown() {
	xxx_messageInfo_Build.DiscardUnknown(m)
}

var xxx_messageInfo_Build proto.InternalMessageInfo

func (m *Build) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Build) GetGithubRepoID() int64 {
	if m != nil {
		return m.GithubRepoID
	}
	return 0
}

func (m *Build) GetIsSuccessfull() bool {
	if m != nil {
		return m.IsSuccessfull
	}
	return false
}

func (m *Build) GetDate() *timestamp.Timestamp {
	if m != nil {
		return m.Date
	}
	return nil
}

func (m *Build) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Build) GetStdout() string {
	if m != nil {
		return m.Stdout
	}
	return ""
}

type BuildsQuery struct {
	RepoID               int64    `protobuf:"varint,1,opt,name=RepoID,proto3" json:"RepoID,omitempty"`
	Branch               string   `protobuf:"bytes,2,opt,name=Branch,proto3" json:"Branch,omitempty"`
	Page                 int64    `protobuf:"varint,3,opt,name=Page,proto3" json:"Page,omitempty"`
	Limit                int64    `protobuf:"varint,4,opt,name=Limit,proto3" json:"Limit,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildsQuery) Reset()         { *m = BuildsQuery{} }
func (m *BuildsQuery) String() string { return proto.CompactTextString(m) }
func (*BuildsQuery) ProtoMessage()    {}
func (*BuildsQuery) Descriptor() ([]byte, []int) {
	return fileDescriptor_CI_fb63f34039493eb6, []int{1}
}
func (m *BuildsQuery) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildsQuery.Unmarshal(m, b)
}
func (m *BuildsQuery) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildsQuery.Marshal(b, m, deterministic)
}
func (dst *BuildsQuery) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildsQuery.Merge(dst, src)
}
func (m *BuildsQuery) XXX_Size() int {
	return xxx_messageInfo_BuildsQuery.Size(m)
}
func (m *BuildsQuery) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildsQuery.DiscardUnknown(m)
}

var xxx_messageInfo_BuildsQuery proto.InternalMessageInfo

func (m *BuildsQuery) GetRepoID() int64 {
	if m != nil {
		return m.RepoID
	}
	return 0
}

func (m *BuildsQuery) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *BuildsQuery) GetPage() int64 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *BuildsQuery) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

type BuildsArray struct {
	Builds               []*Build `protobuf:"bytes,1,rep,name=builds,proto3" json:"builds,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildsArray) Reset()         { *m = BuildsArray{} }
func (m *BuildsArray) String() string { return proto.CompactTextString(m) }
func (*BuildsArray) ProtoMessage()    {}
func (*BuildsArray) Descriptor() ([]byte, []int) {
	return fileDescriptor_CI_fb63f34039493eb6, []int{2}
}
func (m *BuildsArray) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildsArray.Unmarshal(m, b)
}
func (m *BuildsArray) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildsArray.Marshal(b, m, deterministic)
}
func (dst *BuildsArray) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildsArray.Merge(dst, src)
}
func (m *BuildsArray) XXX_Size() int {
	return xxx_messageInfo_BuildsArray.Size(m)
}
func (m *BuildsArray) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildsArray.DiscardUnknown(m)
}

var xxx_messageInfo_BuildsArray proto.InternalMessageInfo

func (m *BuildsArray) GetBuilds() []*Build {
	if m != nil {
		return m.Builds
	}
	return nil
}

type BuildArtifact struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	BuildID              int64    `protobuf:"varint,2,opt,name=BuildID,proto3" json:"BuildID,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuildArtifact) Reset()         { *m = BuildArtifact{} }
func (m *BuildArtifact) String() string { return proto.CompactTextString(m) }
func (*BuildArtifact) ProtoMessage()    {}
func (*BuildArtifact) Descriptor() ([]byte, []int) {
	return fileDescriptor_CI_fb63f34039493eb6, []int{3}
}
func (m *BuildArtifact) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuildArtifact.Unmarshal(m, b)
}
func (m *BuildArtifact) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuildArtifact.Marshal(b, m, deterministic)
}
func (dst *BuildArtifact) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuildArtifact.Merge(dst, src)
}
func (m *BuildArtifact) XXX_Size() int {
	return xxx_messageInfo_BuildArtifact.Size(m)
}
func (m *BuildArtifact) XXX_DiscardUnknown() {
	xxx_messageInfo_BuildArtifact.DiscardUnknown(m)
}

var xxx_messageInfo_BuildArtifact proto.InternalMessageInfo

func (m *BuildArtifact) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *BuildArtifact) GetBuildID() int64 {
	if m != nil {
		return m.BuildID
	}
	return 0
}

func (m *BuildArtifact) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Build)(nil), "types.Build")
	proto.RegisterType((*BuildsQuery)(nil), "types.BuildsQuery")
	proto.RegisterType((*BuildsArray)(nil), "types.BuildsArray")
	proto.RegisterType((*BuildArtifact)(nil), "types.BuildArtifact")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CIAPIClient is the client API for CIAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CIAPIClient interface {
	GetBuild(ctx context.Context, in *Build, opts ...grpc.CallOption) (*Build, error)
	GetBuildArtifact(ctx context.Context, in *BuildArtifact, opts ...grpc.CallOption) (*BuildArtifact, error)
	GetBuildArtifactByID(ctx context.Context, in *BuildArtifact, opts ...grpc.CallOption) (*BuildArtifact, error)
	GetAllBuilds(ctx context.Context, in *BuildsQuery, opts ...grpc.CallOption) (*BuildsArray, error)
}

type cIAPIClient struct {
	cc *grpc.ClientConn
}

func NewCIAPIClient(cc *grpc.ClientConn) CIAPIClient {
	return &cIAPIClient{cc}
}

func (c *cIAPIClient) GetBuild(ctx context.Context, in *Build, opts ...grpc.CallOption) (*Build, error) {
	out := new(Build)
	err := c.cc.Invoke(ctx, "/types.CIAPI/GetBuild", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cIAPIClient) GetBuildArtifact(ctx context.Context, in *BuildArtifact, opts ...grpc.CallOption) (*BuildArtifact, error) {
	out := new(BuildArtifact)
	err := c.cc.Invoke(ctx, "/types.CIAPI/GetBuildArtifact", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cIAPIClient) GetBuildArtifactByID(ctx context.Context, in *BuildArtifact, opts ...grpc.CallOption) (*BuildArtifact, error) {
	out := new(BuildArtifact)
	err := c.cc.Invoke(ctx, "/types.CIAPI/GetBuildArtifactByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cIAPIClient) GetAllBuilds(ctx context.Context, in *BuildsQuery, opts ...grpc.CallOption) (*BuildsArray, error) {
	out := new(BuildsArray)
	err := c.cc.Invoke(ctx, "/types.CIAPI/GetAllBuilds", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CIAPIServer is the server API for CIAPI service.
type CIAPIServer interface {
	GetBuild(context.Context, *Build) (*Build, error)
	GetBuildArtifact(context.Context, *BuildArtifact) (*BuildArtifact, error)
	GetBuildArtifactByID(context.Context, *BuildArtifact) (*BuildArtifact, error)
	GetAllBuilds(context.Context, *BuildsQuery) (*BuildsArray, error)
}

func RegisterCIAPIServer(s *grpc.Server, srv CIAPIServer) {
	s.RegisterService(&_CIAPI_serviceDesc, srv)
}

func _CIAPI_GetBuild_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Build)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CIAPIServer).GetBuild(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.CIAPI/GetBuild",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CIAPIServer).GetBuild(ctx, req.(*Build))
	}
	return interceptor(ctx, in, info, handler)
}

func _CIAPI_GetBuildArtifact_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildArtifact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CIAPIServer).GetBuildArtifact(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.CIAPI/GetBuildArtifact",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CIAPIServer).GetBuildArtifact(ctx, req.(*BuildArtifact))
	}
	return interceptor(ctx, in, info, handler)
}

func _CIAPI_GetBuildArtifactByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildArtifact)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CIAPIServer).GetBuildArtifactByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.CIAPI/GetBuildArtifactByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CIAPIServer).GetBuildArtifactByID(ctx, req.(*BuildArtifact))
	}
	return interceptor(ctx, in, info, handler)
}

func _CIAPI_GetAllBuilds_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildsQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CIAPIServer).GetAllBuilds(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.CIAPI/GetAllBuilds",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CIAPIServer).GetAllBuilds(ctx, req.(*BuildsQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _CIAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.CIAPI",
	HandlerType: (*CIAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBuild",
			Handler:    _CIAPI_GetBuild_Handler,
		},
		{
			MethodName: "GetBuildArtifact",
			Handler:    _CIAPI_GetBuildArtifact_Handler,
		},
		{
			MethodName: "GetBuildArtifactByID",
			Handler:    _CIAPI_GetBuildArtifactByID_Handler,
		},
		{
			MethodName: "GetAllBuilds",
			Handler:    _CIAPI_GetAllBuilds_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "CI.proto",
}

func init() { proto.RegisterFile("CI.proto", fileDescriptor_CI_fb63f34039493eb6) }

var fileDescriptor_CI_fb63f34039493eb6 = []byte{
	// 385 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x0e, 0x9a, 0x30,
	0x1c, 0xc6, 0x2d, 0x08, 0xd3, 0xbf, 0xba, 0x2c, 0x8d, 0x59, 0x1a, 0x2e, 0x23, 0x8d, 0x07, 0x4e,
	0x98, 0xe8, 0x65, 0xc7, 0xa1, 0x24, 0xa6, 0xc9, 0xb6, 0xb8, 0xba, 0x17, 0x00, 0xac, 0x48, 0x02,
	0x83, 0x40, 0x39, 0xf0, 0x78, 0x7b, 0xa6, 0xbd, 0xc0, 0x62, 0x0b, 0x46, 0xb6, 0x5d, 0x76, 0xeb,
	0xf7, 0xf5, 0xcb, 0xbf, 0xbf, 0x7e, 0x2d, 0xcc, 0x8e, 0xcc, 0xaf, 0xea, 0x52, 0x96, 0xd8, 0x92,
	0x5d, 0x25, 0x1a, 0xe7, 0x43, 0x5a, 0x96, 0x69, 0x2e, 0xb6, 0xca, 0x8c, 0xdb, 0xdb, 0x56, 0x66,
	0x85, 0x68, 0x64, 0x54, 0x54, 0x3a, 0x47, 0x7f, 0x22, 0xb0, 0x0e, 0x6d, 0x96, 0x5f, 0xf1, 0x5b,
	0x30, 0x58, 0x48, 0x90, 0x8b, 0x3c, 0x93, 0x1b, 0x2c, 0xc4, 0x14, 0x96, 0xa7, 0x4c, 0xde, 0xdb,
	0x98, 0x8b, 0xaa, 0x64, 0x21, 0x31, 0xd4, 0xce, 0xc8, 0xc3, 0x1b, 0x58, 0xb1, 0xe6, 0xd2, 0x26,
	0x89, 0x68, 0x9a, 0x5b, 0x9b, 0xe7, 0xc4, 0x74, 0x91, 0x37, 0xe3, 0x63, 0x13, 0xfb, 0x30, 0x0d,
	0x23, 0x29, 0xc8, 0xd4, 0x45, 0xde, 0x62, 0xe7, 0xf8, 0x9a, 0xc9, 0x1f, 0x98, 0xfc, 0xef, 0x03,
	0x13, 0x57, 0x39, 0xfc, 0x1e, 0xec, 0x43, 0x1d, 0xfd, 0x48, 0xee, 0xc4, 0x72, 0x91, 0x37, 0xe7,
	0xbd, 0x7a, 0xf8, 0x17, 0x79, 0x2d, 0x5b, 0x49, 0x6c, 0xed, 0x6b, 0x45, 0x53, 0x58, 0xa8, 0x2b,
	0x34, 0xdf, 0x5a, 0x51, 0x77, 0x8f, 0x58, 0x8f, 0xac, 0x2f, 0xd3, 0xab, 0x97, 0xb1, 0xc6, 0x68,
	0x2c, 0x86, 0xe9, 0x39, 0x4a, 0x85, 0x62, 0x37, 0xb9, 0x5a, 0xe3, 0x35, 0x58, 0x9f, 0xb3, 0x22,
	0x93, 0x8a, 0xd9, 0xe4, 0x5a, 0xd0, 0xfd, 0x70, 0x50, 0x50, 0xd7, 0x51, 0x87, 0x37, 0x60, 0xc7,
	0x4a, 0x12, 0xe4, 0x9a, 0xde, 0x62, 0xb7, 0xf4, 0x55, 0xe9, 0xbe, 0xca, 0xf0, 0x7e, 0x8f, 0x7e,
	0x81, 0x95, 0x32, 0x82, 0x5a, 0x66, 0xb7, 0x28, 0x91, 0x7f, 0x15, 0x4d, 0xe0, 0x8d, 0x0a, 0x3c,
	0x3b, 0x1e, 0xe4, 0x83, 0xec, 0x6b, 0x54, 0x68, 0xb2, 0x39, 0x57, 0xeb, 0xdd, 0x2f, 0x04, 0xd6,
	0x91, 0x05, 0x67, 0x86, 0x3d, 0x98, 0x9d, 0x84, 0xd4, 0x8f, 0x37, 0x3a, 0xda, 0x19, 0x29, 0x3a,
	0xc1, 0x9f, 0xe0, 0xdd, 0x90, 0x7c, 0x52, 0xac, 0x5f, 0x33, 0x83, 0xeb, 0xfc, 0xd3, 0xa5, 0x13,
	0x1c, 0xc2, 0xfa, 0xcf, 0x09, 0x87, 0x8e, 0x85, 0xff, 0x39, 0xe5, 0x23, 0x2c, 0x4f, 0x42, 0x06,
	0x79, 0xae, 0x5b, 0xc4, 0xf8, 0x35, 0xa7, 0x5f, 0xcf, 0x19, 0x7b, 0xaa, 0x68, 0x3a, 0x89, 0x6d,
	0xf5, 0x59, 0xf6, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb9, 0x5a, 0xb2, 0x34, 0xe1, 0x02, 0x00,
	0x00,
}
