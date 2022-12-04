package main

import (
	_ "github.com/kyoukaya/hoxy/mods/constructioninfo"
	_ "github.com/kyoukaya/hoxy/mods/packetlogger"
	"github.com/kyoukaya/hoxy/proxy"
)

func main() {
	proxy.Start()
}
