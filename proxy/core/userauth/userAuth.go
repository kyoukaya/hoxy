package userauth

import (
	"fmt"
	"hoxy/defs"
	"strconv"

	"github.com/elazarl/goproxy"
)

var incompletePktErr = fmt.Errorf("Malformed or incomplete packet")

func AuthHandler(kind string, pkt interface{}, pktCtx *goproxy.ProxyCtx) (openID int, UID, sign, longtoken string, err error) {
	dat, ok := pkt.(*defs.SIndexGetUidEnMicaQueue)
	if !ok {
		err = fmt.Errorf("Failed to cast %#v into &defs.SMicaQueue")
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
		err = incompletePktErr
		return
	}
	openID, err = strconv.Atoi(openidStr[0])
	if err != nil {
		err = fmt.Errorf("userauth: " + err.Error())
		return
	}

	longtokenSlice, ok := pktCtx.Req.Form["sid"]
	if !ok {
		err = incompletePktErr
		return
	}
	longtoken = longtokenSlice[0]
	return
}
