package tlsconfig

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
)

func New(opts ...Option) *tls.Config {
	cf := &tls.Config{
		Certificates: []tls.Certificate{},
	}
	for _, opt := range opts {
		opt(cf)
	}
	return cf
}

type Option func(cf *tls.Config)

func CA(cert string) Option {
	return func(cf *tls.Config) {
		if cf.RootCAs == nil {
			cf.RootCAs = x509.NewCertPool()
		}
		if caCrt, err := ioutil.ReadFile(cert); err == nil {
			cf.RootCAs.AppendCertsFromPEM(caCrt)
		}
	}
}

func ClientCA(cert string) Option {
	return func(cf *tls.Config) {
		if cf.ClientCAs == nil {
			cf.ClientCAs = x509.NewCertPool()
		}
		if caCrt, err := ioutil.ReadFile(cert); err == nil {
			cf.ClientCAs.AppendCertsFromPEM(caCrt)
		}
	}
}

func CertKeyPair(cert string, key string) Option {
	return func(cf *tls.Config) {
		if pair, err := tls.LoadX509KeyPair(cert, key); err == nil {
			cf.Certificates = append(cf.Certificates, pair)
		}
	}
}

func ClientAuth(t tls.ClientAuthType) Option {
	return func(cf *tls.Config) {
		cf.ClientAuth = t
	}
}

func ClientAuthVerified() Option {
	return func(cf *tls.Config) {
		cf.ClientAuth = tls.RequireAndVerifyClientCert
	}
}

func ServerAuthSkip() Option {
	return func(cf *tls.Config) {
		cf.InsecureSkipVerify = true
	}
}

func GetCertificate(getter func(*tls.ClientHelloInfo) (*tls.Certificate, error)) Option {
	return func(cf *tls.Config) {
		cf.GetCertificate = getter
	}
}

func GetClientCertificate(getter func(*tls.CertificateRequestInfo) (*tls.Certificate, error)) Option {
	return func(cf *tls.Config) {
		cf.GetClientCertificate = getter
	}
}
