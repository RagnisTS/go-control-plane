// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/admin/v3/certs.proto

package envoy_admin_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Certificates struct {
	Certificates         []*Certificate `protobuf:"bytes,1,rep,name=certificates,proto3" json:"certificates,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *Certificates) Reset()         { *m = Certificates{} }
func (m *Certificates) String() string { return proto.CompactTextString(m) }
func (*Certificates) ProtoMessage()    {}
func (*Certificates) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{0}
}

func (m *Certificates) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificates.Unmarshal(m, b)
}
func (m *Certificates) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificates.Marshal(b, m, deterministic)
}
func (m *Certificates) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificates.Merge(m, src)
}
func (m *Certificates) XXX_Size() int {
	return xxx_messageInfo_Certificates.Size(m)
}
func (m *Certificates) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificates.DiscardUnknown(m)
}

var xxx_messageInfo_Certificates proto.InternalMessageInfo

func (m *Certificates) GetCertificates() []*Certificate {
	if m != nil {
		return m.Certificates
	}
	return nil
}

type Certificate struct {
	CaCert               []*CertificateDetails `protobuf:"bytes,1,rep,name=ca_cert,json=caCert,proto3" json:"ca_cert,omitempty"`
	CertChain            []*CertificateDetails `protobuf:"bytes,2,rep,name=cert_chain,json=certChain,proto3" json:"cert_chain,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *Certificate) Reset()         { *m = Certificate{} }
func (m *Certificate) String() string { return proto.CompactTextString(m) }
func (*Certificate) ProtoMessage()    {}
func (*Certificate) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{1}
}

func (m *Certificate) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Certificate.Unmarshal(m, b)
}
func (m *Certificate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Certificate.Marshal(b, m, deterministic)
}
func (m *Certificate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Certificate.Merge(m, src)
}
func (m *Certificate) XXX_Size() int {
	return xxx_messageInfo_Certificate.Size(m)
}
func (m *Certificate) XXX_DiscardUnknown() {
	xxx_messageInfo_Certificate.DiscardUnknown(m)
}

var xxx_messageInfo_Certificate proto.InternalMessageInfo

func (m *Certificate) GetCaCert() []*CertificateDetails {
	if m != nil {
		return m.CaCert
	}
	return nil
}

func (m *Certificate) GetCertChain() []*CertificateDetails {
	if m != nil {
		return m.CertChain
	}
	return nil
}

type CertificateDetails struct {
	Path                 string                  `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	SerialNumber         string                  `protobuf:"bytes,2,opt,name=serial_number,json=serialNumber,proto3" json:"serial_number,omitempty"`
	SubjectAltNames      []*SubjectAlternateName `protobuf:"bytes,3,rep,name=subject_alt_names,json=subjectAltNames,proto3" json:"subject_alt_names,omitempty"`
	DaysUntilExpiration  uint64                  `protobuf:"varint,4,opt,name=days_until_expiration,json=daysUntilExpiration,proto3" json:"days_until_expiration,omitempty"`
	ValidFrom            *timestamp.Timestamp    `protobuf:"bytes,5,opt,name=valid_from,json=validFrom,proto3" json:"valid_from,omitempty"`
	ExpirationTime       *timestamp.Timestamp    `protobuf:"bytes,6,opt,name=expiration_time,json=expirationTime,proto3" json:"expiration_time,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *CertificateDetails) Reset()         { *m = CertificateDetails{} }
func (m *CertificateDetails) String() string { return proto.CompactTextString(m) }
func (*CertificateDetails) ProtoMessage()    {}
func (*CertificateDetails) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{2}
}

func (m *CertificateDetails) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CertificateDetails.Unmarshal(m, b)
}
func (m *CertificateDetails) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CertificateDetails.Marshal(b, m, deterministic)
}
func (m *CertificateDetails) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CertificateDetails.Merge(m, src)
}
func (m *CertificateDetails) XXX_Size() int {
	return xxx_messageInfo_CertificateDetails.Size(m)
}
func (m *CertificateDetails) XXX_DiscardUnknown() {
	xxx_messageInfo_CertificateDetails.DiscardUnknown(m)
}

var xxx_messageInfo_CertificateDetails proto.InternalMessageInfo

func (m *CertificateDetails) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CertificateDetails) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *CertificateDetails) GetSubjectAltNames() []*SubjectAlternateName {
	if m != nil {
		return m.SubjectAltNames
	}
	return nil
}

func (m *CertificateDetails) GetDaysUntilExpiration() uint64 {
	if m != nil {
		return m.DaysUntilExpiration
	}
	return 0
}

func (m *CertificateDetails) GetValidFrom() *timestamp.Timestamp {
	if m != nil {
		return m.ValidFrom
	}
	return nil
}

func (m *CertificateDetails) GetExpirationTime() *timestamp.Timestamp {
	if m != nil {
		return m.ExpirationTime
	}
	return nil
}

type SubjectAlternateName struct {
	// Types that are valid to be assigned to Name:
	//	*SubjectAlternateName_Dns
	//	*SubjectAlternateName_Uri
	//	*SubjectAlternateName_IpAddress
	Name                 isSubjectAlternateName_Name `protobuf_oneof:"name"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *SubjectAlternateName) Reset()         { *m = SubjectAlternateName{} }
func (m *SubjectAlternateName) String() string { return proto.CompactTextString(m) }
func (*SubjectAlternateName) ProtoMessage()    {}
func (*SubjectAlternateName) Descriptor() ([]byte, []int) {
	return fileDescriptor_51228a64c985b2dc, []int{3}
}

func (m *SubjectAlternateName) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SubjectAlternateName.Unmarshal(m, b)
}
func (m *SubjectAlternateName) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SubjectAlternateName.Marshal(b, m, deterministic)
}
func (m *SubjectAlternateName) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SubjectAlternateName.Merge(m, src)
}
func (m *SubjectAlternateName) XXX_Size() int {
	return xxx_messageInfo_SubjectAlternateName.Size(m)
}
func (m *SubjectAlternateName) XXX_DiscardUnknown() {
	xxx_messageInfo_SubjectAlternateName.DiscardUnknown(m)
}

var xxx_messageInfo_SubjectAlternateName proto.InternalMessageInfo

type isSubjectAlternateName_Name interface {
	isSubjectAlternateName_Name()
}

type SubjectAlternateName_Dns struct {
	Dns string `protobuf:"bytes,1,opt,name=dns,proto3,oneof"`
}

type SubjectAlternateName_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

type SubjectAlternateName_IpAddress struct {
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3,oneof"`
}

func (*SubjectAlternateName_Dns) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_Uri) isSubjectAlternateName_Name() {}

func (*SubjectAlternateName_IpAddress) isSubjectAlternateName_Name() {}

func (m *SubjectAlternateName) GetName() isSubjectAlternateName_Name {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *SubjectAlternateName) GetDns() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Dns); ok {
		return x.Dns
	}
	return ""
}

func (m *SubjectAlternateName) GetUri() string {
	if x, ok := m.GetName().(*SubjectAlternateName_Uri); ok {
		return x.Uri
	}
	return ""
}

func (m *SubjectAlternateName) GetIpAddress() string {
	if x, ok := m.GetName().(*SubjectAlternateName_IpAddress); ok {
		return x.IpAddress
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*SubjectAlternateName) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*SubjectAlternateName_Dns)(nil),
		(*SubjectAlternateName_Uri)(nil),
		(*SubjectAlternateName_IpAddress)(nil),
	}
}

func init() {
	proto.RegisterType((*Certificates)(nil), "envoy.admin.v3.Certificates")
	proto.RegisterType((*Certificate)(nil), "envoy.admin.v3.Certificate")
	proto.RegisterType((*CertificateDetails)(nil), "envoy.admin.v3.CertificateDetails")
	proto.RegisterType((*SubjectAlternateName)(nil), "envoy.admin.v3.SubjectAlternateName")
}

func init() { proto.RegisterFile("envoy/admin/v3/certs.proto", fileDescriptor_51228a64c985b2dc) }

var fileDescriptor_51228a64c985b2dc = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x52, 0xdd, 0xae, 0xd2, 0x40,
	0x10, 0xb6, 0x07, 0xc4, 0x30, 0xe0, 0x39, 0x71, 0xd5, 0xa4, 0x41, 0x23, 0x9c, 0x6a, 0x8e, 0x8d,
	0xd1, 0x36, 0x81, 0x2b, 0xf1, 0xc2, 0x00, 0x6a, 0xbc, 0x3a, 0x21, 0x55, 0xaf, 0x9b, 0xa1, 0x5d,
	0x60, 0x4d, 0xbb, 0xdb, 0xec, 0x6e, 0x09, 0xdc, 0x7a, 0xe5, 0x23, 0x18, 0xdf, 0xc3, 0x47, 0xf0,
	0xbd, 0xcc, 0x76, 0xcf, 0x0f, 0x1c, 0x89, 0x78, 0xb7, 0xf3, 0x7d, 0xf3, 0xed, 0xcc, 0x7c, 0x33,
	0xd0, 0xa1, 0x7c, 0x25, 0x36, 0x21, 0xa6, 0x39, 0xe3, 0xe1, 0x6a, 0x10, 0x26, 0x54, 0x6a, 0x15,
	0x14, 0x52, 0x68, 0x41, 0x8e, 0x2b, 0x2e, 0xa8, 0xb8, 0x60, 0x35, 0xe8, 0x74, 0x17, 0x42, 0x2c,
	0x32, 0x1a, 0x56, 0xec, 0xac, 0x9c, 0x87, 0x9a, 0xe5, 0x54, 0x69, 0xcc, 0x0b, 0x2b, 0xe8, 0x9c,
	0x96, 0x69, 0x81, 0x21, 0x72, 0x2e, 0x34, 0x6a, 0x26, 0xb8, 0x0a, 0x57, 0x54, 0x2a, 0x26, 0x38,
	0xe3, 0x0b, 0x9b, 0xe2, 0xad, 0xa1, 0x3d, 0xa1, 0x52, 0xb3, 0x39, 0x4b, 0x50, 0x53, 0x45, 0xde,
	0x42, 0x3b, 0xd9, 0x8a, 0x5d, 0xa7, 0x57, 0xf3, 0x5b, 0xfd, 0x47, 0xc1, 0x6e, 0xe9, 0x60, 0x4b,
	0x13, 0xed, 0x08, 0x86, 0xcf, 0x7f, 0xfe, 0xfe, 0xfe, 0xc4, 0x83, 0xde, 0x8e, 0xa0, 0x8f, 0x59,
	0xb1, 0xc4, 0x6d, 0x95, 0xf2, 0x7e, 0x39, 0xd0, 0xda, 0x02, 0xc8, 0x1b, 0xb8, 0x93, 0x60, 0x6c,
	0xfe, 0xba, 0x28, 0xea, 0xfd, 0xa3, 0xe8, 0x3b, 0xaa, 0x91, 0x65, 0x2a, 0x6a, 0x24, 0x68, 0x50,
	0x32, 0x02, 0x30, 0xca, 0x38, 0x59, 0x22, 0xe3, 0xee, 0xd1, 0x7f, 0xeb, 0x9b, 0x46, 0x35, 0x31,
	0xa2, 0xe1, 0x99, 0x69, 0xfc, 0x14, 0xba, 0x07, 0x1a, 0xf7, 0xbe, 0xd5, 0x80, 0xfc, 0xfd, 0x13,
	0x21, 0x50, 0x2f, 0x50, 0x2f, 0x5d, 0xa7, 0xe7, 0xf8, 0xcd, 0xa8, 0x7a, 0x93, 0xa7, 0x70, 0x57,
	0x51, 0xc9, 0x30, 0x8b, 0x79, 0x99, 0xcf, 0xa8, 0x74, 0x8f, 0x2a, 0xb2, 0x6d, 0xc1, 0xf3, 0x0a,
	0x23, 0x53, 0xb8, 0xa7, 0xca, 0xd9, 0x57, 0x9a, 0xe8, 0x18, 0x33, 0x1d, 0x73, 0xcc, 0xa9, 0x72,
	0x6b, 0xd5, 0x04, 0xcf, 0x6e, 0x4e, 0xf0, 0xc9, 0x26, 0x8e, 0x32, 0x4d, 0x25, 0x47, 0x4d, 0xcf,
	0x31, 0xa7, 0xd1, 0x89, 0xba, 0x42, 0x4d, 0xac, 0x48, 0x1f, 0x1e, 0xa6, 0xb8, 0x51, 0x71, 0xc9,
	0x35, 0xcb, 0x62, 0xba, 0x2e, 0x98, 0xac, 0xd6, 0xef, 0xd6, 0x7b, 0x8e, 0x5f, 0x8f, 0xee, 0x1b,
	0xf2, 0x8b, 0xe1, 0xde, 0x5f, 0x51, 0xe4, 0x35, 0xc0, 0x0a, 0x33, 0x96, 0xc6, 0x73, 0x29, 0x72,
	0xf7, 0x76, 0xcf, 0xf1, 0x5b, 0xfd, 0x4e, 0x60, 0x0f, 0x2c, 0xb8, 0x3c, 0xb0, 0xe0, 0xf3, 0xe5,
	0x81, 0x45, 0xcd, 0x2a, 0xfb, 0x83, 0x14, 0x39, 0x99, 0xc0, 0xc9, 0x75, 0x8d, 0xd8, 0xdc, 0xa0,
	0xdb, 0x38, 0xa8, 0x3f, 0xbe, 0x96, 0x18, 0x70, 0xf8, 0xca, 0xb8, 0xef, 0xc3, 0xd9, 0x01, 0xf7,
	0x2f, 0xdc, 0xf6, 0x7e, 0x38, 0xf0, 0x60, 0x9f, 0x19, 0x84, 0x40, 0x2d, 0xe5, 0xca, 0x6e, 0xe1,
	0xe3, 0xad, 0xc8, 0x04, 0x06, 0x2b, 0x25, 0xb3, 0xe6, 0x1b, 0xac, 0x94, 0x8c, 0x74, 0x01, 0x58,
	0x11, 0x63, 0x9a, 0x4a, 0xaa, 0x8c, 0xdd, 0x96, 0x6a, 0xb2, 0x62, 0x64, 0xa1, 0x61, 0x68, 0x1a,
	0x7a, 0x01, 0xfe, 0xbe, 0x86, 0xf6, 0x55, 0x1e, 0x37, 0xa0, 0x6e, 0x76, 0x37, 0x7e, 0x09, 0x8f,
	0x99, 0xb0, 0x8b, 0x2b, 0xa4, 0x58, 0x6f, 0x6e, 0xec, 0x70, 0x0c, 0x66, 0x1c, 0x35, 0x35, 0x96,
	0x4c, 0x9d, 0x59, 0xa3, 0xf2, 0x66, 0xf0, 0x27, 0x00, 0x00, 0xff, 0xff, 0xdb, 0xc6, 0x88, 0x69,
	0xf8, 0x03, 0x00, 0x00,
}
