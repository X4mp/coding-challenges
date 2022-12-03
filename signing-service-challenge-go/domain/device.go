package domain

import (
	"fmt"

	"github.com/X4mp/coding-challenges/signing-service-challenge/crypto"
	"github.com/google/uuid"
)

// TODO: signature device domain model ...
var ErrInvalidAlgorithm error = fmt.Errorf("invalid algorithm choosen")

type SignatureDevice struct {
	Label      string
	DeviceId   uuid.UUID
	Algorithm  string
	KeyPairRSA *crypto.RSAKeyPair
	KeyPairECC *crypto.ECCKeyPair
	
	counter uint64
	lastSignature string
}

func NewSignatureDevice(label, algorithm string) (device SignatureDevice, err error) {
	device = SignatureDevice{
		DeviceId:  uuid.New(),
		Label:     label,
		Algorithm: algorithm,
	}

	switch algorithm {
	case "RSA":
		generator := crypto.RSAGenerator{}
		device.KeyPairRSA, err = generator.Generate()
	case "ECC":
		generator := crypto.ECCGenerator{}
		device.KeyPairECC, err = generator.Generate()
	default:
		err = ErrInvalidAlgorithm
	}

	return
}

// SignatureResponse data model
type SignatureResponse struct {
	signature []byte
}

// SignTransaction(deviceId: string, data: string): SignatureResponse
func (d SignatureDevice) Sign(deviceId string, dataToBeSigned string) (response *SignatureResponse, err error) {

	var signer crypto.Signer
	switch (d.Algorithm) {
	case "ECC":
		signer = crypto.NewECCSigner(d.KeyPairECC.Private)
	case "RSA": 
		signer = crypto.NewRSASigner(d.KeyPairRSA.Private)
	default:
		err = ErrInvalidAlgorithm
		return
	}

	signature, err := signer.Sign([]byte(dataToBeSigned))
	if err != nil {
		return
	}

	response = &SignatureResponse{
		signature: signature,
	}
	return
}
