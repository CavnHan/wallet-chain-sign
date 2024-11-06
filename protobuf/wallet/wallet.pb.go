// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.29.0--rc1
// source: protobuf/wallet.proto

package wallet

import (
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

type PublicKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CompressPubkey   string `protobuf:"bytes,1,opt,name=compress_pubkey,json=compressPubkey,proto3" json:"compress_pubkey,omitempty"`
	DecompressPubkey string `protobuf:"bytes,2,opt,name=decompress_pubkey,json=decompressPubkey,proto3" json:"decompress_pubkey,omitempty"`
}

func (x *PublicKey) Reset() {
	*x = PublicKey{}
	mi := &file_protobuf_wallet_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PublicKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicKey) ProtoMessage() {}

func (x *PublicKey) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicKey.ProtoReflect.Descriptor instead.
func (*PublicKey) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *PublicKey) GetCompressPubkey() string {
	if x != nil {
		return x.CompressPubkey
	}
	return ""
}

func (x *PublicKey) GetDecompressPubkey() string {
	if x != nil {
		return x.DecompressPubkey
	}
	return ""
}

type SupportSignWayRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
}

func (x *SupportSignWayRequest) Reset() {
	*x = SupportSignWayRequest{}
	mi := &file_protobuf_wallet_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SupportSignWayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportSignWayRequest) ProtoMessage() {}

func (x *SupportSignWayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportSignWayRequest.ProtoReflect.Descriptor instead.
func (*SupportSignWayRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{1}
}

func (x *SupportSignWayRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SupportSignWayRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

type SupportSignWayResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code    string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg     string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Support bool   `protobuf:"varint,3,opt,name=support,proto3" json:"support,omitempty"`
}

func (x *SupportSignWayResponse) Reset() {
	*x = SupportSignWayResponse{}
	mi := &file_protobuf_wallet_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SupportSignWayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SupportSignWayResponse) ProtoMessage() {}

func (x *SupportSignWayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SupportSignWayResponse.ProtoReflect.Descriptor instead.
func (*SupportSignWayResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{2}
}

func (x *SupportSignWayResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SupportSignWayResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SupportSignWayResponse) GetSupport() bool {
	if x != nil {
		return x.Support
	}
	return false
}

type ExportPublicKeyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	Type          string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Number        uint64 `protobuf:"varint,3,opt,name=number,proto3" json:"number,omitempty"`
}

func (x *ExportPublicKeyRequest) Reset() {
	*x = ExportPublicKeyRequest{}
	mi := &file_protobuf_wallet_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportPublicKeyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportPublicKeyRequest) ProtoMessage() {}

func (x *ExportPublicKeyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportPublicKeyRequest.ProtoReflect.Descriptor instead.
func (*ExportPublicKeyRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{3}
}

func (x *ExportPublicKeyRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *ExportPublicKeyRequest) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *ExportPublicKeyRequest) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

type ExportPublicKeyResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string       `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string       `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	PublicKey []*PublicKey `protobuf:"bytes,3,rep,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
}

func (x *ExportPublicKeyResponse) Reset() {
	*x = ExportPublicKeyResponse{}
	mi := &file_protobuf_wallet_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExportPublicKeyResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExportPublicKeyResponse) ProtoMessage() {}

func (x *ExportPublicKeyResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExportPublicKeyResponse.ProtoReflect.Descriptor instead.
func (*ExportPublicKeyResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{4}
}

func (x *ExportPublicKeyResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *ExportPublicKeyResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ExportPublicKeyResponse) GetPublicKey() []*PublicKey {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

type SignTxMessageRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConsumerToken string `protobuf:"bytes,1,opt,name=consumer_token,json=consumerToken,proto3" json:"consumer_token,omitempty"`
	PublicKey     string `protobuf:"bytes,2,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	MessageHash   string `protobuf:"bytes,3,opt,name=message_hash,json=messageHash,proto3" json:"message_hash,omitempty"`
}

func (x *SignTxMessageRequest) Reset() {
	*x = SignTxMessageRequest{}
	mi := &file_protobuf_wallet_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignTxMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignTxMessageRequest) ProtoMessage() {}

func (x *SignTxMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignTxMessageRequest.ProtoReflect.Descriptor instead.
func (*SignTxMessageRequest) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{5}
}

func (x *SignTxMessageRequest) GetConsumerToken() string {
	if x != nil {
		return x.ConsumerToken
	}
	return ""
}

func (x *SignTxMessageRequest) GetPublicKey() string {
	if x != nil {
		return x.PublicKey
	}
	return ""
}

func (x *SignTxMessageRequest) GetMessageHash() string {
	if x != nil {
		return x.MessageHash
	}
	return ""
}

type SignTxMessageResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code      string `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Msg       string `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Signature string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *SignTxMessageResponse) Reset() {
	*x = SignTxMessageResponse{}
	mi := &file_protobuf_wallet_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignTxMessageResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignTxMessageResponse) ProtoMessage() {}

func (x *SignTxMessageResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_wallet_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignTxMessageResponse.ProtoReflect.Descriptor instead.
func (*SignTxMessageResponse) Descriptor() ([]byte, []int) {
	return file_protobuf_wallet_proto_rawDescGZIP(), []int{6}
}

func (x *SignTxMessageResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *SignTxMessageResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *SignTxMessageResponse) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

var File_protobuf_wallet_proto protoreflect.FileDescriptor

var file_protobuf_wallet_proto_rawDesc = []byte{
	0x0a, 0x15, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x22, 0x61, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65,
	0x79, 0x12, 0x27, 0x0a, 0x0f, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x75,
	0x62, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x63, 0x6f, 0x6d, 0x70,
	0x72, 0x65, 0x73, 0x73, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x2b, 0x0a, 0x11, 0x64, 0x65,
	0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x64, 0x65, 0x63, 0x6f, 0x6d, 0x70, 0x72, 0x65, 0x73,
	0x73, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x22, 0x52, 0x0a, 0x15, 0x53, 0x75, 0x70, 0x70, 0x6f,
	0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d,
	0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22, 0x58, 0x0a, 0x16, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x6b, 0x0a, 0x16, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x50,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74, 0x6f, 0x6b, 0x65,
	0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65,
	0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x22, 0x76, 0x0a, 0x17, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x6d, 0x73, 0x67, 0x12, 0x35, 0x0a, 0x0a, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x2e, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52,
	0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x22, 0x7f, 0x0a, 0x14, 0x53, 0x69,
	0x67, 0x6e, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6d, 0x65, 0x72, 0x5f, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6d, 0x65, 0x72, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x75, 0x62,
	0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x48, 0x61, 0x73, 0x68, 0x22, 0x5b, 0x0a, 0x15, 0x53,
	0x69, 0x67, 0x6e, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x32, 0xad, 0x02, 0x0a, 0x0d, 0x57, 0x61, 0x6c,
	0x6c, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x5e, 0x0a, 0x11, 0x67, 0x65,
	0x74, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79, 0x12,
	0x22, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x53, 0x75,
	0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x23, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x69, 0x67,
	0x6e, 0x2e, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x53, 0x69, 0x67, 0x6e, 0x57, 0x61, 0x79,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x62, 0x0a, 0x13, 0x65, 0x78,
	0x70, 0x6f, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x23, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x2e,
	0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e,
	0x73, 0x69, 0x67, 0x6e, 0x2e, 0x45, 0x78, 0x70, 0x6f, 0x72, 0x74, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x58,
	0x0a, 0x0d, 0x73, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x21, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x69, 0x67, 0x6e, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x22, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x73, 0x69, 0x67, 0x6e,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x54, 0x78, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x13, 0x5a, 0x11, 0x2e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_wallet_proto_rawDescOnce sync.Once
	file_protobuf_wallet_proto_rawDescData = file_protobuf_wallet_proto_rawDesc
)

func file_protobuf_wallet_proto_rawDescGZIP() []byte {
	file_protobuf_wallet_proto_rawDescOnce.Do(func() {
		file_protobuf_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_wallet_proto_rawDescData)
	})
	return file_protobuf_wallet_proto_rawDescData
}

var file_protobuf_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_protobuf_wallet_proto_goTypes = []any{
	(*PublicKey)(nil),               // 0: wallet.sign.PublicKey
	(*SupportSignWayRequest)(nil),   // 1: wallet.sign.SupportSignWayRequest
	(*SupportSignWayResponse)(nil),  // 2: wallet.sign.SupportSignWayResponse
	(*ExportPublicKeyRequest)(nil),  // 3: wallet.sign.ExportPublicKeyRequest
	(*ExportPublicKeyResponse)(nil), // 4: wallet.sign.ExportPublicKeyResponse
	(*SignTxMessageRequest)(nil),    // 5: wallet.sign.SignTxMessageRequest
	(*SignTxMessageResponse)(nil),   // 6: wallet.sign.SignTxMessageResponse
}
var file_protobuf_wallet_proto_depIdxs = []int32{
	0, // 0: wallet.sign.ExportPublicKeyResponse.public_key:type_name -> wallet.sign.PublicKey
	1, // 1: wallet.sign.WalletService.getSupportSignWay:input_type -> wallet.sign.SupportSignWayRequest
	3, // 2: wallet.sign.WalletService.exportPublicKeyList:input_type -> wallet.sign.ExportPublicKeyRequest
	5, // 3: wallet.sign.WalletService.signTxMessage:input_type -> wallet.sign.SignTxMessageRequest
	2, // 4: wallet.sign.WalletService.getSupportSignWay:output_type -> wallet.sign.SupportSignWayResponse
	4, // 5: wallet.sign.WalletService.exportPublicKeyList:output_type -> wallet.sign.ExportPublicKeyResponse
	6, // 6: wallet.sign.WalletService.signTxMessage:output_type -> wallet.sign.SignTxMessageResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_protobuf_wallet_proto_init() }
func file_protobuf_wallet_proto_init() {
	if File_protobuf_wallet_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_protobuf_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protobuf_wallet_proto_goTypes,
		DependencyIndexes: file_protobuf_wallet_proto_depIdxs,
		MessageInfos:      file_protobuf_wallet_proto_msgTypes,
	}.Build()
	File_protobuf_wallet_proto = out.File
	file_protobuf_wallet_proto_rawDesc = nil
	file_protobuf_wallet_proto_goTypes = nil
	file_protobuf_wallet_proto_depIdxs = nil
}
