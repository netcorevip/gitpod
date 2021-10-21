// Copyright (c) 2021 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License-AGPL.txt in the project root for license information.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: workspace.proto

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

type WorkspaceDownloadURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId     string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	WorkspaceId string `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
}

func (x *WorkspaceDownloadURLRequest) Reset() {
	*x = WorkspaceDownloadURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceDownloadURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceDownloadURLRequest) ProtoMessage() {}

func (x *WorkspaceDownloadURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceDownloadURLRequest.ProtoReflect.Descriptor instead.
func (*WorkspaceDownloadURLRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{0}
}

func (x *WorkspaceDownloadURLRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *WorkspaceDownloadURLRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

type WorkspaceDownloadURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *WorkspaceDownloadURLResponse) Reset() {
	*x = WorkspaceDownloadURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceDownloadURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceDownloadURLResponse) ProtoMessage() {}

func (x *WorkspaceDownloadURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceDownloadURLResponse.ProtoReflect.Descriptor instead.
func (*WorkspaceDownloadURLResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{1}
}

func (x *WorkspaceDownloadURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type DeleteWorkspaceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId          string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	WorkspaceId      string `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	IncludeSnapshots bool   `protobuf:"varint,3,opt,name=include_snapshots,json=includeSnapshots,proto3" json:"include_snapshots,omitempty"`
}

func (x *DeleteWorkspaceRequest) Reset() {
	*x = DeleteWorkspaceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceRequest) ProtoMessage() {}

func (x *DeleteWorkspaceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceRequest.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{2}
}

func (x *DeleteWorkspaceRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *DeleteWorkspaceRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *DeleteWorkspaceRequest) GetIncludeSnapshots() bool {
	if x != nil {
		return x.IncludeSnapshots
	}
	return false
}

type DeleteWorkspaceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteWorkspaceResponse) Reset() {
	*x = DeleteWorkspaceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteWorkspaceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteWorkspaceResponse) ProtoMessage() {}

func (x *DeleteWorkspaceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteWorkspaceResponse.ProtoReflect.Descriptor instead.
func (*DeleteWorkspaceResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{3}
}

type WorkspaceSnapshotExistsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OwnerId     string `protobuf:"bytes,1,opt,name=owner_id,json=ownerId,proto3" json:"owner_id,omitempty"`
	WorkspaceId string `protobuf:"bytes,2,opt,name=workspace_id,json=workspaceId,proto3" json:"workspace_id,omitempty"`
	Filename    string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
}

func (x *WorkspaceSnapshotExistsRequest) Reset() {
	*x = WorkspaceSnapshotExistsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceSnapshotExistsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSnapshotExistsRequest) ProtoMessage() {}

func (x *WorkspaceSnapshotExistsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSnapshotExistsRequest.ProtoReflect.Descriptor instead.
func (*WorkspaceSnapshotExistsRequest) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{4}
}

func (x *WorkspaceSnapshotExistsRequest) GetOwnerId() string {
	if x != nil {
		return x.OwnerId
	}
	return ""
}

func (x *WorkspaceSnapshotExistsRequest) GetWorkspaceId() string {
	if x != nil {
		return x.WorkspaceId
	}
	return ""
}

func (x *WorkspaceSnapshotExistsRequest) GetFilename() string {
	if x != nil {
		return x.Filename
	}
	return ""
}

type WorkspaceSnapshotExistsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Exists bool `protobuf:"varint,1,opt,name=exists,proto3" json:"exists,omitempty"`
}

func (x *WorkspaceSnapshotExistsResponse) Reset() {
	*x = WorkspaceSnapshotExistsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_workspace_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorkspaceSnapshotExistsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkspaceSnapshotExistsResponse) ProtoMessage() {}

func (x *WorkspaceSnapshotExistsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_workspace_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkspaceSnapshotExistsResponse.ProtoReflect.Descriptor instead.
func (*WorkspaceSnapshotExistsResponse) Descriptor() ([]byte, []int) {
	return file_workspace_proto_rawDescGZIP(), []int{5}
}

func (x *WorkspaceSnapshotExistsResponse) GetExists() bool {
	if x != nil {
		return x.Exists
	}
	return false
}

var File_workspace_proto protoreflect.FileDescriptor

var file_workspace_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x0e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x22, 0x5b, 0x0a, 0x1b, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x6f,
	0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77,
	0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x22, 0x30,
	0x0a, 0x1c, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c,
	0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x22, 0x83, 0x01, 0x0a, 0x16, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f,
	0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f,
	0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x69, 0x6e, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x5f, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x10, 0x69, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x53, 0x6e, 0x61,
	0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x22, 0x19, 0x0a, 0x17, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x7a, 0x0a, 0x1e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x64, 0x12, 0x21,
	0x0a, 0x0c, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x69, 0x6c, 0x65, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x39, 0x0a,
	0x1f, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x06, 0x65, 0x78, 0x69, 0x73, 0x74, 0x73, 0x32, 0xeb, 0x02, 0x0a, 0x10, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x73, 0x0a,
	0x14, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f,
	0x61, 0x64, 0x55, 0x52, 0x4c, 0x12, 0x2b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x44, 0x6f, 0x77, 0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x44, 0x6f, 0x77,
	0x6e, 0x6c, 0x6f, 0x61, 0x64, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x12, 0x64, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72,
	0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x7c, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x78, 0x69,
	0x73, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x57, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65, 0x53, 0x6e,
	0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x45, 0x78, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x74, 0x70, 0x6f, 0x64, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x69, 0x74, 0x70, 0x6f, 0x64, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2d, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_workspace_proto_rawDescOnce sync.Once
	file_workspace_proto_rawDescData = file_workspace_proto_rawDesc
)

func file_workspace_proto_rawDescGZIP() []byte {
	file_workspace_proto_rawDescOnce.Do(func() {
		file_workspace_proto_rawDescData = protoimpl.X.CompressGZIP(file_workspace_proto_rawDescData)
	})
	return file_workspace_proto_rawDescData
}

var file_workspace_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_workspace_proto_goTypes = []interface{}{
	(*WorkspaceDownloadURLRequest)(nil),     // 0: contentservice.WorkspaceDownloadURLRequest
	(*WorkspaceDownloadURLResponse)(nil),    // 1: contentservice.WorkspaceDownloadURLResponse
	(*DeleteWorkspaceRequest)(nil),          // 2: contentservice.DeleteWorkspaceRequest
	(*DeleteWorkspaceResponse)(nil),         // 3: contentservice.DeleteWorkspaceResponse
	(*WorkspaceSnapshotExistsRequest)(nil),  // 4: contentservice.WorkspaceSnapshotExistsRequest
	(*WorkspaceSnapshotExistsResponse)(nil), // 5: contentservice.WorkspaceSnapshotExistsResponse
}
var file_workspace_proto_depIdxs = []int32{
	0, // 0: contentservice.WorkspaceService.WorkspaceDownloadURL:input_type -> contentservice.WorkspaceDownloadURLRequest
	2, // 1: contentservice.WorkspaceService.DeleteWorkspace:input_type -> contentservice.DeleteWorkspaceRequest
	4, // 2: contentservice.WorkspaceService.WorkspaceSnapshotExists:input_type -> contentservice.WorkspaceSnapshotExistsRequest
	1, // 3: contentservice.WorkspaceService.WorkspaceDownloadURL:output_type -> contentservice.WorkspaceDownloadURLResponse
	3, // 4: contentservice.WorkspaceService.DeleteWorkspace:output_type -> contentservice.DeleteWorkspaceResponse
	5, // 5: contentservice.WorkspaceService.WorkspaceSnapshotExists:output_type -> contentservice.WorkspaceSnapshotExistsResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_workspace_proto_init() }
func file_workspace_proto_init() {
	if File_workspace_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_workspace_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceDownloadURLRequest); i {
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
		file_workspace_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceDownloadURLResponse); i {
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
		file_workspace_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceRequest); i {
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
		file_workspace_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteWorkspaceResponse); i {
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
		file_workspace_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceSnapshotExistsRequest); i {
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
		file_workspace_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorkspaceSnapshotExistsResponse); i {
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
			RawDescriptor: file_workspace_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_workspace_proto_goTypes,
		DependencyIndexes: file_workspace_proto_depIdxs,
		MessageInfos:      file_workspace_proto_msgTypes,
	}.Build()
	File_workspace_proto = out.File
	file_workspace_proto_rawDesc = nil
	file_workspace_proto_goTypes = nil
	file_workspace_proto_depIdxs = nil
}
