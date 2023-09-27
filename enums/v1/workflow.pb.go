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
// source: temporal/api/enums/v1/workflow.proto

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

// Defines how new runs of a workflow with a particular ID may or may not be allowed. Note that
// it is *never* valid to have two actively running instances of the same workflow id.
type WorkflowIdReusePolicy int32

const (
	WorkflowIdReusePolicy_WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED WorkflowIdReusePolicy = 0
	// Allow starting a workflow execution using the same workflow id.
	WorkflowIdReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE WorkflowIdReusePolicy = 1
	// Allow starting a workflow execution using the same workflow id, only when the last
	// execution's final state is one of [terminated, cancelled, timed out, failed].
	WorkflowIdReusePolicy_WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY WorkflowIdReusePolicy = 2
	// Do not permit re-use of the workflow id for this workflow. Future start workflow requests
	// could potentially change the policy, allowing re-use of the workflow id.
	WorkflowIdReusePolicy_WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE WorkflowIdReusePolicy = 3
	// If a workflow is running using the same workflow ID, terminate it and start a new one.
	// If no running workflow, then the behavior is the same as ALLOW_DUPLICATE
	WorkflowIdReusePolicy_WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING WorkflowIdReusePolicy = 4
)

// Enum value maps for WorkflowIdReusePolicy.
var (
	WorkflowIdReusePolicy_name = map[int32]string{
		0: "WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED",
		1: "WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE",
		2: "WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY",
		3: "WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE",
		4: "WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING",
	}
	WorkflowIdReusePolicy_value = map[string]int32{
		"WORKFLOW_ID_REUSE_POLICY_UNSPECIFIED":                 0,
		"WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE":             1,
		"WORKFLOW_ID_REUSE_POLICY_ALLOW_DUPLICATE_FAILED_ONLY": 2,
		"WORKFLOW_ID_REUSE_POLICY_REJECT_DUPLICATE":            3,
		"WORKFLOW_ID_REUSE_POLICY_TERMINATE_IF_RUNNING":        4,
	}
)

func (x WorkflowIdReusePolicy) Enum() *WorkflowIdReusePolicy {
	p := new(WorkflowIdReusePolicy)
	*p = x
	return p
}

func (x WorkflowIdReusePolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkflowIdReusePolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[0].Descriptor()
}

func (WorkflowIdReusePolicy) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[0]
}

func (x WorkflowIdReusePolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowIdReusePolicy.Descriptor instead.
func (WorkflowIdReusePolicy) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{0}
}

// Defines how child workflows will react to their parent completing
type ParentClosePolicy int32

const (
	ParentClosePolicy_PARENT_CLOSE_POLICY_UNSPECIFIED ParentClosePolicy = 0
	// The child workflow will also terminate
	ParentClosePolicy_PARENT_CLOSE_POLICY_TERMINATE ParentClosePolicy = 1
	// The child workflow will do nothing
	ParentClosePolicy_PARENT_CLOSE_POLICY_ABANDON ParentClosePolicy = 2
	// Cancellation will be requested of the child workflow
	ParentClosePolicy_PARENT_CLOSE_POLICY_REQUEST_CANCEL ParentClosePolicy = 3
)

// Enum value maps for ParentClosePolicy.
var (
	ParentClosePolicy_name = map[int32]string{
		0: "PARENT_CLOSE_POLICY_UNSPECIFIED",
		1: "PARENT_CLOSE_POLICY_TERMINATE",
		2: "PARENT_CLOSE_POLICY_ABANDON",
		3: "PARENT_CLOSE_POLICY_REQUEST_CANCEL",
	}
	ParentClosePolicy_value = map[string]int32{
		"PARENT_CLOSE_POLICY_UNSPECIFIED":    0,
		"PARENT_CLOSE_POLICY_TERMINATE":      1,
		"PARENT_CLOSE_POLICY_ABANDON":        2,
		"PARENT_CLOSE_POLICY_REQUEST_CANCEL": 3,
	}
)

func (x ParentClosePolicy) Enum() *ParentClosePolicy {
	p := new(ParentClosePolicy)
	*p = x
	return p
}

func (x ParentClosePolicy) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ParentClosePolicy) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[1].Descriptor()
}

func (ParentClosePolicy) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[1]
}

func (x ParentClosePolicy) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ParentClosePolicy.Descriptor instead.
func (ParentClosePolicy) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{1}
}

type ContinueAsNewInitiator int32

const (
	ContinueAsNewInitiator_CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED ContinueAsNewInitiator = 0
	// The workflow itself requested to continue as new
	ContinueAsNewInitiator_CONTINUE_AS_NEW_INITIATOR_WORKFLOW ContinueAsNewInitiator = 1
	// The workflow continued as new because it is retrying
	ContinueAsNewInitiator_CONTINUE_AS_NEW_INITIATOR_RETRY ContinueAsNewInitiator = 2
	// The workflow continued as new because cron has triggered a new execution
	ContinueAsNewInitiator_CONTINUE_AS_NEW_INITIATOR_CRON_SCHEDULE ContinueAsNewInitiator = 3
)

// Enum value maps for ContinueAsNewInitiator.
var (
	ContinueAsNewInitiator_name = map[int32]string{
		0: "CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED",
		1: "CONTINUE_AS_NEW_INITIATOR_WORKFLOW",
		2: "CONTINUE_AS_NEW_INITIATOR_RETRY",
		3: "CONTINUE_AS_NEW_INITIATOR_CRON_SCHEDULE",
	}
	ContinueAsNewInitiator_value = map[string]int32{
		"CONTINUE_AS_NEW_INITIATOR_UNSPECIFIED":   0,
		"CONTINUE_AS_NEW_INITIATOR_WORKFLOW":      1,
		"CONTINUE_AS_NEW_INITIATOR_RETRY":         2,
		"CONTINUE_AS_NEW_INITIATOR_CRON_SCHEDULE": 3,
	}
)

func (x ContinueAsNewInitiator) Enum() *ContinueAsNewInitiator {
	p := new(ContinueAsNewInitiator)
	*p = x
	return p
}

func (x ContinueAsNewInitiator) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContinueAsNewInitiator) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[2].Descriptor()
}

func (ContinueAsNewInitiator) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[2]
}

func (x ContinueAsNewInitiator) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContinueAsNewInitiator.Descriptor instead.
func (ContinueAsNewInitiator) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{2}
}

// (-- api-linter: core::0216::synonyms=disabled
//
//	aip.dev/not-precedent: There is WorkflowExecutionState already in another package. --)
type WorkflowExecutionStatus int32

const (
	WorkflowExecutionStatus_WORKFLOW_EXECUTION_STATUS_UNSPECIFIED WorkflowExecutionStatus = 0
	// Value 1 is hardcoded in SQL persistence.
	WorkflowExecutionStatus_WORKFLOW_EXECUTION_STATUS_RUNNING          WorkflowExecutionStatus = 1
	WorkflowExecutionStatus_WORKFLOW_EXECUTION_STATUS_COMPLETED        WorkflowExecutionStatus = 2
	WorkflowExecutionStatus_WORKFLOW_EXECUTION_STATUS_FAILED           WorkflowExecutionStatus = 3
	WorkflowExecutionStatus_WORKFLOW_EXECUTION_STATUS_CANCELED         WorkflowExecutionStatus = 4
	WorkflowExecutionStatus_WORKFLOW_EXECUTION_STATUS_TERMINATED       WorkflowExecutionStatus = 5
	WorkflowExecutionStatus_WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW WorkflowExecutionStatus = 6
	WorkflowExecutionStatus_WORKFLOW_EXECUTION_STATUS_TIMED_OUT        WorkflowExecutionStatus = 7
)

// Enum value maps for WorkflowExecutionStatus.
var (
	WorkflowExecutionStatus_name = map[int32]string{
		0: "WORKFLOW_EXECUTION_STATUS_UNSPECIFIED",
		1: "WORKFLOW_EXECUTION_STATUS_RUNNING",
		2: "WORKFLOW_EXECUTION_STATUS_COMPLETED",
		3: "WORKFLOW_EXECUTION_STATUS_FAILED",
		4: "WORKFLOW_EXECUTION_STATUS_CANCELED",
		5: "WORKFLOW_EXECUTION_STATUS_TERMINATED",
		6: "WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW",
		7: "WORKFLOW_EXECUTION_STATUS_TIMED_OUT",
	}
	WorkflowExecutionStatus_value = map[string]int32{
		"WORKFLOW_EXECUTION_STATUS_UNSPECIFIED":      0,
		"WORKFLOW_EXECUTION_STATUS_RUNNING":          1,
		"WORKFLOW_EXECUTION_STATUS_COMPLETED":        2,
		"WORKFLOW_EXECUTION_STATUS_FAILED":           3,
		"WORKFLOW_EXECUTION_STATUS_CANCELED":         4,
		"WORKFLOW_EXECUTION_STATUS_TERMINATED":       5,
		"WORKFLOW_EXECUTION_STATUS_CONTINUED_AS_NEW": 6,
		"WORKFLOW_EXECUTION_STATUS_TIMED_OUT":        7,
	}
)

func (x WorkflowExecutionStatus) Enum() *WorkflowExecutionStatus {
	p := new(WorkflowExecutionStatus)
	*p = x
	return p
}

func (x WorkflowExecutionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WorkflowExecutionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[3].Descriptor()
}

func (WorkflowExecutionStatus) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[3]
}

func (x WorkflowExecutionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WorkflowExecutionStatus.Descriptor instead.
func (WorkflowExecutionStatus) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{3}
}

type PendingActivityState int32

const (
	PendingActivityState_PENDING_ACTIVITY_STATE_UNSPECIFIED      PendingActivityState = 0
	PendingActivityState_PENDING_ACTIVITY_STATE_SCHEDULED        PendingActivityState = 1
	PendingActivityState_PENDING_ACTIVITY_STATE_STARTED          PendingActivityState = 2
	PendingActivityState_PENDING_ACTIVITY_STATE_CANCEL_REQUESTED PendingActivityState = 3
)

// Enum value maps for PendingActivityState.
var (
	PendingActivityState_name = map[int32]string{
		0: "PENDING_ACTIVITY_STATE_UNSPECIFIED",
		1: "PENDING_ACTIVITY_STATE_SCHEDULED",
		2: "PENDING_ACTIVITY_STATE_STARTED",
		3: "PENDING_ACTIVITY_STATE_CANCEL_REQUESTED",
	}
	PendingActivityState_value = map[string]int32{
		"PENDING_ACTIVITY_STATE_UNSPECIFIED":      0,
		"PENDING_ACTIVITY_STATE_SCHEDULED":        1,
		"PENDING_ACTIVITY_STATE_STARTED":          2,
		"PENDING_ACTIVITY_STATE_CANCEL_REQUESTED": 3,
	}
)

func (x PendingActivityState) Enum() *PendingActivityState {
	p := new(PendingActivityState)
	*p = x
	return p
}

func (x PendingActivityState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PendingActivityState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[4].Descriptor()
}

func (PendingActivityState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[4]
}

func (x PendingActivityState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PendingActivityState.Descriptor instead.
func (PendingActivityState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{4}
}

type PendingWorkflowTaskState int32

const (
	PendingWorkflowTaskState_PENDING_WORKFLOW_TASK_STATE_UNSPECIFIED PendingWorkflowTaskState = 0
	PendingWorkflowTaskState_PENDING_WORKFLOW_TASK_STATE_SCHEDULED   PendingWorkflowTaskState = 1
	PendingWorkflowTaskState_PENDING_WORKFLOW_TASK_STATE_STARTED     PendingWorkflowTaskState = 2
)

// Enum value maps for PendingWorkflowTaskState.
var (
	PendingWorkflowTaskState_name = map[int32]string{
		0: "PENDING_WORKFLOW_TASK_STATE_UNSPECIFIED",
		1: "PENDING_WORKFLOW_TASK_STATE_SCHEDULED",
		2: "PENDING_WORKFLOW_TASK_STATE_STARTED",
	}
	PendingWorkflowTaskState_value = map[string]int32{
		"PENDING_WORKFLOW_TASK_STATE_UNSPECIFIED": 0,
		"PENDING_WORKFLOW_TASK_STATE_SCHEDULED":   1,
		"PENDING_WORKFLOW_TASK_STATE_STARTED":     2,
	}
)

func (x PendingWorkflowTaskState) Enum() *PendingWorkflowTaskState {
	p := new(PendingWorkflowTaskState)
	*p = x
	return p
}

func (x PendingWorkflowTaskState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PendingWorkflowTaskState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[5].Descriptor()
}

func (PendingWorkflowTaskState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[5]
}

func (x PendingWorkflowTaskState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PendingWorkflowTaskState.Descriptor instead.
func (PendingWorkflowTaskState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{5}
}

type HistoryEventFilterType int32

const (
	HistoryEventFilterType_HISTORY_EVENT_FILTER_TYPE_UNSPECIFIED HistoryEventFilterType = 0
	HistoryEventFilterType_HISTORY_EVENT_FILTER_TYPE_ALL_EVENT   HistoryEventFilterType = 1
	HistoryEventFilterType_HISTORY_EVENT_FILTER_TYPE_CLOSE_EVENT HistoryEventFilterType = 2
)

// Enum value maps for HistoryEventFilterType.
var (
	HistoryEventFilterType_name = map[int32]string{
		0: "HISTORY_EVENT_FILTER_TYPE_UNSPECIFIED",
		1: "HISTORY_EVENT_FILTER_TYPE_ALL_EVENT",
		2: "HISTORY_EVENT_FILTER_TYPE_CLOSE_EVENT",
	}
	HistoryEventFilterType_value = map[string]int32{
		"HISTORY_EVENT_FILTER_TYPE_UNSPECIFIED": 0,
		"HISTORY_EVENT_FILTER_TYPE_ALL_EVENT":   1,
		"HISTORY_EVENT_FILTER_TYPE_CLOSE_EVENT": 2,
	}
)

func (x HistoryEventFilterType) Enum() *HistoryEventFilterType {
	p := new(HistoryEventFilterType)
	*p = x
	return p
}

func (x HistoryEventFilterType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (HistoryEventFilterType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[6].Descriptor()
}

func (HistoryEventFilterType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[6]
}

func (x HistoryEventFilterType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use HistoryEventFilterType.Descriptor instead.
func (HistoryEventFilterType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{6}
}

type RetryState int32

const (
	RetryState_RETRY_STATE_UNSPECIFIED              RetryState = 0
	RetryState_RETRY_STATE_IN_PROGRESS              RetryState = 1
	RetryState_RETRY_STATE_NON_RETRYABLE_FAILURE    RetryState = 2
	RetryState_RETRY_STATE_TIMEOUT                  RetryState = 3
	RetryState_RETRY_STATE_MAXIMUM_ATTEMPTS_REACHED RetryState = 4
	RetryState_RETRY_STATE_RETRY_POLICY_NOT_SET     RetryState = 5
	RetryState_RETRY_STATE_INTERNAL_SERVER_ERROR    RetryState = 6
	RetryState_RETRY_STATE_CANCEL_REQUESTED         RetryState = 7
)

// Enum value maps for RetryState.
var (
	RetryState_name = map[int32]string{
		0: "RETRY_STATE_UNSPECIFIED",
		1: "RETRY_STATE_IN_PROGRESS",
		2: "RETRY_STATE_NON_RETRYABLE_FAILURE",
		3: "RETRY_STATE_TIMEOUT",
		4: "RETRY_STATE_MAXIMUM_ATTEMPTS_REACHED",
		5: "RETRY_STATE_RETRY_POLICY_NOT_SET",
		6: "RETRY_STATE_INTERNAL_SERVER_ERROR",
		7: "RETRY_STATE_CANCEL_REQUESTED",
	}
	RetryState_value = map[string]int32{
		"RETRY_STATE_UNSPECIFIED":              0,
		"RETRY_STATE_IN_PROGRESS":              1,
		"RETRY_STATE_NON_RETRYABLE_FAILURE":    2,
		"RETRY_STATE_TIMEOUT":                  3,
		"RETRY_STATE_MAXIMUM_ATTEMPTS_REACHED": 4,
		"RETRY_STATE_RETRY_POLICY_NOT_SET":     5,
		"RETRY_STATE_INTERNAL_SERVER_ERROR":    6,
		"RETRY_STATE_CANCEL_REQUESTED":         7,
	}
)

func (x RetryState) Enum() *RetryState {
	p := new(RetryState)
	*p = x
	return p
}

func (x RetryState) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RetryState) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[7].Descriptor()
}

func (RetryState) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[7]
}

func (x RetryState) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RetryState.Descriptor instead.
func (RetryState) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{7}
}

type TimeoutType int32

const (
	TimeoutType_TIMEOUT_TYPE_UNSPECIFIED       TimeoutType = 0
	TimeoutType_TIMEOUT_TYPE_START_TO_CLOSE    TimeoutType = 1
	TimeoutType_TIMEOUT_TYPE_SCHEDULE_TO_START TimeoutType = 2
	TimeoutType_TIMEOUT_TYPE_SCHEDULE_TO_CLOSE TimeoutType = 3
	TimeoutType_TIMEOUT_TYPE_HEARTBEAT         TimeoutType = 4
)

// Enum value maps for TimeoutType.
var (
	TimeoutType_name = map[int32]string{
		0: "TIMEOUT_TYPE_UNSPECIFIED",
		1: "TIMEOUT_TYPE_START_TO_CLOSE",
		2: "TIMEOUT_TYPE_SCHEDULE_TO_START",
		3: "TIMEOUT_TYPE_SCHEDULE_TO_CLOSE",
		4: "TIMEOUT_TYPE_HEARTBEAT",
	}
	TimeoutType_value = map[string]int32{
		"TIMEOUT_TYPE_UNSPECIFIED":       0,
		"TIMEOUT_TYPE_START_TO_CLOSE":    1,
		"TIMEOUT_TYPE_SCHEDULE_TO_START": 2,
		"TIMEOUT_TYPE_SCHEDULE_TO_CLOSE": 3,
		"TIMEOUT_TYPE_HEARTBEAT":         4,
	}
)

func (x TimeoutType) Enum() *TimeoutType {
	p := new(TimeoutType)
	*p = x
	return p
}

func (x TimeoutType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TimeoutType) Descriptor() protoreflect.EnumDescriptor {
	return file_temporal_api_enums_v1_workflow_proto_enumTypes[8].Descriptor()
}

func (TimeoutType) Type() protoreflect.EnumType {
	return &file_temporal_api_enums_v1_workflow_proto_enumTypes[8]
}

func (x TimeoutType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TimeoutType.Descriptor instead.
func (TimeoutType) EnumDescriptor() ([]byte, []int) {
	return file_temporal_api_enums_v1_workflow_proto_rawDescGZIP(), []int{8}
}

var File_temporal_api_enums_v1_workflow_proto protoreflect.FileDescriptor

var file_temporal_api_enums_v1_workflow_proto_rawDesc = []byte{
	0x0a, 0x24, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x2a, 0x8b, 0x02,
	0x0a, 0x15, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x49, 0x64, 0x52, 0x65, 0x75, 0x73,
	0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x28, 0x0a, 0x24, 0x57, 0x4f, 0x52, 0x4b, 0x46,
	0x4c, 0x4f, 0x57, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x55, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c,
	0x49, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x2c, 0x0a, 0x28, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x49, 0x44,
	0x5f, 0x52, 0x45, 0x55, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x41, 0x4c,
	0x4c, 0x4f, 0x57, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12,
	0x38, 0x0a, 0x34, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x49, 0x44, 0x5f, 0x52,
	0x45, 0x55, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x41, 0x4c, 0x4c, 0x4f,
	0x57, 0x5f, 0x44, 0x55, 0x50, 0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c,
	0x45, 0x44, 0x5f, 0x4f, 0x4e, 0x4c, 0x59, 0x10, 0x02, 0x12, 0x2d, 0x0a, 0x29, 0x57, 0x4f, 0x52,
	0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x55, 0x53, 0x45, 0x5f, 0x50,
	0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x5f, 0x44, 0x55, 0x50,
	0x4c, 0x49, 0x43, 0x41, 0x54, 0x45, 0x10, 0x03, 0x12, 0x31, 0x0a, 0x2d, 0x57, 0x4f, 0x52, 0x4b,
	0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x49, 0x44, 0x5f, 0x52, 0x45, 0x55, 0x53, 0x45, 0x5f, 0x50, 0x4f,
	0x4c, 0x49, 0x43, 0x59, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x5f, 0x49,
	0x46, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x04, 0x2a, 0xa4, 0x01, 0x0a, 0x11,
	0x50, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x43, 0x6c, 0x6f, 0x73, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x12, 0x23, 0x0a, 0x1f, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x21, 0x0a, 0x1d, 0x50, 0x41, 0x52, 0x45, 0x4e, 0x54,
	0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x54, 0x45,
	0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x10, 0x01, 0x12, 0x1f, 0x0a, 0x1b, 0x50, 0x41, 0x52,
	0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59,
	0x5f, 0x41, 0x42, 0x41, 0x4e, 0x44, 0x4f, 0x4e, 0x10, 0x02, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x41,
	0x52, 0x45, 0x4e, 0x54, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43,
	0x59, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x10, 0x03, 0x2a, 0xbd, 0x01, 0x0a, 0x16, 0x43, 0x6f, 0x6e, 0x74, 0x69, 0x6e, 0x75, 0x65, 0x41,
	0x73, 0x4e, 0x65, 0x77, 0x49, 0x6e, 0x69, 0x74, 0x69, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x29, 0x0a,
	0x25, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x5f, 0x41, 0x53, 0x5f, 0x4e, 0x45, 0x57,
	0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45,
	0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x26, 0x0a, 0x22, 0x43, 0x4f, 0x4e, 0x54,
	0x49, 0x4e, 0x55, 0x45, 0x5f, 0x41, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x49, 0x4e, 0x49, 0x54,
	0x49, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x10, 0x01,
	0x12, 0x23, 0x0a, 0x1f, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x5f, 0x41, 0x53, 0x5f,
	0x4e, 0x45, 0x57, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x54, 0x4f, 0x52, 0x5f, 0x52, 0x45,
	0x54, 0x52, 0x59, 0x10, 0x02, 0x12, 0x2b, 0x0a, 0x27, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55,
	0x45, 0x5f, 0x41, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x5f, 0x49, 0x4e, 0x49, 0x54, 0x49, 0x41, 0x54,
	0x4f, 0x52, 0x5f, 0x43, 0x52, 0x4f, 0x4e, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45,
	0x10, 0x03, 0x2a, 0xe5, 0x02, 0x0a, 0x17, 0x57, 0x6f, 0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x45,
	0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x29,
	0x0a, 0x25, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55,
	0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x25, 0x0a, 0x21, 0x57, 0x4f, 0x52,
	0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x52, 0x55, 0x4e, 0x4e, 0x49, 0x4e, 0x47, 0x10, 0x01,
	0x12, 0x27, 0x0a, 0x23, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45,
	0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x24, 0x0a, 0x20, 0x57, 0x4f, 0x52,
	0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12,
	0x26, 0x0a, 0x22, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43,
	0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43, 0x41, 0x4e,
	0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x12, 0x28, 0x0a, 0x24, 0x57, 0x4f, 0x52, 0x4b, 0x46,
	0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58, 0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x54, 0x45, 0x52, 0x4d, 0x49, 0x4e, 0x41, 0x54, 0x45, 0x44, 0x10,
	0x05, 0x12, 0x2e, 0x0a, 0x2a, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58,
	0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x43,
	0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45, 0x44, 0x5f, 0x41, 0x53, 0x5f, 0x4e, 0x45, 0x57, 0x10,
	0x06, 0x12, 0x27, 0x0a, 0x23, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x45, 0x58,
	0x45, 0x43, 0x55, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x54,
	0x49, 0x4d, 0x45, 0x44, 0x5f, 0x4f, 0x55, 0x54, 0x10, 0x07, 0x2a, 0xb5, 0x01, 0x0a, 0x14, 0x50,
	0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x41, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x22, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41,
	0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x24, 0x0a, 0x20, 0x50,
	0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x44, 0x10,
	0x01, 0x12, 0x22, 0x0a, 0x1e, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x41, 0x43, 0x54,
	0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52,
	0x54, 0x45, 0x44, 0x10, 0x02, 0x12, 0x2b, 0x0a, 0x27, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47,
	0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x49, 0x54, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44,
	0x10, 0x03, 0x2a, 0x9b, 0x01, 0x0a, 0x18, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x54, 0x61, 0x73, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x2b, 0x0a, 0x27, 0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46,
	0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55,
	0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x29, 0x0a, 0x25,
	0x50, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57,
	0x5f, 0x54, 0x41, 0x53, 0x4b, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45,
	0x44, 0x55, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x27, 0x0a, 0x23, 0x50, 0x45, 0x4e, 0x44, 0x49,
	0x4e, 0x47, 0x5f, 0x57, 0x4f, 0x52, 0x4b, 0x46, 0x4c, 0x4f, 0x57, 0x5f, 0x54, 0x41, 0x53, 0x4b,
	0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x45, 0x44, 0x10, 0x02,
	0x2a, 0x97, 0x01, 0x0a, 0x16, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x29, 0x0a, 0x25, 0x48,
	0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x4c,
	0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x27, 0x0a, 0x23, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52,
	0x59, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x41, 0x4c, 0x4c, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12,
	0x29, 0x0a, 0x25, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54,
	0x5f, 0x46, 0x49, 0x4c, 0x54, 0x45, 0x52, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x43, 0x4c, 0x4f,
	0x53, 0x45, 0x5f, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x10, 0x02, 0x2a, 0x9f, 0x02, 0x0a, 0x0a, 0x52,
	0x65, 0x74, 0x72, 0x79, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x54,
	0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1b, 0x0a, 0x17, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f,
	0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x49, 0x4e, 0x5f, 0x50, 0x52, 0x4f, 0x47, 0x52, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x4e, 0x4f, 0x4e, 0x5f, 0x52, 0x45, 0x54, 0x52, 0x59, 0x41, 0x42, 0x4c, 0x45,
	0x5f, 0x46, 0x41, 0x49, 0x4c, 0x55, 0x52, 0x45, 0x10, 0x02, 0x12, 0x17, 0x0a, 0x13, 0x52, 0x45,
	0x54, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55,
	0x54, 0x10, 0x03, 0x12, 0x28, 0x0a, 0x24, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x4d, 0x41, 0x58, 0x49, 0x4d, 0x55, 0x4d, 0x5f, 0x41, 0x54, 0x54, 0x45, 0x4d,
	0x50, 0x54, 0x53, 0x5f, 0x52, 0x45, 0x41, 0x43, 0x48, 0x45, 0x44, 0x10, 0x04, 0x12, 0x24, 0x0a,
	0x20, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x52, 0x45, 0x54,
	0x52, 0x59, 0x5f, 0x50, 0x4f, 0x4c, 0x49, 0x43, 0x59, 0x5f, 0x4e, 0x4f, 0x54, 0x5f, 0x53, 0x45,
	0x54, 0x10, 0x05, 0x12, 0x25, 0x0a, 0x21, 0x52, 0x45, 0x54, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x49, 0x4e, 0x54, 0x45, 0x52, 0x4e, 0x41, 0x4c, 0x5f, 0x53, 0x45, 0x52, 0x56,
	0x45, 0x52, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x06, 0x12, 0x20, 0x0a, 0x1c, 0x52, 0x45,
	0x54, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c,
	0x5f, 0x52, 0x45, 0x51, 0x55, 0x45, 0x53, 0x54, 0x45, 0x44, 0x10, 0x07, 0x2a, 0xb0, 0x01, 0x0a,
	0x0b, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x18,
	0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x1f, 0x0a, 0x1b, 0x54, 0x49,
	0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54,
	0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x01, 0x12, 0x22, 0x0a, 0x1e, 0x54,
	0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x53, 0x43, 0x48, 0x45,
	0x44, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x53, 0x54, 0x41, 0x52, 0x54, 0x10, 0x02, 0x12,
	0x22, 0x0a, 0x1e, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x53, 0x43, 0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x5f, 0x54, 0x4f, 0x5f, 0x43, 0x4c, 0x4f, 0x53,
	0x45, 0x10, 0x03, 0x12, 0x1a, 0x0a, 0x16, 0x54, 0x49, 0x4d, 0x45, 0x4f, 0x55, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x48, 0x45, 0x41, 0x52, 0x54, 0x42, 0x45, 0x41, 0x54, 0x10, 0x04, 0x42,
	0x85, 0x01, 0x0a, 0x18, 0x69, 0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x57, 0x6f,
	0x72, 0x6b, 0x66, 0x6c, 0x6f, 0x77, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x21, 0x67,
	0x6f, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x65, 0x6e, 0x75, 0x6d, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x65, 0x6e, 0x75, 0x6d, 0x73,
	0xaa, 0x02, 0x17, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x2e, 0x41, 0x70,
	0x69, 0x2e, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x56, 0x31, 0xea, 0x02, 0x1a, 0x54, 0x65, 0x6d,
	0x70, 0x6f, 0x72, 0x61, 0x6c, 0x69, 0x6f, 0x3a, 0x3a, 0x41, 0x70, 0x69, 0x3a, 0x3a, 0x45, 0x6e,
	0x75, 0x6d, 0x73, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_temporal_api_enums_v1_workflow_proto_rawDescOnce sync.Once
	file_temporal_api_enums_v1_workflow_proto_rawDescData = file_temporal_api_enums_v1_workflow_proto_rawDesc
)

func file_temporal_api_enums_v1_workflow_proto_rawDescGZIP() []byte {
	file_temporal_api_enums_v1_workflow_proto_rawDescOnce.Do(func() {
		file_temporal_api_enums_v1_workflow_proto_rawDescData = protoimpl.X.CompressGZIP(file_temporal_api_enums_v1_workflow_proto_rawDescData)
	})
	return file_temporal_api_enums_v1_workflow_proto_rawDescData
}

var file_temporal_api_enums_v1_workflow_proto_enumTypes = make([]protoimpl.EnumInfo, 9)
var file_temporal_api_enums_v1_workflow_proto_goTypes = []interface{}{
	(WorkflowIdReusePolicy)(0),    // 0: temporal.api.enums.v1.WorkflowIdReusePolicy
	(ParentClosePolicy)(0),        // 1: temporal.api.enums.v1.ParentClosePolicy
	(ContinueAsNewInitiator)(0),   // 2: temporal.api.enums.v1.ContinueAsNewInitiator
	(WorkflowExecutionStatus)(0),  // 3: temporal.api.enums.v1.WorkflowExecutionStatus
	(PendingActivityState)(0),     // 4: temporal.api.enums.v1.PendingActivityState
	(PendingWorkflowTaskState)(0), // 5: temporal.api.enums.v1.PendingWorkflowTaskState
	(HistoryEventFilterType)(0),   // 6: temporal.api.enums.v1.HistoryEventFilterType
	(RetryState)(0),               // 7: temporal.api.enums.v1.RetryState
	(TimeoutType)(0),              // 8: temporal.api.enums.v1.TimeoutType
}
var file_temporal_api_enums_v1_workflow_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_temporal_api_enums_v1_workflow_proto_init() }
func file_temporal_api_enums_v1_workflow_proto_init() {
	if File_temporal_api_enums_v1_workflow_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_temporal_api_enums_v1_workflow_proto_rawDesc,
			NumEnums:      9,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_temporal_api_enums_v1_workflow_proto_goTypes,
		DependencyIndexes: file_temporal_api_enums_v1_workflow_proto_depIdxs,
		EnumInfos:         file_temporal_api_enums_v1_workflow_proto_enumTypes,
	}.Build()
	File_temporal_api_enums_v1_workflow_proto = out.File
	file_temporal_api_enums_v1_workflow_proto_rawDesc = nil
	file_temporal_api_enums_v1_workflow_proto_goTypes = nil
	file_temporal_api_enums_v1_workflow_proto_depIdxs = nil
}
