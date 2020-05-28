// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.22.0
// 	protoc        v3.11.4
// source: staff/staff.proto

package staff

import (
	common "github.com/cargod-bj/b2c-ptoto/common"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginName string `protobuf:"bytes,2,opt,name=LoginName,proto3" json:"LoginName,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Icon      string `protobuf:"bytes,4,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Password  string `protobuf:"bytes,5,opt,name=Password,proto3" json:"Password,omitempty"`
	Gender    uint64 `protobuf:"varint,6,opt,name=Gender,proto3" json:"Gender,omitempty"`
	Birthday  string `protobuf:"bytes,7,opt,name=Birthday,proto3" json:"Birthday,omitempty"`
	Telephone string `protobuf:"bytes,10,opt,name=Telephone,proto3" json:"Telephone,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_staff_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_staff_staff_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_staff_staff_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *User) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *User) GetGender() uint64 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *User) GetBirthday() string {
	if x != nil {
		return x.Birthday
	}
	return ""
}

func (x *User) GetTelephone() string {
	if x != nil {
		return x.Telephone
	}
	return ""
}

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LoginName string `protobuf:"bytes,1,opt,name=LoginName,proto3" json:"LoginName,omitempty"`
	Password  string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_staff_staff_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_staff_staff_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_staff_staff_proto_rawDescGZIP(), []int{1}
}

func (x *Request) GetLoginName() string {
	if x != nil {
		return x.LoginName
	}
	return ""
}

func (x *Request) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

var File_staff_staff_proto protoreflect.FileDescriptor

var file_staff_staff_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x74, 0x61, 0x66, 0x66, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xbc, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x49, 0x63, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x49, 0x63, 0x6f, 0x6e,
	0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x47, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x47, 0x65,
	0x6e, 0x64, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x42, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79,
	0x12, 0x1c, 0x0a, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x54, 0x65, 0x6c, 0x65, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x43,
	0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x4c, 0x6f,
	0x67, 0x69, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77,
	0x6f, 0x72, 0x64, 0x32, 0x61, 0x0a, 0x05, 0x53, 0x74, 0x61, 0x66, 0x66, 0x12, 0x2b, 0x0a, 0x05,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x2b, 0x0a, 0x08, 0x52, 0x65, 0x67,
	0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x73, 0x74, 0x61, 0x66, 0x66, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x1a, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x26, 0x5a, 0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x72, 0x67, 0x6f, 0x64, 0x2d, 0x62, 0x6a, 0x2f, 0x62,
	0x32, 0x63, 0x2d, 0x70, 0x74, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x66, 0x66, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_staff_staff_proto_rawDescOnce sync.Once
	file_staff_staff_proto_rawDescData = file_staff_staff_proto_rawDesc
)

func file_staff_staff_proto_rawDescGZIP() []byte {
	file_staff_staff_proto_rawDescOnce.Do(func() {
		file_staff_staff_proto_rawDescData = protoimpl.X.CompressGZIP(file_staff_staff_proto_rawDescData)
	})
	return file_staff_staff_proto_rawDescData
}

var file_staff_staff_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_staff_staff_proto_goTypes = []interface{}{
	(*User)(nil),            // 0: staff.User
	(*Request)(nil),         // 1: staff.Request
	(*common.Response)(nil), // 2: common.Response
}
var file_staff_staff_proto_depIdxs = []int32{
	1, // 0: staff.Staff.Login:input_type -> staff.Request
	0, // 1: staff.Staff.Register:input_type -> staff.User
	2, // 2: staff.Staff.Login:output_type -> common.Response
	2, // 3: staff.Staff.Register:output_type -> common.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_staff_staff_proto_init() }
func file_staff_staff_proto_init() {
	if File_staff_staff_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_staff_staff_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_staff_staff_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
			RawDescriptor: file_staff_staff_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_staff_staff_proto_goTypes,
		DependencyIndexes: file_staff_staff_proto_depIdxs,
		MessageInfos:      file_staff_staff_proto_msgTypes,
	}.Build()
	File_staff_staff_proto = out.File
	file_staff_staff_proto_rawDesc = nil
	file_staff_staff_proto_goTypes = nil
	file_staff_staff_proto_depIdxs = nil
}
