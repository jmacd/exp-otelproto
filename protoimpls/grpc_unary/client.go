package grpc_unary

import (
	"context"
	"log"
	"sync/atomic"

	otlp "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/metrics/v1"
	"github.com/tigrannajaryan/exp-otelproto/core"
	"google.golang.org/grpc"
)

// Client can connect to a server and send a batch of spans.
type Client struct {
	client otlp.MetricsServiceClient
	nextId uint64
}

func (c *Client) Connect(server string) error {
	// Set up a connection to the server.
	conn, err := grpc.Dial(server, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c.client = otlp.NewMetricsServiceClient(conn)
	return nil
}

func (c *Client) Export(batch core.ExportRequest) {
	request := batch.(*otlp.ExportMetricsServiceRequest)
	atomic.AddUint64(&c.nextId, 1)
	c.client.Export(context.Background(), request)
}

func (c *Client) Shutdown() {
}
