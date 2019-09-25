package proxy

import (
	"github.com/kyoukaya/hoxy/log"
	"github.com/kyoukaya/hoxy/utils"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"sync"
	"syscall"

	"github.com/elazarl/goproxy"
)

const (
	// GlobalGameBaseURL is the game's base URL for the global version of GFL.
	GlobalGameBaseURL = "http://gf-game.sunborngame.com/index.php/1001/"
)

// HoxyProxy contains the state relevant to the proxy
type HoxyProxy struct {
	mutex              *sync.Mutex
	authHandler        AuthHandler
	baseURL            string
	server             *goproxy.ProxyHttpServer
	logGamePacketEvent bool
	shuttingDown       bool
	// Filters
	httpsFilter     *regexp.Regexp
	telemetryFilter *regexp.Regexp
	logFilter       *regexp.Regexp
	// users contains a mapping of a user's UID in string form to a UserCtx struct
	// containing the context pertaining to the user.
	users map[string]*UserCtx
}

// AuthHandler is called whenever an authentication packet is received by the proxy in order
// to initialize user state.
type AuthHandler func(kind string, pkt interface{}, pktCtx *goproxy.ProxyCtx) (openID int, UID, sign, longtoken string, err error)

// NewHoxy returns a new initialized HoxyProxy
func NewHoxy(baseURL string, authHandler AuthHandler, filters Filters) *HoxyProxy {
	utils.ParseFlags()
	log.InitLogger(true, true, "")
	// TODO: Init the standard logger based on flags.

	server := goproxy.NewProxyHttpServer()

	hoxy := &HoxyProxy{
		mutex:           &sync.Mutex{},
		authHandler:     authHandler,
		baseURL:         baseURL,
		server:          server,
		shuttingDown:    false,
		httpsFilter:     generateFilter(filters.HTTPSFilter),
		telemetryFilter: generateFilter(filters.TelemetryFilter),
		logFilter:       generateFilter(filters.LogFilter),
		users:           make(map[string]*UserCtx)}
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
		server.OnRequest().HandleConnect(goproxy.FuncHttpsHandler(hoxy.httpsPassthrough))
	}

	return hoxy
}

// LogGamePackets will log the op of any packet sent or received from the servers.
// This is probably only useful for developers and is false by default.
func (hoxy *HoxyProxy) LogGamePackets(b bool) {
	hoxy.logGamePacketEvent = b
}

// Start starts the proxy. This is blocking and does not return.
func (hoxy *HoxyProxy) Start() {
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

	ipstring := utils.GetOutboundIP()
	addrSplit := strings.Split(utils.StringFlags("addr"), ":")
	if len(addrSplit) == 2 {
		ipstring += ":" + addrSplit[1]
	}
	log.Infof("Hoxy started on %s", ipstring)
	log.Fatal(http.ListenAndServe(utils.StringFlags("addr"), hoxy.server))
}
