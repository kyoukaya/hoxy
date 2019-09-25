package main

import (
	"github.com/kyoukaya/hoxy/proxy"
	"github.com/kyoukaya/hoxy/proxy/core/userauth"

	_ "github.com/kyoukaya/hoxy/mods/constructioninfo"
	_ "github.com/kyoukaya/hoxy/mods/packetlogger"
)

func main() {
	filters := proxy.Filters{
		HTTPSFilter:     proxy.DefaultHTTPSFilter,
		TelemetryFilter: proxy.DefaultTelemetryFilter,
		LogFilter:       proxy.DefaultLogFilter,
	}
	hoxy := proxy.NewHoxy(proxy.GlobalGameBaseURL, userauth.AuthHandler, filters)
	hoxy.LogGamePackets(true)
	hoxy.Start()
}
