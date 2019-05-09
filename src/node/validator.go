package node

import (
	"crypto/ecdsa"

	"github.com/mosaicnetworks/babble/src/crypto/keys"
)

type Validator struct {
	Key     *ecdsa.PrivateKey
	Moniker string

	id       uint32
	pubBytes []byte
	pubHex   string
}

func NewValidator(key *ecdsa.PrivateKey, moniker string) *Validator {
	return &Validator{
		Key:     key,
		Moniker: moniker,
	}
}

func (v *Validator) ID() uint32 {
	if v.id == 0 {
		v.id = keys.PublicKeyID(&v.Key.PublicKey)
	}
	return v.id
}

func (v *Validator) PublicKeyBytes() []byte {
	if v.pubBytes == nil || len(v.pubBytes) == 0 {
		v.pubBytes = keys.FromPublicKey(&v.Key.PublicKey)
	}
	return v.pubBytes
}

func (v *Validator) PublicKeyHex() string {
	if len(v.pubHex) == 0 {
		v.pubHex = keys.PublicKeyHex(&v.Key.PublicKey)
	}
	return v.pubHex
}
