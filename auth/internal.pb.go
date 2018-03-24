// Code generated by protoc-gen-go. DO NOT EDIT.
// source: internal.proto

package auth

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Passed inside the authserver when the user completes a challenge
type WebChallengeResponse struct {
	// The challenge ID
	Challenge string `protobuf:"bytes,1,opt,name=challenge" json:"challenge,omitempty"`
	// User responded to the challenge as this user
	User string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	// Source of the user
	Ip   string `protobuf:"bytes,3,opt,name=ip" json:"ip,omitempty"`
	Port uint32 `protobuf:"varint,4,opt,name=port" json:"port,omitempty"`
}

func (m *WebChallengeResponse) Reset()                    { *m = WebChallengeResponse{} }
func (m *WebChallengeResponse) String() string            { return proto.CompactTextString(m) }
func (*WebChallengeResponse) ProtoMessage()               {}
func (*WebChallengeResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *WebChallengeResponse) GetChallenge() string {
	if m != nil {
		return m.Challenge
	}
	return ""
}

func (m *WebChallengeResponse) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *WebChallengeResponse) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *WebChallengeResponse) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func init() {
	proto.RegisterType((*WebChallengeResponse)(nil), "auth.WebChallengeResponse")
}

func init() { proto.RegisterFile("internal.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 132 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcb, 0xcc, 0x2b, 0x49,
	0x2d, 0xca, 0x4b, 0xcc, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x49, 0x2c, 0x2d, 0xc9,
	0x50, 0xca, 0xe1, 0x12, 0x09, 0x4f, 0x4d, 0x72, 0xce, 0x48, 0xcc, 0xc9, 0x49, 0xcd, 0x4b, 0x4f,
	0x0d, 0x4a, 0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe1, 0xe2, 0x4c, 0x86, 0x09, 0x4a,
	0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x21, 0x04, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x8b, 0x53, 0x8b,
	0x24, 0x98, 0xc0, 0x12, 0x60, 0xb6, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x81, 0x04, 0x33, 0x58, 0x84,
	0x29, 0xb3, 0x00, 0xa4, 0xa6, 0x20, 0xbf, 0xa8, 0x44, 0x82, 0x45, 0x81, 0x51, 0x83, 0x37, 0x08,
	0xcc, 0x4e, 0x62, 0x03, 0x5b, 0x6d, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x10, 0xf2, 0x39, 0xa4,
	0x8c, 0x00, 0x00, 0x00,
}