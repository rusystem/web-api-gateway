// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.20.3
// source: proto/supplier/supplier.proto

package supplier

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	SupplierService_Create_FullMethodName  = "/supplier.SupplierService/Create"
	SupplierService_GetById_FullMethodName = "/supplier.SupplierService/GetById"
	SupplierService_Update_FullMethodName  = "/supplier.SupplierService/Update"
	SupplierService_Delete_FullMethodName  = "/supplier.SupplierService/Delete"
	SupplierService_GetList_FullMethodName = "/supplier.SupplierService/GetList"
)

// SupplierServiceClient is the client API for SupplierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupplierServiceClient interface {
	Create(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*SupplierId, error)
	GetById(ctx context.Context, in *SupplierId, opts ...grpc.CallOption) (*Supplier, error)
	Update(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Delete(ctx context.Context, in *SupplierId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetList(ctx context.Context, in *SupplierCompanyId, opts ...grpc.CallOption) (*SupplierList, error)
}

type supplierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupplierServiceClient(cc grpc.ClientConnInterface) SupplierServiceClient {
	return &supplierServiceClient{cc}
}

func (c *supplierServiceClient) Create(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*SupplierId, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SupplierId)
	err := c.cc.Invoke(ctx, SupplierService_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) GetById(ctx context.Context, in *SupplierId, opts ...grpc.CallOption) (*Supplier, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Supplier)
	err := c.cc.Invoke(ctx, SupplierService_GetById_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) Update(ctx context.Context, in *Supplier, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SupplierService_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) Delete(ctx context.Context, in *SupplierId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, SupplierService_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *supplierServiceClient) GetList(ctx context.Context, in *SupplierCompanyId, opts ...grpc.CallOption) (*SupplierList, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SupplierList)
	err := c.cc.Invoke(ctx, SupplierService_GetList_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupplierServiceServer is the server API for SupplierService service.
// All implementations should embed UnimplementedSupplierServiceServer
// for forward compatibility
type SupplierServiceServer interface {
	Create(context.Context, *Supplier) (*SupplierId, error)
	GetById(context.Context, *SupplierId) (*Supplier, error)
	Update(context.Context, *Supplier) (*emptypb.Empty, error)
	Delete(context.Context, *SupplierId) (*emptypb.Empty, error)
	GetList(context.Context, *SupplierCompanyId) (*SupplierList, error)
}

// UnimplementedSupplierServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSupplierServiceServer struct {
}

func (UnimplementedSupplierServiceServer) Create(context.Context, *Supplier) (*SupplierId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedSupplierServiceServer) GetById(context.Context, *SupplierId) (*Supplier, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetById not implemented")
}
func (UnimplementedSupplierServiceServer) Update(context.Context, *Supplier) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedSupplierServiceServer) Delete(context.Context, *SupplierId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedSupplierServiceServer) GetList(context.Context, *SupplierCompanyId) (*SupplierList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetList not implemented")
}

// UnsafeSupplierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupplierServiceServer will
// result in compilation errors.
type UnsafeSupplierServiceServer interface {
	mustEmbedUnimplementedSupplierServiceServer()
}

func RegisterSupplierServiceServer(s grpc.ServiceRegistrar, srv SupplierServiceServer) {
	s.RegisterService(&SupplierService_ServiceDesc, srv)
}

func _SupplierService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupplierService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Create(ctx, req.(*Supplier))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_GetById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplierId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).GetById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupplierService_GetById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).GetById(ctx, req.(*SupplierId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Supplier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupplierService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Update(ctx, req.(*Supplier))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplierId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupplierService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).Delete(ctx, req.(*SupplierId))
	}
	return interceptor(ctx, in, info, handler)
}

func _SupplierService_GetList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SupplierCompanyId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupplierServiceServer).GetList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SupplierService_GetList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupplierServiceServer).GetList(ctx, req.(*SupplierCompanyId))
	}
	return interceptor(ctx, in, info, handler)
}

// SupplierService_ServiceDesc is the grpc.ServiceDesc for SupplierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupplierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "supplier.SupplierService",
	HandlerType: (*SupplierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _SupplierService_Create_Handler,
		},
		{
			MethodName: "GetById",
			Handler:    _SupplierService_GetById_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _SupplierService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _SupplierService_Delete_Handler,
		},
		{
			MethodName: "GetList",
			Handler:    _SupplierService_GetList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/supplier/supplier.proto",
}
