// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/top_level_domain.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TopLevelDomainRole struct {
	Owner string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *TopLevelDomainRole) Reset()         { *m = TopLevelDomainRole{} }
func (m *TopLevelDomainRole) String() string { return proto.CompactTextString(m) }
func (*TopLevelDomainRole) ProtoMessage()    {}
func (*TopLevelDomainRole) Descriptor() ([]byte, []int) {
	return fileDescriptor_0136e389ac8054f7, []int{0}
}
func (m *TopLevelDomainRole) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopLevelDomainRole) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TopLevelDomainRole.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TopLevelDomainRole) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopLevelDomainRole.Merge(m, src)
}
func (m *TopLevelDomainRole) XXX_Size() int {
	return m.Size()
}
func (m *TopLevelDomainRole) XXX_DiscardUnknown() {
	xxx_messageInfo_TopLevelDomainRole.DiscardUnknown(m)
}

var xxx_messageInfo_TopLevelDomainRole proto.InternalMessageInfo

func (m *TopLevelDomainRole) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

type TopLevelDomain struct {
	Name            string                         `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	ExpirationDate  int64                          `protobuf:"varint,2,opt,name=expirationDate,proto3" json:"expirationDate,omitempty"`
	Metadata        map[string]string              `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	SubdomainConfig *SubdomainConfig               `protobuf:"bytes,4,opt,name=subdomainConfig,proto3" json:"subdomainConfig,omitempty"`
	SubdomainCount  uint64                         `protobuf:"varint,5,opt,name=subdomainCount,proto3" json:"subdomainCount,omitempty"`
	AccessControl   map[string]*TopLevelDomainRole `protobuf:"bytes,6,rep,name=accessControl,proto3" json:"accessControl,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *TopLevelDomain) Reset()         { *m = TopLevelDomain{} }
func (m *TopLevelDomain) String() string { return proto.CompactTextString(m) }
func (*TopLevelDomain) ProtoMessage()    {}
func (*TopLevelDomain) Descriptor() ([]byte, []int) {
	return fileDescriptor_0136e389ac8054f7, []int{1}
}
func (m *TopLevelDomain) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TopLevelDomain) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TopLevelDomain.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TopLevelDomain) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TopLevelDomain.Merge(m, src)
}
func (m *TopLevelDomain) XXX_Size() int {
	return m.Size()
}
func (m *TopLevelDomain) XXX_DiscardUnknown() {
	xxx_messageInfo_TopLevelDomain.DiscardUnknown(m)
}

var xxx_messageInfo_TopLevelDomain proto.InternalMessageInfo

func (m *TopLevelDomain) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TopLevelDomain) GetExpirationDate() int64 {
	if m != nil {
		return m.ExpirationDate
	}
	return 0
}

func (m *TopLevelDomain) GetMetadata() map[string]string {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *TopLevelDomain) GetSubdomainConfig() *SubdomainConfig {
	if m != nil {
		return m.SubdomainConfig
	}
	return nil
}

func (m *TopLevelDomain) GetSubdomainCount() uint64 {
	if m != nil {
		return m.SubdomainCount
	}
	return 0
}

func (m *TopLevelDomain) GetAccessControl() map[string]*TopLevelDomainRole {
	if m != nil {
		return m.AccessControl
	}
	return nil
}

func init() {
	proto.RegisterType((*TopLevelDomainRole)(nil), "mycel.registry.TopLevelDomainRole")
	proto.RegisterType((*TopLevelDomain)(nil), "mycel.registry.TopLevelDomain")
	proto.RegisterMapType((map[string]*TopLevelDomainRole)(nil), "mycel.registry.TopLevelDomain.AccessControlEntry")
	proto.RegisterMapType((map[string]string)(nil), "mycel.registry.TopLevelDomain.MetadataEntry")
}

func init() {
	proto.RegisterFile("mycel/registry/top_level_domain.proto", fileDescriptor_0136e389ac8054f7)
}

var fileDescriptor_0136e389ac8054f7 = []byte{
	// 392 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xee, 0x36, 0x69, 0x79, 0xdd, 0xd2, 0xbe, 0xc7, 0xf2, 0x0e, 0xa1, 0x87, 0xbc, 0x50, 0x78,
	0x12, 0x44, 0x13, 0xac, 0x97, 0xa2, 0x27, 0x6d, 0x85, 0x0a, 0x7a, 0x89, 0x82, 0xe0, 0xa5, 0x6c,
	0xd3, 0xb5, 0x06, 0x93, 0xdd, 0x90, 0x6c, 0x6a, 0xf3, 0x03, 0xbc, 0xfb, 0xb3, 0x3c, 0xf6, 0xe8,
	0x51, 0xda, 0x3f, 0x22, 0xd9, 0xad, 0x25, 0x89, 0xa2, 0xb7, 0x99, 0xd9, 0xef, 0xfb, 0x66, 0xbe,
	0xd9, 0x81, 0xff, 0x83, 0xd4, 0x25, 0xbe, 0x1d, 0x91, 0x99, 0x17, 0xf3, 0x28, 0xb5, 0x39, 0x0b,
	0xc7, 0x3e, 0x99, 0x13, 0x7f, 0x3c, 0x65, 0x01, 0xf6, 0xa8, 0x15, 0x46, 0x8c, 0x33, 0xd4, 0x16,
	0x30, 0xeb, 0x03, 0xd6, 0x29, 0xd3, 0xe2, 0x64, 0x22, 0xf1, 0x63, 0x97, 0xd1, 0x3b, 0x6f, 0x26,
	0x69, 0xdd, 0x5d, 0x88, 0xae, 0x59, 0x78, 0x91, 0xe9, 0x0d, 0xc5, 0xb3, 0xc3, 0x7c, 0x82, 0xfe,
	0xc2, 0x1a, 0x7b, 0xa4, 0x24, 0xd2, 0x80, 0x01, 0xcc, 0x86, 0x23, 0x93, 0xee, 0x93, 0x0a, 0xdb,
	0x45, 0x30, 0x42, 0x50, 0xa5, 0x38, 0x20, 0x1b, 0x9c, 0x88, 0xd1, 0x0e, 0x6c, 0x93, 0x45, 0xe8,
	0x45, 0x98, 0x7b, 0x8c, 0x0e, 0x31, 0x27, 0x5a, 0xd5, 0x00, 0xa6, 0xe2, 0x94, 0xaa, 0x68, 0x04,
	0x7f, 0x05, 0x84, 0xe3, 0x29, 0xe6, 0x58, 0x53, 0x0c, 0xc5, 0x6c, 0xf6, 0xf6, 0xac, 0xa2, 0x09,
	0xab, 0xd8, 0xcd, 0xba, 0xdc, 0xc0, 0xcf, 0x28, 0x8f, 0x52, 0x67, 0xcb, 0x46, 0xe7, 0xf0, 0xf7,
	0xd6, 0xde, 0x40, 0xb8, 0xd3, 0x54, 0x03, 0x98, 0xcd, 0xde, 0xbf, 0xb2, 0xe0, 0x55, 0x11, 0xe6,
	0x94, 0x79, 0xd9, 0xf0, 0xb9, 0x52, 0x42, 0xb9, 0x56, 0x33, 0x80, 0xa9, 0x3a, 0xa5, 0x2a, 0xba,
	0x81, 0x2d, 0xec, 0xba, 0x24, 0x8e, 0x07, 0x8c, 0xf2, 0x88, 0xf9, 0x5a, 0x5d, 0x38, 0x38, 0xf8,
	0xc1, 0xc1, 0x49, 0x9e, 0x23, 0x6d, 0x14, 0x75, 0x3a, 0xc7, 0xb0, 0x55, 0xb0, 0x89, 0xfe, 0x40,
	0xe5, 0x81, 0xa4, 0x9b, 0x0d, 0x67, 0x61, 0xf6, 0x3b, 0x73, 0xec, 0x27, 0x72, 0xaf, 0x0d, 0x47,
	0x26, 0x47, 0xd5, 0x3e, 0xe8, 0x4c, 0x21, 0xfa, 0xdc, 0xe1, 0x0b, 0x85, 0x7e, 0x5e, 0xa1, 0xd9,
	0xeb, 0x7e, 0x3f, 0x75, 0x76, 0x12, 0xb9, 0x2e, 0xa7, 0xa3, 0x97, 0x95, 0x0e, 0x96, 0x2b, 0x1d,
	0xbc, 0xad, 0x74, 0xf0, 0xbc, 0xd6, 0x2b, 0xcb, 0xb5, 0x5e, 0x79, 0x5d, 0xeb, 0x95, 0x5b, 0x6b,
	0xe6, 0xf1, 0xfb, 0x64, 0x62, 0xb9, 0x2c, 0xb0, 0x85, 0xe4, 0xbe, 0xdc, 0x99, 0x4c, 0xec, 0x45,
	0xee, 0x8a, 0xd3, 0x90, 0xc4, 0x93, 0xba, 0x38, 0xc2, 0xc3, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xbf, 0x03, 0x9b, 0xe0, 0xe4, 0x02, 0x00, 0x00,
}

func (m *TopLevelDomainRole) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopLevelDomainRole) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopLevelDomainRole) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TopLevelDomain) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TopLevelDomain) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TopLevelDomain) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccessControl) > 0 {
		for k := range m.AccessControl {
			v := m.AccessControl[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintTopLevelDomain(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTopLevelDomain(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTopLevelDomain(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x32
		}
	}
	if m.SubdomainCount != 0 {
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(m.SubdomainCount))
		i--
		dAtA[i] = 0x28
	}
	if m.SubdomainConfig != nil {
		{
			size, err := m.SubdomainConfig.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTopLevelDomain(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Metadata) > 0 {
		for k := range m.Metadata {
			v := m.Metadata[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintTopLevelDomain(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintTopLevelDomain(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintTopLevelDomain(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.ExpirationDate != 0 {
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(m.ExpirationDate))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTopLevelDomain(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTopLevelDomain(dAtA []byte, offset int, v uint64) int {
	offset -= sovTopLevelDomain(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TopLevelDomainRole) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTopLevelDomain(uint64(l))
	}
	return n
}

func (m *TopLevelDomain) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTopLevelDomain(uint64(l))
	}
	if m.ExpirationDate != 0 {
		n += 1 + sovTopLevelDomain(uint64(m.ExpirationDate))
	}
	if len(m.Metadata) > 0 {
		for k, v := range m.Metadata {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovTopLevelDomain(uint64(len(k))) + 1 + len(v) + sovTopLevelDomain(uint64(len(v)))
			n += mapEntrySize + 1 + sovTopLevelDomain(uint64(mapEntrySize))
		}
	}
	if m.SubdomainConfig != nil {
		l = m.SubdomainConfig.Size()
		n += 1 + l + sovTopLevelDomain(uint64(l))
	}
	if m.SubdomainCount != 0 {
		n += 1 + sovTopLevelDomain(uint64(m.SubdomainCount))
	}
	if len(m.AccessControl) > 0 {
		for k, v := range m.AccessControl {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovTopLevelDomain(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovTopLevelDomain(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovTopLevelDomain(uint64(mapEntrySize))
		}
	}
	return n
}

func sovTopLevelDomain(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTopLevelDomain(x uint64) (n int) {
	return sovTopLevelDomain(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TopLevelDomainRole) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopLevelDomain
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TopLevelDomainRole: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopLevelDomainRole: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopLevelDomain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TopLevelDomain) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTopLevelDomain
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TopLevelDomain: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TopLevelDomain: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpirationDate", wireType)
			}
			m.ExpirationDate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpirationDate |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Metadata == nil {
				m.Metadata = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopLevelDomain
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTopLevelDomain
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTopLevelDomain
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTopLevelDomain(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Metadata[mapkey] = mapvalue
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubdomainConfig", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubdomainConfig == nil {
				m.SubdomainConfig = &SubdomainConfig{}
			}
			if err := m.SubdomainConfig.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubdomainCount", wireType)
			}
			m.SubdomainCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SubdomainCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccessControl", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.AccessControl == nil {
				m.AccessControl = make(map[string]*TopLevelDomainRole)
			}
			var mapkey string
			var mapvalue *TopLevelDomainRole
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowTopLevelDomain
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTopLevelDomain
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowTopLevelDomain
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &TopLevelDomainRole{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipTopLevelDomain(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthTopLevelDomain
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.AccessControl[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTopLevelDomain(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTopLevelDomain
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTopLevelDomain(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTopLevelDomain
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
					return 0, ErrIntOverflowTopLevelDomain
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTopLevelDomain
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
			if length < 0 {
				return 0, ErrInvalidLengthTopLevelDomain
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTopLevelDomain
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTopLevelDomain
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTopLevelDomain        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTopLevelDomain          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTopLevelDomain = fmt.Errorf("proto: unexpected end of group")
)
