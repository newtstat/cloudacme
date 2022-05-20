// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: v1/cloudacme/certificates.proto

package cloudacme

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IssueCertificateRequest_Provider int32

const (
	IssueCertificateRequest_none   IssueCertificateRequest_Provider = 0
	IssueCertificateRequest_gcloud IssueCertificateRequest_Provider = 1
)

// Enum value maps for IssueCertificateRequest_Provider.
var (
	IssueCertificateRequest_Provider_name = map[int32]string{
		0: "none",
		1: "gcloud",
	}
	IssueCertificateRequest_Provider_value = map[string]int32{
		"none":   0,
		"gcloud": 1,
	}
)

func (x IssueCertificateRequest_Provider) Enum() *IssueCertificateRequest_Provider {
	p := new(IssueCertificateRequest_Provider)
	*p = x
	return p
}

func (x IssueCertificateRequest_Provider) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IssueCertificateRequest_Provider) Descriptor() protoreflect.EnumDescriptor {
	return file_v1_cloudacme_certificates_proto_enumTypes[0].Descriptor()
}

func (IssueCertificateRequest_Provider) Type() protoreflect.EnumType {
	return &file_v1_cloudacme_certificates_proto_enumTypes[0]
}

func (x IssueCertificateRequest_Provider) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IssueCertificateRequest_Provider.Descriptor instead.
func (IssueCertificateRequest_Provider) EnumDescriptor() ([]byte, []int) {
	return file_v1_cloudacme_certificates_proto_rawDescGZIP(), []int{0, 0}
}

type IssueCertificateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VaultProvider               string   `protobuf:"bytes,1,opt,name=vaultProvider,proto3" json:"vaultProvider,omitempty"`
	AcmeAccountKeyVaultResource string   `protobuf:"bytes,2,opt,name=acmeAccountKeyVaultResource,proto3" json:"acmeAccountKeyVaultResource,omitempty"`
	PrivateKeyVaultResource     string   `protobuf:"bytes,3,opt,name=privateKeyVaultResource,proto3" json:"privateKeyVaultResource,omitempty"`
	CertificateVaultResource    string   `protobuf:"bytes,4,opt,name=certificateVaultResource,proto3" json:"certificateVaultResource,omitempty"`
	RenewPrivateKey             bool     `protobuf:"varint,5,opt,name=renewPrivateKey,proto3" json:"renewPrivateKey,omitempty"`
	KeyAlgorithm                string   `protobuf:"bytes,6,opt,name=keyAlgorithm,proto3" json:"keyAlgorithm,omitempty"`
	DnsProvider                 string   `protobuf:"bytes,7,opt,name=dnsProvider,proto3" json:"dnsProvider,omitempty"`
	DnsProviderID               string   `protobuf:"bytes,8,opt,name=dnsProviderID,proto3" json:"dnsProviderID,omitempty"`
	TermsOfServiceAgreed        bool     `protobuf:"varint,9,opt,name=termsOfServiceAgreed,proto3" json:"termsOfServiceAgreed,omitempty"`
	Email                       string   `protobuf:"bytes,10,opt,name=email,proto3" json:"email,omitempty"`
	ThresholdOfDaysToExpire     int64    `protobuf:"varint,11,opt,name=thresholdOfDaysToExpire,proto3" json:"thresholdOfDaysToExpire,omitempty"`
	Domains                     []string `protobuf:"bytes,12,rep,name=domains,proto3" json:"domains,omitempty"`
	Staging                     bool     `protobuf:"varint,13,opt,name=staging,proto3" json:"staging,omitempty"`
}

func (x *IssueCertificateRequest) Reset() {
	*x = IssueCertificateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cloudacme_certificates_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCertificateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCertificateRequest) ProtoMessage() {}

func (x *IssueCertificateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cloudacme_certificates_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCertificateRequest.ProtoReflect.Descriptor instead.
func (*IssueCertificateRequest) Descriptor() ([]byte, []int) {
	return file_v1_cloudacme_certificates_proto_rawDescGZIP(), []int{0}
}

func (x *IssueCertificateRequest) GetVaultProvider() string {
	if x != nil {
		return x.VaultProvider
	}
	return ""
}

func (x *IssueCertificateRequest) GetAcmeAccountKeyVaultResource() string {
	if x != nil {
		return x.AcmeAccountKeyVaultResource
	}
	return ""
}

func (x *IssueCertificateRequest) GetPrivateKeyVaultResource() string {
	if x != nil {
		return x.PrivateKeyVaultResource
	}
	return ""
}

func (x *IssueCertificateRequest) GetCertificateVaultResource() string {
	if x != nil {
		return x.CertificateVaultResource
	}
	return ""
}

func (x *IssueCertificateRequest) GetRenewPrivateKey() bool {
	if x != nil {
		return x.RenewPrivateKey
	}
	return false
}

func (x *IssueCertificateRequest) GetKeyAlgorithm() string {
	if x != nil {
		return x.KeyAlgorithm
	}
	return ""
}

func (x *IssueCertificateRequest) GetDnsProvider() string {
	if x != nil {
		return x.DnsProvider
	}
	return ""
}

func (x *IssueCertificateRequest) GetDnsProviderID() string {
	if x != nil {
		return x.DnsProviderID
	}
	return ""
}

func (x *IssueCertificateRequest) GetTermsOfServiceAgreed() bool {
	if x != nil {
		return x.TermsOfServiceAgreed
	}
	return false
}

func (x *IssueCertificateRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *IssueCertificateRequest) GetThresholdOfDaysToExpire() int64 {
	if x != nil {
		return x.ThresholdOfDaysToExpire
	}
	return 0
}

func (x *IssueCertificateRequest) GetDomains() []string {
	if x != nil {
		return x.Domains
	}
	return nil
}

func (x *IssueCertificateRequest) GetStaging() bool {
	if x != nil {
		return x.Staging
	}
	return false
}

type IssueCertificateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PrivateKeyVaultVersionResource  string `protobuf:"bytes,1,opt,name=privateKeyVaultVersionResource,proto3" json:"privateKeyVaultVersionResource,omitempty"`
	CertificateVaultVersionResource string `protobuf:"bytes,2,opt,name=certificateVaultVersionResource,proto3" json:"certificateVaultVersionResource,omitempty"`
}

func (x *IssueCertificateResponse) Reset() {
	*x = IssueCertificateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_v1_cloudacme_certificates_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCertificateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCertificateResponse) ProtoMessage() {}

func (x *IssueCertificateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_v1_cloudacme_certificates_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCertificateResponse.ProtoReflect.Descriptor instead.
func (*IssueCertificateResponse) Descriptor() ([]byte, []int) {
	return file_v1_cloudacme_certificates_proto_rawDescGZIP(), []int{1}
}

func (x *IssueCertificateResponse) GetPrivateKeyVaultVersionResource() string {
	if x != nil {
		return x.PrivateKeyVaultVersionResource
	}
	return ""
}

func (x *IssueCertificateResponse) GetCertificateVaultVersionResource() string {
	if x != nil {
		return x.CertificateVaultVersionResource
	}
	return ""
}

var File_v1_cloudacme_certificates_proto protoreflect.FileDescriptor

var file_v1_cloudacme_certificates_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0c, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf4, 0x05, 0x0a, 0x17, 0x49, 0x73, 0x73, 0x75, 0x65,
	0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x33, 0x0a, 0x0d, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x72, 0x08,
	0x52, 0x06, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x52, 0x0d, 0x76, 0x61, 0x75, 0x6c, 0x74, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x49, 0x0a, 0x1b, 0x61, 0x63, 0x6d, 0x65, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x1b, 0x61, 0x63, 0x6d, 0x65, 0x41, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x41, 0x0a, 0x17, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79,
	0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x17, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x18, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01,
	0x52, 0x18, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x75,
	0x6c, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x28, 0x0a, 0x0f, 0x72, 0x65,
	0x6e, 0x65, 0x77, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0f, 0x72, 0x65, 0x6e, 0x65, 0x77, 0x50, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x4b, 0x65, 0x79, 0x12, 0x5a, 0x0a, 0x0c, 0x6b, 0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x72,
	0x69, 0x74, 0x68, 0x6d, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x42, 0x36, 0xfa, 0x42, 0x33, 0x72,
	0x31, 0x52, 0x00, 0x52, 0x07, 0x72, 0x73, 0x61, 0x32, 0x30, 0x34, 0x38, 0x52, 0x07, 0x72, 0x73,
	0x61, 0x34, 0x30, 0x39, 0x36, 0x52, 0x07, 0x72, 0x73, 0x61, 0x38, 0x31, 0x39, 0x32, 0x52, 0x08,
	0x65, 0x63, 0x64, 0x73, 0x61, 0x32, 0x35, 0x36, 0x52, 0x08, 0x65, 0x63, 0x64, 0x73, 0x61, 0x33,
	0x38, 0x34, 0x52, 0x0c, 0x6b, 0x65, 0x79, 0x41, 0x6c, 0x67, 0x6f, 0x72, 0x69, 0x74, 0x68, 0x6d,
	0x12, 0x2f, 0x0a, 0x0b, 0x64, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0xfa, 0x42, 0x0a, 0x72, 0x08, 0x52, 0x06, 0x67, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x52, 0x0b, 0x64, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x12, 0x2d, 0x0a, 0x0d, 0x64, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10,
	0x01, 0x52, 0x0d, 0x64, 0x6e, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44,
	0x12, 0x32, 0x0a, 0x14, 0x74, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x41, 0x67, 0x72, 0x65, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x08, 0x52, 0x14,
	0x74, 0x65, 0x72, 0x6d, 0x73, 0x4f, 0x66, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x41, 0x67,
	0x72, 0x65, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x60, 0x01, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x38, 0x0a, 0x17, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64,
	0x4f, 0x66, 0x44, 0x61, 0x79, 0x73, 0x54, 0x6f, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x17, 0x74, 0x68, 0x72, 0x65, 0x73, 0x68, 0x6f, 0x6c, 0x64, 0x4f,
	0x66, 0x44, 0x61, 0x79, 0x73, 0x54, 0x6f, 0x45, 0x78, 0x70, 0x69, 0x72, 0x65, 0x12, 0x22, 0x0a,
	0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x09, 0x42, 0x08,
	0xfa, 0x42, 0x05, 0x92, 0x01, 0x02, 0x08, 0x01, 0x52, 0x07, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x73, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x74, 0x61, 0x67, 0x69, 0x6e, 0x67, 0x22, 0x20, 0x0a, 0x08, 0x50,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x08, 0x0a, 0x04, 0x6e, 0x6f, 0x6e, 0x65, 0x10,
	0x00, 0x12, 0x0a, 0x0a, 0x06, 0x67, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x10, 0x01, 0x22, 0xac, 0x01,
	0x0a, 0x18, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x46, 0x0a, 0x1e, 0x70, 0x72,
	0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x1e, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74, 0x65, 0x4b, 0x65, 0x79, 0x56, 0x61,
	0x75, 0x6c, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x12, 0x48, 0x0a, 0x1f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1f, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x56, 0x61, 0x75, 0x6c, 0x74, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x32, 0x8d, 0x01, 0x0a,
	0x0c, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x12, 0x7d, 0x0a,
	0x05, 0x49, 0x73, 0x73, 0x75, 0x65, 0x12, 0x25, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63,
	0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e,
	0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x6d, 0x65, 0x2e, 0x76, 0x31, 0x2e, 0x49, 0x73, 0x73,
	0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x25, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1f, 0x22, 0x1a, 0x2f,
	0x72, 0x70, 0x63, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x73, 0x2f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x3f, 0x5a, 0x3d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x6e, 0x6f, 0x6b,
	0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x6d, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x67, 0x6f,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x61, 0x63, 0x6d, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_v1_cloudacme_certificates_proto_rawDescOnce sync.Once
	file_v1_cloudacme_certificates_proto_rawDescData = file_v1_cloudacme_certificates_proto_rawDesc
)

func file_v1_cloudacme_certificates_proto_rawDescGZIP() []byte {
	file_v1_cloudacme_certificates_proto_rawDescOnce.Do(func() {
		file_v1_cloudacme_certificates_proto_rawDescData = protoimpl.X.CompressGZIP(file_v1_cloudacme_certificates_proto_rawDescData)
	})
	return file_v1_cloudacme_certificates_proto_rawDescData
}

var file_v1_cloudacme_certificates_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_v1_cloudacme_certificates_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_v1_cloudacme_certificates_proto_goTypes = []interface{}{
	(IssueCertificateRequest_Provider)(0), // 0: cloudacme.v1.IssueCertificateRequest.Provider
	(*IssueCertificateRequest)(nil),       // 1: cloudacme.v1.IssueCertificateRequest
	(*IssueCertificateResponse)(nil),      // 2: cloudacme.v1.IssueCertificateResponse
}
var file_v1_cloudacme_certificates_proto_depIdxs = []int32{
	1, // 0: cloudacme.v1.Certificates.Issue:input_type -> cloudacme.v1.IssueCertificateRequest
	2, // 1: cloudacme.v1.Certificates.Issue:output_type -> cloudacme.v1.IssueCertificateResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_v1_cloudacme_certificates_proto_init() }
func file_v1_cloudacme_certificates_proto_init() {
	if File_v1_cloudacme_certificates_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_v1_cloudacme_certificates_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCertificateRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_v1_cloudacme_certificates_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCertificateResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_v1_cloudacme_certificates_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_v1_cloudacme_certificates_proto_goTypes,
		DependencyIndexes: file_v1_cloudacme_certificates_proto_depIdxs,
		EnumInfos:         file_v1_cloudacme_certificates_proto_enumTypes,
		MessageInfos:      file_v1_cloudacme_certificates_proto_msgTypes,
	}.Build()
	File_v1_cloudacme_certificates_proto = out.File
	file_v1_cloudacme_certificates_proto_rawDesc = nil
	file_v1_cloudacme_certificates_proto_goTypes = nil
	file_v1_cloudacme_certificates_proto_depIdxs = nil
}
