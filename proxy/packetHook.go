package proxy

import (
	"github.com/elazarl/goproxy"
)

// PacketHook contains information about the hook and allows for execution of the underlying
// Handle and Shutdown methods in the PacketHandler.
type PacketHook struct {
	name     string
	target   string
	priority int
	raw      bool
	handler  PacketHandler
}

// Handle calls the Handle method of the underlying PacketHandler.
func (hook *PacketHook) Handle(kind string, pkt interface{}, user *UserCtx, pktCtx *goproxy.ProxyCtx) (interface{}, error) {
	return hook.handler(kind, pkt, user, pktCtx)
}

// PacketHandler represents handler functions exposed by a module.
type PacketHandler func(kind string, pkt interface{}, user *UserCtx, pktCtx *goproxy.ProxyCtx) (interface{}, error)

// HookGenerator is run once a user's context has been established, after the authentication packet.
type HookGenerator func() *PacketHook

// NewPacketHook returns an initialized PacketHook struct.
func NewPacketHook(name, target string, priority int, raw bool, handler PacketHandler) *PacketHook {
	return &PacketHook{name, target, priority, raw, handler}
}
