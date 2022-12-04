package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
)

var ErrNotVerified error = fmt.Errorf("not able to verify")

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

type ECCVerifier struct {
	PublicKeyECC *ecdsa.PublicKey
}

func (e *ECCVerifier) VerifySignature(dataToBeSigned, signature []byte) (err error) {
	msgHash := sha256.New()
	_, err = msgHash.Write(dataToBeSigned)
	if err != nil {
		return
	}
	msgHashSum := msgHash.Sum(nil)

	verified := ecdsa.VerifyASN1(e.PublicKeyECC, msgHashSum, signature)
	if !verified {
		err = ErrNotVerified
	}
	return
}
