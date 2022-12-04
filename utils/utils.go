package utils

import (
	"flag"
	"log"
	"net"
	"net/http"
	"path"
	"runtime"
	"strings"
)

var PackageRoot = func() string {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		panic("Unable to determine package root")
	}
	return path.Join(path.Dir(filename), "../")
}()

var (
	boolFlags   = make(map[string]bool)
	stringFlags = make(map[string]string)
)

func Check(err error) {
	if err != nil {
		panic(err)
	}
}

func SprintHeaders(header http.Header) string {
	s := ""
	for k, v := range header {
		s += k + ":"
		for _, v2 := range v {
			s += v2
		}
		s += "\n"
	}
	return s
}

func BoolFlags(key string) bool {
	return boolFlags[key]
}

func StringFlags(key string) string {
	return stringFlags[key]
}

func ParseFlags() {
	verbose := flag.Bool("v", false, "log every GoProxy request to stdout")
	logVerbose := flag.Bool("hoxy-verbose", false, "log more Hoxy information")
	addr := flag.String("addr", ":8080", "proxy listen address")
	https := flag.Bool("https", false, "MITM https connections. Requires loading a CA")
	config := flag.String("c", "", "Path of TOML config file")

	flag.Parse()

	stringFlags["addr"] = *addr
	stringFlags["config"] = *config
	boolFlags["v"] = *verbose
	boolFlags["hoxy-verbose"] = *logVerbose
	boolFlags["https"] = *https
}

// GetOutboundIP gets preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// Convert a String into an Array with a delimiter
func StringToArray(dat string, delim string) []string {
	// []string to be returned
	var ret []string
	// counter for ret
	var o = 0

	// main loop
	for {
		bef, aft, f := strings.Cut(dat, delim)
		if !f {
			ret = append(ret, bef)
			break
		}

		ret = append(ret, bef)
		o++
		dat = aft
	}

	return ret
}
