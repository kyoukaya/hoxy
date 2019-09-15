package utils

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"os"

	"github.com/elazarl/goproxy"
)

const (
	CertPath = "/cert.pem"
	KeyPath  = "/key.pem"
)

func LoadCA() error {
	var err error

	certFile, err := os.Open(PackageRoot + CertPath)
	if err != nil {
		return err
	}
	defer certFile.Close()
	keyFile, err := os.Open(PackageRoot + KeyPath)
	if err != nil {
		return err
	}
	defer keyFile.Close()

	caCert, err := ioutil.ReadAll(certFile)
	if err != nil {
		return err
	}
	caKey, err := ioutil.ReadAll(keyFile)
	if err != nil {
		return err
	}

	goproxyCa, err := tls.X509KeyPair(caCert, caKey)
	if err != nil {
		return err
	}
	if goproxyCa.Leaf, err = x509.ParseCertificate(goproxyCa.Certificate[0]); err != nil {
		return err
	}
	goproxy.GoproxyCa = goproxyCa
	goproxy.OkConnect = &goproxy.ConnectAction{Action: goproxy.ConnectAccept, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.MitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.HTTPMitmConnect = &goproxy.ConnectAction{Action: goproxy.ConnectHTTPMitm, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	goproxy.RejectConnect = &goproxy.ConnectAction{Action: goproxy.ConnectReject, TLSConfig: goproxy.TLSConfigFromCA(&goproxyCa)}
	return nil
}
