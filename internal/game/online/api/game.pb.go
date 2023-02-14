// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: game.proto

package api

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

type JoinRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ip string `protobuf:"bytes,1,opt,name=ip,proto3" json:"ip,omitempty"`
}

func (x *JoinRequest) Reset() {
	*x = JoinRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRequest) ProtoMessage() {}

func (x *JoinRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRequest.ProtoReflect.Descriptor instead.
func (*JoinRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{0}
}

func (x *JoinRequest) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

type JoinResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success      bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	ClientPlayer string `protobuf:"bytes,2,opt,name=clientPlayer,proto3" json:"clientPlayer,omitempty"`
}

func (x *JoinResponse) Reset() {
	*x = JoinResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinResponse) ProtoMessage() {}

func (x *JoinResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinResponse.ProtoReflect.Descriptor instead.
func (*JoinResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{1}
}

func (x *JoinResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *JoinResponse) GetClientPlayer() string {
	if x != nil {
		return x.ClientPlayer
	}
	return ""
}

type ServerMoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ServerMoveRequest) Reset() {
	*x = ServerMoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMoveRequest) ProtoMessage() {}

func (x *ServerMoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMoveRequest.ProtoReflect.Descriptor instead.
func (*ServerMoveRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{2}
}

type ServerMoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position       int32    `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
	Board          []string `protobuf:"bytes,2,rep,name=board,proto3" json:"board,omitempty"`
	CurrentPlayer  string   `protobuf:"bytes,3,opt,name=currentPlayer,proto3" json:"currentPlayer,omitempty"`
	IsGameFinished bool     `protobuf:"varint,4,opt,name=isGameFinished,proto3" json:"isGameFinished,omitempty"`
}

func (x *ServerMoveResponse) Reset() {
	*x = ServerMoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ServerMoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServerMoveResponse) ProtoMessage() {}

func (x *ServerMoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServerMoveResponse.ProtoReflect.Descriptor instead.
func (*ServerMoveResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{3}
}

func (x *ServerMoveResponse) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

func (x *ServerMoveResponse) GetBoard() []string {
	if x != nil {
		return x.Board
	}
	return nil
}

func (x *ServerMoveResponse) GetCurrentPlayer() string {
	if x != nil {
		return x.CurrentPlayer
	}
	return ""
}

func (x *ServerMoveResponse) GetIsGameFinished() bool {
	if x != nil {
		return x.IsGameFinished
	}
	return false
}

type ClientMoveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Position int32 `protobuf:"varint,1,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *ClientMoveRequest) Reset() {
	*x = ClientMoveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMoveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMoveRequest) ProtoMessage() {}

func (x *ClientMoveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMoveRequest.ProtoReflect.Descriptor instead.
func (*ClientMoveRequest) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{4}
}

func (x *ClientMoveRequest) GetPosition() int32 {
	if x != nil {
		return x.Position
	}
	return 0
}

type ClientMoveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success        bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Board          []string `protobuf:"bytes,2,rep,name=board,proto3" json:"board,omitempty"`
	CurrentPlayer  string   `protobuf:"bytes,3,opt,name=currentPlayer,proto3" json:"currentPlayer,omitempty"`
	IsGameFinished bool     `protobuf:"varint,4,opt,name=isGameFinished,proto3" json:"isGameFinished,omitempty"`
}

func (x *ClientMoveResponse) Reset() {
	*x = ClientMoveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_game_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClientMoveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClientMoveResponse) ProtoMessage() {}

func (x *ClientMoveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_game_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClientMoveResponse.ProtoReflect.Descriptor instead.
func (*ClientMoveResponse) Descriptor() ([]byte, []int) {
	return file_game_proto_rawDescGZIP(), []int{5}
}

func (x *ClientMoveResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ClientMoveResponse) GetBoard() []string {
	if x != nil {
		return x.Board
	}
	return nil
}

func (x *ClientMoveResponse) GetCurrentPlayer() string {
	if x != nil {
		return x.CurrentPlayer
	}
	return ""
}

func (x *ClientMoveResponse) GetIsGameFinished() bool {
	if x != nil {
		return x.IsGameFinished
	}
	return false
}

var File_game_proto protoreflect.FileDescriptor

var file_game_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x22, 0x1d, 0x0a, 0x0b, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x70, 0x22, 0x4c, 0x0a, 0x0c, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12, 0x22, 0x0a,
	0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x22, 0x13, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x94, 0x01, 0x0a, 0x12, 0x53, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61,
	0x72, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x12,
	0x24, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x69, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x46,
	0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0e, 0x69,
	0x73, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x22, 0x2f, 0x0a,
	0x11, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x92,
	0x01, 0x0a, 0x12, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x6f, 0x61, 0x72, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x26, 0x0a, 0x0e, 0x69,
	0x73, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73, 0x68, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x0e, 0x69, 0x73, 0x47, 0x61, 0x6d, 0x65, 0x46, 0x69, 0x6e, 0x69, 0x73,
	0x68, 0x65, 0x64, 0x32, 0xcc, 0x01, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x4a, 0x6f, 0x69, 0x6e, 0x12, 0x13, 0x2e, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x14, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x4d, 0x6f, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x1a, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4d,
	0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30, 0x01, 0x12, 0x43, 0x0a,
	0x0a, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x12, 0x19, 0x2e, 0x6f, 0x6e,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2e,
	0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4d, 0x6f, 0x76, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0c, 0x5a, 0x0a, 0x6f, 0x6e, 0x6c, 0x69, 0x6e, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_game_proto_rawDescOnce sync.Once
	file_game_proto_rawDescData = file_game_proto_rawDesc
)

func file_game_proto_rawDescGZIP() []byte {
	file_game_proto_rawDescOnce.Do(func() {
		file_game_proto_rawDescData = protoimpl.X.CompressGZIP(file_game_proto_rawDescData)
	})
	return file_game_proto_rawDescData
}

var file_game_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_game_proto_goTypes = []interface{}{
	(*JoinRequest)(nil),        // 0: online.JoinRequest
	(*JoinResponse)(nil),       // 1: online.JoinResponse
	(*ServerMoveRequest)(nil),  // 2: online.ServerMoveRequest
	(*ServerMoveResponse)(nil), // 3: online.ServerMoveResponse
	(*ClientMoveRequest)(nil),  // 4: online.ClientMoveRequest
	(*ClientMoveResponse)(nil), // 5: online.ClientMoveResponse
}
var file_game_proto_depIdxs = []int32{
	0, // 0: online.GameService.Join:input_type -> online.JoinRequest
	2, // 1: online.GameService.ServerMove:input_type -> online.ServerMoveRequest
	4, // 2: online.GameService.ClientMove:input_type -> online.ClientMoveRequest
	1, // 3: online.GameService.Join:output_type -> online.JoinResponse
	3, // 4: online.GameService.ServerMove:output_type -> online.ServerMoveResponse
	5, // 5: online.GameService.ClientMove:output_type -> online.ClientMoveResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_game_proto_init() }
func file_game_proto_init() {
	if File_game_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_game_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRequest); i {
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
		file_game_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinResponse); i {
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
		file_game_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMoveRequest); i {
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
		file_game_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ServerMoveResponse); i {
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
		file_game_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMoveRequest); i {
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
		file_game_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClientMoveResponse); i {
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
			RawDescriptor: file_game_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_game_proto_goTypes,
		DependencyIndexes: file_game_proto_depIdxs,
		MessageInfos:      file_game_proto_msgTypes,
	}.Build()
	File_game_proto = out.File
	file_game_proto_rawDesc = nil
	file_game_proto_goTypes = nil
	file_game_proto_depIdxs = nil
}
