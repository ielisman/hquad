// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cert.proto

package cert

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type PKIName struct {
	Country              string   `protobuf:"bytes,1,opt,name=country,proto3" json:"country,omitempty"`
	Province             string   `protobuf:"bytes,2,opt,name=province,proto3" json:"province,omitempty"`
	City                 string   `protobuf:"bytes,3,opt,name=city,proto3" json:"city,omitempty"`
	Organization         string   `protobuf:"bytes,4,opt,name=organization,proto3" json:"organization,omitempty"`
	OrganizationUnit     string   `protobuf:"bytes,5,opt,name=organizationUnit,proto3" json:"organizationUnit,omitempty"`
	CommonName           string   `protobuf:"bytes,6,opt,name=commonName,proto3" json:"commonName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PKIName) Reset()         { *m = PKIName{} }
func (m *PKIName) String() string { return proto.CompactTextString(m) }
func (*PKIName) ProtoMessage()    {}
func (*PKIName) Descriptor() ([]byte, []int) {
	return fileDescriptor_a142e29cbef9b1cf, []int{0}
}

func (m *PKIName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PKIName.Unmarshal(m, b)
}
func (m *PKIName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PKIName.Marshal(b, m, deterministic)
}
func (m *PKIName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PKIName.Merge(m, src)
}
func (m *PKIName) XXX_Size() int {
	return xxx_messageInfo_PKIName.Size(m)
}
func (m *PKIName) XXX_DiscardUnknown() {
	xxx_messageInfo_PKIName.DiscardUnknown(m)
}

var xxx_messageInfo_PKIName proto.InternalMessageInfo

func (m *PKIName) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *PKIName) GetProvince() string {
	if m != nil {
		return m.Province
	}
	return ""
}

func (m *PKIName) GetCity() string {
	if m != nil {
		return m.City
	}
	return ""
}

func (m *PKIName) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func (m *PKIName) GetOrganizationUnit() string {
	if m != nil {
		return m.OrganizationUnit
	}
	return ""
}

func (m *PKIName) GetCommonName() string {
	if m != nil {
		return m.CommonName
	}
	return ""
}

type CreateServerCertificateRequest struct {
	Subject              *PKIName `protobuf:"bytes,1,opt,name=subject,proto3" json:"subject,omitempty"`
	DomainNames          []string `protobuf:"bytes,2,rep,name=domainNames,proto3" json:"domainNames,omitempty"`
	IPAddrs              []string `protobuf:"bytes,3,rep,name=IPAddrs,proto3" json:"IPAddrs,omitempty"`
	ExpireAfterYears     int32    `protobuf:"varint,4,opt,name=expireAfterYears,proto3" json:"expireAfterYears,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServerCertificateRequest) Reset()         { *m = CreateServerCertificateRequest{} }
func (m *CreateServerCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*CreateServerCertificateRequest) ProtoMessage()    {}
func (*CreateServerCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a142e29cbef9b1cf, []int{1}
}

func (m *CreateServerCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServerCertificateRequest.Unmarshal(m, b)
}
func (m *CreateServerCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServerCertificateRequest.Marshal(b, m, deterministic)
}
func (m *CreateServerCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServerCertificateRequest.Merge(m, src)
}
func (m *CreateServerCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_CreateServerCertificateRequest.Size(m)
}
func (m *CreateServerCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServerCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServerCertificateRequest proto.InternalMessageInfo

func (m *CreateServerCertificateRequest) GetSubject() *PKIName {
	if m != nil {
		return m.Subject
	}
	return nil
}

func (m *CreateServerCertificateRequest) GetDomainNames() []string {
	if m != nil {
		return m.DomainNames
	}
	return nil
}

func (m *CreateServerCertificateRequest) GetIPAddrs() []string {
	if m != nil {
		return m.IPAddrs
	}
	return nil
}

func (m *CreateServerCertificateRequest) GetExpireAfterYears() int32 {
	if m != nil {
		return m.ExpireAfterYears
	}
	return 0
}

type CreateServerCertificateResponse struct {
	ServerCert           string   `protobuf:"bytes,1,opt,name=serverCert,proto3" json:"serverCert,omitempty"`
	ServerKey            string   `protobuf:"bytes,2,opt,name=serverKey,proto3" json:"serverKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateServerCertificateResponse) Reset()         { *m = CreateServerCertificateResponse{} }
func (m *CreateServerCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*CreateServerCertificateResponse) ProtoMessage()    {}
func (*CreateServerCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a142e29cbef9b1cf, []int{2}
}

func (m *CreateServerCertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateServerCertificateResponse.Unmarshal(m, b)
}
func (m *CreateServerCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateServerCertificateResponse.Marshal(b, m, deterministic)
}
func (m *CreateServerCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateServerCertificateResponse.Merge(m, src)
}
func (m *CreateServerCertificateResponse) XXX_Size() int {
	return xxx_messageInfo_CreateServerCertificateResponse.Size(m)
}
func (m *CreateServerCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateServerCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateServerCertificateResponse proto.InternalMessageInfo

func (m *CreateServerCertificateResponse) GetServerCert() string {
	if m != nil {
		return m.ServerCert
	}
	return ""
}

func (m *CreateServerCertificateResponse) GetServerKey() string {
	if m != nil {
		return m.ServerKey
	}
	return ""
}

type SetServerCertificateRequest struct {
	ServerCert           string   `protobuf:"bytes,1,opt,name=serverCert,proto3" json:"serverCert,omitempty"`
	ServerKey            string   `protobuf:"bytes,2,opt,name=serverKey,proto3" json:"serverKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetServerCertificateRequest) Reset()         { *m = SetServerCertificateRequest{} }
func (m *SetServerCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*SetServerCertificateRequest) ProtoMessage()    {}
func (*SetServerCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a142e29cbef9b1cf, []int{3}
}

func (m *SetServerCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetServerCertificateRequest.Unmarshal(m, b)
}
func (m *SetServerCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetServerCertificateRequest.Marshal(b, m, deterministic)
}
func (m *SetServerCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetServerCertificateRequest.Merge(m, src)
}
func (m *SetServerCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_SetServerCertificateRequest.Size(m)
}
func (m *SetServerCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetServerCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetServerCertificateRequest proto.InternalMessageInfo

func (m *SetServerCertificateRequest) GetServerCert() string {
	if m != nil {
		return m.ServerCert
	}
	return ""
}

func (m *SetServerCertificateRequest) GetServerKey() string {
	if m != nil {
		return m.ServerKey
	}
	return ""
}

type SetServerCertificateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetServerCertificateResponse) Reset()         { *m = SetServerCertificateResponse{} }
func (m *SetServerCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*SetServerCertificateResponse) ProtoMessage()    {}
func (*SetServerCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a142e29cbef9b1cf, []int{4}
}

func (m *SetServerCertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetServerCertificateResponse.Unmarshal(m, b)
}
func (m *SetServerCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetServerCertificateResponse.Marshal(b, m, deterministic)
}
func (m *SetServerCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetServerCertificateResponse.Merge(m, src)
}
func (m *SetServerCertificateResponse) XXX_Size() int {
	return xxx_messageInfo_SetServerCertificateResponse.Size(m)
}
func (m *SetServerCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetServerCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetServerCertificateResponse proto.InternalMessageInfo

type SetGatewayCertificateRequest struct {
	GatewayCert          string   `protobuf:"bytes,1,opt,name=gatewayCert,proto3" json:"gatewayCert,omitempty"`
	GatewayKey           string   `protobuf:"bytes,2,opt,name=gatewayKey,proto3" json:"gatewayKey,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGatewayCertificateRequest) Reset()         { *m = SetGatewayCertificateRequest{} }
func (m *SetGatewayCertificateRequest) String() string { return proto.CompactTextString(m) }
func (*SetGatewayCertificateRequest) ProtoMessage()    {}
func (*SetGatewayCertificateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_a142e29cbef9b1cf, []int{5}
}

func (m *SetGatewayCertificateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGatewayCertificateRequest.Unmarshal(m, b)
}
func (m *SetGatewayCertificateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGatewayCertificateRequest.Marshal(b, m, deterministic)
}
func (m *SetGatewayCertificateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGatewayCertificateRequest.Merge(m, src)
}
func (m *SetGatewayCertificateRequest) XXX_Size() int {
	return xxx_messageInfo_SetGatewayCertificateRequest.Size(m)
}
func (m *SetGatewayCertificateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGatewayCertificateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetGatewayCertificateRequest proto.InternalMessageInfo

func (m *SetGatewayCertificateRequest) GetGatewayCert() string {
	if m != nil {
		return m.GatewayCert
	}
	return ""
}

func (m *SetGatewayCertificateRequest) GetGatewayKey() string {
	if m != nil {
		return m.GatewayKey
	}
	return ""
}

type SetGatewayCertificateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGatewayCertificateResponse) Reset()         { *m = SetGatewayCertificateResponse{} }
func (m *SetGatewayCertificateResponse) String() string { return proto.CompactTextString(m) }
func (*SetGatewayCertificateResponse) ProtoMessage()    {}
func (*SetGatewayCertificateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a142e29cbef9b1cf, []int{6}
}

func (m *SetGatewayCertificateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGatewayCertificateResponse.Unmarshal(m, b)
}
func (m *SetGatewayCertificateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGatewayCertificateResponse.Marshal(b, m, deterministic)
}
func (m *SetGatewayCertificateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGatewayCertificateResponse.Merge(m, src)
}
func (m *SetGatewayCertificateResponse) XXX_Size() int {
	return xxx_messageInfo_SetGatewayCertificateResponse.Size(m)
}
func (m *SetGatewayCertificateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGatewayCertificateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetGatewayCertificateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*PKIName)(nil), "cert.PKIName")
	proto.RegisterType((*CreateServerCertificateRequest)(nil), "cert.CreateServerCertificateRequest")
	proto.RegisterType((*CreateServerCertificateResponse)(nil), "cert.CreateServerCertificateResponse")
	proto.RegisterType((*SetServerCertificateRequest)(nil), "cert.SetServerCertificateRequest")
	proto.RegisterType((*SetServerCertificateResponse)(nil), "cert.SetServerCertificateResponse")
	proto.RegisterType((*SetGatewayCertificateRequest)(nil), "cert.SetGatewayCertificateRequest")
	proto.RegisterType((*SetGatewayCertificateResponse)(nil), "cert.SetGatewayCertificateResponse")
}

func init() { proto.RegisterFile("cert.proto", fileDescriptor_a142e29cbef9b1cf) }

var fileDescriptor_a142e29cbef9b1cf = []byte{
	// 456 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x54, 0x51, 0x6f, 0xd3, 0x30,
	0x10, 0x56, 0xda, 0x6e, 0xa5, 0x57, 0x90, 0x90, 0x35, 0x34, 0xab, 0x8c, 0xae, 0x18, 0x10, 0x13,
	0x0f, 0x99, 0x18, 0xbf, 0x60, 0xec, 0x01, 0x4d, 0x93, 0x50, 0x95, 0x89, 0x07, 0x40, 0x88, 0xb9,
	0xce, 0x75, 0x32, 0x28, 0x71, 0xb0, 0x9d, 0x41, 0xf8, 0x53, 0x88, 0xdf, 0xc0, 0x1f, 0x43, 0x76,
	0xd2, 0x25, 0x68, 0x49, 0x78, 0xd8, 0x5b, 0xee, 0xfb, 0xce, 0x77, 0xf7, 0x7d, 0x77, 0x0a, 0x80,
	0x40, 0x6d, 0xc3, 0x4c, 0x2b, 0xab, 0xc8, 0xc8, 0x7d, 0xb3, 0x3f, 0x01, 0x8c, 0x97, 0x67, 0xa7,
	0x6f, 0x79, 0x82, 0x84, 0xc2, 0x58, 0xa8, 0x3c, 0xb5, 0xba, 0xa0, 0xc1, 0x22, 0x38, 0x98, 0x44,
	0x9b, 0x90, 0xcc, 0xe0, 0x4e, 0xa6, 0xd5, 0x95, 0x4c, 0x05, 0xd2, 0x81, 0xa7, 0xae, 0x63, 0x42,
	0x60, 0x24, 0xa4, 0x2d, 0xe8, 0xd0, 0xe3, 0xfe, 0x9b, 0x30, 0xb8, 0xab, 0xf4, 0x25, 0x4f, 0xe5,
	0x4f, 0x6e, 0xa5, 0x4a, 0xe9, 0xc8, 0x73, 0xff, 0x60, 0xe4, 0x05, 0xdc, 0x6f, 0xc6, 0xef, 0x52,
	0x69, 0xe9, 0x96, 0xcf, 0xbb, 0x81, 0x93, 0x39, 0x80, 0x50, 0x49, 0xa2, 0x52, 0x37, 0x27, 0xdd,
	0xf6, 0x59, 0x0d, 0x84, 0xfd, 0x0e, 0x60, 0x7e, 0xa2, 0x91, 0x5b, 0x3c, 0x47, 0x7d, 0x85, 0xfa,
	0x04, 0xb5, 0x95, 0x6b, 0x29, 0xb8, 0xc5, 0x08, 0xbf, 0xe5, 0x68, 0x2c, 0x79, 0x0e, 0x63, 0x93,
	0xaf, 0xbe, 0xa0, 0xb0, 0x5e, 0xdc, 0xf4, 0xe8, 0x5e, 0xe8, 0xcd, 0xa8, 0xc4, 0x47, 0x1b, 0x96,
	0x2c, 0x60, 0x1a, 0xab, 0x84, 0x4b, 0x5f, 0xd9, 0xd0, 0xc1, 0x62, 0x78, 0x30, 0x89, 0x9a, 0x90,
	0xf3, 0xe9, 0x74, 0x79, 0x1c, 0xc7, 0xda, 0xd0, 0xa1, 0x67, 0x37, 0xa1, 0xd3, 0x84, 0x3f, 0x32,
	0xa9, 0xf1, 0x78, 0x6d, 0x51, 0xbf, 0x47, 0xae, 0x8d, 0xd7, 0xbe, 0x15, 0xdd, 0xc0, 0xd9, 0x67,
	0xd8, 0xef, 0x1c, 0xd9, 0x64, 0x2a, 0x35, 0xe8, 0x64, 0x9b, 0x6b, 0xb2, 0xda, 0x49, 0x03, 0x21,
	0x7b, 0x30, 0x29, 0xa3, 0x33, 0x2c, 0xaa, 0xbd, 0xd4, 0x00, 0xfb, 0x08, 0x0f, 0xcf, 0xd1, 0x76,
	0x1a, 0x72, 0xbb, 0xe2, 0x73, 0xd8, 0x6b, 0x2f, 0x5e, 0x8e, 0xce, 0x2e, 0x3c, 0xff, 0x86, 0x5b,
	0xfc, 0xce, 0x8b, 0x96, 0xee, 0x0b, 0x98, 0x5e, 0xd6, 0x64, 0xd5, 0xbe, 0x09, 0xb9, 0xf9, 0xaa,
	0xb0, 0x1e, 0xa0, 0x81, 0xb0, 0x7d, 0x78, 0xd4, 0xd1, 0xa1, 0x1c, 0xe1, 0xe8, 0xd7, 0x00, 0x46,
	0xbe, 0xd2, 0x1a, 0x76, 0x3b, 0x9c, 0x26, 0x4f, 0xcb, 0x23, 0xe8, 0xbf, 0x9d, 0xd9, 0xb3, 0xff,
	0x64, 0x55, 0xeb, 0xfa, 0x04, 0x3b, 0x6d, 0x9e, 0x90, 0xc7, 0xe5, 0xf3, 0x9e, 0x65, 0xcc, 0x58,
	0x5f, 0x4a, 0x55, 0xfe, 0x02, 0x1e, 0xb4, 0x0a, 0x26, 0xf5, 0xe3, 0x4e, 0xbf, 0x67, 0x4f, 0x7a,
	0x73, 0xca, 0x0e, 0xaf, 0x5f, 0xc2, 0xae, 0x50, 0x49, 0x68, 0xf2, 0x4c, 0xa3, 0x3b, 0x77, 0x11,
	0x9a, 0xf8, 0xab, 0x7f, 0xb8, 0x0c, 0x3e, 0xec, 0xac, 0xa4, 0x32, 0x96, 0xeb, 0x43, 0x77, 0x04,
	0x52, 0xe0, 0xa1, 0xc3, 0x57, 0xdb, 0xfe, 0x67, 0xf2, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x22, 0x66, 0x9a, 0x07, 0x5a, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CertClient is the client API for Cert service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CertClient interface {
	// Server Certificate
	CreateServerCertificate(ctx context.Context, in *CreateServerCertificateRequest, opts ...grpc.CallOption) (*CreateServerCertificateResponse, error)
	SetServerCertificate(ctx context.Context, in *SetServerCertificateRequest, opts ...grpc.CallOption) (*SetServerCertificateResponse, error)
	// Set Gatweay Certificate: for Device Gateway only
	SetGatewayCertificate(ctx context.Context, in *SetGatewayCertificateRequest, opts ...grpc.CallOption) (*SetGatewayCertificateResponse, error)
}

type certClient struct {
	cc *grpc.ClientConn
}

func NewCertClient(cc *grpc.ClientConn) CertClient {
	return &certClient{cc}
}

func (c *certClient) CreateServerCertificate(ctx context.Context, in *CreateServerCertificateRequest, opts ...grpc.CallOption) (*CreateServerCertificateResponse, error) {
	out := new(CreateServerCertificateResponse)
	err := c.cc.Invoke(ctx, "/cert.Cert/CreateServerCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certClient) SetServerCertificate(ctx context.Context, in *SetServerCertificateRequest, opts ...grpc.CallOption) (*SetServerCertificateResponse, error) {
	out := new(SetServerCertificateResponse)
	err := c.cc.Invoke(ctx, "/cert.Cert/SetServerCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *certClient) SetGatewayCertificate(ctx context.Context, in *SetGatewayCertificateRequest, opts ...grpc.CallOption) (*SetGatewayCertificateResponse, error) {
	out := new(SetGatewayCertificateResponse)
	err := c.cc.Invoke(ctx, "/cert.Cert/SetGatewayCertificate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CertServer is the server API for Cert service.
type CertServer interface {
	// Server Certificate
	CreateServerCertificate(context.Context, *CreateServerCertificateRequest) (*CreateServerCertificateResponse, error)
	SetServerCertificate(context.Context, *SetServerCertificateRequest) (*SetServerCertificateResponse, error)
	// Set Gatweay Certificate: for Device Gateway only
	SetGatewayCertificate(context.Context, *SetGatewayCertificateRequest) (*SetGatewayCertificateResponse, error)
}

func RegisterCertServer(s *grpc.Server, srv CertServer) {
	s.RegisterService(&_Cert_serviceDesc, srv)
}

func _Cert_CreateServerCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).CreateServerCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cert.Cert/CreateServerCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).CreateServerCertificate(ctx, req.(*CreateServerCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cert_SetServerCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetServerCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).SetServerCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cert.Cert/SetServerCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).SetServerCertificate(ctx, req.(*SetServerCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cert_SetGatewayCertificate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGatewayCertificateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CertServer).SetGatewayCertificate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cert.Cert/SetGatewayCertificate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CertServer).SetGatewayCertificate(ctx, req.(*SetGatewayCertificateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cert_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cert.Cert",
	HandlerType: (*CertServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateServerCertificate",
			Handler:    _Cert_CreateServerCertificate_Handler,
		},
		{
			MethodName: "SetServerCertificate",
			Handler:    _Cert_SetServerCertificate_Handler,
		},
		{
			MethodName: "SetGatewayCertificate",
			Handler:    _Cert_SetGatewayCertificate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cert.proto",
}
