package oidcop

import (
	"crypto/ed25519"

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
