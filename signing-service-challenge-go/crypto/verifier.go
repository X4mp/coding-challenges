package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type Verifier interface {
	VerifySignature([]byte, []byte) error
}

type RSAVerifier struct {
	PublicKeyRSA *rsa.PublicKey
}

func (r *RSAVerifier) VerifySignature(dataToBeSigned, signature []byte) (err error) {
	msgHash := sha256.New()
	_, err = msgHash.Write(dataToBeSigned)
	if err != nil {
		return
	}
	msgHashSum := msgHash.Sum(nil)

	return rsa.VerifyPSS(r.PublicKeyRSA, crypto.SHA256, msgHashSum, signature, nil)
}