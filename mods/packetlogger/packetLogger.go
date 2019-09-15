// Packet Logger - Logs all raw packet information

package packetlogger

import (
	"bufio"
	"fmt"
	hoxyLog "hoxy/log"
	"hoxy/proxy"
	"hoxy/utils"
	"log"
	"os"
	"time"

	"github.com/elazarl/goproxy"
)

type rawPacketLoggerState struct {
	logger  *log.Logger
	buffer  *bufio.Writer
	userCtx *proxy.UserCtx
}

func (state *rawPacketLoggerState) handle(op string, pkt interface{}, userCtx *proxy.UserCtx, pktCtx *goproxy.ProxyCtx) (interface{}, error) {
	state.logger.Printf("[%s] %s\n", op, string(pkt.([]byte)))
	return pkt, nil
}

func (state *rawPacketLoggerState) Shutdown(bool) {
	hoxyLog.Infof("Shutting down packetLogger for %s\n", state.userCtx.UID)
	state.buffer.Flush()
}

// Register hooks with dispatch
func init() {
	modName := "Packet Logger"
	initFunc := func(userCtx *proxy.UserCtx) ([]*proxy.PacketHook, proxy.ShutdownCb, error) {
		dir := utils.PackageRoot + "/logs/packetLogger/" + userCtx.UID + "/"
		if err := os.MkdirAll(dir, 0755); err != nil {
			return nil, nil, err
		}
		now := time.Now()
		f, err := os.Create(fmt.Sprintf("%s%s.log", dir, now.Format("2006-01-02_15.04.05")))
		if err != nil {
			return nil, nil, err
		}
		buffer := bufio.NewWriter(f)
		logger := log.New(buffer, "", log.Ltime)

		mod := &rawPacketLoggerState{logger, buffer, userCtx}
		hooks := []*proxy.PacketHook{
			proxy.NewPacketHook(modName, "*", 0, true, mod.handle),
		}
		return hooks, mod.Shutdown, nil
	}
	proxy.RegisterMod(modName, initFunc)
}
