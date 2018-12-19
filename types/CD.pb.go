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
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Deployment) Reset()         { *m = Deployment{} }
func (m *Deployment) String() string { return proto.CompactTextString(m) }
func (*Deployment) ProtoMessage()    {}
func (*Deployment) Descriptor() ([]byte, []int) {
	return fileDescriptor_CD_c5df9e912b748075, []int{0}
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

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_CD_c5df9e912b748075, []int{1}
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

func init() { proto.RegisterFile("CD.proto", fileDescriptor_CD_c5df9e912b748075) }

var fileDescriptor_CD_c5df9e912b748075 = []byte{
	// 210 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x70, 0x76, 0xd1, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x5a, 0xc4, 0xc8, 0xc5,
	0xe5, 0x92, 0x5a, 0x90, 0x93, 0x5f, 0x99, 0x9b, 0x9a, 0x57, 0x22, 0xc4, 0xc7, 0xc5, 0xe4, 0xe9,
	0x22, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x1c, 0xc4, 0xe4, 0xe9, 0x22, 0x24, 0xc6, 0xc5, 0xe6, 0x54,
	0x94, 0x98, 0x97, 0x9c, 0x21, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe5, 0x81, 0xc4, 0x83,
	0x52, 0x0b, 0xf2, 0x3d, 0x5d, 0x24, 0x98, 0xc1, 0x6a, 0xa1, 0x3c, 0x21, 0x39, 0x2e, 0x2e, 0xc7,
	0xa2, 0x92, 0xcc, 0xb4, 0xc4, 0xe4, 0x12, 0x4f, 0x17, 0x09, 0x16, 0xb0, 0x1c, 0x92, 0x88, 0x90,
	0x04, 0x17, 0xbb, 0xb7, 0x45, 0xb0, 0x5f, 0x62, 0x6e, 0xaa, 0x04, 0x2b, 0xd8, 0x40, 0x18, 0x57,
	0x48, 0x8a, 0x8b, 0xc3, 0x37, 0x31, 0x2f, 0x33, 0x2d, 0xb5, 0xb8, 0x44, 0x82, 0x0d, 0x2c, 0x05,
	0xe7, 0x2b, 0xb1, 0x73, 0xb1, 0xba, 0xe6, 0x16, 0x94, 0x54, 0x1a, 0xd9, 0x71, 0xb1, 0x3a, 0xbb,
	0x38, 0x06, 0x78, 0x0a, 0x99, 0x72, 0x09, 0x38, 0x17, 0xa5, 0x26, 0x96, 0xa4, 0x22, 0xb9, 0x5d,
	0x50, 0x0f, 0xec, 0x25, 0x3d, 0x84, 0x90, 0x14, 0x0f, 0x54, 0x08, 0xac, 0x5b, 0x89, 0x21, 0x89,
	0x0d, 0xec, 0x77, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1f, 0x51, 0x4f, 0x3f, 0x07, 0x01,
	0x00, 0x00,
}
