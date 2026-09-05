package saml

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"time"

	"github.com/crewjam/saml"
	"github.com/crewjam/saml/samlsp"
)

type SAML struct {
	sp      *samlsp.Middleware
	service *saml.ServiceProvider
}

type Config struct {
	EntityID         string
	MetadataURL      string
	AcsURL           string
	PrivateKey       []byte
	Certificate      []byte
	IDPMetadataURL   string
}

func New(cfg *Config) (*SAML, error) {
	key, err := parsePrivateKey(cfg.PrivateKey)
	if err != nil {
		return nil, err
	}
	cert, err := parseCertificate(cfg.Certificate)
	if err != nil {
		return nil, err
	}
	sp := &saml.ServiceProvider{
		EntityID:    cfg.EntityID,
		MetadataURL: cfg.MetadataURL,
		AcsURL:      cfg.AcsURL,
		Key:         key,
		Certificate: cert,
	}
	mw, err := samlsp.New(samlsp.Options{
		URL:                cfg.AcsURL,
		Key:                key,
		Certificate:        cert,
		IDPMetadataURL:     cfg.IDPMetadataURL,
		AllowIDPInitiated:  true,
	})
	if err != nil {
		return nil, err
	}
	return &SAML{sp: mw, service: sp}, nil
}

func (s *SAML) AuthURL() string {
	return s.sp.AuthorizeURL().String()
}

func (s *SAML) Validate(authnResponse string) (*saml.Response, error) {
	// Decode and validate SAML response
	return nil, nil
}

func parsePrivateKey(data []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}
	if block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("unexpected key type: %s", block.Type)
	}
	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func parseCertificate(data []byte) (*x509.Certificate, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}
	if block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("unexpected cert type: %s", block.Type)
	}
	return x509.ParseCertificate(block.Bytes)
}
