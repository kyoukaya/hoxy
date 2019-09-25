package proxy

import (
	"fmt"
	"hoxy/log"
	"hoxy/utils"
	"regexp"
	"strings"

	"github.com/elazarl/goproxy"
)

// Filters are a collection string slices, with each string representing a regexp expression
// to filter out requests. See generateFilter on how exactly they're compiled.
type Filters struct {
	// HTTPSFilter is only used when the https flag is on, any request matched by this
	// is NOT MITM'd.
	HTTPSFilter []string
	// TelemetryFilter blocks all requests that match.
	TelemetryFilter []string
	// LogFilter prevents stuff from being logged unless verbose mode is on.
	LogFilter []string
}

var (
	// DefaultHTTPSFilter prevents the intercepting of appguard requests.
	DefaultHTTPSFilter = []string{
		`appguard\.com\.cn`,
	}
	// DefaultTelemetryFilter blocks a bunch of domains which LDPlayer pings.
	DefaultTelemetryFilter = []string{
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
	}
	// DefaultLogFilter blocks some requests that clog up the log in normal development.
	DefaultLogFilter = []string{
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
	}
)

// generateFilter compiles a regexp expression for a given list of URLs
func generateFilter(list []string) *regexp.Regexp {
	ret := ""
	for _, v := range list {
		ret += fmt.Sprintf(`(^.*%s.*$)|`, v)
	}
	ret = strings.TrimRight(ret, "|")
	return regexp.MustCompile(ret)
}

// httpsPassthrough to allow HTTPS connections to pass through the proxy without being
// MITM'd.
func (hoxy *HoxyProxy) httpsPassthrough(host string, ctx *goproxy.ProxyCtx) (*goproxy.ConnectAction, string) {
	if hoxy.httpsFilter.MatchString(host) {
		log.Verbosef("~~~~ Authguard\n")
		return goproxy.OkConnect, host
	}
	if hoxy.telemetryFilter.MatchString(host) {
		log.Verbosef("==== Rejecting %v", host)
		return goproxy.RejectConnect, host
	}
	if utils.BoolFlags("https") {
		return goproxy.MitmConnect, host
	}
	return goproxy.OkConnect, host
}
