package main

import (
	_ "github.com/kyoukaya/hoxy/mods/constructioninfo"
	_ "github.com/kyoukaya/hoxy/mods/packetlogger"
	"github.com/kyoukaya/hoxy/proxy"
	"github.com/kyoukaya/hoxy/proxy/core/userauth"
)

func main() {
	proxy.Start(userauth.AuthHandler)
}
