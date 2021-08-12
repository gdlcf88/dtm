// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package examples

import (
	context "context"
	dtmgrpc "github.com/yedf/dtm/dtmgrpc"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BusiClient is the client API for Busi service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusiClient interface {
	CanSubmit(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TransIn(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TransOut(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TransInRevert(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TransOutRevert(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TransInConfirm(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TransOutConfirm(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	XaNotify(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	TransInXa(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error)
	TransOutXa(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error)
	TransInTcc(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error)
	TransOutTcc(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error)
	TransInTccNested(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error)
}

type busiClient struct {
	cc grpc.ClientConnInterface
}

func NewBusiClient(cc grpc.ClientConnInterface) BusiClient {
	return &busiClient{cc}
}

func (c *busiClient) CanSubmit(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.Busi/CanSubmit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransIn(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransIn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransOut(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransOut", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransInRevert(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransInRevert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransOutRevert(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransOutRevert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransInConfirm(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransInConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransOutConfirm(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransOutConfirm", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) XaNotify(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/examples.Busi/XaNotify", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransInXa(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error) {
	out := new(dtmgrpc.BusiReply)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransInXa", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransOutXa(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error) {
	out := new(dtmgrpc.BusiReply)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransOutXa", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransInTcc(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error) {
	out := new(dtmgrpc.BusiReply)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransInTcc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransOutTcc(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error) {
	out := new(dtmgrpc.BusiReply)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransOutTcc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *busiClient) TransInTccNested(ctx context.Context, in *dtmgrpc.BusiRequest, opts ...grpc.CallOption) (*dtmgrpc.BusiReply, error) {
	out := new(dtmgrpc.BusiReply)
	err := c.cc.Invoke(ctx, "/examples.Busi/TransInTccNested", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusiServer is the server API for Busi service.
// All implementations must embed UnimplementedBusiServer
// for forward compatibility
type BusiServer interface {
	CanSubmit(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error)
	TransIn(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error)
	TransOut(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error)
	TransInRevert(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error)
	TransOutRevert(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error)
	TransInConfirm(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error)
	TransOutConfirm(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error)
	XaNotify(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error)
	TransInXa(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error)
	TransOutXa(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error)
	TransInTcc(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error)
	TransOutTcc(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error)
	TransInTccNested(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error)
	mustEmbedUnimplementedBusiServer()
}

// UnimplementedBusiServer must be embedded to have forward compatible implementations.
type UnimplementedBusiServer struct {
}

func (UnimplementedBusiServer) CanSubmit(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CanSubmit not implemented")
}
func (UnimplementedBusiServer) TransIn(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransIn not implemented")
}
func (UnimplementedBusiServer) TransOut(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOut not implemented")
}
func (UnimplementedBusiServer) TransInRevert(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransInRevert not implemented")
}
func (UnimplementedBusiServer) TransOutRevert(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOutRevert not implemented")
}
func (UnimplementedBusiServer) TransInConfirm(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransInConfirm not implemented")
}
func (UnimplementedBusiServer) TransOutConfirm(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOutConfirm not implemented")
}
func (UnimplementedBusiServer) XaNotify(context.Context, *dtmgrpc.BusiRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method XaNotify not implemented")
}
func (UnimplementedBusiServer) TransInXa(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransInXa not implemented")
}
func (UnimplementedBusiServer) TransOutXa(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOutXa not implemented")
}
func (UnimplementedBusiServer) TransInTcc(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransInTcc not implemented")
}
func (UnimplementedBusiServer) TransOutTcc(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransOutTcc not implemented")
}
func (UnimplementedBusiServer) TransInTccNested(context.Context, *dtmgrpc.BusiRequest) (*dtmgrpc.BusiReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransInTccNested not implemented")
}
func (UnimplementedBusiServer) mustEmbedUnimplementedBusiServer() {}

// UnsafeBusiServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusiServer will
// result in compilation errors.
type UnsafeBusiServer interface {
	mustEmbedUnimplementedBusiServer()
}

func RegisterBusiServer(s grpc.ServiceRegistrar, srv BusiServer) {
	s.RegisterService(&Busi_ServiceDesc, srv)
}

func _Busi_CanSubmit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).CanSubmit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/CanSubmit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).CanSubmit(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransIn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransIn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransIn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransIn(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransOut_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransOut(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransOut",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransOut(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransInRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransInRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransInRevert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransInRevert(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransOutRevert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransOutRevert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransOutRevert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransOutRevert(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransInConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransInConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransInConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransInConfirm(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransOutConfirm_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransOutConfirm(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransOutConfirm",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransOutConfirm(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_XaNotify_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).XaNotify(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/XaNotify",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).XaNotify(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransInXa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransInXa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransInXa",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransInXa(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransOutXa_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransOutXa(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransOutXa",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransOutXa(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransInTcc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransInTcc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransInTcc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransInTcc(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransOutTcc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransOutTcc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransOutTcc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransOutTcc(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Busi_TransInTccNested_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(dtmgrpc.BusiRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusiServer).TransInTccNested(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/examples.Busi/TransInTccNested",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusiServer).TransInTccNested(ctx, req.(*dtmgrpc.BusiRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Busi_ServiceDesc is the grpc.ServiceDesc for Busi service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Busi_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "examples.Busi",
	HandlerType: (*BusiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CanSubmit",
			Handler:    _Busi_CanSubmit_Handler,
		},
		{
			MethodName: "TransIn",
			Handler:    _Busi_TransIn_Handler,
		},
		{
			MethodName: "TransOut",
			Handler:    _Busi_TransOut_Handler,
		},
		{
			MethodName: "TransInRevert",
			Handler:    _Busi_TransInRevert_Handler,
		},
		{
			MethodName: "TransOutRevert",
			Handler:    _Busi_TransOutRevert_Handler,
		},
		{
			MethodName: "TransInConfirm",
			Handler:    _Busi_TransInConfirm_Handler,
		},
		{
			MethodName: "TransOutConfirm",
			Handler:    _Busi_TransOutConfirm_Handler,
		},
		{
			MethodName: "XaNotify",
			Handler:    _Busi_XaNotify_Handler,
		},
		{
			MethodName: "TransInXa",
			Handler:    _Busi_TransInXa_Handler,
		},
		{
			MethodName: "TransOutXa",
			Handler:    _Busi_TransOutXa_Handler,
		},
		{
			MethodName: "TransInTcc",
			Handler:    _Busi_TransInTcc_Handler,
		},
		{
			MethodName: "TransOutTcc",
			Handler:    _Busi_TransOutTcc_Handler,
		},
		{
			MethodName: "TransInTccNested",
			Handler:    _Busi_TransInTccNested_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "examples/busi.proto",
}
