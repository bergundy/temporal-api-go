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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// plugins:
// - protoc-gen-go-grpc
// - protoc
// source: temporal/api/chasm/activityservice/v1/service.proto

package activityservice

import (
	context "context"

	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ActivityService_ExecuteActivity_FullMethodName  = "/temporal.api.chasm.activityservice.v1.ActivityService/ExecuteActivity"
	ActivityService_DescribeActivity_FullMethodName = "/temporal.api.chasm.activityservice.v1.ActivityService/DescribeActivity"
	ActivityService_ListActivities_FullMethodName   = "/temporal.api.chasm.activityservice.v1.ActivityService/ListActivities"
)

// ActivityServiceClient is the client API for ActivityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ActivityServiceClient interface {
	ExecuteActivity(ctx context.Context, in *ExecuteActivityRequest, opts ...grpc.CallOption) (*ExecuteActivityResponse, error)
	DescribeActivity(ctx context.Context, in *DescribeActivityRequest, opts ...grpc.CallOption) (*DescribeActivityResponse, error)
	ListActivities(ctx context.Context, in *ListActivitiesRequest, opts ...grpc.CallOption) (*ListActivitiesResponse, error)
}

type activityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewActivityServiceClient(cc grpc.ClientConnInterface) ActivityServiceClient {
	return &activityServiceClient{cc}
}

func (c *activityServiceClient) ExecuteActivity(ctx context.Context, in *ExecuteActivityRequest, opts ...grpc.CallOption) (*ExecuteActivityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecuteActivityResponse)
	err := c.cc.Invoke(ctx, ActivityService_ExecuteActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) DescribeActivity(ctx context.Context, in *DescribeActivityRequest, opts ...grpc.CallOption) (*DescribeActivityResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DescribeActivityResponse)
	err := c.cc.Invoke(ctx, ActivityService_DescribeActivity_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *activityServiceClient) ListActivities(ctx context.Context, in *ListActivitiesRequest, opts ...grpc.CallOption) (*ListActivitiesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListActivitiesResponse)
	err := c.cc.Invoke(ctx, ActivityService_ListActivities_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ActivityServiceServer is the server API for ActivityService service.
// All implementations must embed UnimplementedActivityServiceServer
// for forward compatibility.
type ActivityServiceServer interface {
	ExecuteActivity(context.Context, *ExecuteActivityRequest) (*ExecuteActivityResponse, error)
	DescribeActivity(context.Context, *DescribeActivityRequest) (*DescribeActivityResponse, error)
	ListActivities(context.Context, *ListActivitiesRequest) (*ListActivitiesResponse, error)
	mustEmbedUnimplementedActivityServiceServer()
}

// UnimplementedActivityServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedActivityServiceServer struct{}

func (UnimplementedActivityServiceServer) ExecuteActivity(context.Context, *ExecuteActivityRequest) (*ExecuteActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteActivity not implemented")
}
func (UnimplementedActivityServiceServer) DescribeActivity(context.Context, *DescribeActivityRequest) (*DescribeActivityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DescribeActivity not implemented")
}
func (UnimplementedActivityServiceServer) ListActivities(context.Context, *ListActivitiesRequest) (*ListActivitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListActivities not implemented")
}
func (UnimplementedActivityServiceServer) mustEmbedUnimplementedActivityServiceServer() {}
func (UnimplementedActivityServiceServer) testEmbeddedByValue()                         {}

// UnsafeActivityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ActivityServiceServer will
// result in compilation errors.
type UnsafeActivityServiceServer interface {
	mustEmbedUnimplementedActivityServiceServer()
}

func RegisterActivityServiceServer(s grpc.ServiceRegistrar, srv ActivityServiceServer) {
	// If the following call pancis, it indicates UnimplementedActivityServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ActivityService_ServiceDesc, srv)
}

func _ActivityService_ExecuteActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).ExecuteActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_ExecuteActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).ExecuteActivity(ctx, req.(*ExecuteActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_DescribeActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DescribeActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).DescribeActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_DescribeActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).DescribeActivity(ctx, req.(*DescribeActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ActivityService_ListActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ActivityServiceServer).ListActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ActivityService_ListActivities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ActivityServiceServer).ListActivities(ctx, req.(*ListActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ActivityService_ServiceDesc is the grpc.ServiceDesc for ActivityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ActivityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "temporal.api.chasm.activityservice.v1.ActivityService",
	HandlerType: (*ActivityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ExecuteActivity",
			Handler:    _ActivityService_ExecuteActivity_Handler,
		},
		{
			MethodName: "DescribeActivity",
			Handler:    _ActivityService_DescribeActivity_Handler,
		},
		{
			MethodName: "ListActivities",
			Handler:    _ActivityService_ListActivities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "temporal/api/chasm/activityservice/v1/service.proto",
}
