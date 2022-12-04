package proxy

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/kyoukaya/hoxy/log"
	"github.com/kyoukaya/hoxy/utils"

	"github.com/elazarl/goproxy"
)

// HandleReq proccess an outgoing http(s) request.
func (hoxy *HoxyProxy) HandleReq(req *http.Request, ctx *goproxy.ProxyCtx) (*http.Request, *http.Response) {
	var body []byte

	if hoxy.shuttingDown {
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
	if hoxy.telemetryFilter.MatchString(req.Host) {
		log.Verbosef("==== Rejecting %v", req.Host)
		// Use the UserData field as a flag to indicate to the response handler that the
		// request that generated the response was blocked.
		reqCtx.RequestIsBlocked = true
		return req, goproxy.NewResponse(req, goproxy.ContentTypeText, http.StatusOK, "")
	}
	reqURL := req.URL.String()
	body, err := io.ReadAll(req.Body)
	utils.Check(err)
	// Reading from the body consumes all the bytes, so we need to save back a copy.
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	// TODO: Parsing wrecks the request body damnit! There should be a better way of parsing http forms.
	req.ParseForm()
	req.Body = io.NopCloser(bytes.NewBuffer(body))
	// Handle game traffic
	if strings.HasPrefix(req.URL.String(), hoxy.baseURL) {
		return hoxy.dispatchReq(req, ctx)
	}
	// Non-game traffic
	// Block some of the uninteresting requests to prevent them from flooding our logs.
	if hoxy.logFilter.MatchString(reqURL) {
		log.Verbosef(">>>> %s\n%v\n%s", reqURL, utils.SprintHeaders(req.Header), body)
	} else {
		log.Infof(">>>> %s\n%v\n%s", reqURL, utils.SprintHeaders(req.Header), body)
	}
	return req, nil
}

// HandleResp processes an incoming http(s) response.
func (hoxy *HoxyProxy) HandleResp(resp *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
	var body []byte

	if hoxy.shuttingDown {
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
	body, err := io.ReadAll(resp.Body)
	utils.Check(err)
	// Reading from the body consumes all the bytes, so we need to save back a copy.
	resp.Body = io.NopCloser(bytes.NewBuffer(body))
	// Game traffic
	if strings.HasPrefix(reqURL, hoxy.baseURL) {
		return hoxy.dispatchRes(body, resp, ctx)
	}
	// Non-game traffic
	// Block some of the uninteresting requests to prevent them from flooding our logs.
	if hoxy.logFilter.MatchString(reqURL) {
		log.Verbosef("<<<< %s\n%s", reqURL, body)
	} else {
		log.Infof("<<<< %s\n%s", reqURL, body)
	}
	return resp
}

// Shutdown stops processing all packets captured by the proxy and calls Shutdown
// on all modules for all users.
func (hoxy *HoxyProxy) Shutdown() {
	hoxy.shuttingDown = true
	for _, user := range hoxy.users {
		for _, cb := range user.shutdownCBs {
			cb(true)
		}
	}
}

// GetUser returns a UserCtx for the specified UID
func (hoxy *HoxyProxy) GetUser(UID string) *UserCtx {
	hoxy.mutex.Lock()
	defer hoxy.mutex.Unlock()
	return hoxy.users[UID]
}

// addUser records a user's information indexed by their UID, if a record belonging to
// the specified UID already exists, its hooks will be shutdown and the record will be overwritten.
func (hoxy *HoxyProxy) addUser(openID int, UID, sign, longtoken string) {
	hoxy.mutex.Lock()
	defer hoxy.mutex.Unlock()

	if user, exists := hoxy.users[UID]; exists {
		log.Infof("%s reconnecting. Shutting down mods.", UID)
		for _, cb := range user.shutdownCBs {
			cb(false)
		}
	}

	hoxy.users[UID] = &UserCtx{
		mutex:     &sync.Mutex{},
		Key:       sign,
		Longtoken: longtoken,
		UID:       UID,
		Openid:    openID,
		RawHooks:  make(map[string][]*PacketHook),
		Hooks:     make(map[string][]*PacketHook),
	}
}
