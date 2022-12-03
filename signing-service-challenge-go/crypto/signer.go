package crypto

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

// Signer defines a contract for different types of signing implementations.
type Signer interface {
	Sign(dataToBeSigned []byte) ([]byte, error)
}

// TODO: implement RSA and ECDSA signing ...
type RSASigner struct {
	privateKey *rsa.PrivateKey
}

func (r *RSASigner) Sign(dataToBeSigned []byte) (signature []byte, err error) {
	msgHash := sha256.New()
	_, err = msgHash.Write(dataToBeSigned)
	if err != nil {
		return
	}
	msgHashSum := msgHash.Sum(nil)


	signature, err = rsa.SignPSS(rand.Reader, r.privateKey, crypto.SHA256, msgHashSum, nil)
	return
}


type ECCSigner struct {
	eccPrivateKey *ecdsa.PrivateKey
}

func (e *ECCSigner) Sign(dataToBeSigned []byte) (signature []byte, err error) {
	msgHash := sha256.New()
	_, err = msgHash.Write(dataToBeSigned)
	if err != nil {
		return
	}
	msgHashSum := msgHash.Sum(nil)

	signature, err = ecdsa.SignASN1(rand.Reader, e.eccPrivateKey, msgHashSum)
	return
}