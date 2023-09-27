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

// Code generated by proxygenerator; DO NOT EDIT.

package proxy

import (
	"context"

	"go.temporal.io/api/workflowservice/v1"
)

// WorkflowServiceProxyOptions provides options for configuring a WorkflowServiceProxyServer.
// Client is a WorkflowServiceClient used to forward requests received by the server to the
// Temporal Frontend.
type WorkflowServiceProxyOptions struct {
	Client workflowservice.WorkflowServiceClient
}

type workflowServiceProxyServer struct {
	workflowservice.UnimplementedWorkflowServiceServer
	client workflowservice.WorkflowServiceClient
}

// NewWorkflowServiceProxyServer creates a WorkflowServiceServer suitable for registering with a gRPC Server. Requests will
// be forwarded to the passed in WorkflowService Client. gRPC interceptors can be added on the Server or Client to adjust
// requests and responses.
func NewWorkflowServiceProxyServer(options WorkflowServiceProxyOptions) (workflowservice.WorkflowServiceServer, error) {
	return &workflowServiceProxyServer{
		client: options.Client,
	}, nil
}

func (s *workflowServiceProxyServer) CountWorkflowExecutions(ctx context.Context, req *workflowservice.CountWorkflowExecutionsRequest) (*workflowservice.CountWorkflowExecutionsResponse, error) {
	return s.client.CountWorkflowExecutions(ctx, req)
}

func (s *workflowServiceProxyServer) CreateSchedule(ctx context.Context, req *workflowservice.CreateScheduleRequest) (*workflowservice.CreateScheduleResponse, error) {
	return s.client.CreateSchedule(ctx, req)
}

func (s *workflowServiceProxyServer) DeleteSchedule(ctx context.Context, req *workflowservice.DeleteScheduleRequest) (*workflowservice.DeleteScheduleResponse, error) {
	return s.client.DeleteSchedule(ctx, req)
}

func (s *workflowServiceProxyServer) DeleteWorkflowExecution(ctx context.Context, req *workflowservice.DeleteWorkflowExecutionRequest) (*workflowservice.DeleteWorkflowExecutionResponse, error) {
	return s.client.DeleteWorkflowExecution(ctx, req)
}

func (s *workflowServiceProxyServer) DeprecateNamespace(ctx context.Context, req *workflowservice.DeprecateNamespaceRequest) (*workflowservice.DeprecateNamespaceResponse, error) {
	return s.client.DeprecateNamespace(ctx, req)
}

func (s *workflowServiceProxyServer) DescribeBatchOperation(ctx context.Context, req *workflowservice.DescribeBatchOperationRequest) (*workflowservice.DescribeBatchOperationResponse, error) {
	return s.client.DescribeBatchOperation(ctx, req)
}

func (s *workflowServiceProxyServer) DescribeNamespace(ctx context.Context, req *workflowservice.DescribeNamespaceRequest) (*workflowservice.DescribeNamespaceResponse, error) {
	return s.client.DescribeNamespace(ctx, req)
}

func (s *workflowServiceProxyServer) DescribeSchedule(ctx context.Context, req *workflowservice.DescribeScheduleRequest) (*workflowservice.DescribeScheduleResponse, error) {
	return s.client.DescribeSchedule(ctx, req)
}

func (s *workflowServiceProxyServer) DescribeTaskQueue(ctx context.Context, req *workflowservice.DescribeTaskQueueRequest) (*workflowservice.DescribeTaskQueueResponse, error) {
	return s.client.DescribeTaskQueue(ctx, req)
}

func (s *workflowServiceProxyServer) DescribeWorkflowExecution(ctx context.Context, req *workflowservice.DescribeWorkflowExecutionRequest) (*workflowservice.DescribeWorkflowExecutionResponse, error) {
	return s.client.DescribeWorkflowExecution(ctx, req)
}

func (s *workflowServiceProxyServer) GetClusterInfo(ctx context.Context, req *workflowservice.GetClusterInfoRequest) (*workflowservice.GetClusterInfoResponse, error) {
	return s.client.GetClusterInfo(ctx, req)
}

func (s *workflowServiceProxyServer) GetSearchAttributes(ctx context.Context, req *workflowservice.GetSearchAttributesRequest) (*workflowservice.GetSearchAttributesResponse, error) {
	return s.client.GetSearchAttributes(ctx, req)
}

func (s *workflowServiceProxyServer) GetSystemInfo(ctx context.Context, req *workflowservice.GetSystemInfoRequest) (*workflowservice.GetSystemInfoResponse, error) {
	return s.client.GetSystemInfo(ctx, req)
}

func (s *workflowServiceProxyServer) GetWorkerBuildIdCompatibility(ctx context.Context, req *workflowservice.GetWorkerBuildIdCompatibilityRequest) (*workflowservice.GetWorkerBuildIdCompatibilityResponse, error) {
	return s.client.GetWorkerBuildIdCompatibility(ctx, req)
}

func (s *workflowServiceProxyServer) GetWorkerTaskReachability(ctx context.Context, req *workflowservice.GetWorkerTaskReachabilityRequest) (*workflowservice.GetWorkerTaskReachabilityResponse, error) {
	return s.client.GetWorkerTaskReachability(ctx, req)
}

func (s *workflowServiceProxyServer) GetWorkflowExecutionHistory(ctx context.Context, req *workflowservice.GetWorkflowExecutionHistoryRequest) (*workflowservice.GetWorkflowExecutionHistoryResponse, error) {
	return s.client.GetWorkflowExecutionHistory(ctx, req)
}

func (s *workflowServiceProxyServer) GetWorkflowExecutionHistoryReverse(ctx context.Context, req *workflowservice.GetWorkflowExecutionHistoryReverseRequest) (*workflowservice.GetWorkflowExecutionHistoryReverseResponse, error) {
	return s.client.GetWorkflowExecutionHistoryReverse(ctx, req)
}

func (s *workflowServiceProxyServer) ListArchivedWorkflowExecutions(ctx context.Context, req *workflowservice.ListArchivedWorkflowExecutionsRequest) (*workflowservice.ListArchivedWorkflowExecutionsResponse, error) {
	return s.client.ListArchivedWorkflowExecutions(ctx, req)
}

func (s *workflowServiceProxyServer) ListBatchOperations(ctx context.Context, req *workflowservice.ListBatchOperationsRequest) (*workflowservice.ListBatchOperationsResponse, error) {
	return s.client.ListBatchOperations(ctx, req)
}

func (s *workflowServiceProxyServer) ListClosedWorkflowExecutions(ctx context.Context, req *workflowservice.ListClosedWorkflowExecutionsRequest) (*workflowservice.ListClosedWorkflowExecutionsResponse, error) {
	return s.client.ListClosedWorkflowExecutions(ctx, req)
}

func (s *workflowServiceProxyServer) ListNamespaces(ctx context.Context, req *workflowservice.ListNamespacesRequest) (*workflowservice.ListNamespacesResponse, error) {
	return s.client.ListNamespaces(ctx, req)
}

func (s *workflowServiceProxyServer) ListOpenWorkflowExecutions(ctx context.Context, req *workflowservice.ListOpenWorkflowExecutionsRequest) (*workflowservice.ListOpenWorkflowExecutionsResponse, error) {
	return s.client.ListOpenWorkflowExecutions(ctx, req)
}

func (s *workflowServiceProxyServer) ListScheduleMatchingTimes(ctx context.Context, req *workflowservice.ListScheduleMatchingTimesRequest) (*workflowservice.ListScheduleMatchingTimesResponse, error) {
	return s.client.ListScheduleMatchingTimes(ctx, req)
}

func (s *workflowServiceProxyServer) ListSchedules(ctx context.Context, req *workflowservice.ListSchedulesRequest) (*workflowservice.ListSchedulesResponse, error) {
	return s.client.ListSchedules(ctx, req)
}

func (s *workflowServiceProxyServer) ListTaskQueuePartitions(ctx context.Context, req *workflowservice.ListTaskQueuePartitionsRequest) (*workflowservice.ListTaskQueuePartitionsResponse, error) {
	return s.client.ListTaskQueuePartitions(ctx, req)
}

func (s *workflowServiceProxyServer) ListWorkflowExecutions(ctx context.Context, req *workflowservice.ListWorkflowExecutionsRequest) (*workflowservice.ListWorkflowExecutionsResponse, error) {
	return s.client.ListWorkflowExecutions(ctx, req)
}

func (s *workflowServiceProxyServer) PatchSchedule(ctx context.Context, req *workflowservice.PatchScheduleRequest) (*workflowservice.PatchScheduleResponse, error) {
	return s.client.PatchSchedule(ctx, req)
}

func (s *workflowServiceProxyServer) PollActivityTaskQueue(ctx context.Context, req *workflowservice.PollActivityTaskQueueRequest) (*workflowservice.PollActivityTaskQueueResponse, error) {
	return s.client.PollActivityTaskQueue(ctx, req)
}

func (s *workflowServiceProxyServer) PollNexusTaskQueue(ctx context.Context, req *workflowservice.PollNexusTaskQueueRequest) (*workflowservice.PollNexusTaskQueueResponse, error) {
	return s.client.PollNexusTaskQueue(ctx, req)
}

func (s *workflowServiceProxyServer) PollWorkflowExecutionUpdate(ctx context.Context, req *workflowservice.PollWorkflowExecutionUpdateRequest) (*workflowservice.PollWorkflowExecutionUpdateResponse, error) {
	return s.client.PollWorkflowExecutionUpdate(ctx, req)
}

func (s *workflowServiceProxyServer) PollWorkflowTaskQueue(ctx context.Context, req *workflowservice.PollWorkflowTaskQueueRequest) (*workflowservice.PollWorkflowTaskQueueResponse, error) {
	return s.client.PollWorkflowTaskQueue(ctx, req)
}

func (s *workflowServiceProxyServer) QueryWorkflow(ctx context.Context, req *workflowservice.QueryWorkflowRequest) (*workflowservice.QueryWorkflowResponse, error) {
	return s.client.QueryWorkflow(ctx, req)
}

func (s *workflowServiceProxyServer) RecordActivityTaskHeartbeat(ctx context.Context, req *workflowservice.RecordActivityTaskHeartbeatRequest) (*workflowservice.RecordActivityTaskHeartbeatResponse, error) {
	return s.client.RecordActivityTaskHeartbeat(ctx, req)
}

func (s *workflowServiceProxyServer) RecordActivityTaskHeartbeatById(ctx context.Context, req *workflowservice.RecordActivityTaskHeartbeatByIdRequest) (*workflowservice.RecordActivityTaskHeartbeatByIdResponse, error) {
	return s.client.RecordActivityTaskHeartbeatById(ctx, req)
}

func (s *workflowServiceProxyServer) RegisterNamespace(ctx context.Context, req *workflowservice.RegisterNamespaceRequest) (*workflowservice.RegisterNamespaceResponse, error) {
	return s.client.RegisterNamespace(ctx, req)
}

func (s *workflowServiceProxyServer) RequestCancelWorkflowExecution(ctx context.Context, req *workflowservice.RequestCancelWorkflowExecutionRequest) (*workflowservice.RequestCancelWorkflowExecutionResponse, error) {
	return s.client.RequestCancelWorkflowExecution(ctx, req)
}

func (s *workflowServiceProxyServer) ResetStickyTaskQueue(ctx context.Context, req *workflowservice.ResetStickyTaskQueueRequest) (*workflowservice.ResetStickyTaskQueueResponse, error) {
	return s.client.ResetStickyTaskQueue(ctx, req)
}

func (s *workflowServiceProxyServer) ResetWorkflowExecution(ctx context.Context, req *workflowservice.ResetWorkflowExecutionRequest) (*workflowservice.ResetWorkflowExecutionResponse, error) {
	return s.client.ResetWorkflowExecution(ctx, req)
}

func (s *workflowServiceProxyServer) RespondActivityTaskCanceled(ctx context.Context, req *workflowservice.RespondActivityTaskCanceledRequest) (*workflowservice.RespondActivityTaskCanceledResponse, error) {
	return s.client.RespondActivityTaskCanceled(ctx, req)
}

func (s *workflowServiceProxyServer) RespondActivityTaskCanceledById(ctx context.Context, req *workflowservice.RespondActivityTaskCanceledByIdRequest) (*workflowservice.RespondActivityTaskCanceledByIdResponse, error) {
	return s.client.RespondActivityTaskCanceledById(ctx, req)
}

func (s *workflowServiceProxyServer) RespondActivityTaskCompleted(ctx context.Context, req *workflowservice.RespondActivityTaskCompletedRequest) (*workflowservice.RespondActivityTaskCompletedResponse, error) {
	return s.client.RespondActivityTaskCompleted(ctx, req)
}

func (s *workflowServiceProxyServer) RespondActivityTaskCompletedById(ctx context.Context, req *workflowservice.RespondActivityTaskCompletedByIdRequest) (*workflowservice.RespondActivityTaskCompletedByIdResponse, error) {
	return s.client.RespondActivityTaskCompletedById(ctx, req)
}

func (s *workflowServiceProxyServer) RespondActivityTaskFailed(ctx context.Context, req *workflowservice.RespondActivityTaskFailedRequest) (*workflowservice.RespondActivityTaskFailedResponse, error) {
	return s.client.RespondActivityTaskFailed(ctx, req)
}

func (s *workflowServiceProxyServer) RespondActivityTaskFailedById(ctx context.Context, req *workflowservice.RespondActivityTaskFailedByIdRequest) (*workflowservice.RespondActivityTaskFailedByIdResponse, error) {
	return s.client.RespondActivityTaskFailedById(ctx, req)
}

func (s *workflowServiceProxyServer) RespondNexusTaskCompleted(ctx context.Context, req *workflowservice.RespondNexusTaskCompletedRequest) (*workflowservice.RespondNexusTaskCompletedResponse, error) {
	return s.client.RespondNexusTaskCompleted(ctx, req)
}

func (s *workflowServiceProxyServer) RespondQueryTaskCompleted(ctx context.Context, req *workflowservice.RespondQueryTaskCompletedRequest) (*workflowservice.RespondQueryTaskCompletedResponse, error) {
	return s.client.RespondQueryTaskCompleted(ctx, req)
}

func (s *workflowServiceProxyServer) RespondWorkflowTaskCompleted(ctx context.Context, req *workflowservice.RespondWorkflowTaskCompletedRequest) (*workflowservice.RespondWorkflowTaskCompletedResponse, error) {
	return s.client.RespondWorkflowTaskCompleted(ctx, req)
}

func (s *workflowServiceProxyServer) RespondWorkflowTaskFailed(ctx context.Context, req *workflowservice.RespondWorkflowTaskFailedRequest) (*workflowservice.RespondWorkflowTaskFailedResponse, error) {
	return s.client.RespondWorkflowTaskFailed(ctx, req)
}

func (s *workflowServiceProxyServer) ScanWorkflowExecutions(ctx context.Context, req *workflowservice.ScanWorkflowExecutionsRequest) (*workflowservice.ScanWorkflowExecutionsResponse, error) {
	return s.client.ScanWorkflowExecutions(ctx, req)
}

func (s *workflowServiceProxyServer) SignalWithStartWorkflowExecution(ctx context.Context, req *workflowservice.SignalWithStartWorkflowExecutionRequest) (*workflowservice.SignalWithStartWorkflowExecutionResponse, error) {
	return s.client.SignalWithStartWorkflowExecution(ctx, req)
}

func (s *workflowServiceProxyServer) SignalWorkflowExecution(ctx context.Context, req *workflowservice.SignalWorkflowExecutionRequest) (*workflowservice.SignalWorkflowExecutionResponse, error) {
	return s.client.SignalWorkflowExecution(ctx, req)
}

func (s *workflowServiceProxyServer) StartBatchOperation(ctx context.Context, req *workflowservice.StartBatchOperationRequest) (*workflowservice.StartBatchOperationResponse, error) {
	return s.client.StartBatchOperation(ctx, req)
}

func (s *workflowServiceProxyServer) StartWorkflowExecution(ctx context.Context, req *workflowservice.StartWorkflowExecutionRequest) (*workflowservice.StartWorkflowExecutionResponse, error) {
	return s.client.StartWorkflowExecution(ctx, req)
}

func (s *workflowServiceProxyServer) StopBatchOperation(ctx context.Context, req *workflowservice.StopBatchOperationRequest) (*workflowservice.StopBatchOperationResponse, error) {
	return s.client.StopBatchOperation(ctx, req)
}

func (s *workflowServiceProxyServer) TerminateWorkflowExecution(ctx context.Context, req *workflowservice.TerminateWorkflowExecutionRequest) (*workflowservice.TerminateWorkflowExecutionResponse, error) {
	return s.client.TerminateWorkflowExecution(ctx, req)
}

func (s *workflowServiceProxyServer) UpdateNamespace(ctx context.Context, req *workflowservice.UpdateNamespaceRequest) (*workflowservice.UpdateNamespaceResponse, error) {
	return s.client.UpdateNamespace(ctx, req)
}

func (s *workflowServiceProxyServer) UpdateSchedule(ctx context.Context, req *workflowservice.UpdateScheduleRequest) (*workflowservice.UpdateScheduleResponse, error) {
	return s.client.UpdateSchedule(ctx, req)
}

func (s *workflowServiceProxyServer) UpdateWorkerBuildIdCompatibility(ctx context.Context, req *workflowservice.UpdateWorkerBuildIdCompatibilityRequest) (*workflowservice.UpdateWorkerBuildIdCompatibilityResponse, error) {
	return s.client.UpdateWorkerBuildIdCompatibility(ctx, req)
}

func (s *workflowServiceProxyServer) UpdateWorkflowExecution(ctx context.Context, req *workflowservice.UpdateWorkflowExecutionRequest) (*workflowservice.UpdateWorkflowExecutionResponse, error) {
	return s.client.UpdateWorkflowExecution(ctx, req)
}
