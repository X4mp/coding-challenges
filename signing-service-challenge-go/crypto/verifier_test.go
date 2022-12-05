package crypto_test

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"

	cr "github.com/X4mp/coding-challenges/signing-service-challenge/crypto"
)

func TestVerifier_Verify(t *testing.T) {
	dataToBeVerified := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")

	t.Run("RSA", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			keyPair, err := (&cr.RSAGenerator{}).Generate()
			assert.NoError(t, err)

			msgHash, err := cr.HashMessage(dataToBeVerified)
			assert.NoError(t, err)

			signature, err := rsa.SignPSS(rand.Reader, keyPair.Private, crypto.SHA256, msgHash, nil)
			assert.NoError(t, err)

			verifier := cr.NewRSAVerifier(keyPair.Public)
			assert.NoError(t, verifier.VerifySignature(dataToBeVerified, signature))
		})
	})

	t.Run("ECC", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			keyPair, err := (&cr.ECCGenerator{}).Generate()
			assert.NoError(t, err)

			msgHash, err := cr.HashMessage(dataToBeVerified)
			assert.NoError(t, err)

			signature, err := ecdsa.SignASN1(rand.Reader, keyPair.Private, msgHash)
			assert.NoError(t, err)

			verifier := cr.NewECCVerifier(keyPair.Public)
			assert.NoError(t, verifier.VerifySignature(dataToBeVerified, signature))
		})
	})
}
