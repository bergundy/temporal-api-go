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

type ExecuteActivityRequest struct {
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

func (x *ExecuteActivityRequest) Reset() {
	*x = ExecuteActivityRequest{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteActivityRequest) ProtoMessage() {}

func (x *ExecuteActivityRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ExecuteActivityRequest.ProtoReflect.Descriptor instead.
func (*ExecuteActivityRequest) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteActivityRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *ExecuteActivityRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ExecuteActivityRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ExecuteActivityRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *ExecuteActivityRequest) GetOptions() *v1.ActivityOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *ExecuteActivityRequest) GetInputs() *v11.Payloads {
	if x != nil {
		return x.Inputs
	}
	return nil
}

func (x *ExecuteActivityRequest) GetRequestEagerExecution() bool {
	if x != nil {
		return x.RequestEagerExecution
	}
	return false
}

type ExecuteActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result *v11.Payload `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *ExecuteActivityResponse) Reset() {
	*x = ExecuteActivityResponse{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteActivityResponse) ProtoMessage() {}

func (x *ExecuteActivityResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use ExecuteActivityResponse.ProtoReflect.Descriptor instead.
func (*ExecuteActivityResponse) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteActivityResponse) GetResult() *v11.Payload {
	if x != nil {
		return x.Result
	}
	return nil
}

type ExecuteActivityStartResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RunId string `protobuf:"bytes,1,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *ExecuteActivityStartResponse) Reset() {
	*x = ExecuteActivityStartResponse{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExecuteActivityStartResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteActivityStartResponse) ProtoMessage() {}

func (x *ExecuteActivityStartResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteActivityStartResponse.ProtoReflect.Descriptor instead.
func (*ExecuteActivityStartResponse) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{2}
}

func (x *ExecuteActivityStartResponse) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type DescribeActivityRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity  string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Namespace string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	RequestId string `protobuf:"bytes,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	EntityId  string `protobuf:"bytes,4,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	RunId     string `protobuf:"bytes,5,opt,name=run_id,json=runId,proto3" json:"run_id,omitempty"`
}

func (x *DescribeActivityRequest) Reset() {
	*x = DescribeActivityRequest{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeActivityRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeActivityRequest) ProtoMessage() {}

func (x *DescribeActivityRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeActivityRequest.ProtoReflect.Descriptor instead.
func (*DescribeActivityRequest) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{3}
}

func (x *DescribeActivityRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *DescribeActivityRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *DescribeActivityRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *DescribeActivityRequest) GetEntityId() string {
	if x != nil {
		return x.EntityId
	}
	return ""
}

func (x *DescribeActivityRequest) GetRunId() string {
	if x != nil {
		return x.RunId
	}
	return ""
}

type DescribeActivityResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Options *v1.ActivityOptions `protobuf:"bytes,1,opt,name=options,proto3" json:"options,omitempty"`
	Info    *v1.ActivityInfo    `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
}

func (x *DescribeActivityResponse) Reset() {
	*x = DescribeActivityResponse{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DescribeActivityResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DescribeActivityResponse) ProtoMessage() {}

func (x *DescribeActivityResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DescribeActivityResponse.ProtoReflect.Descriptor instead.
func (*DescribeActivityResponse) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{4}
}

func (x *DescribeActivityResponse) GetOptions() *v1.ActivityOptions {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *DescribeActivityResponse) GetInfo() *v1.ActivityInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type ListActivitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity      string `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Namespace     string `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Query         string `protobuf:"bytes,3,opt,name=query,proto3" json:"query,omitempty"`
	NextPageToken []byte `protobuf:"bytes,4,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
}

func (x *ListActivitiesRequest) Reset() {
	*x = ListActivitiesRequest{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListActivitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActivitiesRequest) ProtoMessage() {}

func (x *ListActivitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActivitiesRequest.ProtoReflect.Descriptor instead.
func (*ListActivitiesRequest) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{5}
}

func (x *ListActivitiesRequest) GetIdentity() string {
	if x != nil {
		return x.Identity
	}
	return ""
}

func (x *ListActivitiesRequest) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *ListActivitiesRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *ListActivitiesRequest) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

type ListActivitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NextPageToken []byte                 `protobuf:"bytes,1,opt,name=next_page_token,json=nextPageToken,proto3" json:"next_page_token,omitempty"`
	Items         []*v1.ActivityListInfo `protobuf:"bytes,2,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListActivitiesResponse) Reset() {
	*x = ListActivitiesResponse{}
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListActivitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListActivitiesResponse) ProtoMessage() {}

func (x *ListActivitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListActivitiesResponse.ProtoReflect.Descriptor instead.
func (*ListActivitiesResponse) Descriptor() ([]byte, []int) {
	return file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDescGZIP(), []int{6}
}

func (x *ListActivitiesResponse) GetNextPageToken() []byte {
	if x != nil {
		return x.NextPageToken
	}
	return nil
}

func (x *ListActivitiesResponse) GetItems() []*v1.ActivityListInfo {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_temporal_api_chasm_activityservice_v1_request_response_proto protoreflect.FileDescriptor

var file_temporal_api_chasm_activityservice_v1_request_response_proto_rawDesc = []byte{
	0x0a, 0x3c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x68,
	0x61, 0x73, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x25, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a,
	0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x2c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2f, 0x76,
	0x31, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe7,
	0x02, 0x0a, 0x16, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x42, 0x02, 0x68, 0x00, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61,
	0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x21, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x1f, 0x0a, 0x09, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4d, 0x0a, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61,
	0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x42, 0x02, 0x68,
	0x00, 0x12, 0x3c, 0x0a, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x20, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x73,
	0x52, 0x06, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x3a, 0x0a, 0x17, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x65, 0x61, 0x67, 0x65, 0x72, 0x5f, 0x65, 0x78, 0x65, 0x63,
	0x75, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x15, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x45, 0x61, 0x67, 0x65, 0x72, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x02, 0x68, 0x00, 0x22, 0x56, 0x0a, 0x17, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3b,
	0x0a, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x42, 0x02, 0x68, 0x00, 0x22, 0x39, 0x0a, 0x1c, 0x45, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74, 0x61, 0x72, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x22,
	0xba, 0x01, 0x0a, 0x17, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x42, 0x02, 0x68, 0x00, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65,
	0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x02, 0x68, 0x00, 0x12, 0x21, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x1f, 0x0a, 0x09, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x42, 0x02, 0x68, 0x00, 0x12, 0x19, 0x0a, 0x06, 0x72, 0x75, 0x6e, 0x5f,
	0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x75, 0x6e, 0x49, 0x64, 0x42, 0x02,
	0x68, 0x00, 0x22, 0xaf, 0x01, 0x0a, 0x18, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x62, 0x65, 0x41, 0x63,
	0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4d, 0x0a,
	0x07, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2f, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x73,
	0x6d, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74,
	0x69, 0x76, 0x69, 0x74, 0x79, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x6f, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x42, 0x02, 0x68, 0x00, 0x12, 0x44, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x42, 0x02, 0x68, 0x00, 0x22, 0x9f, 0x01, 0x0a, 0x15, 0x4c,
	0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x02, 0x68,
	0x00, 0x12, 0x20, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x42, 0x02, 0x68,
	0x00, 0x12, 0x18, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x42, 0x02, 0x68, 0x00, 0x12, 0x2a, 0x0a, 0x0f, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x42,
	0x02, 0x68, 0x00, 0x22, 0x90, 0x01, 0x0a, 0x16, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2a, 0x0a, 0x0f,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x42, 0x02, 0x68, 0x00, 0x12, 0x4a, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x4c, 0x69, 0x73, 0x74,
	0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x42, 0x02, 0x68, 0x00, 0x42, 0xd7,
	0x01, 0x0a, 0x28, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x42, 0x14, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x3b, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x68, 0x61, 0x73, 0x6d, 0x2f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0xaa, 0x02, 0x27, 0x54, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x43, 0x68, 0x61, 0x73, 0x6d, 0x2e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31,
	0xea, 0x02, 0x2b, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41,
	0x70, 0x69, 0x3a, 0x3a, 0x43, 0x68, 0x61, 0x73, 0x6d, 0x3a, 0x3a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69,
	0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
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

var file_temporal_api_chasm_activityservice_v1_request_response_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_temporal_api_chasm_activityservice_v1_request_response_proto_goTypes = []any{
	(*ExecuteActivityRequest)(nil),       // 0: temporal.api.chasm.activityservice.v1.ExecuteActivityRequest
	(*ExecuteActivityResponse)(nil),      // 1: temporal.api.chasm.activityservice.v1.ExecuteActivityResponse
	(*ExecuteActivityStartResponse)(nil), // 2: temporal.api.chasm.activityservice.v1.ExecuteActivityStartResponse
	(*DescribeActivityRequest)(nil),      // 3: temporal.api.chasm.activityservice.v1.DescribeActivityRequest
	(*DescribeActivityResponse)(nil),     // 4: temporal.api.chasm.activityservice.v1.DescribeActivityResponse
	(*ListActivitiesRequest)(nil),        // 5: temporal.api.chasm.activityservice.v1.ListActivitiesRequest
	(*ListActivitiesResponse)(nil),       // 6: temporal.api.chasm.activityservice.v1.ListActivitiesResponse
	(*v1.ActivityOptions)(nil),           // 7: temporal.api.chasm.activity.v1.ActivityOptions
	(*v11.Payloads)(nil),                 // 8: temporal.api.common.v1.Payloads
	(*v11.Payload)(nil),                  // 9: temporal.api.common.v1.Payload
	(*v1.ActivityInfo)(nil),              // 10: temporal.api.chasm.activity.v1.ActivityInfo
	(*v1.ActivityListInfo)(nil),          // 11: temporal.api.chasm.activity.v1.ActivityListInfo
}
var file_temporal_api_chasm_activityservice_v1_request_response_proto_depIdxs = []int32{
	7,  // 0: temporal.api.chasm.activityservice.v1.ExecuteActivityRequest.options:type_name -> temporal.api.chasm.activity.v1.ActivityOptions
	8,  // 1: temporal.api.chasm.activityservice.v1.ExecuteActivityRequest.inputs:type_name -> temporal.api.common.v1.Payloads
	9,  // 2: temporal.api.chasm.activityservice.v1.ExecuteActivityResponse.result:type_name -> temporal.api.common.v1.Payload
	7,  // 3: temporal.api.chasm.activityservice.v1.DescribeActivityResponse.options:type_name -> temporal.api.chasm.activity.v1.ActivityOptions
	10, // 4: temporal.api.chasm.activityservice.v1.DescribeActivityResponse.info:type_name -> temporal.api.chasm.activity.v1.ActivityInfo
	11, // 5: temporal.api.chasm.activityservice.v1.ListActivitiesResponse.items:type_name -> temporal.api.chasm.activity.v1.ActivityListInfo
	6,  // [6:6] is the sub-list for method output_type
	6,  // [6:6] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
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
			NumMessages:   7,
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
