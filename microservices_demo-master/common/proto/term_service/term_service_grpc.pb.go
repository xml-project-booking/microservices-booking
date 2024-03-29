// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0--rc3
// source: term_service.proto

package terms

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	TermService_Get_FullMethodName                                = "/terms.TermService/Get"
	TermService_GetAll_FullMethodName                             = "/terms.TermService/GetAll"
	TermService_Create_FullMethodName                             = "/terms.TermService/Create"
	TermService_Update_FullMethodName                             = "/terms.TermService/Update"
	TermService_Delete_FullMethodName                             = "/terms.TermService/Delete"
	TermService_DeleteInPeriod_FullMethodName                     = "/terms.TermService/DeleteInPeriod"
	TermService_UpdateInPeriod_FullMethodName                     = "/terms.TermService/UpdateInPeriod"
	TermService_GetByAccommodationId_FullMethodName               = "/terms.TermService/GetByAccommodationId"
	TermService_GetAvailableAccommodationsInPeriod_FullMethodName = "/terms.TermService/GetAvailableAccommodationsInPeriod"
	TermService_GetAllAccommodationIdsInPriceRange_FullMethodName = "/terms.TermService/GetAllAccommodationIdsInPriceRange"
	TermService_GetAllAccommodationIdsInTimePeriod_FullMethodName = "/terms.TermService/GetAllAccommodationIdsInTimePeriod"
	TermService_GetTermInfoByAccommodationId_FullMethodName       = "/terms.TermService/GetTermInfoByAccommodationId"
	TermService_ChangeUserIdInTermPeriod_FullMethodName           = "/terms.TermService/ChangeUserIdInTermPeriod"
	TermService_GetTermsInPeriod_FullMethodName                   = "/terms.TermService/GetTermsInPeriod"
)

// TermServiceClient is the client API for TermService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TermServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	DeleteInPeriod(ctx context.Context, in *DeleteInPeriodRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	UpdateInPeriod(ctx context.Context, in *UpdateInPeriodRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	GetByAccommodationId(ctx context.Context, in *GetByAccommodationIdRequest, opts ...grpc.CallOption) (*GetByAccommodationIdResponse, error)
	GetAvailableAccommodationsInPeriod(ctx context.Context, in *GetAvailableAccommodationsInPeriodRequest, opts ...grpc.CallOption) (*GetAvailableAccommodationsInPeriodResponse, error)
	GetAllAccommodationIdsInPriceRange(ctx context.Context, in *PriceRangeRequest, opts ...grpc.CallOption) (*PriceRangeResponse, error)
	GetAllAccommodationIdsInTimePeriod(ctx context.Context, in *TimePeriodRequest, opts ...grpc.CallOption) (*TimePeriodResponse, error)
	GetTermInfoByAccommodationId(ctx context.Context, in *TermInfoRequest, opts ...grpc.CallOption) (*TermInfoResponse, error)
	ChangeUserIdInTermPeriod(ctx context.Context, in *BookTermRequest, opts ...grpc.CallOption) (*BookTermResponse, error)
	GetTermsInPeriod(ctx context.Context, in *GetTermsInPeriodRequest, opts ...grpc.CallOption) (*GetTermsInPeriodResponse, error)
}

type termServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTermServiceClient(cc grpc.ClientConnInterface) TermServiceClient {
	return &termServiceClient{cc}
}

func (c *termServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, TermService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, TermService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, TermService_Create_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, TermService_Update_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, TermService_Delete_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) DeleteInPeriod(ctx context.Context, in *DeleteInPeriodRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, TermService_DeleteInPeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) UpdateInPeriod(ctx context.Context, in *UpdateInPeriodRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, TermService_UpdateInPeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) GetByAccommodationId(ctx context.Context, in *GetByAccommodationIdRequest, opts ...grpc.CallOption) (*GetByAccommodationIdResponse, error) {
	out := new(GetByAccommodationIdResponse)
	err := c.cc.Invoke(ctx, TermService_GetByAccommodationId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) GetAvailableAccommodationsInPeriod(ctx context.Context, in *GetAvailableAccommodationsInPeriodRequest, opts ...grpc.CallOption) (*GetAvailableAccommodationsInPeriodResponse, error) {
	out := new(GetAvailableAccommodationsInPeriodResponse)
	err := c.cc.Invoke(ctx, TermService_GetAvailableAccommodationsInPeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) GetAllAccommodationIdsInPriceRange(ctx context.Context, in *PriceRangeRequest, opts ...grpc.CallOption) (*PriceRangeResponse, error) {
	out := new(PriceRangeResponse)
	err := c.cc.Invoke(ctx, TermService_GetAllAccommodationIdsInPriceRange_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) GetAllAccommodationIdsInTimePeriod(ctx context.Context, in *TimePeriodRequest, opts ...grpc.CallOption) (*TimePeriodResponse, error) {
	out := new(TimePeriodResponse)
	err := c.cc.Invoke(ctx, TermService_GetAllAccommodationIdsInTimePeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) GetTermInfoByAccommodationId(ctx context.Context, in *TermInfoRequest, opts ...grpc.CallOption) (*TermInfoResponse, error) {
	out := new(TermInfoResponse)
	err := c.cc.Invoke(ctx, TermService_GetTermInfoByAccommodationId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) ChangeUserIdInTermPeriod(ctx context.Context, in *BookTermRequest, opts ...grpc.CallOption) (*BookTermResponse, error) {
	out := new(BookTermResponse)
	err := c.cc.Invoke(ctx, TermService_ChangeUserIdInTermPeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *termServiceClient) GetTermsInPeriod(ctx context.Context, in *GetTermsInPeriodRequest, opts ...grpc.CallOption) (*GetTermsInPeriodResponse, error) {
	out := new(GetTermsInPeriodResponse)
	err := c.cc.Invoke(ctx, TermService_GetTermsInPeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TermServiceServer is the server API for TermService service.
// All implementations must embed UnimplementedTermServiceServer
// for forward compatibility
type TermServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	Create(context.Context, *CreateRequest) (*GetAllResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	DeleteInPeriod(context.Context, *DeleteInPeriodRequest) (*DeleteResponse, error)
	UpdateInPeriod(context.Context, *UpdateInPeriodRequest) (*UpdateResponse, error)
	GetByAccommodationId(context.Context, *GetByAccommodationIdRequest) (*GetByAccommodationIdResponse, error)
	GetAvailableAccommodationsInPeriod(context.Context, *GetAvailableAccommodationsInPeriodRequest) (*GetAvailableAccommodationsInPeriodResponse, error)
	GetAllAccommodationIdsInPriceRange(context.Context, *PriceRangeRequest) (*PriceRangeResponse, error)
	GetAllAccommodationIdsInTimePeriod(context.Context, *TimePeriodRequest) (*TimePeriodResponse, error)
	GetTermInfoByAccommodationId(context.Context, *TermInfoRequest) (*TermInfoResponse, error)
	ChangeUserIdInTermPeriod(context.Context, *BookTermRequest) (*BookTermResponse, error)
	GetTermsInPeriod(context.Context, *GetTermsInPeriodRequest) (*GetTermsInPeriodResponse, error)
	mustEmbedUnimplementedTermServiceServer()
}

// UnimplementedTermServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTermServiceServer struct {
}

func (UnimplementedTermServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedTermServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedTermServiceServer) Create(context.Context, *CreateRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedTermServiceServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedTermServiceServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedTermServiceServer) DeleteInPeriod(context.Context, *DeleteInPeriodRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteInPeriod not implemented")
}
func (UnimplementedTermServiceServer) UpdateInPeriod(context.Context, *UpdateInPeriodRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateInPeriod not implemented")
}
func (UnimplementedTermServiceServer) GetByAccommodationId(context.Context, *GetByAccommodationIdRequest) (*GetByAccommodationIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetByAccommodationId not implemented")
}
func (UnimplementedTermServiceServer) GetAvailableAccommodationsInPeriod(context.Context, *GetAvailableAccommodationsInPeriodRequest) (*GetAvailableAccommodationsInPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableAccommodationsInPeriod not implemented")
}
func (UnimplementedTermServiceServer) GetAllAccommodationIdsInPriceRange(context.Context, *PriceRangeRequest) (*PriceRangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAccommodationIdsInPriceRange not implemented")
}
func (UnimplementedTermServiceServer) GetAllAccommodationIdsInTimePeriod(context.Context, *TimePeriodRequest) (*TimePeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllAccommodationIdsInTimePeriod not implemented")
}
func (UnimplementedTermServiceServer) GetTermInfoByAccommodationId(context.Context, *TermInfoRequest) (*TermInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTermInfoByAccommodationId not implemented")
}
func (UnimplementedTermServiceServer) ChangeUserIdInTermPeriod(context.Context, *BookTermRequest) (*BookTermResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeUserIdInTermPeriod not implemented")
}
func (UnimplementedTermServiceServer) GetTermsInPeriod(context.Context, *GetTermsInPeriodRequest) (*GetTermsInPeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTermsInPeriod not implemented")
}
func (UnimplementedTermServiceServer) mustEmbedUnimplementedTermServiceServer() {}

// UnsafeTermServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TermServiceServer will
// result in compilation errors.
type UnsafeTermServiceServer interface {
	mustEmbedUnimplementedTermServiceServer()
}

func RegisterTermServiceServer(s grpc.ServiceRegistrar, srv TermServiceServer) {
	s.RegisterService(&TermService_ServiceDesc, srv)
}

func _TermService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_DeleteInPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteInPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).DeleteInPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_DeleteInPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).DeleteInPeriod(ctx, req.(*DeleteInPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_UpdateInPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateInPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).UpdateInPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_UpdateInPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).UpdateInPeriod(ctx, req.(*UpdateInPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_GetByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByAccommodationIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).GetByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_GetByAccommodationId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).GetByAccommodationId(ctx, req.(*GetByAccommodationIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_GetAvailableAccommodationsInPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAvailableAccommodationsInPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).GetAvailableAccommodationsInPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_GetAvailableAccommodationsInPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).GetAvailableAccommodationsInPeriod(ctx, req.(*GetAvailableAccommodationsInPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_GetAllAccommodationIdsInPriceRange_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PriceRangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).GetAllAccommodationIdsInPriceRange(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_GetAllAccommodationIdsInPriceRange_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).GetAllAccommodationIdsInPriceRange(ctx, req.(*PriceRangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_GetAllAccommodationIdsInTimePeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TimePeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).GetAllAccommodationIdsInTimePeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_GetAllAccommodationIdsInTimePeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).GetAllAccommodationIdsInTimePeriod(ctx, req.(*TimePeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_GetTermInfoByAccommodationId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TermInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).GetTermInfoByAccommodationId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_GetTermInfoByAccommodationId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).GetTermInfoByAccommodationId(ctx, req.(*TermInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_ChangeUserIdInTermPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookTermRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).ChangeUserIdInTermPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_ChangeUserIdInTermPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).ChangeUserIdInTermPeriod(ctx, req.(*BookTermRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TermService_GetTermsInPeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTermsInPeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TermServiceServer).GetTermsInPeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: TermService_GetTermsInPeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TermServiceServer).GetTermsInPeriod(ctx, req.(*GetTermsInPeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TermService_ServiceDesc is the grpc.ServiceDesc for TermService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TermService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "terms.TermService",
	HandlerType: (*TermServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _TermService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _TermService_GetAll_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _TermService_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _TermService_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _TermService_Delete_Handler,
		},
		{
			MethodName: "DeleteInPeriod",
			Handler:    _TermService_DeleteInPeriod_Handler,
		},
		{
			MethodName: "UpdateInPeriod",
			Handler:    _TermService_UpdateInPeriod_Handler,
		},
		{
			MethodName: "GetByAccommodationId",
			Handler:    _TermService_GetByAccommodationId_Handler,
		},
		{
			MethodName: "GetAvailableAccommodationsInPeriod",
			Handler:    _TermService_GetAvailableAccommodationsInPeriod_Handler,
		},
		{
			MethodName: "GetAllAccommodationIdsInPriceRange",
			Handler:    _TermService_GetAllAccommodationIdsInPriceRange_Handler,
		},
		{
			MethodName: "GetAllAccommodationIdsInTimePeriod",
			Handler:    _TermService_GetAllAccommodationIdsInTimePeriod_Handler,
		},
		{
			MethodName: "GetTermInfoByAccommodationId",
			Handler:    _TermService_GetTermInfoByAccommodationId_Handler,
		},
		{
			MethodName: "ChangeUserIdInTermPeriod",
			Handler:    _TermService_ChangeUserIdInTermPeriod_Handler,
		},
		{
			MethodName: "GetTermsInPeriod",
			Handler:    _TermService_GetTermsInPeriod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "term_service.proto",
}
