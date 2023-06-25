// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.4
// source: accommodation_service.proto

package accommodations

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
	AccommodationService_Get_FullMethodName                                = "/accommodations.AccommodationService/Get"
	AccommodationService_GetAll_FullMethodName                             = "/accommodations.AccommodationService/GetAll"
	AccommodationService_ChangeAccommodationReservationType_FullMethodName = "/accommodations.AccommodationService/ChangeAccommodationReservationType"
	AccommodationService_CreateAccommodation_FullMethodName                = "/accommodations.AccommodationService/CreateAccommodation"
<<<<<<< HEAD
	AccommodationService_GetAllIdsByHost_FullMethodName                    = "/accommodations.AccommodationService/GetAllIdsByHost"
	AccommodationService_DeleteAllByHost_FullMethodName                    = "/accommodations.AccommodationService/DeleteAllByHost"
=======
	AccommodationService_GetAccommodation_FullMethodName                   = "/accommodations.AccommodationService/GetAccommodation"
>>>>>>> bf2cdc8 (fixed mistakes made in implmentation of resrvation service)
)

// AccommodationServiceClient is the client API for AccommodationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccommodationServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	ChangeAccommodationReservationType(ctx context.Context, in *ChangeReservationTypeRequest, opts ...grpc.CallOption) (*ChangeReservationTypeResponse, error)
	CreateAccommodation(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error)
<<<<<<< HEAD
	GetAllIdsByHost(ctx context.Context, in *GetAllIdsByHostRequest, opts ...grpc.CallOption) (*GetAllIdsByHostResponse, error)
	DeleteAllByHost(ctx context.Context, in *DeleteAllByHostRequest, opts ...grpc.CallOption) (*DeleteAllByHostResponse, error)
=======
	GetAccommodation(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAccommodationResponse, error)
>>>>>>> bf2cdc8 (fixed mistakes made in implmentation of resrvation service)
}

type accommodationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccommodationServiceClient(cc grpc.ClientConnInterface) AccommodationServiceClient {
	return &accommodationServiceClient{cc}
}

func (c *accommodationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, AccommodationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) ChangeAccommodationReservationType(ctx context.Context, in *ChangeReservationTypeRequest, opts ...grpc.CallOption) (*ChangeReservationTypeResponse, error) {
	out := new(ChangeReservationTypeResponse)
	err := c.cc.Invoke(ctx, AccommodationService_ChangeAccommodationReservationType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) CreateAccommodation(ctx context.Context, in *CreateAccommodationRequest, opts ...grpc.CallOption) (*CreateAccommodationResponse, error) {
	out := new(CreateAccommodationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_CreateAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

<<<<<<< HEAD
func (c *accommodationServiceClient) GetAllIdsByHost(ctx context.Context, in *GetAllIdsByHostRequest, opts ...grpc.CallOption) (*GetAllIdsByHostResponse, error) {
	out := new(GetAllIdsByHostResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAllIdsByHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accommodationServiceClient) DeleteAllByHost(ctx context.Context, in *DeleteAllByHostRequest, opts ...grpc.CallOption) (*DeleteAllByHostResponse, error) {
	out := new(DeleteAllByHostResponse)
	err := c.cc.Invoke(ctx, AccommodationService_DeleteAllByHost_FullMethodName, in, out, opts...)
=======
func (c *accommodationServiceClient) GetAccommodation(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetAccommodationResponse, error) {
	out := new(GetAccommodationResponse)
	err := c.cc.Invoke(ctx, AccommodationService_GetAccommodation_FullMethodName, in, out, opts...)
>>>>>>> bf2cdc8 (fixed mistakes made in implmentation of resrvation service)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccommodationServiceServer is the server API for AccommodationService service.
// All implementations must embed UnimplementedAccommodationServiceServer
// for forward compatibility
type AccommodationServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	ChangeAccommodationReservationType(context.Context, *ChangeReservationTypeRequest) (*ChangeReservationTypeResponse, error)
	CreateAccommodation(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error)
<<<<<<< HEAD
	GetAllIdsByHost(context.Context, *GetAllIdsByHostRequest) (*GetAllIdsByHostResponse, error)
	DeleteAllByHost(context.Context, *DeleteAllByHostRequest) (*DeleteAllByHostResponse, error)
=======
	GetAccommodation(context.Context, *GetRequest) (*GetAccommodationResponse, error)
>>>>>>> bf2cdc8 (fixed mistakes made in implmentation of resrvation service)
	mustEmbedUnimplementedAccommodationServiceServer()
}

// UnimplementedAccommodationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedAccommodationServiceServer struct {
}

func (UnimplementedAccommodationServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedAccommodationServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedAccommodationServiceServer) ChangeAccommodationReservationType(context.Context, *ChangeReservationTypeRequest) (*ChangeReservationTypeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeAccommodationReservationType not implemented")
}
func (UnimplementedAccommodationServiceServer) CreateAccommodation(context.Context, *CreateAccommodationRequest) (*CreateAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccommodation not implemented")
}
<<<<<<< HEAD
func (UnimplementedAccommodationServiceServer) GetAllIdsByHost(context.Context, *GetAllIdsByHostRequest) (*GetAllIdsByHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllIdsByHost not implemented")
}
func (UnimplementedAccommodationServiceServer) DeleteAllByHost(context.Context, *DeleteAllByHostRequest) (*DeleteAllByHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAllByHost not implemented")
=======
func (UnimplementedAccommodationServiceServer) GetAccommodation(context.Context, *GetRequest) (*GetAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodation not implemented")
>>>>>>> bf2cdc8 (fixed mistakes made in implmentation of resrvation service)
}
func (UnimplementedAccommodationServiceServer) mustEmbedUnimplementedAccommodationServiceServer() {}

// UnsafeAccommodationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccommodationServiceServer will
// result in compilation errors.
type UnsafeAccommodationServiceServer interface {
	mustEmbedUnimplementedAccommodationServiceServer()
}

func RegisterAccommodationServiceServer(s grpc.ServiceRegistrar, srv AccommodationServiceServer) {
	s.RegisterService(&AccommodationService_ServiceDesc, srv)
}

func _AccommodationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_ChangeAccommodationReservationType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeReservationTypeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).ChangeAccommodationReservationType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_ChangeAccommodationReservationType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).ChangeAccommodationReservationType(ctx, req.(*ChangeReservationTypeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_CreateAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_CreateAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).CreateAccommodation(ctx, req.(*CreateAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

<<<<<<< HEAD
func _AccommodationService_GetAllIdsByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllIdsByHostRequest)
=======
func _AccommodationService_GetAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
>>>>>>> bf2cdc8 (fixed mistakes made in implmentation of resrvation service)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
<<<<<<< HEAD
		return srv.(AccommodationServiceServer).GetAllIdsByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAllIdsByHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAllIdsByHost(ctx, req.(*GetAllIdsByHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccommodationService_DeleteAllByHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAllByHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccommodationServiceServer).DeleteAllByHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_DeleteAllByHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).DeleteAllByHost(ctx, req.(*DeleteAllByHostRequest))
=======
		return srv.(AccommodationServiceServer).GetAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: AccommodationService_GetAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccommodationServiceServer).GetAccommodation(ctx, req.(*GetRequest))
>>>>>>> bf2cdc8 (fixed mistakes made in implmentation of resrvation service)
	}
	return interceptor(ctx, in, info, handler)
}

// AccommodationService_ServiceDesc is the grpc.ServiceDesc for AccommodationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccommodationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "accommodations.AccommodationService",
	HandlerType: (*AccommodationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _AccommodationService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _AccommodationService_GetAll_Handler,
		},
		{
			MethodName: "ChangeAccommodationReservationType",
			Handler:    _AccommodationService_ChangeAccommodationReservationType_Handler,
		},
		{
			MethodName: "CreateAccommodation",
			Handler:    _AccommodationService_CreateAccommodation_Handler,
		},
		{
<<<<<<< HEAD
			MethodName: "GetAllIdsByHost",
			Handler:    _AccommodationService_GetAllIdsByHost_Handler,
		},
		{
			MethodName: "DeleteAllByHost",
			Handler:    _AccommodationService_DeleteAllByHost_Handler,
=======
			MethodName: "GetAccommodation",
			Handler:    _AccommodationService_GetAccommodation_Handler,
>>>>>>> bf2cdc8 (fixed mistakes made in implmentation of resrvation service)
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "accommodation_service.proto",
}
