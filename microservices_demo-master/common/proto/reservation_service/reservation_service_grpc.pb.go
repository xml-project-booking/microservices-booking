// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.4
// source: reservation_service.proto

package reservations

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
	ReservationService_Get_FullMethodName                                   = "/reservations.ReservationService/Get"
	ReservationService_GetAll_FullMethodName                                = "/reservations.ReservationService/GetAll"
	ReservationService_GetAllByAccommodation_FullMethodName                 = "/reservations.ReservationService/GetAllByAccommodation"
	ReservationService_GetAllByAccommodationConfirmed_FullMethodName        = "/reservations.ReservationService/GetAllByAccommodationConfirmed"
	ReservationService_GetAllByGuest_FullMethodName                         = "/reservations.ReservationService/GetAllByGuest"
	ReservationService_GetAllByGuestPending_FullMethodName                  = "/reservations.ReservationService/GetAllByGuestPending"
	ReservationService_MakeRequestForReservation_FullMethodName             = "/reservations.ReservationService/MakeRequestForReservation"
	ReservationService_CancelReservation_FullMethodName                     = "/reservations.ReservationService/CancelReservation"
	ReservationService_ConfirmReservationManually_FullMethodName            = "/reservations.ReservationService/ConfirmReservationManually"
	ReservationService_CancelReservationManually_FullMethodName             = "/reservations.ReservationService/CancelReservationManually"
	ReservationService_ConfirmReservationAutomatically_FullMethodName       = "/reservations.ReservationService/ConfirmReservationAutomatically"
	ReservationService_HasActiveReservations_FullMethodName                 = "/reservations.ReservationService/HasActiveReservations"
	ReservationService_GetAllFuture_FullMethodName                          = "/reservations.ReservationService/GetAllFuture"
	ReservationService_DeleteReservationRequestGuest_FullMethodName         = "/reservations.ReservationService/DeleteReservationRequestGuest"
	ReservationService_TermCheck_FullMethodName                             = "/reservations.ReservationService/TermCheck"
	ReservationService_CheckReservationRequirementsHost_FullMethodName      = "/reservations.ReservationService/CheckReservationRequirementsHost"
	ReservationService_GetAccommodationsReservedInTimePeriod_FullMethodName = "/reservations.ReservationService/GetAccommodationsReservedInTimePeriod"
)

// ReservationServiceClient is the client API for ReservationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReservationServiceClient interface {
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
	GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error)
	GetAllByAccommodation(ctx context.Context, in *GetAllByAccommodationRequest, opts ...grpc.CallOption) (*GetAllByAccommodationResponse, error)
	GetAllByAccommodationConfirmed(ctx context.Context, in *GetAllByAccommodationRequest, opts ...grpc.CallOption) (*GetAllByAccommodationResponse, error)
	GetAllByGuest(ctx context.Context, in *GetAllByGuestRequest, opts ...grpc.CallOption) (*GetAllByAccommodationResponse, error)
	GetAllByGuestPending(ctx context.Context, in *GetAllByGuestRequest, opts ...grpc.CallOption) (*GetAllByAccommodationResponse, error)
	MakeRequestForReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationRequestResponse, error)
	CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error)
	ConfirmReservationManually(ctx context.Context, in *ConfirmReservationManuallyRequest, opts ...grpc.CallOption) (*ConfirmReservationManuallyResponse, error)
	CancelReservationManually(ctx context.Context, in *CancelReservationManuallyRequest, opts ...grpc.CallOption) (*CancelReservationManuallyResponse, error)
	ConfirmReservationAutomatically(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ConfirmReservationAutomaticallyMessage, error)
	HasActiveReservations(ctx context.Context, in *HasActiveReservationsRequest, opts ...grpc.CallOption) (*HasActiveReservationsResponse, error)
	GetAllFuture(ctx context.Context, in *GetAllFutureRequest, opts ...grpc.CallOption) (*GetAllFutureResponse, error)
	DeleteReservationRequestGuest(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error)
	TermCheck(ctx context.Context, in *TermCheckRequest, opts ...grpc.CallOption) (*TermCheckResponse, error)
	CheckReservationRequirementsHost(ctx context.Context, in *ReservationRequirementsHostRequest, opts ...grpc.CallOption) (*ReservationRequirementsHostResponse, error)
	GetAccommodationsReservedInTimePeriod(ctx context.Context, in *GetAccTimePeriodRequest, opts ...grpc.CallOption) (*GetAccTimePeriodResponse, error)
}

type reservationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReservationServiceClient(cc grpc.ClientConnInterface) ReservationServiceClient {
	return &reservationServiceClient{cc}
}

func (c *reservationServiceClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, ReservationService_Get_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAll(ctx context.Context, in *GetAllRequest, opts ...grpc.CallOption) (*GetAllResponse, error) {
	out := new(GetAllResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllByAccommodation(ctx context.Context, in *GetAllByAccommodationRequest, opts ...grpc.CallOption) (*GetAllByAccommodationResponse, error) {
	out := new(GetAllByAccommodationResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllByAccommodation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllByAccommodationConfirmed(ctx context.Context, in *GetAllByAccommodationRequest, opts ...grpc.CallOption) (*GetAllByAccommodationResponse, error) {
	out := new(GetAllByAccommodationResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllByAccommodationConfirmed_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllByGuest(ctx context.Context, in *GetAllByGuestRequest, opts ...grpc.CallOption) (*GetAllByAccommodationResponse, error) {
	out := new(GetAllByAccommodationResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllByGuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllByGuestPending(ctx context.Context, in *GetAllByGuestRequest, opts ...grpc.CallOption) (*GetAllByAccommodationResponse, error) {
	out := new(GetAllByAccommodationResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllByGuestPending_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) MakeRequestForReservation(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ReservationRequestResponse, error) {
	out := new(ReservationRequestResponse)
	err := c.cc.Invoke(ctx, ReservationService_MakeRequestForReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CancelReservation(ctx context.Context, in *CancelReservationRequest, opts ...grpc.CallOption) (*CancelReservationResponse, error) {
	out := new(CancelReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_CancelReservation_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ConfirmReservationManually(ctx context.Context, in *ConfirmReservationManuallyRequest, opts ...grpc.CallOption) (*ConfirmReservationManuallyResponse, error) {
	out := new(ConfirmReservationManuallyResponse)
	err := c.cc.Invoke(ctx, ReservationService_ConfirmReservationManually_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CancelReservationManually(ctx context.Context, in *CancelReservationManuallyRequest, opts ...grpc.CallOption) (*CancelReservationManuallyResponse, error) {
	out := new(CancelReservationManuallyResponse)
	err := c.cc.Invoke(ctx, ReservationService_CancelReservationManually_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) ConfirmReservationAutomatically(ctx context.Context, in *ReservationRequest, opts ...grpc.CallOption) (*ConfirmReservationAutomaticallyMessage, error) {
	out := new(ConfirmReservationAutomaticallyMessage)
	err := c.cc.Invoke(ctx, ReservationService_ConfirmReservationAutomatically_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) HasActiveReservations(ctx context.Context, in *HasActiveReservationsRequest, opts ...grpc.CallOption) (*HasActiveReservationsResponse, error) {
	out := new(HasActiveReservationsResponse)
	err := c.cc.Invoke(ctx, ReservationService_HasActiveReservations_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAllFuture(ctx context.Context, in *GetAllFutureRequest, opts ...grpc.CallOption) (*GetAllFutureResponse, error) {
	out := new(GetAllFutureResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAllFuture_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) DeleteReservationRequestGuest(ctx context.Context, in *DeleteReservationRequest, opts ...grpc.CallOption) (*DeleteReservationResponse, error) {
	out := new(DeleteReservationResponse)
	err := c.cc.Invoke(ctx, ReservationService_DeleteReservationRequestGuest_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) TermCheck(ctx context.Context, in *TermCheckRequest, opts ...grpc.CallOption) (*TermCheckResponse, error) {
	out := new(TermCheckResponse)
	err := c.cc.Invoke(ctx, ReservationService_TermCheck_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) CheckReservationRequirementsHost(ctx context.Context, in *ReservationRequirementsHostRequest, opts ...grpc.CallOption) (*ReservationRequirementsHostResponse, error) {
	out := new(ReservationRequirementsHostResponse)
	err := c.cc.Invoke(ctx, ReservationService_CheckReservationRequirementsHost_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reservationServiceClient) GetAccommodationsReservedInTimePeriod(ctx context.Context, in *GetAccTimePeriodRequest, opts ...grpc.CallOption) (*GetAccTimePeriodResponse, error) {
	out := new(GetAccTimePeriodResponse)
	err := c.cc.Invoke(ctx, ReservationService_GetAccommodationsReservedInTimePeriod_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ReservationServiceServer is the server API for ReservationService service.
// All implementations must embed UnimplementedReservationServiceServer
// for forward compatibility
type ReservationServiceServer interface {
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error)
	GetAllByAccommodation(context.Context, *GetAllByAccommodationRequest) (*GetAllByAccommodationResponse, error)
	GetAllByAccommodationConfirmed(context.Context, *GetAllByAccommodationRequest) (*GetAllByAccommodationResponse, error)
	GetAllByGuest(context.Context, *GetAllByGuestRequest) (*GetAllByAccommodationResponse, error)
	GetAllByGuestPending(context.Context, *GetAllByGuestRequest) (*GetAllByAccommodationResponse, error)
	MakeRequestForReservation(context.Context, *ReservationRequest) (*ReservationRequestResponse, error)
	CancelReservation(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error)
	ConfirmReservationManually(context.Context, *ConfirmReservationManuallyRequest) (*ConfirmReservationManuallyResponse, error)
	CancelReservationManually(context.Context, *CancelReservationManuallyRequest) (*CancelReservationManuallyResponse, error)
	ConfirmReservationAutomatically(context.Context, *ReservationRequest) (*ConfirmReservationAutomaticallyMessage, error)
	HasActiveReservations(context.Context, *HasActiveReservationsRequest) (*HasActiveReservationsResponse, error)
	GetAllFuture(context.Context, *GetAllFutureRequest) (*GetAllFutureResponse, error)
	DeleteReservationRequestGuest(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error)
	TermCheck(context.Context, *TermCheckRequest) (*TermCheckResponse, error)
	CheckReservationRequirementsHost(context.Context, *ReservationRequirementsHostRequest) (*ReservationRequirementsHostResponse, error)
	GetAccommodationsReservedInTimePeriod(context.Context, *GetAccTimePeriodRequest) (*GetAccTimePeriodResponse, error)
	mustEmbedUnimplementedReservationServiceServer()
}

// UnimplementedReservationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReservationServiceServer struct {
}

func (UnimplementedReservationServiceServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedReservationServiceServer) GetAll(context.Context, *GetAllRequest) (*GetAllResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAll not implemented")
}
func (UnimplementedReservationServiceServer) GetAllByAccommodation(context.Context, *GetAllByAccommodationRequest) (*GetAllByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByAccommodation not implemented")
}
func (UnimplementedReservationServiceServer) GetAllByAccommodationConfirmed(context.Context, *GetAllByAccommodationRequest) (*GetAllByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByAccommodationConfirmed not implemented")
}
func (UnimplementedReservationServiceServer) GetAllByGuest(context.Context, *GetAllByGuestRequest) (*GetAllByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByGuest not implemented")
}
func (UnimplementedReservationServiceServer) GetAllByGuestPending(context.Context, *GetAllByGuestRequest) (*GetAllByAccommodationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllByGuestPending not implemented")
}
func (UnimplementedReservationServiceServer) MakeRequestForReservation(context.Context, *ReservationRequest) (*ReservationRequestResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method MakeRequestForReservation not implemented")
}
func (UnimplementedReservationServiceServer) CancelReservation(context.Context, *CancelReservationRequest) (*CancelReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservation not implemented")
}
func (UnimplementedReservationServiceServer) ConfirmReservationManually(context.Context, *ConfirmReservationManuallyRequest) (*ConfirmReservationManuallyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReservationManually not implemented")
}
func (UnimplementedReservationServiceServer) CancelReservationManually(context.Context, *CancelReservationManuallyRequest) (*CancelReservationManuallyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CancelReservationManually not implemented")
}
func (UnimplementedReservationServiceServer) ConfirmReservationAutomatically(context.Context, *ReservationRequest) (*ConfirmReservationAutomaticallyMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmReservationAutomatically not implemented")
}
func (UnimplementedReservationServiceServer) HasActiveReservations(context.Context, *HasActiveReservationsRequest) (*HasActiveReservationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HasActiveReservations not implemented")
}
func (UnimplementedReservationServiceServer) GetAllFuture(context.Context, *GetAllFutureRequest) (*GetAllFutureResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllFuture not implemented")
}
func (UnimplementedReservationServiceServer) DeleteReservationRequestGuest(context.Context, *DeleteReservationRequest) (*DeleteReservationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteReservationRequestGuest not implemented")
}
func (UnimplementedReservationServiceServer) TermCheck(context.Context, *TermCheckRequest) (*TermCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TermCheck not implemented")
}
func (UnimplementedReservationServiceServer) CheckReservationRequirementsHost(context.Context, *ReservationRequirementsHostRequest) (*ReservationRequirementsHostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckReservationRequirementsHost not implemented")
}
func (UnimplementedReservationServiceServer) GetAccommodationsReservedInTimePeriod(context.Context, *GetAccTimePeriodRequest) (*GetAccTimePeriodResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccommodationsReservedInTimePeriod not implemented")
}
func (UnimplementedReservationServiceServer) mustEmbedUnimplementedReservationServiceServer() {}

// UnsafeReservationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReservationServiceServer will
// result in compilation errors.
type UnsafeReservationServiceServer interface {
	mustEmbedUnimplementedReservationServiceServer()
}

func RegisterReservationServiceServer(s grpc.ServiceRegistrar, srv ReservationServiceServer) {
	s.RegisterService(&ReservationService_ServiceDesc, srv)
}

func _ReservationService_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_Get_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAll(ctx, req.(*GetAllRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllByAccommodation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllByAccommodation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllByAccommodation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllByAccommodation(ctx, req.(*GetAllByAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllByAccommodationConfirmed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByAccommodationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllByAccommodationConfirmed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllByAccommodationConfirmed_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllByAccommodationConfirmed(ctx, req.(*GetAllByAccommodationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllByGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllByGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllByGuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllByGuest(ctx, req.(*GetAllByGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllByGuestPending_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllByGuestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllByGuestPending(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllByGuestPending_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllByGuestPending(ctx, req.(*GetAllByGuestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_MakeRequestForReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).MakeRequestForReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_MakeRequestForReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).MakeRequestForReservation(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CancelReservation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CancelReservation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CancelReservation_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CancelReservation(ctx, req.(*CancelReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ConfirmReservationManually_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmReservationManuallyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ConfirmReservationManually(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_ConfirmReservationManually_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ConfirmReservationManually(ctx, req.(*ConfirmReservationManuallyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CancelReservationManually_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CancelReservationManuallyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CancelReservationManually(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CancelReservationManually_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CancelReservationManually(ctx, req.(*CancelReservationManuallyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_ConfirmReservationAutomatically_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).ConfirmReservationAutomatically(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_ConfirmReservationAutomatically_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).ConfirmReservationAutomatically(ctx, req.(*ReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_HasActiveReservations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HasActiveReservationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).HasActiveReservations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_HasActiveReservations_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).HasActiveReservations(ctx, req.(*HasActiveReservationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAllFuture_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllFutureRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAllFuture(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAllFuture_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAllFuture(ctx, req.(*GetAllFutureRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_DeleteReservationRequestGuest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteReservationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).DeleteReservationRequestGuest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_DeleteReservationRequestGuest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).DeleteReservationRequestGuest(ctx, req.(*DeleteReservationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_TermCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TermCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).TermCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_TermCheck_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).TermCheck(ctx, req.(*TermCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_CheckReservationRequirementsHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReservationRequirementsHostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).CheckReservationRequirementsHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_CheckReservationRequirementsHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).CheckReservationRequirementsHost(ctx, req.(*ReservationRequirementsHostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReservationService_GetAccommodationsReservedInTimePeriod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccTimePeriodRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReservationServiceServer).GetAccommodationsReservedInTimePeriod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ReservationService_GetAccommodationsReservedInTimePeriod_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReservationServiceServer).GetAccommodationsReservedInTimePeriod(ctx, req.(*GetAccTimePeriodRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ReservationService_ServiceDesc is the grpc.ServiceDesc for ReservationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReservationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reservations.ReservationService",
	HandlerType: (*ReservationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Get",
			Handler:    _ReservationService_Get_Handler,
		},
		{
			MethodName: "GetAll",
			Handler:    _ReservationService_GetAll_Handler,
		},
		{
			MethodName: "GetAllByAccommodation",
			Handler:    _ReservationService_GetAllByAccommodation_Handler,
		},
		{
			MethodName: "GetAllByAccommodationConfirmed",
			Handler:    _ReservationService_GetAllByAccommodationConfirmed_Handler,
		},
		{
			MethodName: "GetAllByGuest",
			Handler:    _ReservationService_GetAllByGuest_Handler,
		},
		{
			MethodName: "GetAllByGuestPending",
			Handler:    _ReservationService_GetAllByGuestPending_Handler,
		},
		{
			MethodName: "MakeRequestForReservation",
			Handler:    _ReservationService_MakeRequestForReservation_Handler,
		},
		{
			MethodName: "CancelReservation",
			Handler:    _ReservationService_CancelReservation_Handler,
		},
		{
			MethodName: "ConfirmReservationManually",
			Handler:    _ReservationService_ConfirmReservationManually_Handler,
		},
		{
			MethodName: "CancelReservationManually",
			Handler:    _ReservationService_CancelReservationManually_Handler,
		},
		{
			MethodName: "ConfirmReservationAutomatically",
			Handler:    _ReservationService_ConfirmReservationAutomatically_Handler,
		},
		{
			MethodName: "HasActiveReservations",
			Handler:    _ReservationService_HasActiveReservations_Handler,
		},
		{
			MethodName: "GetAllFuture",
			Handler:    _ReservationService_GetAllFuture_Handler,
		},
		{
			MethodName: "DeleteReservationRequestGuest",
			Handler:    _ReservationService_DeleteReservationRequestGuest_Handler,
		},
		{
			MethodName: "TermCheck",
			Handler:    _ReservationService_TermCheck_Handler,
		},
		{
			MethodName: "CheckReservationRequirementsHost",
			Handler:    _ReservationService_CheckReservationRequirementsHost_Handler,
		},
		{
			MethodName: "GetAccommodationsReservedInTimePeriod",
			Handler:    _ReservationService_GetAccommodationsReservedInTimePeriod_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "reservation_service.proto",
}
