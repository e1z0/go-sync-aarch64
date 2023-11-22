// Copyright 2015 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

// If you change or add any fields in this file, update proto_visitors.h and
// potentially proto_enum_conversions.{h, cc}.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.1
// source: entity_metadata.proto

package sync_pb

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

// Sync proto to store entity metadata in model type storage.
type EntityMetadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// A hash based on the client tag and model type.
	// Used for various map lookups. Should always be available.
	// Sent to the server as SyncEntity::client_tag_hash.
	ClientTagHash *string `protobuf:"bytes,1,opt,name=client_tag_hash,json=clientTagHash" json:"client_tag_hash,omitempty"`
	// The entity's server-assigned ID.
	//
	// Prior to the item's first commit, we leave this value as an empty string.
	// The initial ID for a newly created item has to meet certain uniqueness
	// requirements, and we handle those on the sync thread.
	ServerId *string `protobuf:"bytes,2,opt,name=server_id,json=serverId" json:"server_id,omitempty"`
	// Whether or not the entity is deleted.
	IsDeleted *bool `protobuf:"varint,3,opt,name=is_deleted,json=isDeleted" json:"is_deleted,omitempty"`
	// A version number used to track in-progress commits. Each local change
	// increments this number.
	SequenceNumber *int64 `protobuf:"varint,4,opt,name=sequence_number,json=sequenceNumber" json:"sequence_number,omitempty"`
	// The sequence number of the last item known to be successfully committed.
	AckedSequenceNumber *int64 `protobuf:"varint,5,opt,name=acked_sequence_number,json=ackedSequenceNumber" json:"acked_sequence_number,omitempty"`
	// The server version on which this item is based.
	//
	// If there are no local changes, this is the version of the entity as we see
	// it here.
	//
	// If there are local changes, this is the version of the entity on which
	// those changes are based.
	ServerVersion *int64 `protobuf:"varint,6,opt,name=server_version,json=serverVersion,def=-1" json:"server_version,omitempty"`
	// Entity creation and modification timestamps. Assigned by the client and
	// synced by the server, though the server usually doesn't bother to inspect
	// their values. They are encoded as milliseconds since the Unix epoch.
	CreationTime     *int64 `protobuf:"varint,7,opt,name=creation_time,json=creationTime" json:"creation_time,omitempty"`
	ModificationTime *int64 `protobuf:"varint,8,opt,name=modification_time,json=modificationTime" json:"modification_time,omitempty"`
	// A hash of the current entity specifics value. Used to detect whether
	// entity's specifics value has changed without having to keep specifics in
	// memory.
	SpecificsHash *string `protobuf:"bytes,9,opt,name=specifics_hash,json=specificsHash" json:"specifics_hash,omitempty"`
	// A hash of the last specifics known by both the client and server. Used to
	// detect when local commits and remote updates are just for encryption. This
	// value will be the empty string only in the following cases: the entity is
	// in sync with the server, has never been synced, or is deleted.
	BaseSpecificsHash *string `protobuf:"bytes,10,opt,name=base_specifics_hash,json=baseSpecificsHash" json:"base_specifics_hash,omitempty"`
	// Used for positioning entities among their siblings. Relevant only for data
	// types that support positions (e.g bookmarks). Refer to its definition in
	// unique_position.proto for more information about its internal
	// representation.
	UniquePosition *UniquePosition `protobuf:"bytes,11,opt,name=unique_position,json=uniquePosition" json:"unique_position,omitempty"`
	// Used only for bookmarks. It's analogous to |specifics_hash| but it
	// exclusively hashes the content of the favicon image, as represented in
	// proto field BookmarkSpecifics.favicon, using base::PersistentHash().
	BookmarkFaviconHash *uint32 `protobuf:"fixed32,12,opt,name=bookmark_favicon_hash,json=bookmarkFaviconHash" json:"bookmark_favicon_hash,omitempty"`
	// Last specifics known by both the client and server. Used during commits to
	// the server in order to prevent data loss caused by older clients dealing
	// with unknown proto fields (fields that were introduced later). Datatypes
	// (ModelTypeSyncBridge) may implement logic to trim down (or fully clear)
	// this proto prior to caching, to avoid the memory and I/O overhead of
	// dealing with an extra copy of the data. Introduced in M101.
	PossiblyTrimmedBaseSpecifics *EntitySpecifics `protobuf:"bytes,13,opt,name=possibly_trimmed_base_specifics,json=possiblyTrimmedBaseSpecifics" json:"possibly_trimmed_base_specifics,omitempty"`
}

// Default values for EntityMetadata fields.
const (
	Default_EntityMetadata_ServerVersion = int64(-1)
)

func (x *EntityMetadata) Reset() {
	*x = EntityMetadata{}
	if protoimpl.UnsafeEnabled {
		mi := &file_entity_metadata_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityMetadata) ProtoMessage() {}

func (x *EntityMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_entity_metadata_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityMetadata.ProtoReflect.Descriptor instead.
func (*EntityMetadata) Descriptor() ([]byte, []int) {
	return file_entity_metadata_proto_rawDescGZIP(), []int{0}
}

func (x *EntityMetadata) GetClientTagHash() string {
	if x != nil && x.ClientTagHash != nil {
		return *x.ClientTagHash
	}
	return ""
}

func (x *EntityMetadata) GetServerId() string {
	if x != nil && x.ServerId != nil {
		return *x.ServerId
	}
	return ""
}

func (x *EntityMetadata) GetIsDeleted() bool {
	if x != nil && x.IsDeleted != nil {
		return *x.IsDeleted
	}
	return false
}

func (x *EntityMetadata) GetSequenceNumber() int64 {
	if x != nil && x.SequenceNumber != nil {
		return *x.SequenceNumber
	}
	return 0
}

func (x *EntityMetadata) GetAckedSequenceNumber() int64 {
	if x != nil && x.AckedSequenceNumber != nil {
		return *x.AckedSequenceNumber
	}
	return 0
}

func (x *EntityMetadata) GetServerVersion() int64 {
	if x != nil && x.ServerVersion != nil {
		return *x.ServerVersion
	}
	return Default_EntityMetadata_ServerVersion
}

func (x *EntityMetadata) GetCreationTime() int64 {
	if x != nil && x.CreationTime != nil {
		return *x.CreationTime
	}
	return 0
}

func (x *EntityMetadata) GetModificationTime() int64 {
	if x != nil && x.ModificationTime != nil {
		return *x.ModificationTime
	}
	return 0
}

func (x *EntityMetadata) GetSpecificsHash() string {
	if x != nil && x.SpecificsHash != nil {
		return *x.SpecificsHash
	}
	return ""
}

func (x *EntityMetadata) GetBaseSpecificsHash() string {
	if x != nil && x.BaseSpecificsHash != nil {
		return *x.BaseSpecificsHash
	}
	return ""
}

func (x *EntityMetadata) GetUniquePosition() *UniquePosition {
	if x != nil {
		return x.UniquePosition
	}
	return nil
}

func (x *EntityMetadata) GetBookmarkFaviconHash() uint32 {
	if x != nil && x.BookmarkFaviconHash != nil {
		return *x.BookmarkFaviconHash
	}
	return 0
}

func (x *EntityMetadata) GetPossiblyTrimmedBaseSpecifics() *EntitySpecifics {
	if x != nil {
		return x.PossiblyTrimmedBaseSpecifics
	}
	return nil
}

var File_entity_metadata_proto protoreflect.FileDescriptor

var file_entity_metadata_proto_rawDesc = []byte{
	0x0a, 0x15, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x1a, 0x16, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x75, 0x6e, 0x69, 0x71, 0x75, 0x65,
	0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xfc, 0x04, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x12, 0x26, 0x0a, 0x0f, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x61, 0x67,
	0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x54, 0x61, 0x67, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x73, 0x5f, 0x64, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x73, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x12, 0x27, 0x0a, 0x0f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e,
	0x63, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0e, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x32, 0x0a, 0x15, 0x61, 0x63, 0x6b, 0x65, 0x64, 0x5f, 0x73, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x13,
	0x61, 0x63, 0x6b, 0x65, 0x64, 0x53, 0x65, 0x71, 0x75, 0x65, 0x6e, 0x63, 0x65, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x29, 0x0a, 0x0e, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x5f, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x3a, 0x02, 0x2d, 0x31, 0x52,
	0x0d, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x23,
	0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54,
	0x69, 0x6d, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x10,
	0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x2e, 0x0a, 0x13, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x62, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66,
	0x69, 0x63, 0x73, 0x48, 0x61, 0x73, 0x68, 0x12, 0x40, 0x0a, 0x0f, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x5f, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x55, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x75, 0x6e, 0x69, 0x71, 0x75,
	0x65, 0x50, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x32, 0x0a, 0x15, 0x62, 0x6f, 0x6f,
	0x6b, 0x6d, 0x61, 0x72, 0x6b, 0x5f, 0x66, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x5f, 0x68, 0x61,
	0x73, 0x68, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x07, 0x52, 0x13, 0x62, 0x6f, 0x6f, 0x6b, 0x6d, 0x61,
	0x72, 0x6b, 0x46, 0x61, 0x76, 0x69, 0x63, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x5f, 0x0a,
	0x1f, 0x70, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x79, 0x5f, 0x74, 0x72, 0x69, 0x6d, 0x6d, 0x65,
	0x64, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
	0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73,
	0x52, 0x1c, 0x70, 0x6f, 0x73, 0x73, 0x69, 0x62, 0x6c, 0x79, 0x54, 0x72, 0x69, 0x6d, 0x6d, 0x65,
	0x64, 0x42, 0x61, 0x73, 0x65, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x63, 0x73, 0x42, 0x36,
	0x0a, 0x25, 0x6f, 0x72, 0x67, 0x2e, 0x63, 0x68, 0x72, 0x6f, 0x6d, 0x69, 0x75, 0x6d, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x73, 0x79, 0x6e, 0x63, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x48, 0x03, 0x50, 0x01, 0x5a, 0x09, 0x2e, 0x2f, 0x73,
	0x79, 0x6e, 0x63, 0x5f, 0x70, 0x62,
}

var (
	file_entity_metadata_proto_rawDescOnce sync.Once
	file_entity_metadata_proto_rawDescData = file_entity_metadata_proto_rawDesc
)

func file_entity_metadata_proto_rawDescGZIP() []byte {
	file_entity_metadata_proto_rawDescOnce.Do(func() {
		file_entity_metadata_proto_rawDescData = protoimpl.X.CompressGZIP(file_entity_metadata_proto_rawDescData)
	})
	return file_entity_metadata_proto_rawDescData
}

var file_entity_metadata_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_entity_metadata_proto_goTypes = []interface{}{
	(*EntityMetadata)(nil),  // 0: sync_pb.EntityMetadata
	(*UniquePosition)(nil),  // 1: sync_pb.UniquePosition
	(*EntitySpecifics)(nil), // 2: sync_pb.EntitySpecifics
}
var file_entity_metadata_proto_depIdxs = []int32{
	1, // 0: sync_pb.EntityMetadata.unique_position:type_name -> sync_pb.UniquePosition
	2, // 1: sync_pb.EntityMetadata.possibly_trimmed_base_specifics:type_name -> sync_pb.EntitySpecifics
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_entity_metadata_proto_init() }
func file_entity_metadata_proto_init() {
	if File_entity_metadata_proto != nil {
		return
	}
	file_entity_specifics_proto_init()
	file_unique_position_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_entity_metadata_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityMetadata); i {
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
			RawDescriptor: file_entity_metadata_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_entity_metadata_proto_goTypes,
		DependencyIndexes: file_entity_metadata_proto_depIdxs,
		MessageInfos:      file_entity_metadata_proto_msgTypes,
	}.Build()
	File_entity_metadata_proto = out.File
	file_entity_metadata_proto_rawDesc = nil
	file_entity_metadata_proto_goTypes = nil
	file_entity_metadata_proto_depIdxs = nil
}
