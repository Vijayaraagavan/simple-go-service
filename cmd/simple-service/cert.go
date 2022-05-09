package main

import (
	"fmt"
	"crypto/rand"
	"time"
	"math/big"
	"crypto/rsa"
	// "crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	// "io/ioutil"
	"net"
	// "net/http"
	"os"
)


func main() {
	serverCert, serverKey, err := generate()
	if err != nil {
		fmt.Println("error generating sever certificate: %v", err)
		os.Exit(1)
	}
	fmt.Println("Server Certificate")
	fmt.Printf("%s\n",serverCert)
	fmt.Println("Server Key")
	fmt.Printf("%s\n",serverKey)

}

func generate() (cert []byte, privateKey []byte, err error) {

	//serialnumber, ca
	serialNumber, err := rand.Int(rand.Reader, big.NewInt(27))
	if err!=nil {
		return cert, privateKey, err
	}
	notBefore := time.Now()
	ca := &x509.Certificate{
		SerialNumber: serialNumber,
		Subject : pkix.Name{
			Organization: []string{"startup.com"},
		},
		NotBefore: notBefore,
		KeyUsage: x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		BasicConstraintsValid: true,
		IPAddresses:           []net.IP{net.ParseIP("127.0.0.1")},	
	}

	//rsaKey generation
	rsaKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err!=nil {
		return cert, privateKey, err
	}

	// creating DER
	DER, err :=  x509.CreateCertificate(rand.Reader, ca, ca, &rsaKey.PublicKey, rsaKey)
	if err!=nil {
		return cert, privateKey, err
	}
	
	// PEM block
	b := pem.Block{
		Type: "CERTIFICATE",
		Bytes: DER,
	}

	//certifcate 
	cert = pem.EncodeToMemory(&b)

	//Private key 
	privateKey = pem.EncodeToMemory(
		&pem.Block {
			Type: "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(rsaKey),
		})
		return cert, privateKey, nil
}