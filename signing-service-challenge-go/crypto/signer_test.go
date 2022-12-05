package crypto_test

import (
	"crypto"
	"crypto/ecdsa"
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
		t.Run("IncorrectInput", func(t *testing.T) {
			keyPair, err := (&cr.RSAGenerator{}).Generate()
			assert.NoError(t, err)

			signer := cr.NewRSASigner(keyPair.Private)

			signature, err := signer.Sign(dataToBeSigned)
			assert.NoError(t, err)

			incorectInput := "This is not the correct input."
			msgHash, err := cr.HashMessage([]byte(incorectInput))
			assert.NoError(t, err)

			assert.EqualError(t, rsa.VerifyPSS(keyPair.Public, crypto.SHA256, msgHash, signature, nil), rsa.ErrVerification.Error())
		})
	})

	t.Run("ECC", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			keyPair, err := (&cr.ECCGenerator{}).Generate()
			assert.NoError(t, err)

			signer := cr.NewECCSigner(keyPair.Private)

			signature, err := signer.Sign(dataToBeSigned)
			assert.NoError(t, err)

			msgHash, err := cr.HashMessage(dataToBeSigned)
			assert.NoError(t, err)

			assert.True(t, ecdsa.VerifyASN1(keyPair.Public, msgHash, signature))
		})

		t.Run("IncorrectInput", func(t *testing.T) {
			keyPair, err := (&cr.ECCGenerator{}).Generate()
			assert.NoError(t, err)

			signer := cr.NewECCSigner(keyPair.Private)

			signature, err := signer.Sign(dataToBeSigned)
			assert.NoError(t, err)

			incorectInput := "This is not the correct input."
			msgHash, err := cr.HashMessage([]byte(incorectInput))
			assert.NoError(t, err)

			assert.False(t, ecdsa.VerifyASN1(keyPair.Public, msgHash, signature))
		})
	})

}
