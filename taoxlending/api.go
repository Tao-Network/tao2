package taoxlending

import (
	"context"
	"errors"
	"sync"
	"time"
)


// List of errors
var (
	ErrOrderNonceTooLow  = errors.New("OrderNonce too low")
	ErrOrderNonceTooHigh = errors.New("OrderNonce too high")
)

// PublicTaoXLendingAPI provides the taoX RPC service that can be
// use publicly without security implications.
type PublicTaoXLendingAPI struct {
	t        *Lending
	mu       sync.Mutex
	lastUsed map[string]time.Time // keeps track when a filter was polled for the last time.

}

// NewPublicTaoXLendingAPI create a new RPC taoX service.
func NewPublicTaoXLendingAPI(t *Lending) *PublicTaoXLendingAPI {
	api := &PublicTaoXLendingAPI{
		t:        t,
		lastUsed: make(map[string]time.Time),
	}
	return api
}

// Version returns the Lending sub-protocol version.
func (api *PublicTaoXLendingAPI) Version(ctx context.Context) string {
	return ProtocolVersionStr
}
