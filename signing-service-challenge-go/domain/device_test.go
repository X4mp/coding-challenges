package domain_test

import (
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/X4mp/coding-challenges/signing-service-challenge/crypto"
	domain "github.com/X4mp/coding-challenges/signing-service-challenge/domain"
)

func TestSignatureDevice_New(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		device, err := domain.NewSignatureDevice("Device 1", "RSA")
		assert.NoError(t, err)
		assert.EqualValues(t, "Device 1", device.Label)
		assert.EqualValues(t, "RSA", device.Algorithm)

		device, err = domain.NewSignatureDevice("Device 2", "ECC")
		assert.NoError(t, err)
		assert.EqualValues(t, "Device 2", device.Label)
		assert.EqualValues(t, "ECC", device.Algorithm)
	})

	t.Run("Incorrect Algorithm", func(t *testing.T) {
		device, err := domain.NewSignatureDevice("Device 1", "DSA")
		assert.Nil(t, device)
		assert.EqualError(t, err, domain.ErrInvalidAlgorithm.Error())
	})
}

func TestSignatureDevice_Sign(t *testing.T) {
	dataToBeSigned := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."

	t.Run("RSA", func(t *testing.T) {
		t.Run("OK", func(t *testing.T) {
			device, err := domain.NewSignatureDevice("Device 1", "RSA")
			assert.NoError(t, err)

			signatureResponse, err := device.Sign(dataToBeSigned)
			assert.NoError(t, err)

			verifier := crypto.NewRSAVerifier(device.KeyPairRSA.Public)

			rawSignature, err := base64.StdEncoding.DecodeString(signatureResponse.Signature)
			assert.NoError(t, err)
			assert.NoError(t, verifier.VerifySignature([]byte(signatureResponse.DataToBeSigned), rawSignature))
		})
	})
}
