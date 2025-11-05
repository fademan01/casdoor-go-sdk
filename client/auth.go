package client

import (
	"context"

	"github.com/casdoor/casdoor-go-sdk/proto"
)

func (c *GRPCClient) GetOAuthToken(ctx context.Context, code, state string) (*proto.Token, error) {
	req := &proto.GetOAuthTokenRequest{
		Code:  code,
		State: state,
	}
	return c.AuthClient.GetOAuthToken(ctx, req)
}

func (c *GRPCClient) RefreshOAuthToken(ctx context.Context, refreshToken string) (*proto.Token, error) {
	req := &proto.RefreshOAuthTokenRequest{
		RefreshToken: refreshToken,
	}
	return c.AuthClient.RefreshOAuthToken(ctx, req)
}
