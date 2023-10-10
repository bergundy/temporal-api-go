// The MIT License
//
// Copyright (c) 2022 Temporal Technologies Inc.  All rights reserved.
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

// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: temporal/api/enums/v1/failed_cause.proto

package enums

import (
	fmt "fmt"
	math "math"
	strconv "strconv"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Workflow tasks can fail for various reasons. Note that some of these reasons can only originate
// from the server, and some of them can only originate from the SDK/worker.
type WorkflowTaskFailedCause int32

const (
	WORKFLOW_TASK_FAILED_CAUSE_UNSPECIFIED WorkflowTaskFailedCause = 0
	// Between starting and completing the workflow task (with a workflow completion command), some
	// new command (like a signal) was processed into workflow history. The outstanding task will be
	// failed with this reason, and a worker must pick up a new task.
	WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND                                         WorkflowTaskFailedCause = 1
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_ACTIVITY_ATTRIBUTES                          WorkflowTaskFailedCause = 2
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_ACTIVITY_ATTRIBUTES                    WorkflowTaskFailedCause = 3
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_TIMER_ATTRIBUTES                                WorkflowTaskFailedCause = 4
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_TIMER_ATTRIBUTES                               WorkflowTaskFailedCause = 5
	WORKFLOW_TASK_FAILED_CAUSE_BAD_RECORD_MARKER_ATTRIBUTES                              WorkflowTaskFailedCause = 6
	WORKFLOW_TASK_FAILED_CAUSE_BAD_COMPLETE_WORKFLOW_EXECUTION_ATTRIBUTES                WorkflowTaskFailedCause = 7
	WORKFLOW_TASK_FAILED_CAUSE_BAD_FAIL_WORKFLOW_EXECUTION_ATTRIBUTES                    WorkflowTaskFailedCause = 8
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CANCEL_WORKFLOW_EXECUTION_ATTRIBUTES                  WorkflowTaskFailedCause = 9
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_EXTERNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 10
	WORKFLOW_TASK_FAILED_CAUSE_BAD_CONTINUE_AS_NEW_ATTRIBUTES                            WorkflowTaskFailedCause = 11
	WORKFLOW_TASK_FAILED_CAUSE_START_TIMER_DUPLICATE_ID                                  WorkflowTaskFailedCause = 12
	// The worker wishes to fail the task and have the next one be generated on a normal, not sticky
	// queue. Generally workers should prefer to use the explicit `ResetStickyTaskQueue` RPC call.
	WORKFLOW_TASK_FAILED_CAUSE_RESET_STICKY_TASK_QUEUE                  WorkflowTaskFailedCause = 13
	WORKFLOW_TASK_FAILED_CAUSE_WORKFLOW_WORKER_UNHANDLED_FAILURE        WorkflowTaskFailedCause = 14
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_WORKFLOW_EXECUTION_ATTRIBUTES WorkflowTaskFailedCause = 15
	WORKFLOW_TASK_FAILED_CAUSE_BAD_START_CHILD_EXECUTION_ATTRIBUTES     WorkflowTaskFailedCause = 16
	WORKFLOW_TASK_FAILED_CAUSE_FORCE_CLOSE_COMMAND                      WorkflowTaskFailedCause = 17
	WORKFLOW_TASK_FAILED_CAUSE_FAILOVER_CLOSE_COMMAND                   WorkflowTaskFailedCause = 18
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SIGNAL_INPUT_SIZE                    WorkflowTaskFailedCause = 19
	WORKFLOW_TASK_FAILED_CAUSE_RESET_WORKFLOW                           WorkflowTaskFailedCause = 20
	WORKFLOW_TASK_FAILED_CAUSE_BAD_BINARY                               WorkflowTaskFailedCause = 21
	WORKFLOW_TASK_FAILED_CAUSE_SCHEDULE_ACTIVITY_DUPLICATE_ID           WorkflowTaskFailedCause = 22
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SEARCH_ATTRIBUTES                    WorkflowTaskFailedCause = 23
	// The worker encountered a mismatch while replaying history between what was expected, and
	// what the workflow code actually did.
	WORKFLOW_TASK_FAILED_CAUSE_NON_DETERMINISTIC_ERROR                   WorkflowTaskFailedCause = 24
	WORKFLOW_TASK_FAILED_CAUSE_BAD_MODIFY_WORKFLOW_PROPERTIES_ATTRIBUTES WorkflowTaskFailedCause = 25
	// We send the below error codes to users when their requests would violate a size constraint
	// of their workflow. We do this to ensure that the state of their workflow does not become too
	// large because that can cause severe performance degradation. You can modify the thresholds for
	// each of these errors within your dynamic config.
	//
	// Spawning a new child workflow would cause this workflow to exceed its limit of pending child
	// workflows.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_CHILD_WORKFLOWS_LIMIT_EXCEEDED WorkflowTaskFailedCause = 26
	// Starting a new activity would cause this workflow to exceed its limit of pending activities
	// that we track.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_ACTIVITIES_LIMIT_EXCEEDED WorkflowTaskFailedCause = 27
	// A workflow has a buffer of signals that have not yet reached their destination. We return this
	// error when sending a new signal would exceed the capacity of this buffer.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_SIGNALS_LIMIT_EXCEEDED WorkflowTaskFailedCause = 28
	// Similarly, we have a buffer of pending requests to cancel other workflows. We return this error
	// when our capacity for pending cancel requests is already reached.
	WORKFLOW_TASK_FAILED_CAUSE_PENDING_REQUEST_CANCEL_LIMIT_EXCEEDED WorkflowTaskFailedCause = 29
	// Workflow execution update message (update.Acceptance, update.Rejection, or update.Response)
	// has wrong format, or missing required fields.
	WORKFLOW_TASK_FAILED_CAUSE_BAD_UPDATE_WORKFLOW_EXECUTION_MESSAGE WorkflowTaskFailedCause = 30
	// Similar to WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_COMMAND, but for updates.
	WORKFLOW_TASK_FAILED_CAUSE_UNHANDLED_UPDATE                              WorkflowTaskFailedCause = 31
	WORKFLOW_TASK_FAILED_CAUSE_BAD_SCHEDULE_NEXUS_OPERATION_ATTRIBUTES       WorkflowTaskFailedCause = 32
	WORKFLOW_TASK_FAILED_CAUSE_BAD_REQUEST_CANCEL_NEXUS_OPERATION_ATTRIBUTES WorkflowTaskFailedCause = 33
)

var WorkflowTaskFailedCause_name = map[int32]string{
	0:  "Unspecified",
	1:  "UnhandledCommand",
	2:  "BadScheduleActivityAttributes",
	3:  "BadRequestCancelActivityAttributes",
	4:  "BadStartTimerAttributes",
	5:  "BadCancelTimerAttributes",
	6:  "BadRecordMarkerAttributes",
	7:  "BadCompleteWorkflowExecutionAttributes",
	8:  "BadFailWorkflowExecutionAttributes",
	9:  "BadCancelWorkflowExecutionAttributes",
	10: "BadRequestCancelExternalWorkflowExecutionAttributes",
	11: "BadContinueAsNewAttributes",
	12: "StartTimerDuplicateId",
	13: "ResetStickyTaskQueue",
	14: "WorkflowWorkerUnhandledFailure",
	15: "BadSignalWorkflowExecutionAttributes",
	16: "BadStartChildExecutionAttributes",
	17: "ForceCloseCommand",
	18: "FailoverCloseCommand",
	19: "BadSignalInputSize",
	20: "ResetWorkflow",
	21: "BadBinary",
	22: "ScheduleActivityDuplicateId",
	23: "BadSearchAttributes",
	24: "NonDeterministicError",
	25: "BadModifyWorkflowPropertiesAttributes",
	26: "PendingChildWorkflowsLimitExceeded",
	27: "PendingActivitiesLimitExceeded",
	28: "PendingSignalsLimitExceeded",
	29: "PendingRequestCancelLimitExceeded",
	30: "BadUpdateWorkflowExecutionMessage",
	31: "UnhandledUpdate",
	32: "BadScheduleNexusOperationAttributes",
	33: "BadRequestCancelNexusOperationAttributes",
}

var WorkflowTaskFailedCause_value = map[string]int32{
	"Unspecified":                                         0,
	"UnhandledCommand":                                    1,
	"BadScheduleActivityAttributes":                       2,
	"BadRequestCancelActivityAttributes":                  3,
	"BadStartTimerAttributes":                             4,
	"BadCancelTimerAttributes":                            5,
	"BadRecordMarkerAttributes":                           6,
	"BadCompleteWorkflowExecutionAttributes":              7,
	"BadFailWorkflowExecutionAttributes":                  8,
	"BadCancelWorkflowExecutionAttributes":                9,
	"BadRequestCancelExternalWorkflowExecutionAttributes": 10,
	"BadContinueAsNewAttributes":                          11,
	"StartTimerDuplicateId":                               12,
	"ResetStickyTaskQueue":                                13,
	"WorkflowWorkerUnhandledFailure":                      14,
	"BadSignalWorkflowExecutionAttributes":                15,
	"BadStartChildExecutionAttributes":                    16,
	"ForceCloseCommand":                                   17,
	"FailoverCloseCommand":                                18,
	"BadSignalInputSize":                                  19,
	"ResetWorkflow":                                       20,
	"BadBinary":                                           21,
	"ScheduleActivityDuplicateId":                         22,
	"BadSearchAttributes":                                 23,
	"NonDeterministicError":                               24,
	"BadModifyWorkflowPropertiesAttributes":               25,
	"PendingChildWorkflowsLimitExceeded":                  26,
	"PendingActivitiesLimitExceeded":                      27,
	"PendingSignalsLimitExceeded":                         28,
	"PendingRequestCancelLimitExceeded":                   29,
	"BadUpdateWorkflowExecutionMessage":                   30,
	"UnhandledUpdate":                                     31,
	"BadScheduleNexusOperationAttributes":                 32,
	"BadRequestCancelNexusOperationAttributes":            33,
}

func (WorkflowTaskFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{0}
}

type StartChildWorkflowExecutionFailedCause int32

const (
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED             StartChildWorkflowExecutionFailedCause = 0
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_WORKFLOW_ALREADY_EXISTS StartChildWorkflowExecutionFailedCause = 1
	START_CHILD_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND     StartChildWorkflowExecutionFailedCause = 2
)

var StartChildWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "WorkflowAlreadyExists",
	2: "NamespaceNotFound",
}

var StartChildWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":           0,
	"WorkflowAlreadyExists": 1,
	"NamespaceNotFound":     2,
}

func (StartChildWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{1}
}

type CancelExternalWorkflowExecutionFailedCause int32

const (
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           CancelExternalWorkflowExecutionFailedCause = 0
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND CancelExternalWorkflowExecutionFailedCause = 1
	CANCEL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND                   CancelExternalWorkflowExecutionFailedCause = 2
)

var CancelExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
	2: "NamespaceNotFound",
}

var CancelExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
	"NamespaceNotFound":                 2,
}

func (CancelExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{2}
}

type SignalExternalWorkflowExecutionFailedCause int32

const (
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_UNSPECIFIED                           SignalExternalWorkflowExecutionFailedCause = 0
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_EXTERNAL_WORKFLOW_EXECUTION_NOT_FOUND SignalExternalWorkflowExecutionFailedCause = 1
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_NAMESPACE_NOT_FOUND                   SignalExternalWorkflowExecutionFailedCause = 2
	// Signal count limit is per workflow and controlled by server dynamic config "history.maximumSignalsPerExecution"
	SIGNAL_EXTERNAL_WORKFLOW_EXECUTION_FAILED_CAUSE_SIGNAL_COUNT_LIMIT_EXCEEDED SignalExternalWorkflowExecutionFailedCause = 3
)

var SignalExternalWorkflowExecutionFailedCause_name = map[int32]string{
	0: "Unspecified",
	1: "ExternalWorkflowExecutionNotFound",
	2: "NamespaceNotFound",
	3: "SignalCountLimitExceeded",
}

var SignalExternalWorkflowExecutionFailedCause_value = map[string]int32{
	"Unspecified":                       0,
	"ExternalWorkflowExecutionNotFound": 1,
	"NamespaceNotFound":                 2,
	"SignalCountLimitExceeded":          3,
}

func (SignalExternalWorkflowExecutionFailedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{3}
}

type ResourceExhaustedCause int32

const (
	RESOURCE_EXHAUSTED_CAUSE_UNSPECIFIED ResourceExhaustedCause = 0
	// Caller exceeds request per second limit.
	RESOURCE_EXHAUSTED_CAUSE_RPS_LIMIT ResourceExhaustedCause = 1
	// Caller exceeds max concurrent request limit.
	RESOURCE_EXHAUSTED_CAUSE_CONCURRENT_LIMIT ResourceExhaustedCause = 2
	// System overloaded.
	RESOURCE_EXHAUSTED_CAUSE_SYSTEM_OVERLOADED ResourceExhaustedCause = 3
	// Namespace exceeds persistence rate limit.
	RESOURCE_EXHAUSTED_CAUSE_PERSISTENCE_LIMIT ResourceExhaustedCause = 4
	// Workflow is busy
	RESOURCE_EXHAUSTED_CAUSE_BUSY_WORKFLOW ResourceExhaustedCause = 5
	// Caller exceeds action per second limit.
	RESOURCE_EXHAUSTED_CAUSE_APS_LIMIT ResourceExhaustedCause = 6
)

var ResourceExhaustedCause_name = map[int32]string{
	0: "Unspecified",
	1: "RpsLimit",
	2: "ConcurrentLimit",
	3: "SystemOverloaded",
	4: "PersistenceLimit",
	5: "BusyWorkflow",
	6: "ApsLimit",
}

var ResourceExhaustedCause_value = map[string]int32{
	"Unspecified":      0,
	"RpsLimit":         1,
	"ConcurrentLimit":  2,
	"SystemOverloaded": 3,
	"PersistenceLimit": 4,
	"BusyWorkflow":     5,
	"ApsLimit":         6,
}

func (ResourceExhaustedCause) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b293cf8d1d965f2d, []int{4}
}

func init() {
	proto.RegisterEnum("temporal.api.enums.v1.WorkflowTaskFailedCause", WorkflowTaskFailedCause_name, WorkflowTaskFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.StartChildWorkflowExecutionFailedCause", StartChildWorkflowExecutionFailedCause_name, StartChildWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.CancelExternalWorkflowExecutionFailedCause", CancelExternalWorkflowExecutionFailedCause_name, CancelExternalWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.SignalExternalWorkflowExecutionFailedCause", SignalExternalWorkflowExecutionFailedCause_name, SignalExternalWorkflowExecutionFailedCause_value)
	proto.RegisterEnum("temporal.api.enums.v1.ResourceExhaustedCause", ResourceExhaustedCause_name, ResourceExhaustedCause_value)
}

func init() {
	proto.RegisterFile("temporal/api/enums/v1/failed_cause.proto", fileDescriptor_b293cf8d1d965f2d)
}

var fileDescriptor_b293cf8d1d965f2d = []byte{
	// 1075 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x57, 0x4f, 0x73, 0xdb, 0xc4,
	0x1b, 0xb6, 0xdc, 0x3f, 0xbf, 0x1f, 0xcb, 0x3f, 0xb1, 0xd0, 0xa6, 0x14, 0x10, 0x94, 0x81, 0x4c,
	0x6b, 0xc0, 0x69, 0x5a, 0xda, 0x4c, 0x1d, 0x98, 0x74, 0xbd, 0x7a, 0x1d, 0xef, 0x44, 0x5a, 0xa9,
	0xbb, 0xab, 0xc4, 0xee, 0x65, 0x47, 0xa4, 0x6e, 0xab, 0xa9, 0x1b, 0x7b, 0x1c, 0xa7, 0xe4, 0xc8,
	0x47, 0x80, 0x6f, 0xc1, 0xf0, 0x19, 0xf8, 0x00, 0x5c, 0x98, 0xc9, 0xb1, 0x47, 0xe2, 0x5c, 0x18,
	0x4e, 0x9d, 0x61, 0xb8, 0x33, 0x72, 0xec, 0x44, 0x71, 0x1c, 0x49, 0xe6, 0x26, 0x7b, 0x9f, 0xe7,
	0xd9, 0x7d, 0x9f, 0x7d, 0xf7, 0x7d, 0x77, 0xd1, 0xf5, 0x7e, 0xeb, 0x79, 0xb7, 0xd3, 0x0b, 0xdb,
	0x0b, 0x61, 0x37, 0x5a, 0x68, 0x6d, 0xed, 0x3c, 0xdf, 0x5e, 0x78, 0xb1, 0xb8, 0xf0, 0x38, 0x8c,
	0xda, 0xad, 0x47, 0x7a, 0x33, 0xdc, 0xd9, 0x6e, 0x95, 0xbb, 0xbd, 0x4e, 0xbf, 0x83, 0x2f, 0x8d,
	0x91, 0xe5, 0xb0, 0x1b, 0x95, 0x87, 0xc8, 0xf2, 0x8b, 0xc5, 0xd2, 0x3f, 0x26, 0x9a, 0xdb, 0xe8,
	0xf4, 0x9e, 0x3d, 0x6e, 0x77, 0xbe, 0x57, 0xe1, 0xf6, 0xb3, 0xda, 0x90, 0x49, 0x63, 0x22, 0x2e,
	0xa1, 0xf9, 0x0d, 0x4f, 0xac, 0xd5, 0x1c, 0x6f, 0x43, 0x2b, 0x22, 0xd7, 0x74, 0x8d, 0x30, 0x07,
	0x6c, 0x4d, 0x49, 0x20, 0x41, 0x07, 0x5c, 0xfa, 0x40, 0x59, 0x8d, 0x81, 0x6d, 0x16, 0xf0, 0x4d,
	0xf4, 0x65, 0x2a, 0xb6, 0x4e, 0xb8, 0x3d, 0xfc, 0xed, 0xb9, 0x2e, 0xe1, 0xb6, 0x69, 0xe0, 0x15,
	0xb4, 0x9c, 0xc2, 0xa8, 0x12, 0x5b, 0x4b, 0x5a, 0x07, 0x3b, 0x70, 0x40, 0x13, 0xaa, 0xd8, 0x3a,
	0x53, 0x4d, 0x4d, 0x94, 0x12, 0xac, 0x1a, 0x28, 0x90, 0x66, 0x11, 0x03, 0x22, 0x19, 0x02, 0x02,
	0x1e, 0x04, 0x20, 0x95, 0xa6, 0x84, 0x53, 0x70, 0xa6, 0xca, 0x9c, 0xc3, 0xf7, 0xd0, 0x9d, 0xac,
	0x75, 0x28, 0x22, 0x94, 0x56, 0xcc, 0x05, 0x91, 0xa4, 0x9e, 0xc7, 0x15, 0x74, 0x37, 0x83, 0x3a,
	0x9a, 0xf9, 0x14, 0xf7, 0x02, 0x5e, 0x46, 0x4b, 0x99, 0xab, 0xa7, 0x9e, 0xb0, 0xb5, 0x4b, 0xc4,
	0xda, 0x49, 0xf2, 0x45, 0xcc, 0x10, 0x64, 0x4d, 0xec, 0xb9, 0xbe, 0x03, 0x0a, 0xf4, 0x11, 0x0e,
	0x1a, 0x40, 0x03, 0xc5, 0x3c, 0x9e, 0x94, 0xfa, 0x5f, 0x0e, 0x17, 0xe3, 0x3f, 0x32, 0x64, 0xfe,
	0x8f, 0x57, 0x11, 0xcd, 0x67, 0x45, 0xba, 0xd0, 0x6b, 0xb8, 0x81, 0xd4, 0x6c, 0xbb, 0x0a, 0x0d,
	0x05, 0x82, 0x93, 0x2c, 0x65, 0x84, 0xbf, 0x45, 0xf7, 0x32, 0x4d, 0xe3, 0x8a, 0xf1, 0x00, 0x34,
	0x91, 0x9a, 0xc3, 0x46, 0x92, 0xfe, 0x3a, 0x5e, 0x42, 0xb7, 0x53, 0xe8, 0xc9, 0x1c, 0xb1, 0x03,
	0xdf, 0x61, 0x94, 0x28, 0xd0, 0xcc, 0x36, 0xdf, 0xc0, 0x77, 0xd1, 0xad, 0x14, 0xa2, 0x00, 0x09,
	0x4a, 0x4b, 0xc5, 0xe8, 0x5a, 0xf3, 0x70, 0xf8, 0x41, 0x00, 0x01, 0x98, 0x6f, 0xe2, 0xfb, 0xe8,
	0x9b, 0x14, 0xde, 0xd1, 0x50, 0xfc, 0x01, 0x22, 0x71, 0xc4, 0x62, 0x58, 0x20, 0xc0, 0x7c, 0x2b,
	0xc7, 0xa6, 0x48, 0xb6, 0x9a, 0x6d, 0xdd, 0xdb, 0x98, 0xa2, 0x95, 0x5c, 0x67, 0x84, 0xd6, 0x99,
	0x63, 0x4f, 0x17, 0x31, 0xf1, 0x2d, 0x54, 0x4e, 0x11, 0xa9, 0x79, 0x82, 0x82, 0xa6, 0x8e, 0x27,
	0xe1, 0xa8, 0x48, 0xbc, 0x83, 0xef, 0xa0, 0xc5, 0x34, 0x0e, 0x61, 0x8e, 0xb7, 0x0e, 0x62, 0x82,
	0x86, 0xf1, 0xd7, 0xe8, 0x66, 0xbe, 0xc0, 0x19, 0xf7, 0x03, 0xa5, 0x25, 0x7b, 0x08, 0xe6, 0xbb,
	0xf8, 0x2b, 0x74, 0x23, 0x73, 0xa3, 0xc6, 0x00, 0xf3, 0x3d, 0x7c, 0x03, 0x7d, 0x9e, 0x31, 0x49,
	0x95, 0x71, 0x22, 0x9a, 0xe6, 0xa5, 0x8c, 0xd4, 0x3b, 0x5d, 0xe7, 0x4e, 0x64, 0xd0, 0xe5, 0x3c,
	0xe1, 0x00, 0x11, 0xb4, 0x9e, 0xf4, 0x7b, 0x2e, 0x23, 0xef, 0xb8, 0xc7, 0xb5, 0x0d, 0x0a, 0x84,
	0xcb, 0x38, 0x8b, 0xd3, 0x4f, 0x83, 0x10, 0x9e, 0x30, 0xaf, 0xe0, 0x3a, 0xb2, 0x33, 0x66, 0x73,
	0x3d, 0x9b, 0xd5, 0x9a, 0xc7, 0x59, 0xe3, 0x0b, 0xcf, 0x07, 0xa1, 0x18, 0xc8, 0xe4, 0x0a, 0xde,
	0xcf, 0xa8, 0x2d, 0x3e, 0x70, 0x9b, 0xf1, 0xd5, 0x51, 0xd2, 0x8c, 0x81, 0x52, 0x3b, 0xcc, 0x65,
	0x4a, 0x43, 0x83, 0x02, 0xd8, 0x60, 0x9b, 0x57, 0x33, 0x0e, 0xc2, 0x58, 0x66, 0x64, 0x5e, 0xbc,
	0x88, 0x09, 0x85, 0x0f, 0x32, 0xfc, 0x1f, 0x2b, 0x1c, 0xe6, 0xc4, 0x29, 0xfa, 0x87, 0xd8, 0x46,
	0xf7, 0x73, 0xd0, 0x27, 0xea, 0xd2, 0x84, 0xca, 0x47, 0x19, 0x2a, 0xb1, 0xaf, 0x81, 0x6f, 0x93,
	0xe9, 0x25, 0xdb, 0x05, 0x29, 0xc9, 0x2a, 0x98, 0x16, 0x5e, 0x40, 0x5f, 0xe4, 0x6a, 0xb4, 0x87,
	0x5a, 0xe6, 0xc7, 0xb8, 0x86, 0xaa, 0x79, 0xfb, 0x2c, 0x87, 0x46, 0x20, 0x75, 0xbc, 0x97, 0x64,
	0xf2, 0xf8, 0x7e, 0x82, 0x1d, 0x54, 0x9f, 0xad, 0x30, 0xa7, 0xa8, 0x5d, 0x2b, 0xfd, 0x6d, 0xa0,
	0x79, 0xd9, 0x0f, 0x7b, 0x7d, 0xfa, 0x34, 0x6a, 0x3f, 0x1a, 0xdf, 0x40, 0x60, 0xb7, 0xb5, 0xb9,
	0xd3, 0x8f, 0x3a, 0x5b, 0xc9, 0x6b, 0xc8, 0x32, 0x5a, 0x4a, 0x56, 0x97, 0x29, 0xee, 0xa4, 0xdc,
	0x4b, 0x56, 0x11, 0x9d, 0x85, 0x7c, 0x34, 0x4e, 0x1c, 0x01, 0xc4, 0x6e, 0x6a, 0x68, 0x30, 0xa9,
	0xa4, 0x69, 0xc4, 0x25, 0x70, 0x16, 0x21, 0x4e, 0x5c, 0x90, 0x3e, 0xa1, 0xf1, 0x41, 0x53, 0xba,
	0xe6, 0x05, 0xdc, 0x36, 0x8b, 0xa5, 0x9f, 0x8a, 0xa8, 0x44, 0xc3, 0xad, 0xcd, 0x56, 0x1b, 0x76,
	0xfb, 0xad, 0xde, 0x56, 0xd8, 0x4e, 0x8d, 0x7c, 0x05, 0x2d, 0xe7, 0x68, 0x72, 0x29, 0xd1, 0x37,
	0x51, 0x30, 0xab, 0x40, 0x1a, 0xf0, 0x38, 0x14, 0x23, 0x36, 0x76, 0x56, 0xe9, 0xe9, 0x9e, 0x0c,
	0x8a, 0xa8, 0x24, 0xa3, 0x27, 0x5b, 0x61, 0x6e, 0x4f, 0x46, 0xb5, 0xfb, 0xbf, 0x7b, 0x32, 0xab,
	0xc0, 0x0c, 0x9e, 0xcc, 0x2a, 0x3d, 0xd5, 0x13, 0xec, 0xa1, 0xb5, 0x59, 0x85, 0x46, 0x78, 0xea,
	0x05, 0x5c, 0x4d, 0xd6, 0x9e, 0x73, 0xa5, 0xdf, 0x8b, 0xe8, 0xb2, 0x68, 0x6d, 0x77, 0x76, 0x7a,
	0x9b, 0x2d, 0xd8, 0x7d, 0x1a, 0xee, 0x6c, 0xf7, 0xc7, 0x86, 0x5e, 0x47, 0x9f, 0x09, 0x90, 0x5e,
	0x10, 0xb7, 0x5f, 0x68, 0xd4, 0x49, 0x20, 0xd5, 0x19, 0xce, 0xcd, 0xa3, 0x4f, 0xcf, 0x44, 0x0a,
	0x7f, 0x54, 0x37, 0x4d, 0x23, 0xee, 0xa3, 0x67, 0xe2, 0xa8, 0xc7, 0x69, 0x20, 0x04, 0x8c, 0x17,
	0x69, 0x16, 0x71, 0x19, 0x95, 0xce, 0x84, 0xcb, 0xa6, 0x54, 0xe0, 0xea, 0xb8, 0xc9, 0x3b, 0x1e,
	0x19, 0xc6, 0x92, 0x8a, 0xf7, 0x41, 0x48, 0x26, 0x15, 0x70, 0x0a, 0x23, 0xfd, 0xf3, 0xf1, 0x33,
	0xe6, 0x4c, 0x7c, 0x35, 0x90, 0xc7, 0x7d, 0xcc, 0xbc, 0x90, 0x1a, 0x22, 0x39, 0x0a, 0xf1, 0x62,
	0xf5, 0x57, 0x63, 0x6f, 0xdf, 0x2a, 0xbc, 0xdc, 0xb7, 0x0a, 0xaf, 0xf6, 0x2d, 0xe3, 0x87, 0x81,
	0x65, 0xfc, 0x3c, 0xb0, 0x8c, 0xdf, 0x06, 0x96, 0xb1, 0x37, 0xb0, 0x8c, 0x3f, 0x06, 0x96, 0xf1,
	0xe7, 0xc0, 0x2a, 0xbc, 0x1a, 0x58, 0xc6, 0x8f, 0x07, 0x56, 0x61, 0xef, 0xc0, 0x2a, 0xbc, 0x3c,
	0xb0, 0x0a, 0xe8, 0x4a, 0xd4, 0x29, 0x4f, 0x7d, 0x87, 0x55, 0xcd, 0x44, 0x8e, 0xfb, 0xf1, 0x83,
	0xcd, 0x37, 0x1e, 0x5e, 0x7b, 0x92, 0x40, 0x47, 0x9d, 0x13, 0x4f, 0xbc, 0xe5, 0xe1, 0xc7, 0x2f,
	0xc5, 0x39, 0x35, 0x02, 0x44, 0x9d, 0x32, 0xe9, 0x46, 0x65, 0x18, 0x0a, 0xae, 0x2f, 0xfe, 0x55,
	0xbc, 0x7a, 0x3c, 0x52, 0xa9, 0x90, 0x6e, 0x54, 0xa9, 0x0c, 0xc7, 0x2a, 0x95, 0xf5, 0xc5, 0xef,
	0x2e, 0x0e, 0xdf, 0x84, 0xb7, 0xff, 0x0d, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x88, 0xbb, 0x41, 0x3f,
	0x0e, 0x00, 0x00,
}

func (x WorkflowTaskFailedCause) String() string {
	s, ok := WorkflowTaskFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x StartChildWorkflowExecutionFailedCause) String() string {
	s, ok := StartChildWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x CancelExternalWorkflowExecutionFailedCause) String() string {
	s, ok := CancelExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x SignalExternalWorkflowExecutionFailedCause) String() string {
	s, ok := SignalExternalWorkflowExecutionFailedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
func (x ResourceExhaustedCause) String() string {
	s, ok := ResourceExhaustedCause_name[int32(x)]
	if ok {
		return s
	}
	return strconv.Itoa(int(x))
}
