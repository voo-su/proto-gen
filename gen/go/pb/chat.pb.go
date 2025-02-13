// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.1
// source: chat.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetChatListRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChatListRequest) Reset() {
	*x = GetChatListRequest{}
	mi := &file_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChatListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatListRequest) ProtoMessage() {}

func (x *GetChatListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatListRequest.ProtoReflect.Descriptor instead.
func (*GetChatListRequest) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{0}
}

type GetChatListResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Items         []*ChatItem            `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetChatListResponse) Reset() {
	*x = GetChatListResponse{}
	mi := &file_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetChatListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetChatListResponse) ProtoMessage() {}

func (x *GetChatListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetChatListResponse.ProtoReflect.Descriptor instead.
func (*GetChatListResponse) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{1}
}

func (x *GetChatListResponse) GetItems() []*ChatItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type ChatItem struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ChatType      int32                  `protobuf:"varint,2,opt,name=chat_type,json=chatType,proto3" json:"chat_type,omitempty"`
	ReceiverId    int64                  `protobuf:"varint,3,opt,name=receiver_id,json=receiverId,proto3" json:"receiver_id,omitempty"`
	Username      string                 `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Avatar        string                 `protobuf:"bytes,5,opt,name=avatar,proto3" json:"avatar,omitempty"`
	Name          string                 `protobuf:"bytes,6,opt,name=name,proto3" json:"name,omitempty"`
	Surname       string                 `protobuf:"bytes,7,opt,name=surname,proto3" json:"surname,omitempty"`
	MsgText       string                 `protobuf:"bytes,8,opt,name=msg_text,json=msgText,proto3" json:"msg_text,omitempty"`
	UnreadNum     int64                  `protobuf:"varint,9,opt,name=unread_num,json=unreadNum,proto3" json:"unread_num,omitempty"`
	UpdatedAt     string                 `protobuf:"bytes,10,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	IsOnline      bool                   `protobuf:"varint,11,opt,name=is_online,json=isOnline,proto3" json:"is_online,omitempty"`
	IsDisturb     bool                   `protobuf:"varint,12,opt,name=is_disturb,json=isDisturb,proto3" json:"is_disturb,omitempty"`
	IsBot         bool                   `protobuf:"varint,13,opt,name=is_bot,json=isBot,proto3" json:"is_bot,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatItem) Reset() {
	*x = ChatItem{}
	mi := &file_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatItem) ProtoMessage() {}

func (x *ChatItem) ProtoReflect() protoreflect.Message {
	mi := &file_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatItem.ProtoReflect.Descriptor instead.
func (*ChatItem) Descriptor() ([]byte, []int) {
	return file_chat_proto_rawDescGZIP(), []int{2}
}

func (x *ChatItem) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ChatItem) GetChatType() int32 {
	if x != nil {
		return x.ChatType
	}
	return 0
}

func (x *ChatItem) GetReceiverId() int64 {
	if x != nil {
		return x.ReceiverId
	}
	return 0
}

func (x *ChatItem) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *ChatItem) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *ChatItem) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChatItem) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *ChatItem) GetMsgText() string {
	if x != nil {
		return x.MsgText
	}
	return ""
}

func (x *ChatItem) GetUnreadNum() int64 {
	if x != nil {
		return x.UnreadNum
	}
	return 0
}

func (x *ChatItem) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *ChatItem) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

func (x *ChatItem) GetIsDisturb() bool {
	if x != nil {
		return x.IsDisturb
	}
	return false
}

func (x *ChatItem) GetIsBot() bool {
	if x != nil {
		return x.IsBot
	}
	return false
}

var File_chat_proto protoreflect.FileDescriptor

var file_chat_proto_rawDesc = string([]byte{
	0x0a, 0x0a, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x63, 0x68,
	0x61, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x3b, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43,
	0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x24, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x43, 0x68, 0x61, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05,
	0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xe6, 0x02, 0x0a, 0x08, 0x43, 0x68, 0x61, 0x74, 0x49, 0x74,
	0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x68, 0x61, 0x74, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x63, 0x68, 0x61, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x72, 0x65, 0x63, 0x65, 0x69, 0x76, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x6d, 0x73, 0x67, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x73, 0x67, 0x54, 0x65, 0x78, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x09, 0x75, 0x6e, 0x72, 0x65, 0x61, 0x64, 0x4e, 0x75, 0x6d, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69,
	0x73, 0x5f, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08,
	0x69, 0x73, 0x4f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64,
	0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73,
	0x44, 0x69, 0x73, 0x74, 0x75, 0x72, 0x62, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f, 0x62, 0x6f,
	0x74, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x42, 0x6f, 0x74, 0x32, 0x4c,
	0x0a, 0x0b, 0x43, 0x68, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a,
	0x04, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x68, 0x61, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x19, 0x2e, 0x63, 0x68, 0x61, 0x74, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x68, 0x61, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09,
	0x2e, 0x2f, 0x63, 0x68, 0x61, 0x74, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
})

var (
	file_chat_proto_rawDescOnce sync.Once
	file_chat_proto_rawDescData []byte
)

func file_chat_proto_rawDescGZIP() []byte {
	file_chat_proto_rawDescOnce.Do(func() {
		file_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chat_proto_rawDesc), len(file_chat_proto_rawDesc)))
	})
	return file_chat_proto_rawDescData
}

var file_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chat_proto_goTypes = []any{
	(*GetChatListRequest)(nil),  // 0: chat.GetChatListRequest
	(*GetChatListResponse)(nil), // 1: chat.GetChatListResponse
	(*ChatItem)(nil),            // 2: chat.ChatItem
}
var file_chat_proto_depIdxs = []int32{
	2, // 0: chat.GetChatListResponse.items:type_name -> chat.ChatItem
	0, // 1: chat.ChatService.List:input_type -> chat.GetChatListRequest
	1, // 2: chat.ChatService.List:output_type -> chat.GetChatListResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_chat_proto_init() }
func file_chat_proto_init() {
	if File_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chat_proto_rawDesc), len(file_chat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_proto_goTypes,
		DependencyIndexes: file_chat_proto_depIdxs,
		MessageInfos:      file_chat_proto_msgTypes,
	}.Build()
	File_chat_proto = out.File
	file_chat_proto_goTypes = nil
	file_chat_proto_depIdxs = nil
}
