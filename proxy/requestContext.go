package proxy

import (
	"errors"

	"github.com/elazarl/goproxy"
)

// DispatchContext contains additional data to be stored in goproxy.ProxyCtx.UserData to
// provide request context to the response handler.
type DispatchContext struct {
	// Was the request blocked by the proxy
	RequestIsBlocked bool
	// Decoded request body
	RequestBody []byte
	RequestOp   string
	// An error that occurred during request unmarshalling, if any
	RequestUnmarshalErr error
	// Contains the unmarshalled request packet. Nil if DispatchContext is for a request.
	RequestPkt interface{}
	// An error that occurred during response unmarshalling, if any
	ResponseUnmarshalErr error
}

// GetDispatchContext returns the dispatch context for a goproxy.ProxyCtx, will panic if
// called on a goproxy.ProxyCtx not associated with game data. I.e., the request was not
// handled by proxy.HandleReq.
func GetDispatchContext(ctx *goproxy.ProxyCtx) *DispatchContext {
	reqCtx, ok := ctx.UserData.(*DispatchContext)
	if reqCtx == nil || !ok {
		panic(errors.New("failed to cast ctx.UserData into *RequestContext"))
	}
	return reqCtx
}
