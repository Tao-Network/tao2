package taox

import (
	"context"
	"errors"
	"sync"
	"time"
)

const (
	LimitThresholdOrderNonceInQueue = 100
)

// List of errors
var (
	ErrNoTopics          = errors.New("missing topic(s)")
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicTaoXAPI provides the tomoX RPC service that can be
// use publicly without security implications.
type PublicTaoXAPI struct {
	t        *TaoX
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicTaoXAPI create a new RPC tomoX service.
func NewPublicTaoXAPI(t *TaoX) *PublicTaoXAPI {
	api := &PublicTaoXAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the TaoX sub-protocol version.
func (api *PublicTaoXAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
