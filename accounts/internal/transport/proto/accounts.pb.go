// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: accounts.transport

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NewUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *NewUserRequest) Reset() {
	*x = NewUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewUserRequest) ProtoMessage() {}

func (x *NewUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewUserRequest.ProtoReflect.Descriptor instead.
func (*NewUserRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{0}
}

func (x *NewUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type NewWalletRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *NewWalletRequest) Reset() {
	*x = NewWalletRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewWalletRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewWalletRequest) ProtoMessage() {}

func (x *NewWalletRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewWalletRequest.ProtoReflect.Descriptor instead.
func (*NewWalletRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{1}
}

func (x *NewWalletRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type ChangeBalanceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WalletId uint64 `protobuf:"varint,2,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	Change   string `protobuf:"bytes,3,opt,name=change,proto3" json:"change,omitempty"`
}

func (x *ChangeBalanceRequest) Reset() {
	*x = ChangeBalanceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeBalanceRequest) ProtoMessage() {}

func (x *ChangeBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeBalanceRequest.ProtoReflect.Descriptor instead.
func (*ChangeBalanceRequest) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{2}
}

func (x *ChangeBalanceRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ChangeBalanceRequest) GetWalletId() uint64 {
	if x != nil {
		return x.WalletId
	}
	return 0
}

func (x *ChangeBalanceRequest) GetChange() string {
	if x != nil {
		return x.Change
	}
	return ""
}

type AccountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	WalletId uint64 `protobuf:"varint,2,opt,name=wallet_id,json=walletId,proto3" json:"wallet_id,omitempty"`
	Balance  string `protobuf:"bytes,3,opt,name=balance,proto3" json:"balance,omitempty"`
}

func (x *AccountResponse) Reset() {
	*x = AccountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accounts_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccountResponse) ProtoMessage() {}

func (x *AccountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accounts_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccountResponse.ProtoReflect.Descriptor instead.
func (*AccountResponse) Descriptor() ([]byte, []int) {
	return file_accounts_proto_rawDescGZIP(), []int{3}
}

func (x *AccountResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *AccountResponse) GetWalletId() uint64 {
	if x != nil {
		return x.WalletId
	}
	return 0
}

func (x *AccountResponse) GetBalance() string {
	if x != nil {
		return x.Balance
	}
	return ""
}

var File_accounts_proto protoreflect.FileDescriptor

var file_accounts_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x03, 0x61, 0x70, 0x69, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x26, 0x0a, 0x0e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x22, 0x0a, 0x10, 0x4e, 0x65,
	0x77, 0x57, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x5b,
	0x0a, 0x14, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65,
	0x74, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x22, 0x58, 0x0a, 0x0f, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b,
	0x0a, 0x09, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x08, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x62,
	0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x61,
	0x6c, 0x61, 0x6e, 0x63, 0x65, 0x32, 0xca, 0x02, 0x0a, 0x08, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x73, 0x12, 0x3c, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3f, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x73, 0x12,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30,
	0x01, 0x12, 0x40, 0x0a, 0x0f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x12, 0x15, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4e, 0x65, 0x77, 0x57, 0x61,
	0x6c, 0x6c, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x3c, 0x0a, 0x07, 0x44, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x12, 0x19,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x61, 0x6c, 0x61, 0x6e,
	0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x3f, 0x0a, 0x0a, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x61, 0x6c, 0x12,
	0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x42, 0x61, 0x6c, 0x61,
	0x6e, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accounts_proto_rawDescOnce sync.Once
	file_accounts_proto_rawDescData = file_accounts_proto_rawDesc
)

func file_accounts_proto_rawDescGZIP() []byte {
	file_accounts_proto_rawDescOnce.Do(func() {
		file_accounts_proto_rawDescData = protoimpl.X.CompressGZIP(file_accounts_proto_rawDescData)
	})
	return file_accounts_proto_rawDescData
}

var file_accounts_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_accounts_proto_goTypes = []interface{}{
	(*NewUserRequest)(nil),       // 0: transport.NewUserRequest
	(*NewWalletRequest)(nil),     // 1: transport.NewWalletRequest
	(*ChangeBalanceRequest)(nil), // 2: transport.ChangeBalanceRequest
	(*AccountResponse)(nil),      // 3: transport.AccountResponse
	(*emptypb.Empty)(nil),        // 4: google.protobuf.Empty
}
var file_accounts_proto_depIdxs = []int32{
	0, // 0: transport.Accounts.CreateAccount:input_type -> transport.NewUserRequest
	4, // 1: transport.Accounts.GetAccounts:input_type -> google.protobuf.Empty
	1, // 2: transport.Accounts.GenerateAddress:input_type -> transport.NewWalletRequest
	2, // 3: transport.Accounts.Deposit:input_type -> transport.ChangeBalanceRequest
	2, // 4: transport.Accounts.Withdrawal:input_type -> transport.ChangeBalanceRequest
	3, // 5: transport.Accounts.CreateAccount:output_type -> transport.AccountResponse
	3, // 6: transport.Accounts.GetAccounts:output_type -> transport.AccountResponse
	3, // 7: transport.Accounts.GenerateAddress:output_type -> transport.AccountResponse
	3, // 8: transport.Accounts.Deposit:output_type -> transport.AccountResponse
	3, // 9: transport.Accounts.Withdrawal:output_type -> transport.AccountResponse
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_accounts_proto_init() }
func file_accounts_proto_init() {
	if File_accounts_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accounts_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewUserRequest); i {
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
		file_accounts_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewWalletRequest); i {
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
		file_accounts_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeBalanceRequest); i {
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
		file_accounts_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccountResponse); i {
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
			RawDescriptor: file_accounts_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accounts_proto_goTypes,
		DependencyIndexes: file_accounts_proto_depIdxs,
		MessageInfos:      file_accounts_proto_msgTypes,
	}.Build()
	File_accounts_proto = out.File
	file_accounts_proto_rawDesc = nil
	file_accounts_proto_goTypes = nil
	file_accounts_proto_depIdxs = nil
}
