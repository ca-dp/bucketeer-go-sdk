// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: proto/feature/feature_last_used_info.proto

package feature

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

type FeatureLastUsedInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FeatureId           string `protobuf:"bytes,1,opt,name=feature_id,json=featureId,proto3" json:"feature_id,omitempty"`
	Version             int32  `protobuf:"varint,2,opt,name=version,proto3" json:"version,omitempty"`
	LastUsedAt          int64  `protobuf:"varint,3,opt,name=last_used_at,json=lastUsedAt,proto3" json:"last_used_at,omitempty"`
	CreatedAt           int64  `protobuf:"varint,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ClientOldestVersion string `protobuf:"bytes,5,opt,name=client_oldest_version,json=clientOldestVersion,proto3" json:"client_oldest_version,omitempty"`
	ClientLatestVersion string `protobuf:"bytes,6,opt,name=client_latest_version,json=clientLatestVersion,proto3" json:"client_latest_version,omitempty"`
}

func (x *FeatureLastUsedInfo) Reset() {
	*x = FeatureLastUsedInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_feature_feature_last_used_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FeatureLastUsedInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FeatureLastUsedInfo) ProtoMessage() {}

func (x *FeatureLastUsedInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_feature_feature_last_used_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FeatureLastUsedInfo.ProtoReflect.Descriptor instead.
func (*FeatureLastUsedInfo) Descriptor() ([]byte, []int) {
	return file_proto_feature_feature_last_used_info_proto_rawDescGZIP(), []int{0}
}

func (x *FeatureLastUsedInfo) GetFeatureId() string {
	if x != nil {
		return x.FeatureId
	}
	return ""
}

func (x *FeatureLastUsedInfo) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *FeatureLastUsedInfo) GetLastUsedAt() int64 {
	if x != nil {
		return x.LastUsedAt
	}
	return 0
}

func (x *FeatureLastUsedInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *FeatureLastUsedInfo) GetClientOldestVersion() string {
	if x != nil {
		return x.ClientOldestVersion
	}
	return ""
}

func (x *FeatureLastUsedInfo) GetClientLatestVersion() string {
	if x != nil {
		return x.ClientLatestVersion
	}
	return ""
}

var File_proto_feature_feature_last_used_info_proto protoreflect.FileDescriptor

var file_proto_feature_feature_last_used_info_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2f,
	0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x5f, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65,
	0x64, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2e, 0x66, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22,
	0xf7, 0x01, 0x0a, 0x13, 0x46, 0x65, 0x61, 0x74, 0x75, 0x72, 0x65, 0x4c, 0x61, 0x73, 0x74, 0x55,
	0x73, 0x65, 0x64, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x65, 0x61, 0x74, 0x75,
	0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x65, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x12, 0x20, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x75, 0x73, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x6c, 0x61, 0x73, 0x74, 0x55, 0x73, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x6f, 0x6c, 0x64, 0x65,
	0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4f, 0x6c, 0x64, 0x65, 0x73, 0x74, 0x56, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x4c, 0x61, 0x74, 0x65,
	0x73, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x2a, 0x5a, 0x28, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x61, 0x2d, 0x64, 0x70, 0x2f, 0x62, 0x75,
	0x63, 0x6b, 0x65, 0x74, 0x65, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x65,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_feature_feature_last_used_info_proto_rawDescOnce sync.Once
	file_proto_feature_feature_last_used_info_proto_rawDescData = file_proto_feature_feature_last_used_info_proto_rawDesc
)

func file_proto_feature_feature_last_used_info_proto_rawDescGZIP() []byte {
	file_proto_feature_feature_last_used_info_proto_rawDescOnce.Do(func() {
		file_proto_feature_feature_last_used_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_feature_feature_last_used_info_proto_rawDescData)
	})
	return file_proto_feature_feature_last_used_info_proto_rawDescData
}

var file_proto_feature_feature_last_used_info_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_proto_feature_feature_last_used_info_proto_goTypes = []interface{}{
	(*FeatureLastUsedInfo)(nil), // 0: bucketeer.feature.FeatureLastUsedInfo
}
var file_proto_feature_feature_last_used_info_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_feature_feature_last_used_info_proto_init() }
func file_proto_feature_feature_last_used_info_proto_init() {
	if File_proto_feature_feature_last_used_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_feature_feature_last_used_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FeatureLastUsedInfo); i {
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
			RawDescriptor: file_proto_feature_feature_last_used_info_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_feature_feature_last_used_info_proto_goTypes,
		DependencyIndexes: file_proto_feature_feature_last_used_info_proto_depIdxs,
		MessageInfos:      file_proto_feature_feature_last_used_info_proto_msgTypes,
	}.Build()
	File_proto_feature_feature_last_used_info_proto = out.File
	file_proto_feature_feature_last_used_info_proto_rawDesc = nil
	file_proto_feature_feature_last_used_info_proto_goTypes = nil
	file_proto_feature_feature_last_used_info_proto_depIdxs = nil
}
