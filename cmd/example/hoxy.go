package main

import (
	"hoxy/proxy"

	_ "hoxy/mods/constructioninfo"
	_ "hoxy/mods/packetlogger"
)

func main() {
	hoxy := proxy.NewHoxy()
	hoxy.Start()
}
