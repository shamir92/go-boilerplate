// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.24.3
// source: protocol/grpc/presenters/proto/member.proto

package compiled

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

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Firstname string `protobuf:"bytes,2,opt,name=firstname,proto3" json:"firstname,omitempty"`
	Lastname  string `protobuf:"bytes,3,opt,name=lastname,proto3" json:"lastname,omitempty"`
	Email     string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_grpc_presenters_proto_member_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_grpc_presenters_proto_member_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_protocol_grpc_presenters_proto_member_proto_rawDescGZIP(), []int{0}
}

func (x *Member) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Member) GetFirstname() string {
	if x != nil {
		return x.Firstname
	}
	return ""
}

func (x *Member) GetLastname() string {
	if x != nil {
		return x.Lastname
	}
	return ""
}

func (x *Member) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type MemberRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member *Member `protobuf:"bytes,1,opt,name=member,proto3" json:"member,omitempty"`
}

func (x *MemberRequest) Reset() {
	*x = MemberRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_grpc_presenters_proto_member_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberRequest) ProtoMessage() {}

func (x *MemberRequest) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_grpc_presenters_proto_member_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberRequest.ProtoReflect.Descriptor instead.
func (*MemberRequest) Descriptor() ([]byte, []int) {
	return file_protocol_grpc_presenters_proto_member_proto_rawDescGZIP(), []int{1}
}

func (x *MemberRequest) GetMember() *Member {
	if x != nil {
		return x.Member
	}
	return nil
}

type MemberResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Member []*Member `protobuf:"bytes,1,rep,name=member,proto3" json:"member,omitempty"`
}

func (x *MemberResponse) Reset() {
	*x = MemberResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protocol_grpc_presenters_proto_member_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MemberResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MemberResponse) ProtoMessage() {}

func (x *MemberResponse) ProtoReflect() protoreflect.Message {
	mi := &file_protocol_grpc_presenters_proto_member_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MemberResponse.ProtoReflect.Descriptor instead.
func (*MemberResponse) Descriptor() ([]byte, []int) {
	return file_protocol_grpc_presenters_proto_member_proto_rawDescGZIP(), []int{2}
}

func (x *MemberResponse) GetMember() []*Member {
	if x != nil {
		return x.Member
	}
	return nil
}

var File_protocol_grpc_presenters_proto_member_proto protoreflect.FileDescriptor

var file_protocol_grpc_presenters_proto_member_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67,
	0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x68, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x1c, 0x0a, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69,
	0x6c, 0x22, 0x39, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x28, 0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x22, 0x3a, 0x0a, 0x0e,
	0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x28,
	0x0a, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x06, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x32, 0x4e, 0x0a, 0x0d, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3d, 0x0a, 0x08, 0x46, 0x65, 0x74,
	0x63, 0x68, 0x41, 0x6c, 0x6c, 0x12, 0x17, 0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18,
	0x2e, 0x67, 0x6f, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e,
	0x74, 0x65, 0x72, 0x73, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x69, 0x6c, 0x65, 0x64, 0x2f, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protocol_grpc_presenters_proto_member_proto_rawDescOnce sync.Once
	file_protocol_grpc_presenters_proto_member_proto_rawDescData = file_protocol_grpc_presenters_proto_member_proto_rawDesc
)

func file_protocol_grpc_presenters_proto_member_proto_rawDescGZIP() []byte {
	file_protocol_grpc_presenters_proto_member_proto_rawDescOnce.Do(func() {
		file_protocol_grpc_presenters_proto_member_proto_rawDescData = protoimpl.X.CompressGZIP(file_protocol_grpc_presenters_proto_member_proto_rawDescData)
	})
	return file_protocol_grpc_presenters_proto_member_proto_rawDescData
}

var file_protocol_grpc_presenters_proto_member_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_protocol_grpc_presenters_proto_member_proto_goTypes = []interface{}{
	(*Member)(nil),         // 0: go_proto.Member
	(*MemberRequest)(nil),  // 1: go_proto.MemberRequest
	(*MemberResponse)(nil), // 2: go_proto.MemberResponse
}
var file_protocol_grpc_presenters_proto_member_proto_depIdxs = []int32{
	0, // 0: go_proto.MemberRequest.member:type_name -> go_proto.Member
	0, // 1: go_proto.MemberResponse.member:type_name -> go_proto.Member
	1, // 2: go_proto.MemberService.FetchAll:input_type -> go_proto.MemberRequest
	2, // 3: go_proto.MemberService.FetchAll:output_type -> go_proto.MemberResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protocol_grpc_presenters_proto_member_proto_init() }
func file_protocol_grpc_presenters_proto_member_proto_init() {
	if File_protocol_grpc_presenters_proto_member_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protocol_grpc_presenters_proto_member_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_protocol_grpc_presenters_proto_member_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberRequest); i {
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
		file_protocol_grpc_presenters_proto_member_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MemberResponse); i {
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
			RawDescriptor: file_protocol_grpc_presenters_proto_member_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protocol_grpc_presenters_proto_member_proto_goTypes,
		DependencyIndexes: file_protocol_grpc_presenters_proto_member_proto_depIdxs,
		MessageInfos:      file_protocol_grpc_presenters_proto_member_proto_msgTypes,
	}.Build()
	File_protocol_grpc_presenters_proto_member_proto = out.File
	file_protocol_grpc_presenters_proto_member_proto_rawDesc = nil
	file_protocol_grpc_presenters_proto_member_proto_goTypes = nil
	file_protocol_grpc_presenters_proto_member_proto_depIdxs = nil
}
