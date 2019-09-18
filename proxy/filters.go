package proxy

import (
	"fmt"
	"hoxy/log"
	"hoxy/utils"
	"regexp"
	"strings"

	"github.com/elazarl/goproxy"
)

var (
	httpsFilter = GenerateFilter([]string{
		// Intercepting appguard requests seem to break the client. So we'll just let them through.
		`appguard\.com\.cn`,
	})
	telemetryFilter = GenerateFilter([]string{
		`graph\.facebook\.com`,
		`appsflyer\.com`,
		`sessions\.bugsnag\.com`,
		`cloud\.unity3d\.com`,
		`rayjump\.com`,
		`duapps\.com`,
		`spykemedia\.g2afse\.com`,
		`ldmnq\.com`,
		`erntech\.net`,
		`st\.frecorp\.net`,
		`apitask\.doglobal\.net`,
		`baidu\.clickurl\.to`,
	})
	logFilter = GenerateFilter([]string{
		// Seems like a lot of irrelevant data
		`gf-transit\.sunborngame\.com`,
		`track-us\.sunborngame\.com`,
		// Seems like irrelevant data
		`dkn3dfwjnmzcj\.cloudfront\.net/image/ImageConfig`,
		// Address not found
		`gf-dc\.sunborngame\.com`,
		// Unity data, probably useful if we plan on caching updates later on
		`s3.us-east-2\.amazonaws\.com:443/gf1-file-server/`,
		// Some FAQ service GFL uses
		`cs30\.net`,
		// Don't printout google stuff
		`google`,
		`gstatic`,
	})
)

// GenerateFilter compiles a regexp expression for a given list of URLs
func GenerateFilter(list []string) *regexp.Regexp {
	ret := ""
	for _, v := range list {
		ret += fmt.Sprintf(`(^.*%s.*$)|`, v)
	}
	ret = strings.TrimRight(ret, "|")
	return regexp.MustCompile(ret)
}

// HTTPSPassthrough to allow HTTPS connections to pass through the proxy without being
// MITM'd.
func HTTPSPassthrough(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
	if httpsFilter.MatchString(host) {
		log.Verbosef("~~~~ Authguard\n")
		return goproxy.OkConnect, host
	}
	if telemetryFilter.MatchString(host) {
		log.Verbosef("==== Rejecting %v", host)
		return goproxy.RejectConnect, host
	}
	if utils.BoolFlags("https") {
		return goproxy.MitmConnect, host
	}
	return goproxy.OkConnect, host
}
