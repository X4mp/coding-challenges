package domain

import (
	b64 "encoding/base64"
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

	counter       uint64
	lastSignature string
}

func NewSignatureDevice(label, algorithm string) (device SignatureDevice, err error) {
	deviceID := uuid.New()
	initLastSignature := b64.StdEncoding.EncodeToString([]byte(deviceID.String()))
	device = SignatureDevice{
		DeviceId:      uuid.New(),
		Label:         label,
		Algorithm:     algorithm,
		lastSignature: initLastSignature,
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
	Signature      string `json:"signature"`
	DataToBeSigned string `json:"signed_data"`
}

// Sign takes the user-data to be signed as input, generates the signature based on the algorithm of the device and returns the
// base64 encoded signature together with the concrete input string
func (d SignatureDevice) Sign(dataToBeSigned string) (response *SignatureResponse, err error) {

	var signer crypto.Signer
	switch d.Algorithm {
	case "ECC":
		signer = crypto.NewECCSigner(d.KeyPairECC.Private)
	case "RSA":
		signer = crypto.NewRSASigner(d.KeyPairRSA.Private)
	default:
		err = ErrInvalidAlgorithm
		return
	}

	input := fmt.Sprintf("%d_%s_%s", d.counter, dataToBeSigned, d.lastSignature)
	signature, err := signer.Sign([]byte(input))
	if err != nil {
		return
	}

	d.lastSignature = b64.StdEncoding.EncodeToString(signature)

	response = &SignatureResponse{
		Signature:      d.lastSignature,
		DataToBeSigned: input,
	}
	return
}

func (d SignatureDevice) Verify(data, signature string) bool {
	var err error
	var verifier crypto.Verifier
	switch d.Algorithm {
	case "ECC":
		verifier = &crypto.RSAVerifier{PublicKeyRSA: d.KeyPairRSA.Public}
	case "RSA":
		// TO BE IMPLEMENTED
	default:
		err = ErrInvalidAlgorithm
		return false
	}

	err = verifier.VerifySignature([]byte(data), []byte(signature))
	return err == nil
}
