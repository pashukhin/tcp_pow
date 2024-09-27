package client

import (
	"context"
	"io"
	"sync"

	"go.uber.org/zap"
)

type ClientRequester func(ctx context.Context, logger *zap.Logger, rw io.ReadWriter)

type Client interface {
	Run(ctx context.Context, wg *sync.WaitGroup, numRequests int)
}
