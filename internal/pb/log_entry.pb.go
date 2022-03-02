// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: internal/pb/protobuf-spec/log_entry.proto

package pb

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

type LogEntry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MethodName   string `protobuf:"bytes,1,opt,name=methodName,proto3" json:"methodName,omitempty"`
	RequestBytes []byte `protobuf:"bytes,2,opt,name=requestBytes,proto3" json:"requestBytes,omitempty"`
}

func (x *LogEntry) Reset() {
	*x = LogEntry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_protobuf_spec_log_entry_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogEntry) ProtoMessage() {}

func (x *LogEntry) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_protobuf_spec_log_entry_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogEntry.ProtoReflect.Descriptor instead.
func (*LogEntry) Descriptor() ([]byte, []int) {
	return file_internal_pb_protobuf_spec_log_entry_proto_rawDescGZIP(), []int{0}
}

func (x *LogEntry) GetMethodName() string {
	if x != nil {
		return x.MethodName
	}
	return ""
}

func (x *LogEntry) GetRequestBytes() []byte {
	if x != nil {
		return x.RequestBytes
	}
	return nil
}

var File_internal_pb_protobuf_spec_log_entry_proto protoreflect.FileDescriptor

var file_internal_pb_protobuf_spec_log_entry_proto_rawDesc = []byte{
	0x0a, 0x29, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2d, 0x73, 0x70, 0x65, 0x63, 0x2f, 0x6c, 0x6f, 0x67, 0x5f,
	0x65, 0x6e, 0x74, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x4e, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1e,
	0x0a, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x74, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x79, 0x74,
	0x65, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70,
	0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_protobuf_spec_log_entry_proto_rawDescOnce sync.Once
	file_internal_pb_protobuf_spec_log_entry_proto_rawDescData = file_internal_pb_protobuf_spec_log_entry_proto_rawDesc
)

func file_internal_pb_protobuf_spec_log_entry_proto_rawDescGZIP() []byte {
	file_internal_pb_protobuf_spec_log_entry_proto_rawDescOnce.Do(func() {
		file_internal_pb_protobuf_spec_log_entry_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_protobuf_spec_log_entry_proto_rawDescData)
	})
	return file_internal_pb_protobuf_spec_log_entry_proto_rawDescData
}

var file_internal_pb_protobuf_spec_log_entry_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_internal_pb_protobuf_spec_log_entry_proto_goTypes = []interface{}{
	(*LogEntry)(nil), // 0: proto.LogEntry
}
var file_internal_pb_protobuf_spec_log_entry_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_pb_protobuf_spec_log_entry_proto_init() }
func file_internal_pb_protobuf_spec_log_entry_proto_init() {
	if File_internal_pb_protobuf_spec_log_entry_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_protobuf_spec_log_entry_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogEntry); i {
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
			RawDescriptor: file_internal_pb_protobuf_spec_log_entry_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_pb_protobuf_spec_log_entry_proto_goTypes,
		DependencyIndexes: file_internal_pb_protobuf_spec_log_entry_proto_depIdxs,
		MessageInfos:      file_internal_pb_protobuf_spec_log_entry_proto_msgTypes,
	}.Build()
	File_internal_pb_protobuf_spec_log_entry_proto = out.File
	file_internal_pb_protobuf_spec_log_entry_proto_rawDesc = nil
	file_internal_pb_protobuf_spec_log_entry_proto_goTypes = nil
	file_internal_pb_protobuf_spec_log_entry_proto_depIdxs = nil
}