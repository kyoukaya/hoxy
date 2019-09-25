package main

import (
	"hoxy/proxy"
	"hoxy/proxy/core/userauth"

	_ "hoxy/mods/constructioninfo"
	_ "hoxy/mods/packetlogger"
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
