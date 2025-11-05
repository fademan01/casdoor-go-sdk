package client

import (
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
