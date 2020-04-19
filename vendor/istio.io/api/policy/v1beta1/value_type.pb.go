// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: policy/v1beta1/value_type.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	math "math"
	strconv "strconv"
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

// ValueType describes the types that values in the Istio system can take. These
// are used to describe the type of Attributes at run time, describe the type of
// the result of evaluating an expression, and to describe the runtime type of
// fields of other descriptors.
type ValueType int32

const (
	// Invalid, default value.
	VALUE_TYPE_UNSPECIFIED ValueType = 0
	// An undiscriminated variable-length string.
	STRING ValueType = 1
	// An undiscriminated 64-bit signed integer.
	INT64 ValueType = 2
	// An undiscriminated 64-bit floating-point value.
	DOUBLE ValueType = 3
	// An undiscriminated boolean value.
	BOOL ValueType = 4
	// A point in time.
	TIMESTAMP ValueType = 5
	// An IP address.
	IP_ADDRESS ValueType = 6
	// An email address.
	EMAIL_ADDRESS ValueType = 7
	// A URI.
	URI ValueType = 8
	// A DNS name.
	DNS_NAME ValueType = 9
	// A span between two points in time.
	DURATION ValueType = 10
	// A map string -> string, typically used by headers.
	STRING_MAP ValueType = 11
)

var ValueType_name = map[int32]string{
	0:  "VALUE_TYPE_UNSPECIFIED",
	1:  "STRING",
	2:  "INT64",
	3:  "DOUBLE",
	4:  "BOOL",
	5:  "TIMESTAMP",
	6:  "IP_ADDRESS",
	7:  "EMAIL_ADDRESS",
	8:  "URI",
	9:  "DNS_NAME",
	10: "DURATION",
	11: "STRING_MAP",
}

var ValueType_value = map[string]int32{
	"VALUE_TYPE_UNSPECIFIED": 0,
	"STRING":                 1,
	"INT64":                  2,
	"DOUBLE":                 3,
	"BOOL":                   4,
	"TIMESTAMP":              5,
	"IP_ADDRESS":             6,
	"EMAIL_ADDRESS":          7,
	"URI":                    8,
	"DNS_NAME":               9,
	"DURATION":               10,
	"STRING_MAP":             11,
}

func (ValueType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_72649b1dc07496e9, []int{0}
}

func init() {
	proto.RegisterEnum("istio.policy.v1beta1.ValueType", ValueType_name, ValueType_value)
}

func init() { proto.RegisterFile("policy/v1beta1/value_type.proto", fileDescriptor_72649b1dc07496e9) }

var fileDescriptor_72649b1dc07496e9 = []byte{
	// 309 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x90, 0xcd, 0x4a, 0xfb, 0x40,
	0x14, 0x47, 0x67, 0xfe, 0xfd, 0xcc, 0xfd, 0x5b, 0x19, 0x07, 0x71, 0xa1, 0x70, 0xdd, 0xbb, 0x68,
	0x28, 0x8a, 0xfb, 0xa9, 0x19, 0x65, 0x20, 0x5f, 0x26, 0x93, 0x82, 0x6e, 0x42, 0x2b, 0x59, 0x04,
	0x0a, 0x09, 0x1a, 0x0b, 0xdd, 0xf9, 0x08, 0x3e, 0x86, 0xef, 0xe0, 0x0b, 0xb8, 0xec, 0xb2, 0x4b,
	0x3b, 0xdd, 0xb8, 0xec, 0x23, 0x48, 0x1b, 0x11, 0x5c, 0xde, 0xc3, 0xe1, 0xf2, 0xe3, 0xc0, 0x69,
	0x59, 0x4c, 0xf3, 0x87, 0xb9, 0x3d, 0x1b, 0x4c, 0xb2, 0x6a, 0x3c, 0xb0, 0x67, 0xe3, 0xe9, 0x73,
	0x96, 0x56, 0xf3, 0x32, 0xeb, 0x97, 0x8f, 0x45, 0x55, 0xf0, 0xc3, 0xfc, 0xa9, 0xca, 0x8b, 0x7e,
	0xad, 0xf5, 0x7f, 0xb4, 0xb3, 0x77, 0x0a, 0xd6, 0x68, 0xab, 0xea, 0x79, 0x99, 0xf1, 0x63, 0x38,
	0x1a, 0x09, 0x37, 0x91, 0xa9, 0xbe, 0x0b, 0x65, 0x9a, 0xf8, 0x71, 0x28, 0xaf, 0xd4, 0xb5, 0x92,
	0x0e, 0x23, 0x1c, 0xa0, 0x1d, 0xeb, 0x48, 0xf9, 0x37, 0x8c, 0x72, 0x0b, 0x5a, 0xca, 0xd7, 0x97,
	0x17, 0xec, 0xdf, 0x16, 0x3b, 0x41, 0x32, 0x74, 0x25, 0x6b, 0xf0, 0x2e, 0x34, 0x87, 0x41, 0xe0,
	0xb2, 0x26, 0xef, 0x81, 0xa5, 0x95, 0x27, 0x63, 0x2d, 0xbc, 0x90, 0xb5, 0xf8, 0x3e, 0x80, 0x0a,
	0x53, 0xe1, 0x38, 0x91, 0x8c, 0x63, 0xd6, 0xe6, 0x07, 0xd0, 0x93, 0x9e, 0x50, 0xee, 0x2f, 0xea,
	0xf0, 0x0e, 0x34, 0x92, 0x48, 0xb1, 0x2e, 0xdf, 0x83, 0xae, 0xe3, 0xc7, 0xa9, 0x2f, 0x3c, 0xc9,
	0xac, 0xdd, 0x95, 0x44, 0x42, 0xab, 0xc0, 0x67, 0xb0, 0xfd, 0x53, 0x6f, 0x48, 0x3d, 0x11, 0xb2,
	0xff, 0xc3, 0xdb, 0xc5, 0x0a, 0xc9, 0x72, 0x85, 0x64, 0xb3, 0x42, 0xfa, 0x62, 0x90, 0xbe, 0x19,
	0xa4, 0x1f, 0x06, 0xe9, 0xc2, 0x20, 0xfd, 0x34, 0x48, 0xbf, 0x0c, 0x92, 0x8d, 0x41, 0xfa, 0xba,
	0x46, 0xb2, 0x58, 0x23, 0x59, 0xae, 0x91, 0xdc, 0x9f, 0xd4, 0x25, 0xf2, 0xc2, 0x1e, 0x97, 0xb9,
	0xfd, 0xb7, 0xdb, 0xa4, 0xbd, 0xab, 0x75, 0xfe, 0x1d, 0x00, 0x00, 0xff, 0xff, 0xa4, 0x0e, 0x39,
	0x92, 0x50, 0x01, 0x00, 0x00,
}

func (x ValueType) String() string {
	s, ok := ValueType_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}