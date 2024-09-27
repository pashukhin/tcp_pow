package client

import (
	"context"
	"net"
	"sync"

	"go.uber.org/zap"
)

type client struct {
	serverAddr string
	logger     *zap.Logger
	requester  ClientRequester
}

func New(logger *zap.Logger, serverAddr string, requester ClientRequester) Client {
	return &client{
		serverAddr: serverAddr,
		logger:     logger,
		requester:  requester,
	}
}

func (c *client) request(ctx context.Context) {
	conn, err := net.Dial("tcp", c.serverAddr)
	if err != nil {
		c.logger.Error("Error connecting to server:", zap.Error(err))
		return
	}
	defer conn.Close()

	c.requester(ctx, c.logger, conn)
}

func (c *client) Run(ctx context.Context, wg *sync.WaitGroup, numRequests int) {
	defer wg.Done()

	for i := 0; i < numRequests; i++ {
		select {
		case <-ctx.Done():
			c.logger.Info("Client received shutdown signal, stopping...")
			return
		default:
			c.request(ctx)
		}
	}
}
