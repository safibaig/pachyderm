// Code generated by protoc-gen-go.
// source: server/pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	server/pps/pps.proto

It has these top-level messages:
	StartJobRequest
	StartJobResponse
	FinishJobRequest
	FinishJobResponse
	ContinueJobRequest
	ContinueJobResponse
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/go/google/protobuf"
import fuse "github.com/pachyderm/pachyderm/src/server/pfs/fuse"
import _ "github.com/pachyderm/pachyderm/src/client/pfs"
import pachyderm_pps "github.com/pachyderm/pachyderm/src/client/pps"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type StartJobRequest struct {
	Job     *pachyderm_pps.Job `protobuf:"bytes,1,opt,name=job" json:"job,omitempty"`
	PodName string             `protobuf:"bytes,2,opt,name=pod_name,json=podName" json:"pod_name,omitempty"`
}

func (m *StartJobRequest) Reset()                    { *m = StartJobRequest{} }
func (m *StartJobRequest) String() string            { return proto.CompactTextString(m) }
func (*StartJobRequest) ProtoMessage()               {}
func (*StartJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *StartJobRequest) GetJob() *pachyderm_pps.Job {
	if m != nil {
		return m.Job
	}
	return nil
}

type StartJobResponse struct {
	ChunkID      string                   `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId" json:"chunk_id,omitempty"`
	Transform    *pachyderm_pps.Transform `protobuf:"bytes,2,opt,name=transform" json:"transform,omitempty"`
	CommitMounts []*fuse.CommitMount      `protobuf:"bytes,3,rep,name=commit_mounts,json=commitMounts" json:"commit_mounts,omitempty"`
}

func (m *StartJobResponse) Reset()                    { *m = StartJobResponse{} }
func (m *StartJobResponse) String() string            { return proto.CompactTextString(m) }
func (*StartJobResponse) ProtoMessage()               {}
func (*StartJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *StartJobResponse) GetTransform() *pachyderm_pps.Transform {
	if m != nil {
		return m.Transform
	}
	return nil
}

func (m *StartJobResponse) GetCommitMounts() []*fuse.CommitMount {
	if m != nil {
		return m.CommitMounts
	}
	return nil
}

type FinishJobRequest struct {
	ChunkID string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId" json:"chunk_id,omitempty"`
	PodName string `protobuf:"bytes,2,opt,name=pod_name,json=podName" json:"pod_name,omitempty"`
	Success bool   `protobuf:"varint,3,opt,name=success" json:"success,omitempty"`
}

func (m *FinishJobRequest) Reset()                    { *m = FinishJobRequest{} }
func (m *FinishJobRequest) String() string            { return proto.CompactTextString(m) }
func (*FinishJobRequest) ProtoMessage()               {}
func (*FinishJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type FinishJobResponse struct {
	// If fail is true, the pod is expected to exit with a non-zero code
	// so that k8s knows to reschedule the pod.
	Fail bool `protobuf:"varint,1,opt,name=fail" json:"fail,omitempty"`
}

func (m *FinishJobResponse) Reset()                    { *m = FinishJobResponse{} }
func (m *FinishJobResponse) String() string            { return proto.CompactTextString(m) }
func (*FinishJobResponse) ProtoMessage()               {}
func (*FinishJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type ContinueJobRequest struct {
	ChunkID string `protobuf:"bytes,1,opt,name=chunk_id,json=chunkId" json:"chunk_id,omitempty"`
	PodName string `protobuf:"bytes,2,opt,name=pod_name,json=podName" json:"pod_name,omitempty"`
}

func (m *ContinueJobRequest) Reset()                    { *m = ContinueJobRequest{} }
func (m *ContinueJobRequest) String() string            { return proto.CompactTextString(m) }
func (*ContinueJobRequest) ProtoMessage()               {}
func (*ContinueJobRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ContinueJobResponse struct {
	// if exit is true, the pod is expected to abandon its work and exit.
	// this happens if the chunk that the pod is working on has been assigned
	// to another pod.
	Exit bool `protobuf:"varint,1,opt,name=exit" json:"exit,omitempty"`
}

func (m *ContinueJobResponse) Reset()                    { *m = ContinueJobResponse{} }
func (m *ContinueJobResponse) String() string            { return proto.CompactTextString(m) }
func (*ContinueJobResponse) ProtoMessage()               {}
func (*ContinueJobResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*StartJobRequest)(nil), "pachyderm.pps.StartJobRequest")
	proto.RegisterType((*StartJobResponse)(nil), "pachyderm.pps.StartJobResponse")
	proto.RegisterType((*FinishJobRequest)(nil), "pachyderm.pps.FinishJobRequest")
	proto.RegisterType((*FinishJobResponse)(nil), "pachyderm.pps.FinishJobResponse")
	proto.RegisterType((*ContinueJobRequest)(nil), "pachyderm.pps.ContinueJobRequest")
	proto.RegisterType((*ContinueJobResponse)(nil), "pachyderm.pps.ContinueJobResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for InternalJobAPI service

type InternalJobAPIClient interface {
	StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error)
	ContinueJob(ctx context.Context, in *ContinueJobRequest, opts ...grpc.CallOption) (*ContinueJobResponse, error)
	FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*FinishJobResponse, error)
}

type internalJobAPIClient struct {
	cc *grpc.ClientConn
}

func NewInternalJobAPIClient(cc *grpc.ClientConn) InternalJobAPIClient {
	return &internalJobAPIClient{cc}
}

func (c *internalJobAPIClient) StartJob(ctx context.Context, in *StartJobRequest, opts ...grpc.CallOption) (*StartJobResponse, error) {
	out := new(StartJobResponse)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/StartJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalJobAPIClient) ContinueJob(ctx context.Context, in *ContinueJobRequest, opts ...grpc.CallOption) (*ContinueJobResponse, error) {
	out := new(ContinueJobResponse)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/ContinueJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *internalJobAPIClient) FinishJob(ctx context.Context, in *FinishJobRequest, opts ...grpc.CallOption) (*FinishJobResponse, error) {
	out := new(FinishJobResponse)
	err := grpc.Invoke(ctx, "/pachyderm.pps.InternalJobAPI/FinishJob", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for InternalJobAPI service

type InternalJobAPIServer interface {
	StartJob(context.Context, *StartJobRequest) (*StartJobResponse, error)
	ContinueJob(context.Context, *ContinueJobRequest) (*ContinueJobResponse, error)
	FinishJob(context.Context, *FinishJobRequest) (*FinishJobResponse, error)
}

func RegisterInternalJobAPIServer(s *grpc.Server, srv InternalJobAPIServer) {
	s.RegisterService(&_InternalJobAPI_serviceDesc, srv)
}

func _InternalJobAPI_StartJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalJobAPIServer).StartJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.InternalJobAPI/StartJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalJobAPIServer).StartJob(ctx, req.(*StartJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalJobAPI_ContinueJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ContinueJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalJobAPIServer).ContinueJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.InternalJobAPI/ContinueJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalJobAPIServer).ContinueJob(ctx, req.(*ContinueJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _InternalJobAPI_FinishJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FinishJobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(InternalJobAPIServer).FinishJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pachyderm.pps.InternalJobAPI/FinishJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(InternalJobAPIServer).FinishJob(ctx, req.(*FinishJobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _InternalJobAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pachyderm.pps.InternalJobAPI",
	HandlerType: (*InternalJobAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartJob",
			Handler:    _InternalJobAPI_StartJob_Handler,
		},
		{
			MethodName: "ContinueJob",
			Handler:    _InternalJobAPI_ContinueJob_Handler,
		},
		{
			MethodName: "FinishJob",
			Handler:    _InternalJobAPI_FinishJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor0,
}

func init() { proto.RegisterFile("server/pps/pps.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xa4, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0xc5, 0x0d, 0x90, 0x64, 0x42, 0xa1, 0x5d, 0x7a, 0x30, 0x46, 0x82, 0xb0, 0x42, 0x22, 0x5c,
	0x1c, 0x29, 0x48, 0xbd, 0x43, 0x25, 0xa4, 0x44, 0x2a, 0xaa, 0x16, 0xc4, 0x81, 0x4b, 0xe4, 0x8f,
	0x71, 0xb3, 0x90, 0xfd, 0x60, 0x67, 0x8d, 0xe8, 0x4f, 0xe0, 0x47, 0xf0, 0x5f, 0x91, 0xb7, 0x75,
	0x93, 0x1a, 0xa5, 0x97, 0x1e, 0x6c, 0xed, 0xf8, 0x3d, 0xbf, 0xf7, 0x66, 0xc6, 0x86, 0x23, 0x42,
	0xf7, 0x0b, 0xdd, 0xd4, 0x5a, 0x6a, 0xae, 0xd4, 0x3a, 0xe3, 0x0d, 0xdb, 0xb7, 0x59, 0xb1, 0xba,
	0x28, 0xd1, 0xa9, 0xd4, 0x5a, 0x4a, 0x9e, 0x9f, 0x1b, 0x73, 0xbe, 0xc6, 0x69, 0x00, 0xf3, 0xba,
	0x9a, 0xa2, 0xb2, 0xfe, 0xe2, 0x92, 0x9b, 0x24, 0xad, 0x42, 0x45, 0xd3, 0xaa, 0x26, 0x0c, 0xb7,
	0x2b, 0xec, 0xa8, 0x58, 0x4b, 0xd4, 0x3e, 0x60, 0xb6, 0xa2, 0xee, 0xd3, 0x6d, 0x4f, 0x2e, 0xe0,
	0xc9, 0x67, 0x9f, 0x39, 0xbf, 0x30, 0xb9, 0xc0, 0x9f, 0x35, 0x92, 0x67, 0xaf, 0xa1, 0xf7, 0xdd,
	0xe4, 0x71, 0x34, 0x8e, 0x26, 0xa3, 0x19, 0x4b, 0x6f, 0x84, 0x4a, 0x1b, 0x5e, 0x03, 0xb3, 0x67,
	0x30, 0xb0, 0xa6, 0x5c, 0xea, 0x4c, 0x61, 0xbc, 0x37, 0x8e, 0x26, 0x43, 0xd1, 0xb7, 0xa6, 0xfc,
	0x94, 0x29, 0xe4, 0x7f, 0x23, 0x38, 0xd8, 0x88, 0x92, 0x35, 0x9a, 0xb0, 0xe1, 0x17, 0xab, 0x5a,
	0xff, 0x58, 0xca, 0x32, 0x48, 0x0f, 0x45, 0x3f, 0xd4, 0xf3, 0x92, 0x1d, 0xc3, 0xd0, 0xbb, 0x4c,
	0x53, 0x65, 0x9c, 0x0a, 0x5a, 0xa3, 0x59, 0xdc, 0xb1, 0xfd, 0xd2, 0xe2, 0x62, 0x43, 0x65, 0xc7,
	0xb0, 0x5f, 0x18, 0xa5, 0xa4, 0x5f, 0x2a, 0x53, 0x6b, 0x4f, 0x71, 0x6f, 0xdc, 0x9b, 0x8c, 0x66,
	0x87, 0x69, 0x98, 0xc5, 0x49, 0x80, 0x4e, 0x1b, 0x44, 0x3c, 0x2a, 0x36, 0x05, 0xf1, 0x1c, 0x0e,
	0x3e, 0x4a, 0x2d, 0x69, 0xb5, 0xd5, 0xf4, 0x2d, 0xf1, 0x76, 0x77, 0xca, 0x62, 0xe8, 0x53, 0x5d,
	0x14, 0x48, 0x8d, 0x77, 0x34, 0x19, 0x88, 0xb6, 0xe4, 0x6f, 0xe0, 0x70, 0xcb, 0xe3, 0x6a, 0x06,
	0x0c, 0xee, 0x57, 0x99, 0x5c, 0x07, 0x83, 0x81, 0x08, 0x67, 0xbe, 0x00, 0x76, 0x62, 0xb4, 0x97,
	0xba, 0xc6, 0xbb, 0xc6, 0xe1, 0x6f, 0xe1, 0xe9, 0x0d, 0xad, 0x8d, 0x2d, 0xfe, 0x96, 0xbe, 0xb5,
	0x6d, 0xce, 0xb3, 0x3f, 0x7b, 0xf0, 0x78, 0xae, 0x3d, 0x3a, 0x9d, 0xad, 0x17, 0x26, 0x7f, 0x7f,
	0x36, 0x67, 0xa7, 0x30, 0x68, 0xb7, 0xc6, 0x5e, 0x74, 0xe6, 0xdf, 0xf9, 0x46, 0x92, 0x97, 0x3b,
	0xf1, 0x4b, 0x4f, 0x7e, 0x8f, 0x7d, 0x85, 0xd1, 0x56, 0x18, 0xf6, 0xaa, 0xf3, 0xc6, 0xff, 0x4d,
	0x27, 0xfc, 0x36, 0xca, 0xb5, 0xee, 0x19, 0x0c, 0xaf, 0x27, 0xcb, 0xba, 0x39, 0xba, 0x7b, 0x4d,
	0xc6, 0xbb, 0x09, 0xad, 0xe2, 0x87, 0x07, 0xdf, 0x7a, 0xd6, 0x52, 0xfe, 0x30, 0xfc, 0x11, 0xef,
	0xfe, 0x05, 0x00, 0x00, 0xff, 0xff, 0x79, 0xad, 0xfe, 0x92, 0x9d, 0x03, 0x00, 0x00,
}
