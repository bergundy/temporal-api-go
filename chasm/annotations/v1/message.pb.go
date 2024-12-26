// The MIT License
//
// Copyright (c) 2024 Temporal Technologies Inc.  All rights reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Code generated by protoc-gen-go. DO NOT EDIT.
// plugins:
// 	protoc-gen-go
// 	protoc
// source: temporal/api/chasm/annotations/v1/message.proto

package annotations

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// RoutingOptions define how a request is routed internally within the Temporal cluster to the appropriate host.
type RoutingOptions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Routing is custom and must be manually implemented.
	Custom bool `protobuf:"varint,1,opt,name=custom,proto3" json:"custom,omitempty"`
	// Request will be routed to a random host.
	AnyHost bool `protobuf:"varint,2,opt,name=any_host,json=anyHost,proto3" json:"any_host,omitempty"`
	// Request will be routed by resolving the entity ID to a given shard.
	EntityId string `protobuf:"bytes,3,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	// Request will be routed by resolving the namespace ID and the workflow ID from this task token to a given shard.
	TaskToken string `protobuf:"bytes,4,opt,name=task_token,json=taskToken,proto3" json:"task_token,omitempty"`
}

func (x *RoutingOptions) Reset() {
	*x = RoutingOptions{}
	mi := &file_temporal_api_chasm_annotations_v1_message_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoutingOptions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutingOptions) ProtoMessage() {}

func (x *RoutingOptions) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_chasm_annotations_v1_message_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutingOptions.ProtoReflect.Descriptor instead.
func (*RoutingOptions) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_annotations_v1_message_proto_rawDescGZIP(), []int{0}
}

func (x *RoutingOptions) GetCustom() bool {
	if x != nil {
		return x.Custom
	}
	return false
}

func (x *RoutingOptions) GetAnyHost() bool {
	if x != nil {
		return x.AnyHost
	}
	return false
}

func (x *RoutingOptions) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *RoutingOptions) GetTaskToken() string {
	if x != nil {
		return x.TaskToken
	}
	return ""
}

var file_temporal_api_chasm_annotations_v1_message_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.MethodOptions)(nil),
		ExtensionType: (*RoutingOptions)(nil),
		Field:         7243,
		Name:          "temporal.api.chasm.annotations.v1.routing",
		Tag:           "bytes,7243,opt,name=routing",
		Filename:      "temporal/api/chasm/annotations/v1/message.proto",
	},
}

// Extension fields to descriptorpb.MethodOptions.
var (
	// optional temporal.api.chasm.annotations.v1.RoutingOptions routing = 7243;
	E_Routing = &file_temporal_api_chasm_annotations_v1_message_proto_extTypes[0]
)

var File_temporal_api_chasm_annotations_v1_message_proto protoreflect.FileDescriptor

var file_temporal_api_chasm_annotations_v1_message_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68,
	0x61, 0x73, 0x6d, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f,
	0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x21, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68,
	0x61, 0x73, 0x6d, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x76,
	0x31, 0x1a, 0x20, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x8f, 0x01, 0x0a, 0x0e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1a, 0x0a, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x42, 0x02, 0x68, 0x00,
	0x12, 0x1d, 0x0a, 0x08, 0x61, 0x6e, 0x79, 0x5f, 0x68, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x07, 0x61, 0x6e, 0x79, 0x48, 0x6f, 0x73, 0x74, 0x42, 0x02, 0x68, 0x00, 0x12, 0x1f,
	0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x21,
	0x0a, 0x0a, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x74, 0x61, 0x73, 0x6b, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42, 0x02, 0x68, 0x00,
	0x3a, 0x6f, 0x0a, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x1e, 0x2e, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x4d, 0x65, 0x74,
	0x68, 0x6f, 0x64, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0xcb, 0x38, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x31, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x4f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x69, 0x6e, 0x67, 0x88, 0x01, 0x01, 0x42, 0xbb, 0x01,
	0x0a, 0x24, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0xaa, 0x02, 0x23, 0x54, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x41,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x27,
	0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a,
	0x43, 0x68, 0x61, 0x73, 0x6d, 0x3a, 0x3a, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_chasm_annotations_v1_message_proto_rawDescOnce sync.Once
	file_temporal_api_chasm_annotations_v1_message_proto_rawDescData = file_temporal_api_chasm_annotations_v1_message_proto_rawDesc
)

func file_temporal_api_chasm_annotations_v1_message_proto_rawDescGZIP() []byte {
	file_temporal_api_chasm_annotations_v1_message_proto_rawDescOnce.Do(func() {
		file_temporal_api_chasm_annotations_v1_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_chasm_annotations_v1_message_proto_rawDescData)
	})
	return file_temporal_api_chasm_annotations_v1_message_proto_rawDescData
}

var file_temporal_api_chasm_annotations_v1_message_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_temporal_api_chasm_annotations_v1_message_proto_goTypes = []any{
	(*RoutingOptions)(nil),             // 0: temporal.api.chasm.annotations.v1.RoutingOptions
	(*descriptorpb.MethodOptions)(nil), // 1: google.protobuf.MethodOptions
}
var file_temporal_api_chasm_annotations_v1_message_proto_depIdxs = []int32{
	1, // 0: temporal.api.chasm.annotations.v1.routing:extendee -> google.protobuf.MethodOptions
	0, // 1: temporal.api.chasm.annotations.v1.routing:type_name -> temporal.api.chasm.annotations.v1.RoutingOptions
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	1, // [1:2] is the sub-list for extension type_name
	0, // [0:1] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_chasm_annotations_v1_message_proto_init() }
func file_temporal_api_chasm_annotations_v1_message_proto_init() {
	if File_temporal_api_chasm_annotations_v1_message_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_chasm_annotations_v1_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_chasm_annotations_v1_message_proto_goTypes,
		DependencyIndexes: file_temporal_api_chasm_annotations_v1_message_proto_depIdxs,
		MessageInfos:      file_temporal_api_chasm_annotations_v1_message_proto_msgTypes,
		ExtensionInfos:    file_temporal_api_chasm_annotations_v1_message_proto_extTypes,
	}.Build()
	File_temporal_api_chasm_annotations_v1_message_proto = out.File
	file_temporal_api_chasm_annotations_v1_message_proto_rawDesc = nil
	file_temporal_api_chasm_annotations_v1_message_proto_goTypes = nil
	file_temporal_api_chasm_annotations_v1_message_proto_depIdxs = nil
}
