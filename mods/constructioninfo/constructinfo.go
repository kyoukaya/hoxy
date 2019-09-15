package constructioninfo

import (
	"fmt"
	"hoxy/log"
	"hoxy/proxy"
	"hoxy/proxy/defs"
	"hoxy/utils/dollinfo"
	"strconv"
	"strings"

	"github.com/elazarl/goproxy"
)

type constructInfo struct{}

func (*constructInfo) handleSingle(op string, pkt interface{}, userCtx *proxy.UserCtx, pktCtx *goproxy.ProxyCtx) (interface{}, error) {
	data, ok := pkt.(*defs.SGunDevelopGun)
	if !ok {
		return pkt, fmt.Errorf("construction info: wrong packet, expected 'SGunDevelopGun', got: %#v", data)
	}
	dispatchCtx := proxy.GetDispatchContext(pktCtx)
	gunID, err := strconv.Atoi(data.GunID)
	if err != nil {
		return pkt, err
	}
	var rscUsed string
	if dispatchCtx.RequestUnmarshalErr == nil {
		data, ok := dispatchCtx.RequestPkt.(*defs.CGunDevelopGun)
		if !ok {
			return pkt, fmt.Errorf("construction info: reqCtx contains unmarshalled packet")
		}
		rscUsed = fmt.Sprintf("[%d/%d/%d/%d]", data.Mp, data.Ammo, data.Mre, data.Part)
	}
	log.Infof("Gun crafting with %s yields a %s.", rscUsed, dollinfo.Get(gunID).Name)
	return pkt, nil
}

func (*constructInfo) handleMulti(op string, pkt interface{}, userCtx *proxy.UserCtx, pktCtx *goproxy.ProxyCtx) (interface{}, error) {
	data, ok := pkt.(*defs.SGunDevelopMultiGun)
	if !ok {
		return pkt, fmt.Errorf("construction info: wrong packet, expected 'SGunDevelopMultiGun', got: %#v", data)
	}
	dispatchCtx := proxy.GetDispatchContext(pktCtx)
	var gunNames []string
	for _, v := range data.GunIds {
		id, err := strconv.Atoi(v.ID)
		if err != nil {
			return pkt, fmt.Errorf("Error while parsing %#v", v)
		}
		gunNames = append(gunNames, dollinfo.Get(id).Name)
	}
	var rscUsed string
	if dispatchCtx.RequestUnmarshalErr == nil {
		data, ok := dispatchCtx.RequestPkt.(*defs.CGunDevelopMultiGun)
		if !ok {
			return pkt, fmt.Errorf("construction Info: reqCtx contains unmarshalled packet")
		}
		rscUsed = fmt.Sprintf("[%d/%d/%d/%d] ", data.Mp, data.Ammo, data.Mre, data.Part)
	}
	log.Infof("Gun crafting with %s yields [%s].", rscUsed, strings.Join(gunNames, ", "))
	return pkt, nil
}

// Register hooks with dispatch
func init() {
	const modName = "Construction Info"
	initFunc := func(userCtx *proxy.UserCtx) ([]*proxy.PacketHook, proxy.ShutdownCb, error) {
		mod := &constructInfo{}
		hooks := []*proxy.PacketHook{
			proxy.NewPacketHook(modName, "SGun/developGun", 0, false, mod.handleSingle),
			proxy.NewPacketHook(modName, "SGun/developMultiGun", 0, false, mod.handleMulti),
		}
		return hooks, func(bool) {}, nil
	}
	proxy.RegisterMod(modName, initFunc)
}
