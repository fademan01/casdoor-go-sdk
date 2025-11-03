package client

import (
	"context"
	"crypto/tls"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/casdoor/casdoor-go-sdk/proto" // replace "your-project" with your module name
)

// GRPCClient wraps both auth and user gRPC clients
type GRPCClient struct {
	Conn       *grpc.ClientConn
	AuthClient pb.AuthServiceClient
	UserClient pb.UserServiceClient
}

// NewGRPCClient creates a new gRPC client connected to the given address.
// If useTLS is true, it uses TLS; otherwise, insecure connection.
func NewGRPCClient(addr string, useTLS bool, opts ...grpc.DialOption) (*GRPCClient, error) {
	var grpcOpts []grpc.DialOption

	if useTLS {
		// Use system default TLS config
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})))
	} else {
		grpcOpts = append(grpcOpts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// Add custom options (e.g., timeouts, interceptors)
	grpcOpts = append(grpcOpts, opts...)

	// Set default timeout if not provided
	if len(opts) == 0 {
		grpcOpts = append(grpcOpts)
	}

	conn, err := grpc.NewClient(addr, grpcOpts...)
	if err != nil {
		return nil, err
	}

	return &GRPCClient{
		Conn:       conn,
		AuthClient: pb.NewAuthServiceClient(conn),
		UserClient: pb.NewUserServiceClient(conn),
	}, nil
}

// Close closes the gRPC connection
func (c *GRPCClient) Close() error {
	return c.Conn.Close()
}

// --- Auth Methods ---

func (c *GRPCClient) GetOAuthToken(ctx context.Context, code, state string) (*pb.Token, error) {
	req := &pb.GetOAuthTokenRequest{
		Code:  code,
		State: state,
	}
	return c.AuthClient.GetOAuthToken(ctx, req)
}

func (c *GRPCClient) RefreshOAuthToken(ctx context.Context, refreshToken string) (*pb.Token, error) {
	req := &pb.RefreshOAuthTokenRequest{
		RefreshToken: refreshToken,
	}
	return c.AuthClient.RefreshOAuthToken(ctx, req)
}

// --- User Methods ---

func (c *GRPCClient) GetUser(ctx context.Context, id string) (*pb.User, error) {
	req := &pb.GetUserRequest{Id: id}
	resp, err := c.UserClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

func (c *GRPCClient) GetUserByEmail(ctx context.Context, owner, email string) (*pb.User, error) {
	req := &pb.GetUserByEmailRequest{
		Owner: owner,
		Email: email,
	}
	resp, err := c.UserClient.GetUserByEmail(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

func (c *GRPCClient) GetUserByPhone(ctx context.Context, owner, phone string) (*pb.User, error) {
	req := &pb.GetUserByPhoneRequest{
		Owner: owner,
		Phone: phone,
	}
	resp, err := c.UserClient.GetUserByPhone(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

func (c *GRPCClient) GetUserByUserId(ctx context.Context, owner, userId string) (*pb.User, error) {
	req := &pb.GetUserByUserIdRequest{
		Owner:  owner,
		UserId: userId,
	}
	resp, err := c.UserClient.GetUserByUserId(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

func (c *GRPCClient) GetUsers(ctx context.Context, owner string) ([]*pb.User, error) {
	req := &pb.GetUsersRequest{Owner: owner}
	resp, err := c.UserClient.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}

func (c *GRPCClient) GetSortedUsers(ctx context.Context, owner, sorter string, limit int32) ([]*pb.User, error) {
	req := &pb.GetSortedUsersRequest{
		Owner:  owner,
		Sorter: sorter,
		Limit:  limit,
	}
	resp, err := c.UserClient.GetSortedUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp.Users, nil
}

func (c *GRPCClient) GetPaginationUsers(
	ctx context.Context,
	owner string,
	page, pageSize int32,
	query map[string]string,
) ([]*pb.User, int32, error) {
	req := &pb.GetPaginationUsersRequest{
		Owner:    owner,
		Page:     page,
		PageSize: pageSize,
		Query:    query,
	}
	resp, err := c.UserClient.GetPaginationUsers(ctx, req)
	if err != nil {
		return nil, 0, err
	}
	return resp.Users, resp.TotalCount, nil
}

func (c *GRPCClient) GetUserCount(ctx context.Context, owner, isOnline string) (int32, error) {
	req := &pb.GetUserCountRequest{
		Owner:    owner,
		IsOnline: isOnline,
	}
	resp, err := c.UserClient.GetUserCount(ctx, req)
	if err != nil {
		return 0, err
	}
	return resp.Count, nil
}

func (c *GRPCClient) AddUser(ctx context.Context, user *pb.User) (bool, error) {
	req := &pb.AddUserRequest{User: user}
	resp, err := c.UserClient.AddUser(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}

func (c *GRPCClient) UpdateUser(ctx context.Context, user *pb.User, columns []string) (bool, error) {
	req := &pb.UpdateUserRequest{
		User:    user,
		Columns: columns,
	}
	resp, err := c.UserClient.UpdateUser(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}

func (c *GRPCClient) DeleteUser(ctx context.Context, user *pb.User) (bool, error) {
	req := &pb.DeleteUserRequest{User: user}
	resp, err := c.UserClient.DeleteUser(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}

func (c *GRPCClient) SetPassword(
	ctx context.Context,
	userOwner, userName, oldPassword, newPassword string,
) (bool, error) {
	req := &pb.SetPasswordRequest{
		UserOwner:   userOwner,
		UserName:    userName,
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}
	resp, err := c.UserClient.SetPassword(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}

func (c *GRPCClient) CheckUserPassword(ctx context.Context, user *pb.User) (bool, error) {
	req := &pb.CheckUserPasswordRequest{User: user}
	resp, err := c.UserClient.CheckUserPassword(ctx, req)
	if err != nil {
		return false, err
	}
	return resp.Ok, nil
}
