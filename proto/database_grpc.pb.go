// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v6.30.0--rc1
// source: database.proto

package proto

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
	DatabaseService_CreateUser_FullMethodName     = "/proto.DatabaseService/CreateUser"
	DatabaseService_CreateDatabase_FullMethodName = "/proto.DatabaseService/CreateDatabase"
	DatabaseService_CreateTable_FullMethodName    = "/proto.DatabaseService/CreateTable"
	DatabaseService_InsertRecord_FullMethodName   = "/proto.DatabaseService/InsertRecord"
	DatabaseService_QueryData_FullMethodName      = "/proto.DatabaseService/QueryData"
	DatabaseService_UpdateRecord_FullMethodName   = "/proto.DatabaseService/UpdateRecord"
	DatabaseService_DeleteRecord_FullMethodName   = "/proto.DatabaseService/DeleteRecord"
	DatabaseService_UpdateTable_FullMethodName    = "/proto.DatabaseService/UpdateTable"
	DatabaseService_AddIndex_FullMethodName       = "/proto.DatabaseService/AddIndex"
	DatabaseService_DeleteIndex_FullMethodName    = "/proto.DatabaseService/DeleteIndex"
	DatabaseService_ListIndexes_FullMethodName    = "/proto.DatabaseService/ListIndexes"
)

// DatabaseServiceClient is the client API for DatabaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DatabaseServiceClient interface {
	CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error)
	CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*CreateDatabaseResponse, error)
	CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error)
	InsertRecord(ctx context.Context, in *InsertRecordRequest, opts ...grpc.CallOption) (*InsertRecordResponse, error)
	QueryData(ctx context.Context, in *QueryDataRequest, opts ...grpc.CallOption) (*QueryDataResponse, error)
	UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordResponse, error)
	DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error)
	UpdateTable(ctx context.Context, in *UpdateTableRequest, opts ...grpc.CallOption) (*UpdateTableResponse, error)
	AddIndex(ctx context.Context, in *AddIndexRequest, opts ...grpc.CallOption) (*AddIndexResponse, error)
	DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error)
	ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error)
}

type databaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseServiceClient(cc grpc.ClientConnInterface) DatabaseServiceClient {
	return &databaseServiceClient{cc}
}

func (c *databaseServiceClient) CreateUser(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*CreateUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateUserResponse)
	err := c.cc.Invoke(ctx, DatabaseService_CreateUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) CreateDatabase(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*CreateDatabaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDatabaseResponse)
	err := c.cc.Invoke(ctx, DatabaseService_CreateDatabase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) CreateTable(ctx context.Context, in *CreateTableRequest, opts ...grpc.CallOption) (*CreateTableResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateTableResponse)
	err := c.cc.Invoke(ctx, DatabaseService_CreateTable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) InsertRecord(ctx context.Context, in *InsertRecordRequest, opts ...grpc.CallOption) (*InsertRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(InsertRecordResponse)
	err := c.cc.Invoke(ctx, DatabaseService_InsertRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) QueryData(ctx context.Context, in *QueryDataRequest, opts ...grpc.CallOption) (*QueryDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(QueryDataResponse)
	err := c.cc.Invoke(ctx, DatabaseService_QueryData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) UpdateRecord(ctx context.Context, in *UpdateRecordRequest, opts ...grpc.CallOption) (*UpdateRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateRecordResponse)
	err := c.cc.Invoke(ctx, DatabaseService_UpdateRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) DeleteRecord(ctx context.Context, in *DeleteRecordRequest, opts ...grpc.CallOption) (*DeleteRecordResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteRecordResponse)
	err := c.cc.Invoke(ctx, DatabaseService_DeleteRecord_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) UpdateTable(ctx context.Context, in *UpdateTableRequest, opts ...grpc.CallOption) (*UpdateTableResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateTableResponse)
	err := c.cc.Invoke(ctx, DatabaseService_UpdateTable_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) AddIndex(ctx context.Context, in *AddIndexRequest, opts ...grpc.CallOption) (*AddIndexResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AddIndexResponse)
	err := c.cc.Invoke(ctx, DatabaseService_AddIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) DeleteIndex(ctx context.Context, in *DeleteIndexRequest, opts ...grpc.CallOption) (*DeleteIndexResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteIndexResponse)
	err := c.cc.Invoke(ctx, DatabaseService_DeleteIndex_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseServiceClient) ListIndexes(ctx context.Context, in *ListIndexesRequest, opts ...grpc.CallOption) (*ListIndexesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListIndexesResponse)
	err := c.cc.Invoke(ctx, DatabaseService_ListIndexes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServiceServer is the server API for DatabaseService service.
// All implementations must embed UnimplementedDatabaseServiceServer
// for forward compatibility.
type DatabaseServiceServer interface {
	CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error)
	CreateDatabase(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error)
	CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error)
	InsertRecord(context.Context, *InsertRecordRequest) (*InsertRecordResponse, error)
	QueryData(context.Context, *QueryDataRequest) (*QueryDataResponse, error)
	UpdateRecord(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error)
	DeleteRecord(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error)
	UpdateTable(context.Context, *UpdateTableRequest) (*UpdateTableResponse, error)
	AddIndex(context.Context, *AddIndexRequest) (*AddIndexResponse, error)
	DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error)
	ListIndexes(context.Context, *ListIndexesRequest) (*ListIndexesResponse, error)
	mustEmbedUnimplementedDatabaseServiceServer()
}

// UnimplementedDatabaseServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDatabaseServiceServer struct{}

func (UnimplementedDatabaseServiceServer) CreateUser(context.Context, *CreateUserRequest) (*CreateUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedDatabaseServiceServer) CreateDatabase(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabase not implemented")
}
func (UnimplementedDatabaseServiceServer) CreateTable(context.Context, *CreateTableRequest) (*CreateTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateTable not implemented")
}
func (UnimplementedDatabaseServiceServer) InsertRecord(context.Context, *InsertRecordRequest) (*InsertRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InsertRecord not implemented")
}
func (UnimplementedDatabaseServiceServer) QueryData(context.Context, *QueryDataRequest) (*QueryDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryData not implemented")
}
func (UnimplementedDatabaseServiceServer) UpdateRecord(context.Context, *UpdateRecordRequest) (*UpdateRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRecord not implemented")
}
func (UnimplementedDatabaseServiceServer) DeleteRecord(context.Context, *DeleteRecordRequest) (*DeleteRecordResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteRecord not implemented")
}
func (UnimplementedDatabaseServiceServer) UpdateTable(context.Context, *UpdateTableRequest) (*UpdateTableResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateTable not implemented")
}
func (UnimplementedDatabaseServiceServer) AddIndex(context.Context, *AddIndexRequest) (*AddIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddIndex not implemented")
}
func (UnimplementedDatabaseServiceServer) DeleteIndex(context.Context, *DeleteIndexRequest) (*DeleteIndexResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteIndex not implemented")
}
func (UnimplementedDatabaseServiceServer) ListIndexes(context.Context, *ListIndexesRequest) (*ListIndexesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListIndexes not implemented")
}
func (UnimplementedDatabaseServiceServer) mustEmbedUnimplementedDatabaseServiceServer() {}
func (UnimplementedDatabaseServiceServer) testEmbeddedByValue()                         {}

// UnsafeDatabaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServiceServer will
// result in compilation errors.
type UnsafeDatabaseServiceServer interface {
	mustEmbedUnimplementedDatabaseServiceServer()
}

func RegisterDatabaseServiceServer(s grpc.ServiceRegistrar, srv DatabaseServiceServer) {
	// If the following call pancis, it indicates UnimplementedDatabaseServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&DatabaseService_ServiceDesc, srv)
}

func _DatabaseService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_CreateUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).CreateUser(ctx, req.(*CreateUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_CreateDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).CreateDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_CreateDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).CreateDatabase(ctx, req.(*CreateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_CreateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).CreateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_CreateTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).CreateTable(ctx, req.(*CreateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_InsertRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InsertRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).InsertRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_InsertRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).InsertRecord(ctx, req.(*InsertRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_QueryData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).QueryData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_QueryData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).QueryData(ctx, req.(*QueryDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_UpdateRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).UpdateRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_UpdateRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).UpdateRecord(ctx, req.(*UpdateRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_DeleteRecord_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRecordRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).DeleteRecord(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_DeleteRecord_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).DeleteRecord(ctx, req.(*DeleteRecordRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_UpdateTable_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateTableRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).UpdateTable(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_UpdateTable_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).UpdateTable(ctx, req.(*UpdateTableRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_AddIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).AddIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_AddIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).AddIndex(ctx, req.(*AddIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_DeleteIndex_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteIndexRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).DeleteIndex(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_DeleteIndex_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).DeleteIndex(ctx, req.(*DeleteIndexRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DatabaseService_ListIndexes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListIndexesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServiceServer).ListIndexes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: DatabaseService_ListIndexes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServiceServer).ListIndexes(ctx, req.(*ListIndexesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DatabaseService_ServiceDesc is the grpc.ServiceDesc for DatabaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DatabaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.DatabaseService",
	HandlerType: (*DatabaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _DatabaseService_CreateUser_Handler,
		},
		{
			MethodName: "CreateDatabase",
			Handler:    _DatabaseService_CreateDatabase_Handler,
		},
		{
			MethodName: "CreateTable",
			Handler:    _DatabaseService_CreateTable_Handler,
		},
		{
			MethodName: "InsertRecord",
			Handler:    _DatabaseService_InsertRecord_Handler,
		},
		{
			MethodName: "QueryData",
			Handler:    _DatabaseService_QueryData_Handler,
		},
		{
			MethodName: "UpdateRecord",
			Handler:    _DatabaseService_UpdateRecord_Handler,
		},
		{
			MethodName: "DeleteRecord",
			Handler:    _DatabaseService_DeleteRecord_Handler,
		},
		{
			MethodName: "UpdateTable",
			Handler:    _DatabaseService_UpdateTable_Handler,
		},
		{
			MethodName: "AddIndex",
			Handler:    _DatabaseService_AddIndex_Handler,
		},
		{
			MethodName: "DeleteIndex",
			Handler:    _DatabaseService_DeleteIndex_Handler,
		},
		{
			MethodName: "ListIndexes",
			Handler:    _DatabaseService_ListIndexes_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "database.proto",
}
