package utils

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"net"
	"os"
	"time"
)

// helper function to create a cert template with a serial number and other required fields
func certTemplate() (*x509.Certificate, error) {
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, errors.New("failed to generate serial number: " + err.Error())
	}

	issuerName := pkix.Name{
		CommonName:         "github.com/kyoukaya/hoxy",
		Organization:       []string{"Hoxy"},
		OrganizationalUnit: []string{"Hoxy"},
	}
	tmpl := x509.Certificate{
		BasicConstraintsValid: true,
		Issuer:                issuerName,
		NotAfter:              time.Now().AddDate(5, 0, 0), // CA expires in 5 years
		NotBefore:             time.Now(),
		SerialNumber:          serialNumber,
		SignatureAlgorithm:    x509.SHA256WithRSA,
		Subject:               issuerName,
	}
	return &tmpl, nil
}

func createCert(template, parent *x509.Certificate, pub interface{}, parentPriv interface{}) (
	cert *x509.Certificate, certPEM []byte, err error) {

	certDER, err := x509.CreateCertificate(rand.Reader, template, parent, pub, parentPriv)
	if err != nil {
		return
	}
	// parse the resulting certificate so we can use it again
	cert, err = x509.ParseCertificate(certDER)
	if err != nil {
		return
	}
	// PEM encode the certificate (this is a standard TLS encoding)
	b := pem.Block{Type: "CERTIFICATE", Bytes: certDER}
	certPEM = pem.EncodeToMemory(&b)
	return
}

func GenerateCA() error {
	// generate a new key-pair
	rootKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("generating random key: %v", err)
	}

	rootCertTmpl, err := certTemplate()
	if err != nil {
		return fmt.Errorf("creating cert template: %v", err)
	}
	// describe what the certificate will be used for
	rootCertTmpl.IsCA = true
	rootCertTmpl.KeyUsage = x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature
	rootCertTmpl.ExtKeyUsage = []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth}
	rootCertTmpl.IPAddresses = []net.IP{net.ParseIP("127.0.0.1")}
	_, rootCertPEM, err := createCert(rootCertTmpl, rootCertTmpl, &rootKey.PublicKey, rootKey)
	if err != nil {
		return fmt.Errorf("error creating cert: %v", err)
	}
	rootKeyPEM := pem.EncodeToMemory(&pem.Block{
		Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(rootKey),
	})
	f, err := os.Create(PackageRoot + CertPath)
	if err != nil {
		return fmt.Errorf("error writing cert to file: %v", err)
	}
	f.Write(rootCertPEM)
	f, err = os.Create(PackageRoot + KeyPath)
	if err != nil {
		return fmt.Errorf("error writing key to file: %v", err)
	}
	f.Write(rootKeyPEM)
	return nil
}