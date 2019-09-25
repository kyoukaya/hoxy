// Construction Info - Prints to console what gun/equip a user has crafted immediately after
// they have sent the construction order.

package constructioninfo

import (
	"fmt"
	"hoxy/defs"
	"hoxy/log"
	"hoxy/proxy"
	"hoxy/utils/dollinfo"
	"hoxy/utils/equipinfo"

	"github.com/elazarl/goproxy"
)

type constructInfo struct{}

func (*constructInfo) handleGunConstruct(op string, pkt interface{}, userCtx *proxy.UserCtx, pktCtx *goproxy.ProxyCtx) (interface{}, error) {
	dispatchCtx := proxy.GetDispatchContext(pktCtx)
	if dispatchCtx.RequestUnmarshalErr != nil {
		return pkt, fmt.Errorf("construction info: reqCtx contains unmarshalled packet")
	}
	reqPkt := dispatchCtx.RequestPkt

	var rscUsed, guns string
	var err error
	switch pkt.(type) {
	case *defs.SGunDevelopGun:
		rscUsed, guns, err = handleSingleGun(pkt.(*defs.SGunDevelopGun), reqPkt.(*defs.CGunDevelopGun))
	case *defs.SGunDevelopMultiGun:
		rscUsed, guns, err = handleMultiGun(pkt.(*defs.SGunDevelopMultiGun), reqPkt.(*defs.CGunDevelopMultiGun))
	default:
		return pkt, fmt.Errorf("constructioninfo: wrong packet, expected 'SGunDevelopGun/SGunDevelopMultiGun', got: %#v", pkt)
	}

	if err != nil {
		return pkt, err
	}

	log.Infof("Gun crafting with %s yields [%s].", rscUsed, guns)
	return pkt, nil
}

func (*constructInfo) handleEquipConstruct(op string, pkt interface{}, userCtx *proxy.UserCtx, pktCtx *goproxy.ProxyCtx) (interface{}, error) {
	dispatchCtx := proxy.GetDispatchContext(pktCtx)
	if dispatchCtx.RequestUnmarshalErr != nil {
		return pkt, fmt.Errorf("construction info: reqCtx contains unmarshalled packet")
	}
	reqPkt := dispatchCtx.RequestPkt

	var rscUsed, equips string
	var err error
	switch pkt.(type) {
	case *defs.SEquipDevelop:
		rscUsed, equips, err = handleSingleEquip(pkt.(*defs.SEquipDevelop), reqPkt.(*defs.CEquipDevelop))
	case *defs.SEquipDevelopMulti:
		rscUsed, equips, err = handleMultiEquip(pkt.(*defs.SEquipDevelopMulti), reqPkt.(*defs.CEquipDevelopMulti))
	default:
		return pkt, fmt.Errorf("constructioninfo: wrong packet, expected 'SEquipDevelop/SEquipDevelopMulti', got: %#v", pkt)
	}

	if err != nil {
		return pkt, err
	}

	log.Infof("Equip crafting with %s yields [%s].", rscUsed, equips)
	return pkt, nil
}

// Register hooks with dispatch
func init() {
	const modName = "Construction Info"
	initFunc := func(userCtx *proxy.UserCtx) ([]*proxy.PacketHook, proxy.ShutdownCb, error) {
		dollinfo.Init()
		equipinfo.Init()
		mod := &constructInfo{}
		hooks := []*proxy.PacketHook{
			proxy.NewPacketHook(modName, "SGun/developGun", 0, false, mod.handleGunConstruct),
			proxy.NewPacketHook(modName, "SGun/developMultiGun", 0, false, mod.handleGunConstruct),
			proxy.NewPacketHook(modName, "SEquip/develop", 0, false, mod.handleEquipConstruct),
			proxy.NewPacketHook(modName, "SEquip/developMulti", 0, false, mod.handleEquipConstruct),
		}
		return hooks, func(bool) {}, nil
	}
	proxy.RegisterMod(modName, initFunc)
}
