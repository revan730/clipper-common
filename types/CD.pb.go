// Code generated by protoc-gen-go. DO NOT EDIT.
// source: CD.proto

package types

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

type Deployment struct {
	ID                   int64    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Branch               string   `protobuf:"bytes,2,opt,name=Branch,proto3" json:"Branch,omitempty"`
	RepoID               int64    `protobuf:"varint,3,opt,name=RepoID,proto3" json:"RepoID,omitempty"`
	ArtifactID           int64    `protobuf:"varint,4,opt,name=ArtifactID,proto3" json:"ArtifactID,omitempty"`
	K8SName              string   `protobuf:"bytes,5,opt,name=K8SName,proto3" json:"K8SName,omitempty"`
	Manifest             string   `protobuf:"bytes,6,opt,name=Manifest,proto3" json:"Manifest,omitempty"`
	Replicas             int64    `protobuf:"varint,7,opt,name=Replicas,proto3" json:"Replicas,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_CD_5cbdb7b3db38a64d, []int{0}
}
func (m *Deployment) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Deployment.Unmarshal(m, b)
}
func (m *Deployment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Deployment.Marshal(b, m, deterministic)
}
func (dst *Deployment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Deployment.Merge(dst, src)
}
func (m *Deployment) XXX_Size() int {
	return xxx_messageInfo_Deployment.Size(m)
}
func (m *Deployment) XXX_DiscardUnknown() {
	xxx_messageInfo_Deployment.DiscardUnknown(m)
}

var xxx_messageInfo_Deployment proto.InternalMessageInfo

func (m *Deployment) GetID() int64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Deployment) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *Deployment) GetRepoID() int64 {
	if m != nil {
		return m.RepoID
	}
	return 0
}

func (m *Deployment) GetArtifactID() int64 {
	if m != nil {
		return m.ArtifactID
	}
	return 0
}

func (m *Deployment) GetK8SName() string {
	if m != nil {
		return m.K8SName
	}
	return ""
}

func (m *Deployment) GetManifest() string {
	if m != nil {
		return m.Manifest
	}
	return ""
}

func (m *Deployment) GetReplicas() int64 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_CD_5cbdb7b3db38a64d, []int{1}
}
func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (dst *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(dst, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Deployment)(nil), "types.Deployment")
	proto.RegisterType((*Empty)(nil), "types.Empty")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CDAPIClient is the client API for CDAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CDAPIClient interface {
	CreateDeployment(ctx context.Context, in *Deployment, opts ...grpc.CallOption) (*Empty, error)
}

type cDAPIClient struct {
	cc *grpc.ClientConn
}

func NewCDAPIClient(cc *grpc.ClientConn) CDAPIClient {
	return &cDAPIClient{cc}
}

func (c *cDAPIClient) CreateDeployment(ctx context.Context, in *Deployment, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/types.CDAPI/CreateDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CDAPIServer is the server API for CDAPI service.
type CDAPIServer interface {
	CreateDeployment(context.Context, *Deployment) (*Empty, error)
}

func RegisterCDAPIServer(s *grpc.Server, srv CDAPIServer) {
	s.RegisterService(&_CDAPI_serviceDesc, srv)
}

func _CDAPI_CreateDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Deployment)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CDAPIServer).CreateDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/types.CDAPI/CreateDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CDAPIServer).CreateDeployment(ctx, req.(*Deployment))
	}
	return interceptor(ctx, in, info, handler)
}

var _CDAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "types.CDAPI",
	HandlerType: (*CDAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateDeployment",
			Handler:    _CDAPI_CreateDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "CD.proto",
}

func init() { proto.RegisterFile("CD.proto", fileDescriptor_CD_5cbdb7b3db38a64d) }

var fileDescriptor_CD_5cbdb7b3db38a64d = []byte{
	// 224 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x4d, 0x6a, 0x92, 0x3a, 0x88, 0xe8, 0x1e, 0x64, 0xe8, 0x41, 0x4a, 0x4e, 0x3d, 0xe5,
	0xa0, 0x08, 0x9e, 0x84, 0x9a, 0xf5, 0xb0, 0x88, 0x22, 0xeb, 0x13, 0xac, 0x61, 0x8a, 0x81, 0x26,
	0xbb, 0x6c, 0xe6, 0x92, 0x97, 0xf3, 0xd9, 0xa4, 0x43, 0x6c, 0x73, 0xfc, 0xbe, 0x7f, 0xf7, 0x87,
	0x7f, 0x60, 0x59, 0xeb, 0x2a, 0x44, 0xcf, 0x5e, 0x65, 0x3c, 0x06, 0x1a, 0xca, 0xdf, 0x04, 0x40,
	0x53, 0xd8, 0xfb, 0xb1, 0xa3, 0x9e, 0xd5, 0x15, 0xa4, 0x46, 0x63, 0xb2, 0x4e, 0x36, 0x0b, 0x9b,
	0x1a, 0xad, 0x6e, 0x21, 0x7f, 0x89, 0xae, 0x6f, 0x7e, 0x30, 0x5d, 0x27, 0x9b, 0x0b, 0x3b, 0xd1,
	0xc1, 0x5b, 0x0a, 0xde, 0x68, 0x5c, 0xc8, 0xdb, 0x89, 0xd4, 0x1d, 0xc0, 0x36, 0x72, 0xbb, 0x73,
	0x0d, 0x1b, 0x8d, 0xe7, 0x92, 0xcd, 0x8c, 0x42, 0x28, 0xde, 0x9e, 0xbe, 0x3e, 0x5c, 0x47, 0x98,
	0x49, 0xe1, 0x3f, 0xaa, 0x15, 0x2c, 0xdf, 0x5d, 0xdf, 0xee, 0x68, 0x60, 0xcc, 0x25, 0x3a, 0xf2,
	0x21, 0xb3, 0x14, 0xf6, 0x6d, 0xe3, 0x06, 0x2c, 0xa4, 0xf3, 0xc8, 0x65, 0x01, 0xd9, 0x6b, 0x17,
	0x78, 0xbc, 0x7f, 0x86, 0xac, 0xd6, 0xdb, 0x4f, 0xa3, 0x1e, 0xe1, 0xba, 0x8e, 0xe4, 0x98, 0x66,
	0xbb, 0x6e, 0x2a, 0x99, 0x5b, 0x9d, 0xd4, 0xea, 0x72, 0x52, 0xf2, 0xbb, 0x3c, 0xfb, 0xce, 0xe5,
	0x2e, 0x0f, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xb7, 0xbc, 0x7c, 0x87, 0x23, 0x01, 0x00, 0x00,
}