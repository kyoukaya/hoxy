package userauth

import (
	"fmt"
	"github.com/kyoukaya/hoxy/defs"
	"strconv"

	"github.com/elazarl/goproxy"
)

var errIncompletePkt = fmt.Errorf("Malformed or incomplete packet")

// AuthHandler handles user authentication for the global server
func AuthHandler(kind string, pkt interface{}, pktCtx *goproxy.ProxyCtx) (openID int, UID, sign, longtoken string, err error) {
	dat, ok := pkt.(*defs.SIndexGetUidEnMicaQueue)
	if !ok {
		err = fmt.Errorf("Failed to cast %#v into &defs.SMicaQueue", pkt)
		return
	}
	if dat.ErrNo != 0 || dat.ErrMsg != "" {
		err = fmt.Errorf("Failed to login: [%d] %s", dat.ErrNo, dat.ErrMsg)
		return
	}

	UID = dat.UID
	sign = dat.Sign

	openidStr, ok := pktCtx.Req.Form["openid"]
	if !ok {
		err = errIncompletePkt
		return
	}
	openID, err = strconv.Atoi(openidStr[0])
	if err != nil {
		err = fmt.Errorf("userauth: " + err.Error())
		return
	}

	longtokenSlice, ok := pktCtx.Req.Form["sid"]
	if !ok {
		err = errIncompletePkt
		return
	}
	longtoken = longtokenSlice[0]
	return
}
