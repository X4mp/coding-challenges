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
	publicKeyRSA *rsa.PublicKey
}

func NewRSAVerifier(publicKey *rsa.PublicKey) RSAVerifier {
	return RSAVerifier{
		publicKeyRSA: publicKey,
	}
}

func (r RSAVerifier) VerifySignature(dataToBeSigned, signature []byte) (err error) {
	msgHash := sha256.New()
	_, err = msgHash.Write(dataToBeSigned)
	if err != nil {
		return
	}
	msgHashSum := msgHash.Sum(nil)

	return rsa.VerifyPSS(r.publicKeyRSA, crypto.SHA256, msgHashSum, signature, nil)
}

type ECCVerifier struct {
	publicKeyECC *ecdsa.PublicKey
}

func NewECCVerifier(publicKey *ecdsa.PublicKey) ECCVerifier {
	return ECCVerifier{
		publicKeyECC: publicKey,
	}
}

func (e ECCVerifier) VerifySignature(dataToBeSigned, signature []byte) (err error) {
	msgHash := sha256.New()
	_, err = msgHash.Write(dataToBeSigned)
	if err != nil {
		return
	}
	msgHashSum := msgHash.Sum(nil)

	verified := ecdsa.VerifyASN1(e.publicKeyECC, msgHashSum, signature)
	if !verified {
		err = ErrNotVerified
	}
	return
}
