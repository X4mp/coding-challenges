package crypto_test

import (
	"crypto"
	"crypto/rsa"
	"testing"

	"github.com/stretchr/testify/assert"

	cr "github.com/X4mp/coding-challenges/signing-service-challenge/crypto"
)

func TestSigner_Sign(t *testing.T) {
	dataToBeSigned := []byte("Lorem ipsum dolor sit amet, consectetur adipiscing elit.")

	t.Run("RSA", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			keyPair, err := (&cr.RSAGenerator{}).Generate()
			assert.NoError(t, err)

			signer := cr.NewRSASigner(keyPair.Private)

			signature, err := signer.Sign(dataToBeSigned)
			assert.NoError(t, err)

			msgHash, err := cr.HashMessage(dataToBeSigned)
			assert.NoError(t, err)

			assert.NoError(t, rsa.VerifyPSS(keyPair.Public, crypto.SHA256, msgHash, signature, nil))
		})
	})

}
