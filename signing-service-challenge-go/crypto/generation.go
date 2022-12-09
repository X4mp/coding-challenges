package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/rsa"
)

// RSAGenerator generates a RSA key pair.
type RSAGenerator struct{}

// Generate generates a new RSAKeyPair.
func (g *RSAGenerator) Generate() (*RSAKeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := rsa.GenerateKey(rand.Reader, 512)
	if err != nil {
		return nil, err
	}

	return &RSAKeyPair{
		Public:  &key.PublicKey,
		Private: key,
	}, nil
}

// ECCGenerator generates an ECC key pair.
type ECCGenerator struct{}

// Generate generates a new ECCKeyPair.
func (g *ECCGenerator) Generate() (*ECCKeyPair, error) {
	// Security has been ignored for the sake of simplicity.
	key, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		return nil, err
	}

	return &ECCKeyPair{
		Public:  &key.PublicKey,
		Private: key,
	}, nil
}

func NewAbstractTools(algorithm string) (signer Signer, verifier Verifier, err error) {
	switch algorithm {
	case "RSA":
		generator := RSAGenerator{}
		var keyPair *RSAKeyPair
		keyPair, err = generator.Generate()
		if err != nil {
			break
		}
		signer = NewRSASigner(keyPair.Private)
		verifier = NewRSAVerifier(keyPair.Public)
	case "ECC":
		generator := ECCGenerator{}
		var keyPair *ECCKeyPair
		keyPair, err = generator.Generate()
		if err != nil {
			break
		}
		signer = NewECCSigner(keyPair.Private)
		verifier = NewECCVerifier(keyPair.Public)
	default:
		err = ErrInvalidAlgorithm
		return
	}
	return
}
