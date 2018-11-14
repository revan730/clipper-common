// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: CIJobMessage.proto

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CIJob struct {
	RepoURL              string   `protobuf:"bytes,1,opt,name=repoURL,proto3" json:"repoURL,omitempty"`
	Branch               string   `protobuf:"bytes,2,opt,name=branch,proto3" json:"branch,omitempty"`
	HeadSHA              string   `protobuf:"bytes,3,opt,name=headSHA,proto3" json:"headSHA,omitempty"`
	User                 string   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	AccessToken          string   `protobuf:"bytes,5,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RepoID               int64    `protobuf:"varint,6,opt,name=repoID,proto3" json:"repoID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CIJob) Reset()         { *m = CIJob{} }
func (m *CIJob) String() string { return proto.CompactTextString(m) }
func (*CIJob) ProtoMessage()    {}
func (*CIJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_CIJobMessage_007da5173ca81869, []int{0}
}
func (m *CIJob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CIJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CIJob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *CIJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CIJob.Merge(dst, src)
}
func (m *CIJob) XXX_Size() int {
	return m.Size()
}
func (m *CIJob) XXX_DiscardUnknown() {
	xxx_messageInfo_CIJob.DiscardUnknown(m)
}

var xxx_messageInfo_CIJob proto.InternalMessageInfo

func (m *CIJob) GetRepoURL() string {
	if m != nil {
		return m.RepoURL
	}
	return ""
}

func (m *CIJob) GetBranch() string {
	if m != nil {
		return m.Branch
	}
	return ""
}

func (m *CIJob) GetHeadSHA() string {
	if m != nil {
		return m.HeadSHA
	}
	return ""
}

func (m *CIJob) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *CIJob) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *CIJob) GetRepoID() int64 {
	if m != nil {
		return m.RepoID
	}
	return 0
}

func init() {
	proto.RegisterType((*CIJob)(nil), "types.CIJob")
}
func (m *CIJob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CIJob) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.RepoURL) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintCIJobMessage(dAtA, i, uint64(len(m.RepoURL)))
		i += copy(dAtA[i:], m.RepoURL)
	}
	if len(m.Branch) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintCIJobMessage(dAtA, i, uint64(len(m.Branch)))
		i += copy(dAtA[i:], m.Branch)
	}
	if len(m.HeadSHA) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintCIJobMessage(dAtA, i, uint64(len(m.HeadSHA)))
		i += copy(dAtA[i:], m.HeadSHA)
	}
	if len(m.User) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintCIJobMessage(dAtA, i, uint64(len(m.User)))
		i += copy(dAtA[i:], m.User)
	}
	if len(m.AccessToken) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintCIJobMessage(dAtA, i, uint64(len(m.AccessToken)))
		i += copy(dAtA[i:], m.AccessToken)
	}
	if m.RepoID != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintCIJobMessage(dAtA, i, uint64(m.RepoID))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintCIJobMessage(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *CIJob) Size() (n int) {
	var l int
	_ = l
	l = len(m.RepoURL)
	if l > 0 {
		n += 1 + l + sovCIJobMessage(uint64(l))
	}
	l = len(m.Branch)
	if l > 0 {
		n += 1 + l + sovCIJobMessage(uint64(l))
	}
	l = len(m.HeadSHA)
	if l > 0 {
		n += 1 + l + sovCIJobMessage(uint64(l))
	}
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovCIJobMessage(uint64(l))
	}
	l = len(m.AccessToken)
	if l > 0 {
		n += 1 + l + sovCIJobMessage(uint64(l))
	}
	if m.RepoID != 0 {
		n += 1 + sovCIJobMessage(uint64(m.RepoID))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovCIJobMessage(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozCIJobMessage(x uint64) (n int) {
	return sovCIJobMessage(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CIJob) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCIJobMessage
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: CIJob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CIJob: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepoURL", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCIJobMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCIJobMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.RepoURL = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Branch", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCIJobMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCIJobMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Branch = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HeadSHA", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCIJobMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCIJobMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HeadSHA = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCIJobMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCIJobMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessToken", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCIJobMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthCIJobMessage
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccessToken = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RepoID", wireType)
			}
			m.RepoID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCIJobMessage
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RepoID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipCIJobMessage(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCIJobMessage
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipCIJobMessage(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCIJobMessage
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCIJobMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowCIJobMessage
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthCIJobMessage
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowCIJobMessage
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipCIJobMessage(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthCIJobMessage = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCIJobMessage   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("CIJobMessage.proto", fileDescriptor_CIJobMessage_007da5173ca81869) }

var fileDescriptor_CIJobMessage_007da5173ca81869 = []byte{
	// 182 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x72, 0xf6, 0xf4, 0xca,
	0x4f, 0xf2, 0x4d, 0x2d, 0x2e, 0x4e, 0x4c, 0x4f, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x2d, 0xa9, 0x2c, 0x48, 0x2d, 0x56, 0x5a, 0xc8, 0xc8, 0xc5, 0x0a, 0x96, 0x15, 0x92, 0xe0, 0x62,
	0x2f, 0x4a, 0x2d, 0xc8, 0x0f, 0x0d, 0xf2, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71,
	0x85, 0xc4, 0xb8, 0xd8, 0x92, 0x8a, 0x12, 0xf3, 0x92, 0x33, 0x24, 0x98, 0xc0, 0x12, 0x50, 0x1e,
	0x48, 0x47, 0x46, 0x6a, 0x62, 0x4a, 0xb0, 0x87, 0xa3, 0x04, 0x33, 0x44, 0x07, 0x94, 0x2b, 0x24,
	0xc4, 0xc5, 0x52, 0x5a, 0x9c, 0x5a, 0x24, 0xc1, 0x02, 0x16, 0x06, 0xb3, 0x85, 0x14, 0xb8, 0xb8,
	0x13, 0x93, 0x93, 0x53, 0x8b, 0x8b, 0x43, 0xf2, 0xb3, 0x53, 0xf3, 0x24, 0x58, 0xc1, 0x52, 0xc8,
	0x42, 0x20, 0x7b, 0x40, 0x56, 0x7a, 0xba, 0x48, 0xb0, 0x29, 0x30, 0x6a, 0x30, 0x07, 0x41, 0x79,
	0x4e, 0x02, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x8c,
	0xc7, 0x72, 0x0c, 0x49, 0x6c, 0x60, 0x3f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xff, 0x0e,
	0xda, 0x1a, 0xd9, 0x00, 0x00, 0x00,
}
