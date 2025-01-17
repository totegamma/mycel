// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: mycel/registry/dns_record.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DnsRecordType int32

const (
	DnsRecordType_A     DnsRecordType = 0
	DnsRecordType_AAAA  DnsRecordType = 1
	DnsRecordType_CNAME DnsRecordType = 2
	DnsRecordType_NS    DnsRecordType = 3
	DnsRecordType_MX    DnsRecordType = 4
	DnsRecordType_PTR   DnsRecordType = 5
	DnsRecordType_SOA   DnsRecordType = 6
	DnsRecordType_SRV   DnsRecordType = 7
	DnsRecordType_TXT   DnsRecordType = 8
)

var DnsRecordType_name = map[int32]string{
	0: "A",
	1: "AAAA",
	2: "CNAME",
	3: "NS",
	4: "MX",
	5: "PTR",
	6: "SOA",
	7: "SRV",
	8: "TXT",
}

var DnsRecordType_value = map[string]int32{
	"A":     0,
	"AAAA":  1,
	"CNAME": 2,
	"NS":    3,
	"MX":    4,
	"PTR":   5,
	"SOA":   6,
	"SRV":   7,
	"TXT":   8,
}

func (x DnsRecordType) String() string {
	return proto.EnumName(DnsRecordType_name, int32(x))
}

func (DnsRecordType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_84f206db7f2bcc4f, []int{0}
}

func init() {
	proto.RegisterEnum("mycel.registry.DnsRecordType", DnsRecordType_name, DnsRecordType_value)
}

func init() { proto.RegisterFile("mycel/registry/dns_record.proto", fileDescriptor_84f206db7f2bcc4f) }

var fileDescriptor_84f206db7f2bcc4f = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0xad, 0x4c, 0x4e,
	0xcd, 0xd1, 0x2f, 0x4a, 0x4d, 0xcf, 0x2c, 0x2e, 0x29, 0xaa, 0xd4, 0x4f, 0xc9, 0x2b, 0x8e, 0x2f,
	0x4a, 0x4d, 0xce, 0x2f, 0x4a, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03, 0x2b, 0xd0,
	0x83, 0x29, 0xd0, 0x8a, 0xe7, 0xe2, 0x75, 0xc9, 0x2b, 0x0e, 0x02, 0x2b, 0x09, 0xa9, 0x2c, 0x48,
	0x15, 0x62, 0xe5, 0x62, 0x74, 0x14, 0x60, 0x10, 0xe2, 0xe0, 0x62, 0x71, 0x74, 0x74, 0x74, 0x14,
	0x60, 0x14, 0xe2, 0xe4, 0x62, 0x75, 0xf6, 0x73, 0xf4, 0x75, 0x15, 0x60, 0x12, 0x62, 0xe3, 0x62,
	0xf2, 0x0b, 0x16, 0x60, 0x06, 0xd1, 0xbe, 0x11, 0x02, 0x2c, 0x42, 0xec, 0x5c, 0xcc, 0x01, 0x21,
	0x41, 0x02, 0xac, 0x20, 0x46, 0xb0, 0xbf, 0xa3, 0x00, 0x1b, 0x98, 0x11, 0x14, 0x26, 0xc0, 0x0e,
	0x62, 0x84, 0x44, 0x84, 0x08, 0x70, 0x38, 0x79, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c,
	0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1,
	0x1c, 0x43, 0x94, 0x5e, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x3e, 0xd8,
	0x55, 0xba, 0x29, 0xf9, 0xb9, 0x89, 0x99, 0x79, 0x10, 0x8e, 0x7e, 0x05, 0xc2, 0x17, 0x25, 0x95,
	0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x1f, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x8a, 0x45,
	0x85, 0x11, 0xe4, 0x00, 0x00, 0x00,
}
