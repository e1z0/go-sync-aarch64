// Copyright (c) 2020 The Brave Authors. All rights reserved.
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this file,
// You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: brave_device_info_specifics.proto

package sync_pb

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

type BraveSpecificFields struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsSelfDeleteSupported *bool `protobuf:"varint,1,opt,name=is_self_delete_supported,json=isSelfDeleteSupported" json:"is_self_delete_supported,omitempty"`
}

func (x *BraveSpecificFields) Reset() {
	*x = BraveSpecificFields{}
	if protoimpl.UnsafeEnabled {
		mi := &file_brave_device_info_specifics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BraveSpecificFields) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BraveSpecificFields) ProtoMessage() {}

func (x *BraveSpecificFields) ProtoReflect() protoreflect.Message {
	mi := &file_brave_device_info_specifics_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BraveSpecificFields.ProtoReflect.Descriptor instead.
func (*BraveSpecificFields) Descriptor() ([]byte, []int) {
	return file_brave_device_info_specifics_proto_rawDescGZIP(), []int{0}
}

func (x *BraveSpecificFields) GetIsSelfDeleteSupported() bool {
	if x != nil && x.IsSelfDeleteSupported != nil {
		return *x.IsSelfDeleteSupported
	}
	return false
}

var File_brave_device_info_specifics_proto protoreflect.FileDescriptor

var file_brave_device_info_specifics_proto_rawDesc = []byte{
	0x0a, 0x21, 0x62, 0x72, 0x61, 0x76, 0x65, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x69,
	0x6e, 0x66, 0x6f, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x4e, 0x0a, 0x13,
	0x42, 0x72, 0x61, 0x76, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x46, 0x69, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x37, 0x0a, 0x18, 0x69, 0x73, 0x5f, 0x73, 0x65, 0x6c, 0x66, 0x5f, 0x64,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x69, 0x73, 0x53, 0x65, 0x6c, 0x66, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x42, 0x2b, 0x0a, 0x25,
	0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d,
	0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_brave_device_info_specifics_proto_rawDescOnce sync.Once
	file_brave_device_info_specifics_proto_rawDescData = file_brave_device_info_specifics_proto_rawDesc
)

func file_brave_device_info_specifics_proto_rawDescGZIP() []byte {
	file_brave_device_info_specifics_proto_rawDescOnce.Do(func() {
		file_brave_device_info_specifics_proto_rawDescData = protoimpl.X.CompressGZIP(file_brave_device_info_specifics_proto_rawDescData)
	})
	return file_brave_device_info_specifics_proto_rawDescData
}

var file_brave_device_info_specifics_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_brave_device_info_specifics_proto_goTypes = []interface{}{
	(*BraveSpecificFields)(nil), // 0: sync_pb.BraveSpecificFields
}
var file_brave_device_info_specifics_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_brave_device_info_specifics_proto_init() }
func file_brave_device_info_specifics_proto_init() {
	if File_brave_device_info_specifics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_brave_device_info_specifics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BraveSpecificFields); i {
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
			RawDescriptor: file_brave_device_info_specifics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_brave_device_info_specifics_proto_goTypes,
		DependencyIndexes: file_brave_device_info_specifics_proto_depIdxs,
		MessageInfos:      file_brave_device_info_specifics_proto_msgTypes,
	}.Build()
	File_brave_device_info_specifics_proto = out.File
	file_brave_device_info_specifics_proto_rawDesc = nil
	file_brave_device_info_specifics_proto_goTypes = nil
	file_brave_device_info_specifics_proto_depIdxs = nil
}
