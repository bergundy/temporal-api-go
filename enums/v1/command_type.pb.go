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
// source: temporal/api/enums/v1/command_type.proto

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

// Whenever this list of command types is changed do change the function shouldBufferEvent in mutableStateBuilder.go to make sure to do the correct event ordering.
type CommandType int32

const (
	CommandType_COMMAND_TYPE_UNSPECIFIED                                CommandType = 0
	CommandType_COMMAND_TYPE_SCHEDULE_ACTIVITY_TASK                     CommandType = 1
	CommandType_COMMAND_TYPE_REQUEST_CANCEL_ACTIVITY_TASK               CommandType = 2
	CommandType_COMMAND_TYPE_START_TIMER                                CommandType = 3
	CommandType_COMMAND_TYPE_COMPLETE_WORKFLOW_EXECUTION                CommandType = 4
	CommandType_COMMAND_TYPE_FAIL_WORKFLOW_EXECUTION                    CommandType = 5
	CommandType_COMMAND_TYPE_CANCEL_TIMER                               CommandType = 6
	CommandType_COMMAND_TYPE_CANCEL_WORKFLOW_EXECUTION                  CommandType = 7
	CommandType_COMMAND_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION CommandType = 8
	CommandType_COMMAND_TYPE_RECORD_MARKER                              CommandType = 9
	CommandType_COMMAND_TYPE_CONTINUE_AS_NEW_WORKFLOW_EXECUTION         CommandType = 10
	CommandType_COMMAND_TYPE_START_CHILD_WORKFLOW_EXECUTION             CommandType = 11
	CommandType_COMMAND_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION         CommandType = 12
	CommandType_COMMAND_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES          CommandType = 13
	CommandType_COMMAND_TYPE_PROTOCOL_MESSAGE                           CommandType = 14
	CommandType_COMMAND_TYPE_MODIFY_WORKFLOW_PROPERTIES                 CommandType = 16
)

// Enum value maps for CommandType.
var (
	CommandType_name = map[int32]string{
		0:  "COMMAND_TYPE_UNSPECIFIED",
		1:  "COMMAND_TYPE_SCHEDULE_ACTIVITY_TASK",
		2:  "COMMAND_TYPE_REQUEST_CANCEL_ACTIVITY_TASK",
		3:  "COMMAND_TYPE_START_TIMER",
		4:  "COMMAND_TYPE_COMPLETE_WORKFLOW_EXECUTION",
		5:  "COMMAND_TYPE_FAIL_WORKFLOW_EXECUTION",
		6:  "COMMAND_TYPE_CANCEL_TIMER",
		7:  "COMMAND_TYPE_CANCEL_WORKFLOW_EXECUTION",
		8:  "COMMAND_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION",
		9:  "COMMAND_TYPE_RECORD_MARKER",
		10: "COMMAND_TYPE_CONTINUE_AS_NEW_WORKFLOW_EXECUTION",
		11: "COMMAND_TYPE_START_CHILD_WORKFLOW_EXECUTION",
		12: "COMMAND_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION",
		13: "COMMAND_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES",
		14: "COMMAND_TYPE_PROTOCOL_MESSAGE",
		16: "COMMAND_TYPE_MODIFY_WORKFLOW_PROPERTIES",
	}
	CommandType_value = map[string]int32{
		"COMMAND_TYPE_UNSPECIFIED":                                0,
		"COMMAND_TYPE_SCHEDULE_ACTIVITY_TASK":                     1,
		"COMMAND_TYPE_REQUEST_CANCEL_ACTIVITY_TASK":               2,
		"COMMAND_TYPE_START_TIMER":                                3,
		"COMMAND_TYPE_COMPLETE_WORKFLOW_EXECUTION":                4,
		"COMMAND_TYPE_FAIL_WORKFLOW_EXECUTION":                    5,
		"COMMAND_TYPE_CANCEL_TIMER":                               6,
		"COMMAND_TYPE_CANCEL_WORKFLOW_EXECUTION":                  7,
		"COMMAND_TYPE_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION": 8,
		"COMMAND_TYPE_RECORD_MARKER":                              9,
		"COMMAND_TYPE_CONTINUE_AS_NEW_WORKFLOW_EXECUTION":         10,
		"COMMAND_TYPE_START_CHILD_WORKFLOW_EXECUTION":             11,
		"COMMAND_TYPE_SIGNAL_EXTERNAL_WORKFLOW_EXECUTION":         12,
		"COMMAND_TYPE_UPSERT_WORKFLOW_SEARCH_ATTRIBUTES":          13,
		"COMMAND_TYPE_PROTOCOL_MESSAGE":                           14,
		"COMMAND_TYPE_MODIFY_WORKFLOW_PROPERTIES":                 16,
	}
)

func (x CommandType) Enum() *CommandType {
	p := new(CommandType)
	*p = x
	return p
}

func (x CommandType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CommandType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_command_type_proto_enumTypes[0].Descriptor()
}

func (CommandType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_command_type_proto_enumTypes[0]
}

func (x CommandType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CommandType.Descriptor instead.
func (CommandType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_command_type_proto_rawDescGZIP(), []int{0}
}

var File_temporal_api_enums_v1_command_type_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_command_type_proto_rawDesc = []byte{
	0x0a, 0x28, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x2a, 0xc0, 0x05, 0x0a, 0x0b, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50,
	0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12,
	0x27, 0x0a, 0x23, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54,
	0x59, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x01, 0x12, 0x2d, 0x0a, 0x29, 0x43, 0x4f, 0x4d, 0x4d,
	0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x10, 0x02, 0x12, 0x1c, 0x0a, 0x18, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x5f, 0x54, 0x49,
	0x4d, 0x45, 0x52, 0x10, 0x03, 0x12, 0x2c, 0x0a, 0x28, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f,
	0x4e, 0x10, 0x04, 0x12, 0x28, 0x0a, 0x24, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f,
	0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x1d, 0x0a,
	0x19, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x41,
	0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x52, 0x10, 0x06, 0x12, 0x2a, 0x0a, 0x26,
	0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x07, 0x12, 0x3b, 0x0a, 0x37, 0x43, 0x4f, 0x4d, 0x4d,
	0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54,
	0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c,
	0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54,
	0x49, 0x4f, 0x4e, 0x10, 0x08, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x52, 0x45, 0x43, 0x4f, 0x52, 0x44, 0x5f, 0x4d, 0x41, 0x52,
	0x4b, 0x45, 0x52, 0x10, 0x09, 0x12, 0x33, 0x0a, 0x2f, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x5f, 0x41,
	0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45,
	0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0a, 0x12, 0x2f, 0x0a, 0x2b, 0x43, 0x4f,
	0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x5f, 0x43, 0x48, 0x49, 0x4c, 0x44, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f,
	0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0b, 0x12, 0x33, 0x0a, 0x2f, 0x43,
	0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x49, 0x47, 0x4e,
	0x41, 0x4c, 0x5f, 0x45, 0x58, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x0c,
	0x12, 0x32, 0x0a, 0x2e, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x50, 0x53, 0x45, 0x52, 0x54, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57,
	0x5f, 0x53, 0x45, 0x41, 0x52, 0x43, 0x48, 0x5f, 0x41, 0x54, 0x54, 0x52, 0x49, 0x42, 0x55, 0x54,
	0x45, 0x53, 0x10, 0x0d, 0x12, 0x21, 0x0a, 0x1d, 0x43, 0x4f, 0x4d, 0x4d, 0x41, 0x4e, 0x44, 0x5f,
	0x54, 0x59, 0x50, 0x45, 0x5f, 0x50, 0x52, 0x4f, 0x54, 0x4f, 0x43, 0x4f, 0x4c, 0x5f, 0x4d, 0x45,
	0x53, 0x53, 0x41, 0x47, 0x45, 0x10, 0x0e, 0x12, 0x2b, 0x0a, 0x27, 0x43, 0x4f, 0x4d, 0x4d, 0x41,
	0x4e, 0x44, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4d, 0x4f, 0x44, 0x49, 0x46, 0x59, 0x5f, 0x57,
	0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x50, 0x52, 0x4f, 0x50, 0x45, 0x52, 0x54, 0x49,
	0x45, 0x53, 0x10, 0x10, 0x42, 0x88, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70,
	0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76,
	0x31, 0x42, 0x10, 0x43, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x54, 0x79, 0x70, 0x65, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72,
	0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f,
	0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f,
	0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e,
	0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a,
	0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_command_type_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_command_type_proto_rawDescData = file_temporal_api_enums_v1_command_type_proto_rawDesc
)

func file_temporal_api_enums_v1_command_type_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_command_type_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_command_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_command_type_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_command_type_proto_rawDescData
}

var file_temporal_api_enums_v1_command_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_temporal_api_enums_v1_command_type_proto_goTypes = []interface{}{
	(CommandType)(0), // 0: temporal.api.enums.v1.CommandType
}
var file_temporal_api_enums_v1_command_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_command_type_proto_init() }
func file_temporal_api_enums_v1_command_type_proto_init() {
	if File_temporal_api_enums_v1_command_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_command_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_command_type_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_command_type_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_command_type_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_command_type_proto = out.File
	file_temporal_api_enums_v1_command_type_proto_rawDesc = nil
	file_temporal_api_enums_v1_command_type_proto_goTypes = nil
	file_temporal_api_enums_v1_command_type_proto_depIdxs = nil
}
