// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: proto/task.proto

package task

import (
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

type Record struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`     // task id
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` //task name
	Desc    string `protobuf:"bytes,3,opt,name=desc,proto3" json:"desc,omitempty"`
	Created int64  `protobuf:"varint,4,opt,name=created,proto3" json:"created,omitempty"`
	Updated int64  `protobuf:"varint,5,opt,name=updated,proto3" json:"updated,omitempty"`
}

func (x *Record) Reset() {
	*x = Record{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Record) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Record) ProtoMessage() {}

func (x *Record) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Record.ProtoReflect.Descriptor instead.
func (*Record) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{0}
}

func (x *Record) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Record) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Record) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

func (x *Record) GetCreated() int64 {
	if x != nil {
		return x.Created
	}
	return 0
}

func (x *Record) GetUpdated() int64 {
	if x != nil {
		return x.Updated
	}
	return 0
}

type RecordRequst struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Desc string `protobuf:"bytes,2,opt,name=desc,proto3" json:"desc,omitempty"`
}

func (x *RecordRequst) Reset() {
	*x = RecordRequst{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordRequst) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordRequst) ProtoMessage() {}

func (x *RecordRequst) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordRequst.ProtoReflect.Descriptor instead.
func (*RecordRequst) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{1}
}

func (x *RecordRequst) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *RecordRequst) GetDesc() string {
	if x != nil {
		return x.Desc
	}
	return ""
}

type RecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RecordResponse) Reset() {
	*x = RecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RecordResponse) ProtoMessage() {}

func (x *RecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RecordResponse.ProtoReflect.Descriptor instead.
func (*RecordResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{2}
}

var File_proto_task_proto protoreflect.FileDescriptor

var file_proto_task_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x74, 0x61, 0x73, 0x6b, 0x22, 0x74, 0x0a, 0x06, 0x52, 0x65, 0x63, 0x6f,
	0x72, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x22, 0x36,
	0x0a, 0x0c, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x73, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x65, 0x73, 0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x64, 0x65, 0x73, 0x63, 0x22, 0x10, 0x0a, 0x0e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x40, 0x0a, 0x04, 0x54, 0x61, 0x73, 0x6b,
	0x12, 0x38, 0x0a, 0x0a, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x12,
	0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x73, 0x74, 0x1a, 0x14, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x3b, 0x74, 0x61, 0x73, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_task_proto_rawDescOnce sync.Once
	file_proto_task_proto_rawDescData = file_proto_task_proto_rawDesc
)

func file_proto_task_proto_rawDescGZIP() []byte {
	file_proto_task_proto_rawDescOnce.Do(func() {
		file_proto_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_task_proto_rawDescData)
	})
	return file_proto_task_proto_rawDescData
}

var file_proto_task_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_task_proto_goTypes = []interface{}{
	(*Record)(nil),         // 0: task.Record
	(*RecordRequst)(nil),   // 1: task.RecordRequst
	(*RecordResponse)(nil), // 2: task.RecordResponse
}
var file_proto_task_proto_depIdxs = []int32{
	1, // 0: task.Task.RecordTask:input_type -> task.RecordRequst
	2, // 1: task.Task.RecordTask:output_type -> task.RecordResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_task_proto_init() }
func file_proto_task_proto_init() {
	if File_proto_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_task_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Record); i {
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
		file_proto_task_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordRequst); i {
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
		file_proto_task_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RecordResponse); i {
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
			RawDescriptor: file_proto_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_task_proto_goTypes,
		DependencyIndexes: file_proto_task_proto_depIdxs,
		MessageInfos:      file_proto_task_proto_msgTypes,
	}.Build()
	File_proto_task_proto = out.File
	file_proto_task_proto_rawDesc = nil
	file_proto_task_proto_goTypes = nil
	file_proto_task_proto_depIdxs = nil
}
