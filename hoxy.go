package main

import (
	"hoxy/log"
	_ "hoxy/mods"
	"hoxy/proxy"
	"hoxy/utils"
	"hoxy/utils/dollinfo"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/elazarl/goproxy"
)

func main() {
	utils.ParseFlags()
	log.InitLogger(true, true, "")
	// TODO: Init the standard logger based on flags.

	dollinfo.InitDollInfo()
	server := goproxy.NewProxyHttpServer()
	hoxy := proxy.NewHoxy()
	server.OnRequest().DoFunc(hoxy.HandleReq)
	server.OnResponse().DoFunc(hoxy.HandleResp)

	if utils.BoolFlags("https") {
		_, certStatErr := os.Stat(utils.PackageRoot + "/cert.pem")
		_, keyStatErr := os.Stat(utils.PackageRoot + "/key.pem")
		// Generate CA if it doesn't exist
		if os.IsNotExist(certStatErr) || os.IsNotExist(keyStatErr) {
			log.Infof("Generating CA...")
			if err := utils.GenerateCA(); err != nil {
				log.Fatal(err)
			}
			log.Infof("Copy and register the created 'cert.pem' with your client.")
		}
		if err := utils.LoadCA(); err != nil {
			log.Fatal(err)
		}
		server.OnRequest().HandleConnect(goproxy.FuncHttpsHandler(proxy.HTTPSPassthrough))
	}

	sigs := make(chan os.Signal)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// Catch sigint/sigterm and cleanly exit
	go func() {
		sig := <-sigs
		log.Infof("%v received. Shutting down gracefully.\n", sig)
		log.Flush()
		hoxy.Shutdown()
		os.Exit(0)
	}()

	ipstring := ""
	if utils.StringFlags("addr")[0] != '"' {
		ipstring = utils.GetOutboundIP()
	}
	log.Infof("Hoxy started on %s%s", ipstring, utils.StringFlags("addr"))
	log.Fatal(http.ListenAndServe(utils.StringFlags("addr"), server))
}
