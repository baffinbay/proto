// Code generated by protoc-gen-go. DO NOT EDIT.
// source: auth.proto

package auth

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

type SshCertificateRequest struct {
	PublicKey            string   `protobuf:"bytes,1,opt,name=public_key,json=publicKey" json:"public_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SshCertificateRequest) Reset()         { *m = SshCertificateRequest{} }
func (m *SshCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*SshCertificateRequest) ProtoMessage()    {}
func (*SshCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d75b026a43dbb6cb, []int{0}
}
func (m *SshCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshCertificateRequest.Unmarshal(m, b)
}
func (m *SshCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshCertificateRequest.Marshal(b, m, deterministic)
}
func (dst *SshCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshCertificateRequest.Merge(dst, src)
}
func (m *SshCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_SshCertificateRequest.Size(m)
}
func (m *SshCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SshCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SshCertificateRequest proto.InternalMessageInfo

func (m *SshCertificateRequest) GetPublicKey() string {
	if m != nil {
		return m.PublicKey
	}
	return ""
}

type SshCertificate struct {
	Certificate          string   `protobuf:"bytes,1,opt,name=certificate" json:"certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SshCertificate) Reset()         { *m = SshCertificate{} }
func (m *SshCertificate) String() string { return proto.CompactTextString(m) }
func (*SshCertificate) ProtoMessage()    {}
func (*SshCertificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d75b026a43dbb6cb, []int{1}
}
func (m *SshCertificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SshCertificate.Unmarshal(m, b)
}
func (m *SshCertificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SshCertificate.Marshal(b, m, deterministic)
}
func (dst *SshCertificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SshCertificate.Merge(dst, src)
}
func (m *SshCertificate) XXX_Size() int {
	return xxx_messageInfo_SshCertificate.Size(m)
}
func (m *SshCertificate) XXX_DiscardUnknown() {
	xxx_messageInfo_SshCertificate.DiscardUnknown(m)
}

var xxx_messageInfo_SshCertificate proto.InternalMessageInfo

func (m *SshCertificate) GetCertificate() string {
	if m != nil {
		return m.Certificate
	}
	return ""
}

type ClientValidation struct {
	// To make it harder to simply send a URL to another user and lure them into
	// logging in with the goal of stealing their session, we do a validation
	// by querying localhost:1215. If we get back the value presented in this
	// field the request is considered to be genuine.
	Ident                string   `protobuf:"bytes,1,opt,name=ident" json:"ident,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ClientValidation) Reset()         { *m = ClientValidation{} }
func (m *ClientValidation) String() string { return proto.CompactTextString(m) }
func (*ClientValidation) ProtoMessage()    {}
func (*ClientValidation) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d75b026a43dbb6cb, []int{2}
}
func (m *ClientValidation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClientValidation.Unmarshal(m, b)
}
func (m *ClientValidation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClientValidation.Marshal(b, m, deterministic)
}
func (dst *ClientValidation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientValidation.Merge(dst, src)
}
func (m *ClientValidation) XXX_Size() int {
	return xxx_messageInfo_ClientValidation.Size(m)
}
func (m *ClientValidation) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientValidation.DiscardUnknown(m)
}

var xxx_messageInfo_ClientValidation proto.InternalMessageInfo

func (m *ClientValidation) GetIdent() string {
	if m != nil {
		return m.Ident
	}
	return ""
}

type UserCredentialRequest struct {
	// Identification to verify on the requester
	ClientValidation *ClientValidation `protobuf:"bytes,1,opt,name=client_validation,json=clientValidation" json:"client_validation,omitempty"`
	// Request for a short-lived SSH certificate
	SshCertificateRequest *SshCertificateRequest `protobuf:"bytes,2,opt,name=ssh_certificate_request,json=sshCertificateRequest" json:"ssh_certificate_request,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}               `json:"-"`
	XXX_unrecognized      []byte                 `json:"-"`
	XXX_sizecache         int32                  `json:"-"`
}

func (m *UserCredentialRequest) Reset()         { *m = UserCredentialRequest{} }
func (m *UserCredentialRequest) String() string { return proto.CompactTextString(m) }
func (*UserCredentialRequest) ProtoMessage()    {}
func (*UserCredentialRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d75b026a43dbb6cb, []int{3}
}
func (m *UserCredentialRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserCredentialRequest.Unmarshal(m, b)
}
func (m *UserCredentialRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserCredentialRequest.Marshal(b, m, deterministic)
}
func (dst *UserCredentialRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserCredentialRequest.Merge(dst, src)
}
func (m *UserCredentialRequest) XXX_Size() int {
	return xxx_messageInfo_UserCredentialRequest.Size(m)
}
func (m *UserCredentialRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UserCredentialRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UserCredentialRequest proto.InternalMessageInfo

func (m *UserCredentialRequest) GetClientValidation() *ClientValidation {
	if m != nil {
		return m.ClientValidation
	}
	return nil
}

func (m *UserCredentialRequest) GetSshCertificateRequest() *SshCertificateRequest {
	if m != nil {
		return m.SshCertificateRequest
	}
	return nil
}

type UserAction struct {
	// Interactive URL needs to be visited and acted on
	Url                  string   `protobuf:"bytes,1,opt,name=url" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UserAction) Reset()         { *m = UserAction{} }
func (m *UserAction) String() string { return proto.CompactTextString(m) }
func (*UserAction) ProtoMessage()    {}
func (*UserAction) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d75b026a43dbb6cb, []int{4}
}
func (m *UserAction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UserAction.Unmarshal(m, b)
}
func (m *UserAction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UserAction.Marshal(b, m, deterministic)
}
func (dst *UserAction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UserAction.Merge(dst, src)
}
func (m *UserAction) XXX_Size() int {
	return xxx_messageInfo_UserAction.Size(m)
}
func (m *UserAction) XXX_DiscardUnknown() {
	xxx_messageInfo_UserAction.DiscardUnknown(m)
}

var xxx_messageInfo_UserAction proto.InternalMessageInfo

func (m *UserAction) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type CredentialResponse struct {
	// If set, the user needs to do something
	RequiredAction       *UserAction     `protobuf:"bytes,1,opt,name=required_action,json=requiredAction" json:"required_action,omitempty"`
	SshCertificate       *SshCertificate `protobuf:"bytes,2,opt,name=ssh_certificate,json=sshCertificate" json:"ssh_certificate,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CredentialResponse) Reset()         { *m = CredentialResponse{} }
func (m *CredentialResponse) String() string { return proto.CompactTextString(m) }
func (*CredentialResponse) ProtoMessage()    {}
func (*CredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_auth_d75b026a43dbb6cb, []int{5}
}
func (m *CredentialResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CredentialResponse.Unmarshal(m, b)
}
func (m *CredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CredentialResponse.Marshal(b, m, deterministic)
}
func (dst *CredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialResponse.Merge(dst, src)
}
func (m *CredentialResponse) XXX_Size() int {
	return xxx_messageInfo_CredentialResponse.Size(m)
}
func (m *CredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialResponse proto.InternalMessageInfo

func (m *CredentialResponse) GetRequiredAction() *UserAction {
	if m != nil {
		return m.RequiredAction
	}
	return nil
}

func (m *CredentialResponse) GetSshCertificate() *SshCertificate {
	if m != nil {
		return m.SshCertificate
	}
	return nil
}

func init() {
	proto.RegisterType((*SshCertificateRequest)(nil), "auth.SshCertificateRequest")
	proto.RegisterType((*SshCertificate)(nil), "auth.SshCertificate")
	proto.RegisterType((*ClientValidation)(nil), "auth.ClientValidation")
	proto.RegisterType((*UserCredentialRequest)(nil), "auth.UserCredentialRequest")
	proto.RegisterType((*UserAction)(nil), "auth.UserAction")
	proto.RegisterType((*CredentialResponse)(nil), "auth.CredentialResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for AuthenticationService service

type AuthenticationServiceClient interface {
	RequestUserCredential(ctx context.Context, in *UserCredentialRequest, opts ...grpc.CallOption) (AuthenticationService_RequestUserCredentialClient, error)
}

type authenticationServiceClient struct {
	cc *grpc.ClientConn
}

func NewAuthenticationServiceClient(cc *grpc.ClientConn) AuthenticationServiceClient {
	return &authenticationServiceClient{cc}
}

func (c *authenticationServiceClient) RequestUserCredential(ctx context.Context, in *UserCredentialRequest, opts ...grpc.CallOption) (AuthenticationService_RequestUserCredentialClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_AuthenticationService_serviceDesc.Streams[0], c.cc, "/auth.AuthenticationService/RequestUserCredential", opts...)
	if err != nil {
		return nil, err
	}
	x := &authenticationServiceRequestUserCredentialClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type AuthenticationService_RequestUserCredentialClient interface {
	Recv() (*CredentialResponse, error)
	grpc.ClientStream
}

type authenticationServiceRequestUserCredentialClient struct {
	grpc.ClientStream
}

func (x *authenticationServiceRequestUserCredentialClient) Recv() (*CredentialResponse, error) {
	m := new(CredentialResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for AuthenticationService service

type AuthenticationServiceServer interface {
	RequestUserCredential(*UserCredentialRequest, AuthenticationService_RequestUserCredentialServer) error
}

func RegisterAuthenticationServiceServer(s *grpc.Server, srv AuthenticationServiceServer) {
	s.RegisterService(&_AuthenticationService_serviceDesc, srv)
}

func _AuthenticationService_RequestUserCredential_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UserCredentialRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(AuthenticationServiceServer).RequestUserCredential(m, &authenticationServiceRequestUserCredentialServer{stream})
}

type AuthenticationService_RequestUserCredentialServer interface {
	Send(*CredentialResponse) error
	grpc.ServerStream
}

type authenticationServiceRequestUserCredentialServer struct {
	grpc.ServerStream
}

func (x *authenticationServiceRequestUserCredentialServer) Send(m *CredentialResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _AuthenticationService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "auth.AuthenticationService",
	HandlerType: (*AuthenticationServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "RequestUserCredential",
			Handler:       _AuthenticationService_RequestUserCredential_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "auth.proto",
}

func init() { proto.RegisterFile("auth.proto", fileDescriptor_auth_d75b026a43dbb6cb) }

var fileDescriptor_auth_d75b026a43dbb6cb = []byte{
	// 326 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0xdf, 0x6a, 0xc2, 0x30,
	0x14, 0xc6, 0xd7, 0xfd, 0x03, 0x8f, 0xa0, 0x5d, 0xb0, 0x9b, 0x4c, 0x36, 0x24, 0x57, 0x5e, 0xc9,
	0x70, 0x30, 0xd8, 0xc5, 0x2e, 0xa4, 0x97, 0xbb, 0x8b, 0x6c, 0xb7, 0x25, 0xc6, 0x33, 0x1a, 0x2c,
	0xad, 0x4b, 0x52, 0xc1, 0x97, 0xd8, 0xbb, 0xec, 0x0d, 0x47, 0xd2, 0x38, 0x6b, 0xed, 0x5d, 0xce,
	0x21, 0xe7, 0xfb, 0x7e, 0xdf, 0x49, 0x00, 0x78, 0x69, 0xd2, 0xe9, 0x46, 0x15, 0xa6, 0x20, 0x97,
	0xf6, 0x4c, 0x5f, 0x20, 0x5a, 0xe8, 0x34, 0x46, 0x65, 0xe4, 0x97, 0x14, 0xdc, 0x20, 0xc3, 0xef,
	0x12, 0xb5, 0x21, 0x0f, 0x00, 0x9b, 0x72, 0x99, 0x49, 0x91, 0xac, 0x71, 0x37, 0x0c, 0xc6, 0xc1,
	0xa4, 0xc3, 0x3a, 0x55, 0xe7, 0x1d, 0x77, 0x74, 0x06, 0xbd, 0xe3, 0x39, 0x32, 0x86, 0xae, 0x38,
	0x94, 0x7e, 0xa2, 0xde, 0xa2, 0x13, 0x08, 0xe3, 0x4c, 0x62, 0x6e, 0x3e, 0x79, 0x26, 0x57, 0xdc,
	0xc8, 0x22, 0x27, 0x03, 0xb8, 0x92, 0x2b, 0xcc, 0x8d, 0xbf, 0x5f, 0x15, 0xf4, 0x37, 0x80, 0xe8,
	0x43, 0xa3, 0x8a, 0x15, 0xda, 0x5a, 0xf2, 0x6c, 0x8f, 0x15, 0xc3, 0x8d, 0x70, 0x1a, 0xc9, 0xf6,
	0x5f, 0xc4, 0xcd, 0x76, 0x67, 0xb7, 0x53, 0x97, 0xae, 0x69, 0xc1, 0x42, 0xd1, 0x34, 0x5d, 0xc0,
	0x9d, 0xd6, 0x69, 0x52, 0x63, 0x4b, 0x54, 0xa5, 0x3f, 0x3c, 0x77, 0x52, 0xa3, 0x4a, 0xaa, 0x75,
	0x33, 0x2c, 0xd2, 0x6d, 0x6d, 0xfa, 0x08, 0x60, 0x91, 0xe7, 0xc2, 0x59, 0x84, 0x70, 0x51, 0xaa,
	0xcc, 0xa7, 0xb2, 0x47, 0xfa, 0x13, 0x00, 0xa9, 0xe7, 0xd1, 0x9b, 0x22, 0xd7, 0x48, 0x5e, 0xa1,
	0x6f, 0xbd, 0xa5, 0xc2, 0x55, 0xc2, 0x45, 0x2d, 0x4e, 0x58, 0x31, 0x1c, 0x34, 0x59, 0x6f, 0x7f,
	0xd1, 0x7b, 0xbc, 0x41, 0xbf, 0x11, 0xc3, 0xe3, 0x0f, 0x5a, 0xf1, 0x7b, 0xc7, 0xdc, 0xb3, 0x35,
	0x44, 0xf3, 0xd2, 0xa4, 0x96, 0x47, 0xb8, 0xbd, 0x2c, 0x50, 0x6d, 0xa5, 0x40, 0xc2, 0x20, 0xf2,
	0xa1, 0x8e, 0xdf, 0x80, 0x8c, 0x0e, 0x48, 0x27, 0x2f, 0x73, 0x3f, 0xf4, 0xeb, 0x3f, 0x89, 0x48,
	0xcf, 0x9e, 0x82, 0xe5, 0xb5, 0xfb, 0x74, 0xcf, 0x7f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x2a, 0x1a,
	0x48, 0x96, 0x82, 0x02, 0x00, 0x00,
}
