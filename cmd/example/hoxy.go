package main

import (
	"flag"
	"fmt"

	"github.com/kyoukaya/hoxy/proxy"
	"github.com/kyoukaya/hoxy/proxy/core/userauth"
	"github.com/kyoukaya/hoxy/utils"

	_ "github.com/kyoukaya/hoxy/mods/constructioninfo"
	_ "github.com/kyoukaya/hoxy/mods/packetlogger"
)

func main() {
	var floc string
	var help bool
	flag.StringVar(&floc, "c", "", "ini config file location")
	flag.BoolVar(&help, "h", false, "Shows the help menu")
	flag.Parse()
	if help {
		fmt.Println("  -h    This Menu\n  -c    INI Config file location")
	}

	utils.ReadFileIntoList(floc, ' ')
	filters := proxy.Filters{
		HTTPSFilter:     proxy.DefaultHTTPSFilter,
		TelemetryFilter: proxy.DefaultTelemetryFilter,
		LogFilter:       proxy.DefaultLogFilter,
	}

	hoxy := proxy.NewHoxy(proxy.GlobalGameBaseURL, userauth.AuthHandler, filters)
	hoxy.LogGamePackets(true)
	hoxy.Start()
}
