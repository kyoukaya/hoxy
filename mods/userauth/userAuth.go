package userauth

import (
	"fmt"
	"hoxy/proxy"
	"hoxy/proxy/defs"
	"strconv"

	"github.com/elazarl/goproxy"
)

type userAuthState struct{}

var (
	incompletePktErr = fmt.Errorf("Malformed or incomplete packet")
)

func Handle(kind string, pkt interface{}, pktCtx *goproxy.ProxyCtx, prxy *proxy.HoxyProxy) (interface{}, error) {
	if kind != "SIndex/getUidEnMicaQueue" {
		panic("Wrong kind passed to handleMicaQueue " + kind)
	}
	// If we're passed in an empty packet, which could be caused by the server sending
	// us an "error:102-3#1", just ignore it.
	if pkt == nil {
		return pkt, nil
	}

	dat, ok := pkt.(*defs.SIndexGetUidEnMicaQueue)
	if !ok {
		return pkt, fmt.Errorf("Failed to cast %#v into &defs.SMicaQueue", pkt)
	}
	openid, ok := pktCtx.Req.Form["openid"]
	if !ok {
		return pkt, incompletePktErr
	}
	longtoken, ok := pktCtx.Req.Form["sid"]
	if !ok {
		return pkt, incompletePktErr
	}
	id, err := strconv.Atoi(openid[0])
	if err != nil {
		return pkt, incompletePktErr
	}
	prxy.AddUser(id, dat.UID, dat.Sign, longtoken[0])
	return pkt, nil
}

// Register hooks with dispatch
func init() {
	proxy.AuthHook = Handle
}
