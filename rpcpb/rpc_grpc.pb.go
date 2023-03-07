// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: rpcpb/rpc.proto

package rpcpb

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
	PingService_Ping_FullMethodName = "/rpcpb.PingService/Ping"
)

// PingServiceClient is the client API for PingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PingServiceClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
}

type pingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPingServiceClient(cc grpc.ClientConnInterface) PingServiceClient {
	return &pingServiceClient{cc}
}

func (c *pingServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := c.cc.Invoke(ctx, PingService_Ping_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PingServiceServer is the server API for PingService service.
// All implementations must embed UnimplementedPingServiceServer
// for forward compatibility
type PingServiceServer interface {
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	mustEmbedUnimplementedPingServiceServer()
}

// UnimplementedPingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPingServiceServer struct {
}

func (UnimplementedPingServiceServer) Ping(context.Context, *PingRequest) (*PingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (UnimplementedPingServiceServer) mustEmbedUnimplementedPingServiceServer() {}

// UnsafePingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingServiceServer will
// result in compilation errors.
type UnsafePingServiceServer interface {
	mustEmbedUnimplementedPingServiceServer()
}

func RegisterPingServiceServer(s grpc.ServiceRegistrar, srv PingServiceServer) {
	s.RegisterService(&PingService_ServiceDesc, srv)
}

func _PingService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PingService_Ping_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PingService_ServiceDesc is the grpc.ServiceDesc for PingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.PingService",
	HandlerType: (*PingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _PingService_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rpcpb/rpc.proto",
}

const (
	ControlService_RPCVersion_FullMethodName                = "/rpcpb.ControlService/RPCVersion"
	ControlService_Start_FullMethodName                     = "/rpcpb.ControlService/Start"
	ControlService_CreateBlockchains_FullMethodName         = "/rpcpb.ControlService/CreateBlockchains"
	ControlService_CreateSubnets_FullMethodName             = "/rpcpb.ControlService/CreateSubnets"
	ControlService_Health_FullMethodName                    = "/rpcpb.ControlService/Health"
	ControlService_URIs_FullMethodName                      = "/rpcpb.ControlService/URIs"
	ControlService_WaitForHealthy_FullMethodName            = "/rpcpb.ControlService/WaitForHealthy"
	ControlService_Status_FullMethodName                    = "/rpcpb.ControlService/Status"
	ControlService_StreamStatus_FullMethodName              = "/rpcpb.ControlService/StreamStatus"
	ControlService_RemoveNode_FullMethodName                = "/rpcpb.ControlService/RemoveNode"
	ControlService_AddNode_FullMethodName                   = "/rpcpb.ControlService/AddNode"
	ControlService_RestartNode_FullMethodName               = "/rpcpb.ControlService/RestartNode"
	ControlService_Stop_FullMethodName                      = "/rpcpb.ControlService/Stop"
	ControlService_AttachPeer_FullMethodName                = "/rpcpb.ControlService/AttachPeer"
	ControlService_SendOutboundMessage_FullMethodName       = "/rpcpb.ControlService/SendOutboundMessage"
	ControlService_SaveSnapshot_FullMethodName              = "/rpcpb.ControlService/SaveSnapshot"
	ControlService_LoadSnapshot_FullMethodName              = "/rpcpb.ControlService/LoadSnapshot"
	ControlService_RemoveSnapshot_FullMethodName            = "/rpcpb.ControlService/RemoveSnapshot"
	ControlService_GetSnapshotNames_FullMethodName          = "/rpcpb.ControlService/GetSnapshotNames"
	ControlService_CreateSpecificBlockchains_FullMethodName = "/rpcpb.ControlService/CreateSpecificBlockchains"
)

// ControlServiceClient is the client API for ControlService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ControlServiceClient interface {
	RPCVersion(ctx context.Context, in *RPCVersionRequest, opts ...grpc.CallOption) (*RPCVersionResponse, error)
	Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error)
	CreateBlockchains(ctx context.Context, in *CreateBlockchainsRequest, opts ...grpc.CallOption) (*CreateBlockchainsResponse, error)
	CreateSubnets(ctx context.Context, in *CreateSubnetsRequest, opts ...grpc.CallOption) (*CreateSubnetsResponse, error)
	Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error)
	URIs(ctx context.Context, in *URIsRequest, opts ...grpc.CallOption) (*URIsResponse, error)
	WaitForHealthy(ctx context.Context, in *WaitForHealthyRequest, opts ...grpc.CallOption) (*WaitForHealthyResponse, error)
	Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error)
	StreamStatus(ctx context.Context, in *StreamStatusRequest, opts ...grpc.CallOption) (ControlService_StreamStatusClient, error)
	RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...grpc.CallOption) (*RemoveNodeResponse, error)
	AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*AddNodeResponse, error)
	RestartNode(ctx context.Context, in *RestartNodeRequest, opts ...grpc.CallOption) (*RestartNodeResponse, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error)
	AttachPeer(ctx context.Context, in *AttachPeerRequest, opts ...grpc.CallOption) (*AttachPeerResponse, error)
	SendOutboundMessage(ctx context.Context, in *SendOutboundMessageRequest, opts ...grpc.CallOption) (*SendOutboundMessageResponse, error)
	SaveSnapshot(ctx context.Context, in *SaveSnapshotRequest, opts ...grpc.CallOption) (*SaveSnapshotResponse, error)
	LoadSnapshot(ctx context.Context, in *LoadSnapshotRequest, opts ...grpc.CallOption) (*LoadSnapshotResponse, error)
	RemoveSnapshot(ctx context.Context, in *RemoveSnapshotRequest, opts ...grpc.CallOption) (*RemoveSnapshotResponse, error)
	GetSnapshotNames(ctx context.Context, in *GetSnapshotNamesRequest, opts ...grpc.CallOption) (*GetSnapshotNamesResponse, error)
	CreateSpecificBlockchains(ctx context.Context, in *CreateSpecificBlockchainsRequest, opts ...grpc.CallOption) (*CreateSpecificBlockchainsResponse, error)
}

type controlServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewControlServiceClient(cc grpc.ClientConnInterface) ControlServiceClient {
	return &controlServiceClient{cc}
}

func (c *controlServiceClient) RPCVersion(ctx context.Context, in *RPCVersionRequest, opts ...grpc.CallOption) (*RPCVersionResponse, error) {
	out := new(RPCVersionResponse)
	err := c.cc.Invoke(ctx, ControlService_RPCVersion_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) Start(ctx context.Context, in *StartRequest, opts ...grpc.CallOption) (*StartResponse, error) {
	out := new(StartResponse)
	err := c.cc.Invoke(ctx, ControlService_Start_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) CreateBlockchains(ctx context.Context, in *CreateBlockchainsRequest, opts ...grpc.CallOption) (*CreateBlockchainsResponse, error) {
	out := new(CreateBlockchainsResponse)
	err := c.cc.Invoke(ctx, ControlService_CreateBlockchains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) CreateSubnets(ctx context.Context, in *CreateSubnetsRequest, opts ...grpc.CallOption) (*CreateSubnetsResponse, error) {
	out := new(CreateSubnetsResponse)
	err := c.cc.Invoke(ctx, ControlService_CreateSubnets_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) Health(ctx context.Context, in *HealthRequest, opts ...grpc.CallOption) (*HealthResponse, error) {
	out := new(HealthResponse)
	err := c.cc.Invoke(ctx, ControlService_Health_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) URIs(ctx context.Context, in *URIsRequest, opts ...grpc.CallOption) (*URIsResponse, error) {
	out := new(URIsResponse)
	err := c.cc.Invoke(ctx, ControlService_URIs_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) WaitForHealthy(ctx context.Context, in *WaitForHealthyRequest, opts ...grpc.CallOption) (*WaitForHealthyResponse, error) {
	out := new(WaitForHealthyResponse)
	err := c.cc.Invoke(ctx, ControlService_WaitForHealthy_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) Status(ctx context.Context, in *StatusRequest, opts ...grpc.CallOption) (*StatusResponse, error) {
	out := new(StatusResponse)
	err := c.cc.Invoke(ctx, ControlService_Status_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) StreamStatus(ctx context.Context, in *StreamStatusRequest, opts ...grpc.CallOption) (ControlService_StreamStatusClient, error) {
	stream, err := c.cc.NewStream(ctx, &ControlService_ServiceDesc.Streams[0], ControlService_StreamStatus_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &controlServiceStreamStatusClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ControlService_StreamStatusClient interface {
	Recv() (*StreamStatusResponse, error)
	grpc.ClientStream
}

type controlServiceStreamStatusClient struct {
	grpc.ClientStream
}

func (x *controlServiceStreamStatusClient) Recv() (*StreamStatusResponse, error) {
	m := new(StreamStatusResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *controlServiceClient) RemoveNode(ctx context.Context, in *RemoveNodeRequest, opts ...grpc.CallOption) (*RemoveNodeResponse, error) {
	out := new(RemoveNodeResponse)
	err := c.cc.Invoke(ctx, ControlService_RemoveNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) AddNode(ctx context.Context, in *AddNodeRequest, opts ...grpc.CallOption) (*AddNodeResponse, error) {
	out := new(AddNodeResponse)
	err := c.cc.Invoke(ctx, ControlService_AddNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) RestartNode(ctx context.Context, in *RestartNodeRequest, opts ...grpc.CallOption) (*RestartNodeResponse, error) {
	out := new(RestartNodeResponse)
	err := c.cc.Invoke(ctx, ControlService_RestartNode_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (*StopResponse, error) {
	out := new(StopResponse)
	err := c.cc.Invoke(ctx, ControlService_Stop_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) AttachPeer(ctx context.Context, in *AttachPeerRequest, opts ...grpc.CallOption) (*AttachPeerResponse, error) {
	out := new(AttachPeerResponse)
	err := c.cc.Invoke(ctx, ControlService_AttachPeer_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) SendOutboundMessage(ctx context.Context, in *SendOutboundMessageRequest, opts ...grpc.CallOption) (*SendOutboundMessageResponse, error) {
	out := new(SendOutboundMessageResponse)
	err := c.cc.Invoke(ctx, ControlService_SendOutboundMessage_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) SaveSnapshot(ctx context.Context, in *SaveSnapshotRequest, opts ...grpc.CallOption) (*SaveSnapshotResponse, error) {
	out := new(SaveSnapshotResponse)
	err := c.cc.Invoke(ctx, ControlService_SaveSnapshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) LoadSnapshot(ctx context.Context, in *LoadSnapshotRequest, opts ...grpc.CallOption) (*LoadSnapshotResponse, error) {
	out := new(LoadSnapshotResponse)
	err := c.cc.Invoke(ctx, ControlService_LoadSnapshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) RemoveSnapshot(ctx context.Context, in *RemoveSnapshotRequest, opts ...grpc.CallOption) (*RemoveSnapshotResponse, error) {
	out := new(RemoveSnapshotResponse)
	err := c.cc.Invoke(ctx, ControlService_RemoveSnapshot_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) GetSnapshotNames(ctx context.Context, in *GetSnapshotNamesRequest, opts ...grpc.CallOption) (*GetSnapshotNamesResponse, error) {
	out := new(GetSnapshotNamesResponse)
	err := c.cc.Invoke(ctx, ControlService_GetSnapshotNames_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *controlServiceClient) CreateSpecificBlockchains(ctx context.Context, in *CreateSpecificBlockchainsRequest, opts ...grpc.CallOption) (*CreateSpecificBlockchainsResponse, error) {
	out := new(CreateSpecificBlockchainsResponse)
	err := c.cc.Invoke(ctx, ControlService_CreateSpecificBlockchains_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ControlServiceServer is the server API for ControlService service.
// All implementations must embed UnimplementedControlServiceServer
// for forward compatibility
type ControlServiceServer interface {
	RPCVersion(context.Context, *RPCVersionRequest) (*RPCVersionResponse, error)
	Start(context.Context, *StartRequest) (*StartResponse, error)
	CreateBlockchains(context.Context, *CreateBlockchainsRequest) (*CreateBlockchainsResponse, error)
	CreateSubnets(context.Context, *CreateSubnetsRequest) (*CreateSubnetsResponse, error)
	Health(context.Context, *HealthRequest) (*HealthResponse, error)
	URIs(context.Context, *URIsRequest) (*URIsResponse, error)
	WaitForHealthy(context.Context, *WaitForHealthyRequest) (*WaitForHealthyResponse, error)
	Status(context.Context, *StatusRequest) (*StatusResponse, error)
	StreamStatus(*StreamStatusRequest, ControlService_StreamStatusServer) error
	RemoveNode(context.Context, *RemoveNodeRequest) (*RemoveNodeResponse, error)
	AddNode(context.Context, *AddNodeRequest) (*AddNodeResponse, error)
	RestartNode(context.Context, *RestartNodeRequest) (*RestartNodeResponse, error)
	Stop(context.Context, *StopRequest) (*StopResponse, error)
	AttachPeer(context.Context, *AttachPeerRequest) (*AttachPeerResponse, error)
	SendOutboundMessage(context.Context, *SendOutboundMessageRequest) (*SendOutboundMessageResponse, error)
	SaveSnapshot(context.Context, *SaveSnapshotRequest) (*SaveSnapshotResponse, error)
	LoadSnapshot(context.Context, *LoadSnapshotRequest) (*LoadSnapshotResponse, error)
	RemoveSnapshot(context.Context, *RemoveSnapshotRequest) (*RemoveSnapshotResponse, error)
	GetSnapshotNames(context.Context, *GetSnapshotNamesRequest) (*GetSnapshotNamesResponse, error)
	CreateSpecificBlockchains(context.Context, *CreateSpecificBlockchainsRequest) (*CreateSpecificBlockchainsResponse, error)
	mustEmbedUnimplementedControlServiceServer()
}

// UnimplementedControlServiceServer must be embedded to have forward compatible implementations.
type UnimplementedControlServiceServer struct {
}

func (UnimplementedControlServiceServer) RPCVersion(context.Context, *RPCVersionRequest) (*RPCVersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RPCVersion not implemented")
}
func (UnimplementedControlServiceServer) Start(context.Context, *StartRequest) (*StartResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Start not implemented")
}
func (UnimplementedControlServiceServer) CreateBlockchains(context.Context, *CreateBlockchainsRequest) (*CreateBlockchainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBlockchains not implemented")
}
func (UnimplementedControlServiceServer) CreateSubnets(context.Context, *CreateSubnetsRequest) (*CreateSubnetsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubnets not implemented")
}
func (UnimplementedControlServiceServer) Health(context.Context, *HealthRequest) (*HealthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Health not implemented")
}
func (UnimplementedControlServiceServer) URIs(context.Context, *URIsRequest) (*URIsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method URIs not implemented")
}
func (UnimplementedControlServiceServer) WaitForHealthy(context.Context, *WaitForHealthyRequest) (*WaitForHealthyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WaitForHealthy not implemented")
}
func (UnimplementedControlServiceServer) Status(context.Context, *StatusRequest) (*StatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedControlServiceServer) StreamStatus(*StreamStatusRequest, ControlService_StreamStatusServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamStatus not implemented")
}
func (UnimplementedControlServiceServer) RemoveNode(context.Context, *RemoveNodeRequest) (*RemoveNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveNode not implemented")
}
func (UnimplementedControlServiceServer) AddNode(context.Context, *AddNodeRequest) (*AddNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddNode not implemented")
}
func (UnimplementedControlServiceServer) RestartNode(context.Context, *RestartNodeRequest) (*RestartNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RestartNode not implemented")
}
func (UnimplementedControlServiceServer) Stop(context.Context, *StopRequest) (*StopResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedControlServiceServer) AttachPeer(context.Context, *AttachPeerRequest) (*AttachPeerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AttachPeer not implemented")
}
func (UnimplementedControlServiceServer) SendOutboundMessage(context.Context, *SendOutboundMessageRequest) (*SendOutboundMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendOutboundMessage not implemented")
}
func (UnimplementedControlServiceServer) SaveSnapshot(context.Context, *SaveSnapshotRequest) (*SaveSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveSnapshot not implemented")
}
func (UnimplementedControlServiceServer) LoadSnapshot(context.Context, *LoadSnapshotRequest) (*LoadSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LoadSnapshot not implemented")
}
func (UnimplementedControlServiceServer) RemoveSnapshot(context.Context, *RemoveSnapshotRequest) (*RemoveSnapshotResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSnapshot not implemented")
}
func (UnimplementedControlServiceServer) GetSnapshotNames(context.Context, *GetSnapshotNamesRequest) (*GetSnapshotNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSnapshotNames not implemented")
}
func (UnimplementedControlServiceServer) CreateSpecificBlockchains(context.Context, *CreateSpecificBlockchainsRequest) (*CreateSpecificBlockchainsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSpecificBlockchains not implemented")
}
func (UnimplementedControlServiceServer) mustEmbedUnimplementedControlServiceServer() {}

// UnsafeControlServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ControlServiceServer will
// result in compilation errors.
type UnsafeControlServiceServer interface {
	mustEmbedUnimplementedControlServiceServer()
}

func RegisterControlServiceServer(s grpc.ServiceRegistrar, srv ControlServiceServer) {
	s.RegisterService(&ControlService_ServiceDesc, srv)
}

func _ControlService_RPCVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RPCVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).RPCVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_RPCVersion_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).RPCVersion(ctx, req.(*RPCVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_Start_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Start(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_Start_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Start(ctx, req.(*StartRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_CreateBlockchains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBlockchainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).CreateBlockchains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_CreateBlockchains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).CreateBlockchains(ctx, req.(*CreateBlockchainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_CreateSubnets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubnetsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).CreateSubnets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_CreateSubnets_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).CreateSubnets(ctx, req.(*CreateSubnetsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_Health_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Health(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_Health_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Health(ctx, req.(*HealthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_URIs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(URIsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).URIs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_URIs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).URIs(ctx, req.(*URIsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_WaitForHealthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WaitForHealthyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).WaitForHealthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_WaitForHealthy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).WaitForHealthy(ctx, req.(*WaitForHealthyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_Status_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Status(ctx, req.(*StatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_StreamStatus_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StreamStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ControlServiceServer).StreamStatus(m, &controlServiceStreamStatusServer{stream})
}

type ControlService_StreamStatusServer interface {
	Send(*StreamStatusResponse) error
	grpc.ServerStream
}

type controlServiceStreamStatusServer struct {
	grpc.ServerStream
}

func (x *controlServiceStreamStatusServer) Send(m *StreamStatusResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ControlService_RemoveNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).RemoveNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_RemoveNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).RemoveNode(ctx, req.(*RemoveNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_AddNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).AddNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_AddNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).AddNode(ctx, req.(*AddNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_RestartNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RestartNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).RestartNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_RestartNode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).RestartNode(ctx, req.(*RestartNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_Stop_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).Stop(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_Stop_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).Stop(ctx, req.(*StopRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_AttachPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttachPeerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).AttachPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_AttachPeer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).AttachPeer(ctx, req.(*AttachPeerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_SendOutboundMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendOutboundMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).SendOutboundMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_SendOutboundMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).SendOutboundMessage(ctx, req.(*SendOutboundMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_SaveSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).SaveSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_SaveSnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).SaveSnapshot(ctx, req.(*SaveSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_LoadSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoadSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).LoadSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_LoadSnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).LoadSnapshot(ctx, req.(*LoadSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_RemoveSnapshot_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveSnapshotRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).RemoveSnapshot(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_RemoveSnapshot_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).RemoveSnapshot(ctx, req.(*RemoveSnapshotRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_GetSnapshotNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSnapshotNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).GetSnapshotNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_GetSnapshotNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).GetSnapshotNames(ctx, req.(*GetSnapshotNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ControlService_CreateSpecificBlockchains_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSpecificBlockchainsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ControlServiceServer).CreateSpecificBlockchains(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ControlService_CreateSpecificBlockchains_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ControlServiceServer).CreateSpecificBlockchains(ctx, req.(*CreateSpecificBlockchainsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ControlService_ServiceDesc is the grpc.ServiceDesc for ControlService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ControlService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rpcpb.ControlService",
	HandlerType: (*ControlServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RPCVersion",
			Handler:    _ControlService_RPCVersion_Handler,
		},
		{
			MethodName: "Start",
			Handler:    _ControlService_Start_Handler,
		},
		{
			MethodName: "CreateBlockchains",
			Handler:    _ControlService_CreateBlockchains_Handler,
		},
		{
			MethodName: "CreateSubnets",
			Handler:    _ControlService_CreateSubnets_Handler,
		},
		{
			MethodName: "Health",
			Handler:    _ControlService_Health_Handler,
		},
		{
			MethodName: "URIs",
			Handler:    _ControlService_URIs_Handler,
		},
		{
			MethodName: "WaitForHealthy",
			Handler:    _ControlService_WaitForHealthy_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _ControlService_Status_Handler,
		},
		{
			MethodName: "RemoveNode",
			Handler:    _ControlService_RemoveNode_Handler,
		},
		{
			MethodName: "AddNode",
			Handler:    _ControlService_AddNode_Handler,
		},
		{
			MethodName: "RestartNode",
			Handler:    _ControlService_RestartNode_Handler,
		},
		{
			MethodName: "Stop",
			Handler:    _ControlService_Stop_Handler,
		},
		{
			MethodName: "AttachPeer",
			Handler:    _ControlService_AttachPeer_Handler,
		},
		{
			MethodName: "SendOutboundMessage",
			Handler:    _ControlService_SendOutboundMessage_Handler,
		},
		{
			MethodName: "SaveSnapshot",
			Handler:    _ControlService_SaveSnapshot_Handler,
		},
		{
			MethodName: "LoadSnapshot",
			Handler:    _ControlService_LoadSnapshot_Handler,
		},
		{
			MethodName: "RemoveSnapshot",
			Handler:    _ControlService_RemoveSnapshot_Handler,
		},
		{
			MethodName: "GetSnapshotNames",
			Handler:    _ControlService_GetSnapshotNames_Handler,
		},
		{
			MethodName: "CreateSpecificBlockchains",
			Handler:    _ControlService_CreateSpecificBlockchains_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamStatus",
			Handler:       _ControlService_StreamStatus_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "rpcpb/rpc.proto",
}
