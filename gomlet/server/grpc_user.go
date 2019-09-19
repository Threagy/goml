package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/goml/gomlet/db"
	pbUser "github.com/goml/gomlet/proto/v1/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// UserServer ...
type UserServer struct {
	GrpcServer *grpc.Server
}

// NewUserServer ...
func NewUserServer() *UserServer {
	// create server
	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(ensureValidToken))
	server := &UserServer{GrpcServer: grpcServer}
	pbUser.RegisterUserManagerServer(grpcServer, server)

	return server
}

// GetUsers ...
func (srv *UserServer) GetUsers(ctx context.Context, req *pbUser.GetUsersRequest) (*pbUser.GetUsersResponse, error) {
	users, err := db.GetUsers()
	if err != nil {
		return nil, err
	}

	convertedUsers := convertToUser(users)
	response := &pbUser.GetUsersResponse{
		Users: convertedUsers,
	}

	return response, nil
}

// ModifyUser ...
func (srv *UserServer) ModifyUser(ctx context.Context, req *pbUser.ModifyUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ModifyUser not implemented")
}

// DeleteUser ...
func (srv *UserServer) DeleteUser(ctx context.Context, req *pbUser.DeleteUserRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
