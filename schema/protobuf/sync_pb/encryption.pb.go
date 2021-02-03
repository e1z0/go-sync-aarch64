// Copyright (c) 2012 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.
//
// Common sync protocol for encrypted data.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.24.0
// 	protoc        v3.12.2
// source: encryption.proto

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

// Encrypted sync data consists of two parts: a key name and a blob. Key name is
// the name of the key that was used to encrypt blob and blob is encrypted data
// itself.
//
// The reason we need to keep track of the key name is that a sync user can
// change their passphrase (and thus their encryption key) at any time. When
// that happens, we make a best effort to reencrypt all nodes with the new
// passphrase, but since we don't have transactions on the server-side, we
// cannot guarantee that every node will be reencrypted. As a workaround, we
// keep track of all keys, assign each key a name (by using that key to encrypt
// a well known string) and keep track of which key was used to encrypt each
// node.
type EncryptedData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	KeyName *string `protobuf:"bytes,1,opt,name=key_name,json=keyName" json:"key_name,omitempty"`
	Blob    *string `protobuf:"bytes,2,opt,name=blob" json:"blob,omitempty"` // base64-encoded.
}

func (x *EncryptedData) Reset() {
	*x = EncryptedData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encryption_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncryptedData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncryptedData) ProtoMessage() {}

func (x *EncryptedData) ProtoReflect() protoreflect.Message {
	mi := &file_encryption_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncryptedData.ProtoReflect.Descriptor instead.
func (*EncryptedData) Descriptor() ([]byte, []int) {
	return file_encryption_proto_rawDescGZIP(), []int{0}
}

func (x *EncryptedData) GetKeyName() string {
	if x != nil && x.KeyName != nil {
		return *x.KeyName
	}
	return ""
}

func (x *EncryptedData) GetBlob() string {
	if x != nil && x.Blob != nil {
		return *x.Blob
	}
	return ""
}

var File_encryption_proto protoreflect.FileDescriptor

var file_encryption_proto_rawDesc = []byte{
	0x0a, 0x10, 0x65, 0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x22, 0x3e, 0x0a, 0x0d, 0x45,
	0x6e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x65, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x19, 0x0a, 0x08,
	0x6b, 0x65, 0x79, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6b, 0x65, 0x79, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6c, 0x6f, 0x62, 0x42, 0x2b, 0x0a, 0x25, 0x6f,
	0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70,
	0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01,
}

var (
	file_encryption_proto_rawDescOnce sync.Once
	file_encryption_proto_rawDescData = file_encryption_proto_rawDesc
)

func file_encryption_proto_rawDescGZIP() []byte {
	file_encryption_proto_rawDescOnce.Do(func() {
		file_encryption_proto_rawDescData = protoimpl.X.CompressGZIP(file_encryption_proto_rawDescData)
	})
	return file_encryption_proto_rawDescData
}

var file_encryption_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_encryption_proto_goTypes = []interface{}{
	(*EncryptedData)(nil), // 0: sync_pb.EncryptedData
}
var file_encryption_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_encryption_proto_init() }
func file_encryption_proto_init() {
	if File_encryption_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_encryption_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncryptedData); i {
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
			RawDescriptor: file_encryption_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encryption_proto_goTypes,
		DependencyIndexes: file_encryption_proto_depIdxs,
		MessageInfos:      file_encryption_proto_msgTypes,
	}.Build()
	File_encryption_proto = out.File
	file_encryption_proto_rawDesc = nil
	file_encryption_proto_goTypes = nil
	file_encryption_proto_depIdxs = nil
}
