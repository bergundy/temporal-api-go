// The MIT License
//
// Copyright (c) 2020 Temporal Technologies Inc.  All rights reserved.
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
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: temporal/api/enums/v1/task_queue.proto

package enums

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TaskQueueKind int32

const (
	TASK_QUEUE_KIND_UNSPECIFIED TaskQueueKind = 0
	// Tasks from a normal workflow task queue always include complete workflow history
	//
	// The task queue specified by the user is always a normal task queue. There can be as many
	// workers as desired for a single normal task queue. All those workers may pick up tasks from
	// that queue.
	TASK_QUEUE_KIND_NORMAL TaskQueueKind = 1
	// A sticky queue only includes new history since the last workflow task, and they are
	// per-worker.
	//
	// Sticky queues are created dynamically by each worker during their start up. They only exist
	// for the lifetime of the worker process. Tasks in a sticky task queue are only available to
	// the worker that created the sticky queue.
	//
	// Sticky queues are only for workflow tasks. There are no sticky task queues for activities.
	TASK_QUEUE_KIND_STICKY TaskQueueKind = 2
)

// Enum value maps for TaskQueueKind.
var (
	TaskQueueKind_name = map[int32]string{
		0: "TASK_QUEUE_KIND_UNSPECIFIED",
		1: "TASK_QUEUE_KIND_NORMAL",
		2: "TASK_QUEUE_KIND_STICKY",
	}
	TaskQueueKind_value = map[string]int32{
		"TASK_QUEUE_KIND_UNSPECIFIED": 0,
		"TASK_QUEUE_KIND_NORMAL":      1,
		"TASK_QUEUE_KIND_STICKY":      2,
	}
)

func (x TaskQueueKind) Enum() *TaskQueueKind {
	p := new(TaskQueueKind)
	*p = x
	return p
}

func (x TaskQueueKind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskQueueKind) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_task_queue_proto_enumTypes[0].Descriptor()
}

func (TaskQueueKind) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_task_queue_proto_enumTypes[0]
}

func (x TaskQueueKind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskQueueKind.Descriptor instead.
func (TaskQueueKind) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP(), []int{0}
}

type TaskQueueType int32

const (
	TASK_QUEUE_TYPE_UNSPECIFIED TaskQueueType = 0
	// Workflow type of task queue.
	TASK_QUEUE_TYPE_WORKFLOW TaskQueueType = 1
	// Activity type of task queue.
	TASK_QUEUE_TYPE_ACTIVITY TaskQueueType = 2
)

// Enum value maps for TaskQueueType.
var (
	TaskQueueType_name = map[int32]string{
		0: "TASK_QUEUE_TYPE_UNSPECIFIED",
		1: "TASK_QUEUE_TYPE_WORKFLOW",
		2: "TASK_QUEUE_TYPE_ACTIVITY",
	}
	TaskQueueType_value = map[string]int32{
		"TASK_QUEUE_TYPE_UNSPECIFIED": 0,
		"TASK_QUEUE_TYPE_WORKFLOW":    1,
		"TASK_QUEUE_TYPE_ACTIVITY":    2,
	}
)

func (x TaskQueueType) Enum() *TaskQueueType {
	p := new(TaskQueueType)
	*p = x
	return p
}

func (x TaskQueueType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskQueueType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_task_queue_proto_enumTypes[1].Descriptor()
}

func (TaskQueueType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_task_queue_proto_enumTypes[1]
}

func (x TaskQueueType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskQueueType.Descriptor instead.
func (TaskQueueType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP(), []int{1}
}

// Specifies which category of tasks may reach a worker on a versioned task queue.
// Used both in a reachability query and its response.
type TaskReachability int32

const (
	TASK_REACHABILITY_UNSPECIFIED TaskReachability = 0
	// There's a possiblity for a worker to receive new workflow tasks. Workers should *not* be retired.
	TASK_REACHABILITY_NEW_WORKFLOWS TaskReachability = 1
	// There's a possiblity for a worker to receive existing workflow and activity tasks from existing workflows. Workers
	// should *not* be retired.
	// This enum value does not distinguish between open and closed workflows.
	TASK_REACHABILITY_EXISTING_WORKFLOWS TaskReachability = 2
	// There's a possiblity for a worker to receive existing workflow and activity tasks from open workflows. Workers
	// should *not* be retired.
	TASK_REACHABILITY_OPEN_WORKFLOWS TaskReachability = 3
	// There's a possiblity for a worker to receive existing workflow tasks from closed workflows. Workers may be
	// retired dependending on application requirements. For example, if there's no need to query closed workflows.
	TASK_REACHABILITY_CLOSED_WORKFLOWS TaskReachability = 4
)

// Enum value maps for TaskReachability.
var (
	TaskReachability_name = map[int32]string{
		0: "TASK_REACHABILITY_UNSPECIFIED",
		1: "TASK_REACHABILITY_NEW_WORKFLOWS",
		2: "TASK_REACHABILITY_EXISTING_WORKFLOWS",
		3: "TASK_REACHABILITY_OPEN_WORKFLOWS",
		4: "TASK_REACHABILITY_CLOSED_WORKFLOWS",
	}
	TaskReachability_value = map[string]int32{
		"TASK_REACHABILITY_UNSPECIFIED":        0,
		"TASK_REACHABILITY_NEW_WORKFLOWS":      1,
		"TASK_REACHABILITY_EXISTING_WORKFLOWS": 2,
		"TASK_REACHABILITY_OPEN_WORKFLOWS":     3,
		"TASK_REACHABILITY_CLOSED_WORKFLOWS":   4,
	}
)

func (x TaskReachability) Enum() *TaskReachability {
	p := new(TaskReachability)
	*p = x
	return p
}

func (x TaskReachability) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TaskReachability) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_task_queue_proto_enumTypes[2].Descriptor()
}

func (TaskReachability) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_task_queue_proto_enumTypes[2]
}

func (x TaskReachability) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TaskReachability.Descriptor instead.
func (TaskReachability) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP(), []int{2}
}

var File_temporal_api_enums_v1_task_queue_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_task_queue_proto_rawDesc = []byte{
	0x0a, 0x26, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a,
	0x68, 0x0a, 0x0d, 0x54, 0x61, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x4b, 0x69, 0x6e, 0x64,
	0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4b,
	0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f,
	0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x4e, 0x4f, 0x52, 0x4d, 0x41, 0x4c, 0x10, 0x01, 0x12, 0x1a, 0x0a,
	0x16, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x4b, 0x49, 0x4e, 0x44,
	0x5f, 0x53, 0x54, 0x49, 0x43, 0x4b, 0x59, 0x10, 0x02, 0x2a, 0x6c, 0x0a, 0x0d, 0x54, 0x61, 0x73,
	0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x41,
	0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1c, 0x0a, 0x18, 0x54,
	0x41, 0x53, 0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x54, 0x41, 0x53,
	0x4b, 0x5f, 0x51, 0x55, 0x45, 0x55, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x10, 0x02, 0x2a, 0xd2, 0x01, 0x0a, 0x10, 0x54, 0x61, 0x73, 0x6b,
	0x52, 0x65, 0x61, 0x63, 0x68, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x12, 0x21, 0x0a, 0x1d,
	0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54,
	0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x23, 0x0a, 0x1f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49,
	0x4c, 0x49, 0x54, 0x59, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f,
	0x57, 0x53, 0x10, 0x01, 0x12, 0x28, 0x0a, 0x24, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41,
	0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x45, 0x58, 0x49, 0x53, 0x54, 0x49,
	0x4e, 0x47, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x53, 0x10, 0x02, 0x12, 0x24,
	0x0a, 0x20, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x41, 0x42, 0x49, 0x4c,
	0x49, 0x54, 0x59, 0x5f, 0x4f, 0x50, 0x45, 0x4e, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f,
	0x57, 0x53, 0x10, 0x03, 0x12, 0x26, 0x0a, 0x22, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x52, 0x45, 0x41,
	0x43, 0x48, 0x41, 0x42, 0x49, 0x4c, 0x49, 0x54, 0x59, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x44,
	0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x53, 0x10, 0x04, 0x42, 0x86, 0x01, 0x0a,
	0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x51,
	0x75, 0x65, 0x75, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa, 0x02,
	0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e,
	0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d,
	0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_task_queue_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_task_queue_proto_rawDescData = file_temporal_api_enums_v1_task_queue_proto_rawDesc
)

func file_temporal_api_enums_v1_task_queue_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_task_queue_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_task_queue_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_task_queue_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_task_queue_proto_rawDescData
}

var file_temporal_api_enums_v1_task_queue_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_temporal_api_enums_v1_task_queue_proto_goTypes = []interface{}{
	(TaskQueueKind)(0),    // 0: temporal.api.enums.v1.TaskQueueKind
	(TaskQueueType)(0),    // 1: temporal.api.enums.v1.TaskQueueType
	(TaskReachability)(0), // 2: temporal.api.enums.v1.TaskReachability
}
var file_temporal_api_enums_v1_task_queue_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_task_queue_proto_init() }
func file_temporal_api_enums_v1_task_queue_proto_init() {
	if File_temporal_api_enums_v1_task_queue_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_task_queue_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_task_queue_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_task_queue_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_task_queue_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_task_queue_proto = out.File
	file_temporal_api_enums_v1_task_queue_proto_rawDesc = nil
	file_temporal_api_enums_v1_task_queue_proto_goTypes = nil
	file_temporal_api_enums_v1_task_queue_proto_depIdxs = nil
}
