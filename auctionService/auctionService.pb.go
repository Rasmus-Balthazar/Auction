// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: auctionService/auctionService.proto

package auctionService

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

type BidState int32

const (
	BidState_SUCCESS BidState = 0
	BidState_FAIL    BidState = 1
)

// Enum value maps for BidState.
var (
	BidState_name = map[int32]string{
		0: "SUCCESS",
		1: "FAIL",
	}
	BidState_value = map[string]int32{
		"SUCCESS": 0,
		"FAIL":    1,
	}
)

func (x BidState) Enum() *BidState {
	p := new(BidState)
	*p = x
	return p
}

func (x BidState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (BidState) Descriptor() protoreflect.EnumDescriptor {
	return file_auctionService_auctionService_proto_enumTypes[0].Descriptor()
}

func (BidState) Type() protoreflect.EnumType {
	return &file_auctionService_auctionService_proto_enumTypes[0]
}

func (x BidState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use BidState.Descriptor instead.
func (BidState) EnumDescriptor() ([]byte, []int) {
	return file_auctionService_auctionService_proto_rawDescGZIP(), []int{0}
}

type AuctionState int32

const (
	AuctionState_OVER   AuctionState = 0
	AuctionState_GOING  AuctionState = 1
	AuctionState_FAILED AuctionState = 2
)

// Enum value maps for AuctionState.
var (
	AuctionState_name = map[int32]string{
		0: "OVER",
		1: "GOING",
		2: "FAILED",
	}
	AuctionState_value = map[string]int32{
		"OVER":   0,
		"GOING":  1,
		"FAILED": 2,
	}
)

func (x AuctionState) Enum() *AuctionState {
	p := new(AuctionState)
	*p = x
	return p
}

func (x AuctionState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AuctionState) Descriptor() protoreflect.EnumDescriptor {
	return file_auctionService_auctionService_proto_enumTypes[1].Descriptor()
}

func (AuctionState) Type() protoreflect.EnumType {
	return &file_auctionService_auctionService_proto_enumTypes[1]
}

func (x AuctionState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use AuctionState.Descriptor instead.
func (AuctionState) EnumDescriptor() ([]byte, []int) {
	return file_auctionService_auctionService_proto_rawDescGZIP(), []int{1}
}

type Message struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pid     uint32 `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Message) Reset() {
	*x = Message{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auctionService_auctionService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Message) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Message) ProtoMessage() {}

func (x *Message) ProtoReflect() protoreflect.Message {
	mi := &file_auctionService_auctionService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Message.ProtoReflect.Descriptor instead.
func (*Message) Descriptor() ([]byte, []int) {
	return file_auctionService_auctionService_proto_rawDescGZIP(), []int{0}
}

func (x *Message) GetPid() uint32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *Message) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type BidMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BidderId  uint32 `protobuf:"varint,1,opt,name=bidderId,proto3" json:"bidderId,omitempty"`
	BidAmount int64  `protobuf:"varint,2,opt,name=bidAmount,proto3" json:"bidAmount,omitempty"`
}

func (x *BidMessage) Reset() {
	*x = BidMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auctionService_auctionService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BidMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BidMessage) ProtoMessage() {}

func (x *BidMessage) ProtoReflect() protoreflect.Message {
	mi := &file_auctionService_auctionService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BidMessage.ProtoReflect.Descriptor instead.
func (*BidMessage) Descriptor() ([]byte, []int) {
	return file_auctionService_auctionService_proto_rawDescGZIP(), []int{1}
}

func (x *BidMessage) GetBidderId() uint32 {
	if x != nil {
		return x.BidderId
	}
	return 0
}

func (x *BidMessage) GetBidAmount() int64 {
	if x != nil {
		return x.BidAmount
	}
	return 0
}

type Outcome struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	State    AuctionState `protobuf:"varint,1,opt,name=state,proto3,enum=auctionService.AuctionState" json:"state,omitempty"`
	Amount   string       `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	BidderId uint32       `protobuf:"varint,3,opt,name=bidderId,proto3" json:"bidderId,omitempty"`
}

func (x *Outcome) Reset() {
	*x = Outcome{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auctionService_auctionService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Outcome) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Outcome) ProtoMessage() {}

func (x *Outcome) ProtoReflect() protoreflect.Message {
	mi := &file_auctionService_auctionService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Outcome.ProtoReflect.Descriptor instead.
func (*Outcome) Descriptor() ([]byte, []int) {
	return file_auctionService_auctionService_proto_rawDescGZIP(), []int{2}
}

func (x *Outcome) GetState() AuctionState {
	if x != nil {
		return x.State
	}
	return AuctionState_OVER
}

func (x *Outcome) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *Outcome) GetBidderId() uint32 {
	if x != nil {
		return x.BidderId
	}
	return 0
}

var File_auctionService_auctionService_proto protoreflect.FileDescriptor

var file_auctionService_auctionService_proto_rawDesc = []byte{
	0x0a, 0x23, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x35, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12,
	0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x46, 0x0a, 0x0a, 0x42, 0x69, 0x64,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x62, 0x69, 0x64, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x22, 0x71, 0x0a, 0x07, 0x4f, 0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x32, 0x0a, 0x05,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1c, 0x2e, 0x61, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x69, 0x64, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08, 0x62, 0x69, 0x64, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x2a, 0x21, 0x0a, 0x08, 0x42, 0x69, 0x64, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00, 0x12, 0x08, 0x0a,
	0x04, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x01, 0x2a, 0x2f, 0x0a, 0x0c, 0x41, 0x75, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4f, 0x56, 0x45, 0x52, 0x10,
	0x00, 0x12, 0x09, 0x0a, 0x05, 0x47, 0x4f, 0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06,
	0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x32, 0xce, 0x01, 0x0a, 0x0e, 0x41, 0x75, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x17, 0x2e,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f,
	0x75, 0x74, 0x63, 0x6f, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x03, 0x42, 0x69, 0x64, 0x12, 0x1a, 0x2e,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42,
	0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x2e, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x75, 0x74, 0x63, 0x6f,
	0x6d, 0x65, 0x12, 0x45, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x12, 0x1a, 0x2e,
	0x61, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42,
	0x69, 0x64, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x1a, 0x2e, 0x61, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x42, 0x69, 0x64, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3d, 0x5a, 0x3b, 0x68, 0x74, 0x74,
	0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x52, 0x61, 0x73, 0x6d, 0x75, 0x73, 0x2d, 0x42, 0x61, 0x6c, 0x74, 0x68, 0x61, 0x7a, 0x61, 0x72,
	0x2f, 0x41, 0x75, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x3b, 0x61, 0x75, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auctionService_auctionService_proto_rawDescOnce sync.Once
	file_auctionService_auctionService_proto_rawDescData = file_auctionService_auctionService_proto_rawDesc
)

func file_auctionService_auctionService_proto_rawDescGZIP() []byte {
	file_auctionService_auctionService_proto_rawDescOnce.Do(func() {
		file_auctionService_auctionService_proto_rawDescData = protoimpl.X.CompressGZIP(file_auctionService_auctionService_proto_rawDescData)
	})
	return file_auctionService_auctionService_proto_rawDescData
}

var file_auctionService_auctionService_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_auctionService_auctionService_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_auctionService_auctionService_proto_goTypes = []interface{}{
	(BidState)(0),         // 0: auctionService.BidState
	(AuctionState)(0),     // 1: auctionService.AuctionState
	(*Message)(nil),       // 2: auctionService.Message
	(*BidMessage)(nil),    // 3: auctionService.BidMessage
	(*Outcome)(nil),       // 4: auctionService.Outcome
	(*emptypb.Empty)(nil), // 5: google.protobuf.Empty
}
var file_auctionService_auctionService_proto_depIdxs = []int32{
	1, // 0: auctionService.Outcome.state:type_name -> auctionService.AuctionState
	5, // 1: auctionService.AuctionService.Result:input_type -> google.protobuf.Empty
	3, // 2: auctionService.AuctionService.Bid:input_type -> auctionService.BidMessage
	3, // 3: auctionService.AuctionService.Connect:input_type -> auctionService.BidMessage
	4, // 4: auctionService.AuctionService.Result:output_type -> auctionService.Outcome
	4, // 5: auctionService.AuctionService.Bid:output_type -> auctionService.Outcome
	3, // 6: auctionService.AuctionService.Connect:output_type -> auctionService.BidMessage
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_auctionService_auctionService_proto_init() }
func file_auctionService_auctionService_proto_init() {
	if File_auctionService_auctionService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auctionService_auctionService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Message); i {
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
		file_auctionService_auctionService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BidMessage); i {
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
		file_auctionService_auctionService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Outcome); i {
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
			RawDescriptor: file_auctionService_auctionService_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_auctionService_auctionService_proto_goTypes,
		DependencyIndexes: file_auctionService_auctionService_proto_depIdxs,
		EnumInfos:         file_auctionService_auctionService_proto_enumTypes,
		MessageInfos:      file_auctionService_auctionService_proto_msgTypes,
	}.Build()
	File_auctionService_auctionService_proto = out.File
	file_auctionService_auctionService_proto_rawDesc = nil
	file_auctionService_auctionService_proto_goTypes = nil
	file_auctionService_auctionService_proto_depIdxs = nil
}