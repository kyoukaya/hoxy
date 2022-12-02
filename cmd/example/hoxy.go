package main

import (
	"github.com/kyoukaya/hoxy/proxy"

	_ "github.com/kyoukaya/hoxy/mods/constructioninfo"
	_ "github.com/kyoukaya/hoxy/mods/packetlogger"
)

func main() {
	proxy.Start()
}
