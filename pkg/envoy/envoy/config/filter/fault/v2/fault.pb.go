// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/fault/v2/fault.proto

/*
Package v2 is a generated protocol buffer package.

It is generated from these files:
	envoy/config/filter/fault/v2/fault.proto

It has these top-level messages:
	FaultDelay
*/
package v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import envoy_type "github.com/cilium/cilium/pkg/envoy/envoy/type"
import google_protobuf1 "github.com/golang/protobuf/ptypes/duration"
import _ "github.com/lyft/protoc-gen-validate/validate"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FaultDelay_FaultDelayType int32

const (
	// Fixed delay (step function).
	FaultDelay_FIXED FaultDelay_FaultDelayType = 0
)

var FaultDelay_FaultDelayType_name = map[int32]string{
	0: "FIXED",
}
var FaultDelay_FaultDelayType_value = map[string]int32{
	"FIXED": 0,
}

func (x FaultDelay_FaultDelayType) String() string {
	return proto.EnumName(FaultDelay_FaultDelayType_name, int32(x))
}
func (FaultDelay_FaultDelayType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0, 0} }

// Delay specification is used to inject latency into the
// HTTP/gRPC/Mongo/Redis operation or delay proxying of TCP connections.
type FaultDelay struct {
	// Delay type to use (fixed|exponential|..). Currently, only fixed delay (step function) is
	// supported.
	Type FaultDelay_FaultDelayType `protobuf:"varint,1,opt,name=type,enum=envoy.config.filter.fault.v2.FaultDelay_FaultDelayType" json:"type,omitempty"`
	// An integer between 0-100 indicating the percentage of operations/connection requests
	// on which the delay will be injected.
	//
	// .. attention::
	//
	//   Use of integer `percent` value is deprecated. Use fractional `percentage` field instead.
	Percent uint32 `protobuf:"varint,2,opt,name=percent" json:"percent,omitempty"`
	// Types that are valid to be assigned to FaultDelaySecifier:
	//	*FaultDelay_FixedDelay
	FaultDelaySecifier isFaultDelay_FaultDelaySecifier `protobuf_oneof:"fault_delay_secifier"`
	// The percentage of operations/connection requests on which the delay will be injected.
	Percentage *envoy_type.FractionalPercent `protobuf:"bytes,4,opt,name=percentage" json:"percentage,omitempty"`
}

func (m *FaultDelay) Reset()                    { *m = FaultDelay{} }
func (m *FaultDelay) String() string            { return proto.CompactTextString(m) }
func (*FaultDelay) ProtoMessage()               {}
func (*FaultDelay) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type isFaultDelay_FaultDelaySecifier interface {
	isFaultDelay_FaultDelaySecifier()
}

type FaultDelay_FixedDelay struct {
	FixedDelay *google_protobuf1.Duration `protobuf:"bytes,3,opt,name=fixed_delay,json=fixedDelay,oneof"`
}

func (*FaultDelay_FixedDelay) isFaultDelay_FaultDelaySecifier() {}

func (m *FaultDelay) GetFaultDelaySecifier() isFaultDelay_FaultDelaySecifier {
	if m != nil {
		return m.FaultDelaySecifier
	}
	return nil
}

func (m *FaultDelay) GetType() FaultDelay_FaultDelayType {
	if m != nil {
		return m.Type
	}
	return FaultDelay_FIXED
}

func (m *FaultDelay) GetPercent() uint32 {
	if m != nil {
		return m.Percent
	}
	return 0
}

func (m *FaultDelay) GetFixedDelay() *google_protobuf1.Duration {
	if x, ok := m.GetFaultDelaySecifier().(*FaultDelay_FixedDelay); ok {
		return x.FixedDelay
	}
	return nil
}

func (m *FaultDelay) GetPercentage() *envoy_type.FractionalPercent {
	if m != nil {
		return m.Percentage
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*FaultDelay) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _FaultDelay_OneofMarshaler, _FaultDelay_OneofUnmarshaler, _FaultDelay_OneofSizer, []interface{}{
		(*FaultDelay_FixedDelay)(nil),
	}
}

func _FaultDelay_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*FaultDelay)
	// fault_delay_secifier
	switch x := m.FaultDelaySecifier.(type) {
	case *FaultDelay_FixedDelay:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.FixedDelay); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("FaultDelay.FaultDelaySecifier has unexpected type %T", x)
	}
	return nil
}

func _FaultDelay_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*FaultDelay)
	switch tag {
	case 3: // fault_delay_secifier.fixed_delay
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(google_protobuf1.Duration)
		err := b.DecodeMessage(msg)
		m.FaultDelaySecifier = &FaultDelay_FixedDelay{msg}
		return true, err
	default:
		return false, nil
	}
}

func _FaultDelay_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*FaultDelay)
	// fault_delay_secifier
	switch x := m.FaultDelaySecifier.(type) {
	case *FaultDelay_FixedDelay:
		s := proto.Size(x.FixedDelay)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*FaultDelay)(nil), "envoy.config.filter.fault.v2.FaultDelay")
	proto.RegisterEnum("envoy.config.filter.fault.v2.FaultDelay_FaultDelayType", FaultDelay_FaultDelayType_name, FaultDelay_FaultDelayType_value)
}

func init() { proto.RegisterFile("envoy/config/filter/fault/v2/fault.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0x4f, 0x4a, 0xc3, 0x40,
	0x18, 0xc5, 0x3b, 0x69, 0xaa, 0x38, 0xc5, 0x52, 0x42, 0xc1, 0xb1, 0x5a, 0x2d, 0x05, 0x21, 0x74,
	0x31, 0x03, 0x71, 0xe1, 0xca, 0x4d, 0xa8, 0x45, 0xc1, 0x85, 0x04, 0x41, 0x71, 0x53, 0xa6, 0xc9,
	0x24, 0x0c, 0x84, 0x4c, 0x88, 0x69, 0x30, 0x5b, 0x4f, 0xe1, 0x19, 0x3c, 0x81, 0xb8, 0xea, 0x25,
	0x5c, 0xbb, 0xee, 0x2d, 0x64, 0xfe, 0x04, 0xeb, 0xc6, 0xdd, 0x4b, 0xbe, 0xf7, 0xbd, 0xef, 0xfd,
	0x06, 0xba, 0x2c, 0xab, 0x44, 0x4d, 0x42, 0x91, 0xc5, 0x3c, 0x21, 0x31, 0x4f, 0x4b, 0x56, 0x90,
	0x98, 0xae, 0xd2, 0x92, 0x54, 0x9e, 0x16, 0x38, 0x2f, 0x44, 0x29, 0x9c, 0x63, 0xe5, 0xc4, 0xda,
	0x89, 0xb5, 0x13, 0x6b, 0x43, 0xe5, 0x0d, 0x91, 0xce, 0x29, 0xeb, 0x9c, 0x91, 0x9c, 0x15, 0x21,
	0xcb, 0xcc, 0xde, 0xf0, 0x24, 0x11, 0x22, 0x49, 0x19, 0x51, 0x5f, 0xcb, 0x55, 0x4c, 0xa2, 0x55,
	0x41, 0x4b, 0x2e, 0x32, 0x33, 0x3f, 0xa8, 0x68, 0xca, 0x23, 0x5a, 0x32, 0xd2, 0x08, 0x33, 0x18,
	0x24, 0x22, 0x11, 0x4a, 0x12, 0xa9, 0xf4, 0xdf, 0xc9, 0x97, 0x05, 0xe1, 0x5c, 0x5e, 0x9d, 0xb1,
	0x94, 0xd6, 0xce, 0x03, 0xb4, 0xe5, 0x4d, 0x04, 0xc6, 0xc0, 0xed, 0x79, 0x17, 0xf8, 0xbf, 0x92,
	0xf8, 0x77, 0x6f, 0x4b, 0xde, 0xd7, 0x39, 0xf3, 0xe1, 0xe7, 0x66, 0xdd, 0xee, 0xbc, 0x02, 0xab,
	0x0f, 0x02, 0x15, 0xe8, 0x9c, 0xc1, 0x5d, 0xc3, 0x81, 0xac, 0x31, 0x70, 0xf7, 0xfd, 0xae, 0xb4,
	0xd8, 0x53, 0x0b, 0x45, 0x08, 0x04, 0xcd, 0xcc, 0xb9, 0x85, 0xdd, 0x98, 0xbf, 0xb0, 0x68, 0x11,
	0xc9, 0x2c, 0xd4, 0x1e, 0x03, 0xb7, 0xeb, 0x1d, 0x62, 0xcd, 0x8c, 0x1b, 0x66, 0x3c, 0x33, 0xcc,
	0x7e, 0xef, 0xed, 0xfb, 0x14, 0xa8, 0x63, 0xef, 0xc0, 0x9a, 0xb6, 0xae, 0x5b, 0x01, 0x54, 0xfb,
	0x9a, 0xe6, 0x12, 0x42, 0x13, 0x4c, 0x13, 0x86, 0x6c, 0x15, 0x36, 0x32, 0x4c, 0xb2, 0x15, 0x9e,
	0x17, 0x34, 0x94, 0x39, 0x34, 0xbd, 0xd3, 0xbe, 0x60, 0x6b, 0x61, 0x72, 0x04, 0x7b, 0x7f, 0xb9,
	0x9c, 0x3d, 0xd8, 0x99, 0xdf, 0x3c, 0x5e, 0xcd, 0xfa, 0x2d, 0x7f, 0x04, 0x07, 0xea, 0x21, 0x74,
	0xd3, 0xc5, 0x33, 0x0b, 0x79, 0xcc, 0x59, 0xe1, 0x74, 0x3e, 0x36, 0xeb, 0x36, 0xf0, 0xed, 0x27,
	0xab, 0xf2, 0x96, 0x3b, 0xaa, 0xf1, 0xf9, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0x76, 0x49,
	0x42, 0x17, 0x02, 0x00, 0x00,
}
