package oidcop

import jose "github.com/go-jose/go-jose/v4"

type PublicKey struct {
	*SigningKey
}

func (s *PublicKey) ID() string {
	return s.id
}

func (s *PublicKey) Algorithm() jose.SignatureAlgorithm {
	return s.algorithm
}

func (s *PublicKey) Use() string {
	return "sig"
}

func (s *PublicKey) Key() any {
	return s.pubkey
}
