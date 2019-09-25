package proxy

import (
	"bytes"
	"hoxy/log"
	"hoxy/utils"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/elazarl/goproxy"
)

const (
	gameBaseURL = "http://gf-game.sunborngame.com/index.php/1001/"
)

// HandleReq proccess an outgoing http(s) request.
func (proxy *HoxyProxy) HandleReq(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	var body []byte

	if proxy.shuttingDown {
		return nil, nil
	}

	defer func() {
		log.Flush()
		if rec := recover(); rec != nil {
			log.Infoln(body)
			panic(rec)
		}
	}()

	reqCtx := &DispatchContext{}
	ctx.UserData = reqCtx
	// Block telemetry requests
	if telemetryFilter.MatchString(req.Host) {
		log.Verbosef("==== Rejecting %v", req.Host)
		// Use the UserData field as a flag to indicate to the response handler that the
		// request that generated the response was blocked.
		reqCtx.RequestIsBlocked = true
		return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusOK, "")
	}
	reqURL := req.URL.String()
	body, err := ioutil.ReadAll(req.Body)
	utils.Check(err)
	// Reading from the body consumes all the bytes, so we need to save back a copy.
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	// TODO: Parsing wrecks the request body damnit! There should be a better way of parsing http forms.
	req.ParseForm()
	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	// Handle game traffic
	if strings.HasPrefix(req.URL.String(), gameBaseURL) {
		return proxy.dispatchReq(req, ctx)
	}
	// Non-game traffic
	// Block some of the uninteresting requests to prevent them from flooding our logs.
	if logFilter.MatchString(reqURL) {
		log.Verbosef(">>>> %s\n%v\n%s", reqURL, utils.SprintHeaders(req.Header), body)
	} else {
		log.Infof(">>>> %s\n%v\n%s", reqURL, utils.SprintHeaders(req.Header), body)
	}
	return req, nil
}

// HandleResp processes an incoming http(s) response.
func (proxy *HoxyProxy) HandleResp(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	var body []byte

	if proxy.shuttingDown {
		return nil
	}

	defer func() {
		log.Flush()
		if rec := recover(); rec != nil {
			log.Infoln(string(body))
			panic(rec)
		}
	}()

	reqCtx := GetDispatchContext(ctx)
	// If request that generated response was blocked or response not OK.
	if reqCtx == nil || resp == nil {
		return resp
	}
	if reqCtx.RequestIsBlocked || resp == nil || resp.StatusCode != 200 {
		return resp
	}
	reqURL := ctx.Req.URL.String()
	body, err := ioutil.ReadAll(resp.Body)
	utils.Check(err)
	// Reading from the body consumes all the bytes, so we need to save back a copy.
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	// Game traffic
	if strings.HasPrefix(reqURL, gameBaseURL) {
		return proxy.dispatchRes(body, resp, ctx)
	}
	// Non-game traffic
	// Block some of the uninteresting requests to prevent them from flooding our logs.
	if logFilter.MatchString(reqURL) {
		log.Verbosef("<<<< %s\n%s", reqURL, body)
	} else {
		log.Infof("<<<< %s\n%s", reqURL, body)
	}
	return resp
}

// Shutdown stops processing all packets captured by the proxy and calls Shutdown
// on all modules for all users.
func (proxy *HoxyProxy) Shutdown() {
	proxy.shuttingDown = true
	for _, user := range proxy.users {
		for _, cb := range user.shutdownCBs {
			cb(true)
		}
	}
}

// GetUser returns a UserCtx for the specified UID
func (proxy *HoxyProxy) GetUser(UID string) *UserCtx {
	proxy.mutex.Lock()
	defer proxy.mutex.Unlock()
	return proxy.users[UID]
}

// addUser records a user's information indexed by their UID, if a record belonging to
// the specified UID already exists, its hooks will be shutdown and the record will be overwritten.
func (proxy *HoxyProxy) addUser(openID int, UID, sign, longtoken string) {
	proxy.mutex.Lock()
	defer proxy.mutex.Unlock()

	if user, exists := proxy.users[UID]; exists {
		log.Infof("%s reconnecting. Shutting down mods.", UID)
		for _, cb := range user.shutdownCBs {
			cb(false)
		}
	}

	proxy.users[UID] = &UserCtx{
		mutex:     &sync.Mutex{},
		Key:       sign,
		Longtoken: longtoken,
		UID:       UID,
		Openid:    openID,
		RawHooks:  make(map[string][]*PacketHook),
		Hooks:     make(map[string][]*PacketHook),
	}
}
