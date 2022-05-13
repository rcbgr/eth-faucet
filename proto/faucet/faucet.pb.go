// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: faucet/faucet.proto

package faucet

import (
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

type FundingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	WalletAddress string `protobuf:"bytes,1,opt,name=wallet_address,json=walletAddress,proto3" json:"wallet_address,omitempty"`
}

func (x *FundingRequest) Reset() {
	*x = FundingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faucet_faucet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundingRequest) ProtoMessage() {}

func (x *FundingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_faucet_faucet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundingRequest.ProtoReflect.Descriptor instead.
func (*FundingRequest) Descriptor() ([]byte, []int) {
	return file_faucet_faucet_proto_rawDescGZIP(), []int{0}
}

func (x *FundingRequest) GetWalletAddress() string {
	if x != nil {
		return x.WalletAddress
	}
	return ""
}

type FundingResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Amount          string `protobuf:"bytes,1,opt,name=amount,proto3" json:"amount,omitempty"`
	TransactionHash string `protobuf:"bytes,2,opt,name=transaction_hash,json=transactionHash,proto3" json:"transaction_hash,omitempty"`
}

func (x *FundingResponse) Reset() {
	*x = FundingResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_faucet_faucet_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundingResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundingResponse) ProtoMessage() {}

func (x *FundingResponse) ProtoReflect() protoreflect.Message {
	mi := &file_faucet_faucet_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FundingResponse.ProtoReflect.Descriptor instead.
func (*FundingResponse) Descriptor() ([]byte, []int) {
	return file_faucet_faucet_proto_rawDescGZIP(), []int{1}
}

func (x *FundingResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *FundingResponse) GetTransactionHash() string {
	if x != nil {
		return x.TransactionHash
	}
	return ""
}

var File_faucet_faucet_proto protoreflect.FileDescriptor

var file_faucet_faucet_proto_rawDesc = []byte{
	0x0a, 0x13, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x37, 0x0a, 0x0e, 0x46,
	0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x25, 0x0a,
	0x0e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x22, 0x54, 0x0a, 0x0f, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x29, 0x0a, 0x10, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x68,
	0x61, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x74, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x32, 0x6c, 0x0a, 0x06, 0x46, 0x61,
	0x75, 0x63, 0x65, 0x74, 0x12, 0x62, 0x0a, 0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x46,
	0x75, 0x6e, 0x64, 0x73, 0x12, 0x16, 0x2e, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x46, 0x75,
	0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x66,
	0x61, 0x75, 0x63, 0x65, 0x74, 0x2e, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x21, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1b, 0x22, 0x16, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x3a, 0x01, 0x2a, 0x42, 0x2f, 0x5a, 0x2d, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x61, 0x75, 0x6c, 0x6a, 0x6f, 0x72, 0x64, 0x61,
	0x6e, 0x2f, 0x65, 0x74, 0x68, 0x2d, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x66, 0x61, 0x75, 0x63, 0x65, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_faucet_faucet_proto_rawDescOnce sync.Once
	file_faucet_faucet_proto_rawDescData = file_faucet_faucet_proto_rawDesc
)

func file_faucet_faucet_proto_rawDescGZIP() []byte {
	file_faucet_faucet_proto_rawDescOnce.Do(func() {
		file_faucet_faucet_proto_rawDescData = protoimpl.X.CompressGZIP(file_faucet_faucet_proto_rawDescData)
	})
	return file_faucet_faucet_proto_rawDescData
}

var file_faucet_faucet_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_faucet_faucet_proto_goTypes = []interface{}{
	(*FundingRequest)(nil),  // 0: faucet.FundingRequest
	(*FundingResponse)(nil), // 1: faucet.FundingResponse
}
var file_faucet_faucet_proto_depIdxs = []int32{
	0, // 0: faucet.Faucet.RequestFunds:input_type -> faucet.FundingRequest
	1, // 1: faucet.Faucet.RequestFunds:output_type -> faucet.FundingResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_faucet_faucet_proto_init() }
func file_faucet_faucet_proto_init() {
	if File_faucet_faucet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_faucet_faucet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundingRequest); i {
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
		file_faucet_faucet_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundingResponse); i {
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
			RawDescriptor: file_faucet_faucet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_faucet_faucet_proto_goTypes,
		DependencyIndexes: file_faucet_faucet_proto_depIdxs,
		MessageInfos:      file_faucet_faucet_proto_msgTypes,
	}.Build()
	File_faucet_faucet_proto = out.File
	file_faucet_faucet_proto_rawDesc = nil
	file_faucet_faucet_proto_goTypes = nil
	file_faucet_faucet_proto_depIdxs = nil
}
