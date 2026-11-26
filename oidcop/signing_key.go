package oidcop

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"os"

	jose "github.com/go-jose/go-jose/v4"
)

type SigningKey struct {
	id        string
	algorithm jose.SignatureAlgorithm
	key       ed25519.PrivateKey
	pubkey    ed25519.PublicKey
}

func NewSigningKey(id string, key ed25519.PrivateKey, pubkey ed25519.PublicKey) *SigningKey {
	return &SigningKey{
		id:        id,
		algorithm: jose.EdDSA,
		key:       key,
		pubkey:    pubkey,
	}
}

// ID implements [op.SigningKey].
func (s *SigningKey) ID() string {
	return s.id
}

// Key implements [op.SigningKey].
func (s *SigningKey) Key() any {
	return s.key
}

// SignatureAlgorithm implements [op.SigningKey].
func (s *SigningKey) SignatureAlgorithm() jose.SignatureAlgorithm {
	return s.algorithm
}

// LoadSigningKey loads an ed25519 private key from a PEM (PKCS#8) file.
func LoadSigningKey(id string, path string) (*SigningKey, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, ErrInvalidSigningKey
	}

	parsed, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	privKey, ok := parsed.(ed25519.PrivateKey)
	if !ok {
		return nil, ErrInvalidSigningKey
	}

	pubKey, ok := privKey.Public().(ed25519.PublicKey)
	if !ok {
		return nil, ErrInvalidSigningKey
	}

	return NewSigningKey(id, privKey, pubKey), nil
}
