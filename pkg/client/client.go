package client

import (
	"context"
	"fmt"
	"log"

	core "github.com/flux-framework/flux-go/pkg/flux"
	pb "github.com/flux-framework/flux-go/pkg/flux-grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

// FluxionClient interacts with Fluxion
type FluxClient struct {
	ctx        context.Context
	host       string
	connection *grpc.ClientConn
	service    pb.FluxServiceClient
}

var _ Client = (*FluxClient)(nil)

type Client interface {
	Submit(ctx context.Context, jobspec *core.JobSpec) (*pb.SubmitResponse, error)
	Jobs(ctx context.Context) (*pb.JobsResponse, error)
	Close() error
	GetHost() string
	Connected() bool
}

// NewClient creates a new FluxionClient
func NewClient(host string) (Client, error) {
	if host == "" {
		return nil, fmt.Errorf("host is required")
	}

	// We have to have a connection to a flux handle
	log.Printf("🦩️ starting client (%s)...", host)
	c := &FluxClient{host: host, ctx: context.Background()}

	// Set up a connection to the server.
	conn, err := grpc.NewClient(
		c.GetHost(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to %s: %s", host, err)
	}
	c.connection = conn
	c.service = pb.NewFluxServiceClient(conn)
	return c, nil
}

// Close closes the created resources (e.g. connection).
func (c *FluxClient) Close() error {
	if c.connection != nil {
		return c.connection.Close()
	}
	return nil
}

// Connected returns  true if we are connected and the connection is ready
func (c *FluxClient) Connected() bool {
	return c.service != nil && c.connection != nil && c.connection.GetState() == connectivity.Ready
}

// GetHost returns the private hostn name
func (c *FluxClient) GetHost() string {
	return c.host
}
