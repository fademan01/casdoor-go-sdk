package client

import (
	"context"

	pb "github.com/casdoor/casdoor-go-sdk/proto"
)

func (c *GRPCClient) GetUser(ctx context.Context, user_name string, organization_name string) (*pb.UserResponse, error) {
	req := &pb.GetUserRequest{
		UserName:  user_name,
		UserOwner: organization_name,
	}

	resp, err := c.UserClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
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
