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
// source: temporal/api/chasm/activityservice/v1/request_response.proto

package activityservice

import (
	reflect "reflect"
	sync "sync"

	v1 "go.temporal.io/api/chasm/activity/v1"
	v11 "go.temporal.io/api/common/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type StartActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity  string              `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Namespace string              `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	RequestId string              `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	EntityId  string              `protobuf:"bytes,4,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	Options   *v1.ActivityOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	Inputs    *v11.Payloads       `protobuf:"bytes,6,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// @ui(disabled)
	// or @contexts(["worker"])
	RequestEagerExecution bool `protobuf:"varint,7,opt,name=request_eager_execution,json=requestEagerExecution,proto3" json:"request_eager_execution,omitempty"`
}

func (x *StartActivityRequest) Reset() {
	*x = StartActivityRequest{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartActivityRequest) ProtoMessage() {}

func (x *StartActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartActivityRequest.ProtoReflect.Descriptor instead.
func (*StartActivityRequest) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{0}
}

func (x *StartActivityRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *StartActivityRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *StartActivityRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *StartActivityRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *StartActivityRequest) GetOptions() *v1.ActivityOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *StartActivityRequest) GetInputs() *v11.Payloads {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *StartActivityRequest) GetRequestEagerExecution() bool {
	if x != nil {
		return x.RequestEagerExecution
	}
	return false
}

type StartActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *StartActivityResponse) Reset() {
	*x = StartActivityResponse{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *StartActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartActivityResponse) ProtoMessage() {}

func (x *StartActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartActivityResponse.ProtoReflect.Descriptor instead.
func (*StartActivityResponse) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{1}
}

func (x *StartActivityResponse) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

var File_temporal_api_chasm_activityservice_v1_request_response_proto protoreflect.FileDescriptor

var file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68,
	0x61, 0x73, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x1a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe5, 0x02, 0x0a, 0x14, 0x53, 0x74, 0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x42, 0x02, 0x68, 0x00, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x21, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x1f, 0x0a, 0x09, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4d, 0x0a, 0x07, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e,
	0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x3c, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73, 0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x42, 0x02,
	0x68, 0x00, 0x12, 0x3a, 0x0a, 0x17, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x61, 0x67,
	0x65, 0x72, 0x5f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x15, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x45, 0x61, 0x67, 0x65, 0x72, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x22, 0x32, 0x0a, 0x15, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x42, 0xd7, 0x01,
	0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x3b, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xaa, 0x02, 0x27, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x73, 0x6d,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x56, 0x31, 0xea, 0x02, 0x2b, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x43, 0x68, 0x61, 0x73, 0x6d, 0x3a, 0x3a, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescOnce sync.Once
	file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescData = file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDesc
)

func file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP() []byte {
	file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescOnce.Do(func() {
		file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescData)
	})
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescData
}

var file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_temporal_api_chasm_activityservice_v1_request_response_proto_goTypes = []any{
	(*StartActivityRequest)(nil),  // 0: temporal.api.chasm.activityservice.v1.StartActivityRequest
	(*StartActivityResponse)(nil), // 1: temporal.api.chasm.activityservice.v1.StartActivityResponse
	(*v1.ActivityOptions)(nil),    // 2: temporal.api.chasm.activity.v1.ActivityOptions
	(*v11.Payloads)(nil),          // 3: temporal.api.common.v1.Payloads
}
var file_temporal_api_chasm_activityservice_v1_request_response_proto_depIdxs = []int32{
	2, // 0: temporal.api.chasm.activityservice.v1.StartActivityRequest.options:type_name -> temporal.api.chasm.activity.v1.ActivityOptions
	3, // 1: temporal.api.chasm.activityservice.v1.StartActivityRequest.inputs:type_name -> temporal.api.common.v1.Payloads
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_temporal_api_chasm_activityservice_v1_request_response_proto_init() }
func file_temporal_api_chasm_activityservice_v1_request_response_proto_init() {
	if File_temporal_api_chasm_activityservice_v1_request_response_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_chasm_activityservice_v1_request_response_proto_goTypes,
		DependencyIndexes: file_temporal_api_chasm_activityservice_v1_request_response_proto_depIdxs,
		MessageInfos:      file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes,
	}.Build()
	File_temporal_api_chasm_activityservice_v1_request_response_proto = out.File
	file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDesc = nil
	file_temporal_api_chasm_activityservice_v1_request_response_proto_goTypes = nil
	file_temporal_api_chasm_activityservice_v1_request_response_proto_depIdxs = nil
}
