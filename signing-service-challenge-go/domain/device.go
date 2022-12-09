package domain

import (
	b64 "encoding/base64"
	"fmt"

	"github.com/X4mp/coding-challenges/signing-service-challenge/crypto"
	"github.com/google/uuid"
)

type SignatureDevice struct {
	Label     string
	DeviceId  uuid.UUID
	Algorithm string

	Signer   crypto.Signer   // exposed for testing purpose
	Verifier crypto.Verifier // exposed for testing purpose

	counter       uint64
	lastSignature string
}

func NewSignatureDevice(label, algorithm string) (device *SignatureDevice, err error) {
	deviceID := uuid.New()
	initLastSignature := b64.StdEncoding.EncodeToString([]byte(deviceID.String()))
	device = &SignatureDevice{
		DeviceId:      uuid.New(),
		Label:         label,
		Algorithm:     algorithm,
		lastSignature: initLastSignature,
	}

	device.Signer, device.Verifier, err = crypto.NewAbstractTools(algorithm)
	if err != nil {
		device = nil
		return
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
func (d *SignatureDevice) Sign(dataToBeSigned string) (response *SignatureResponse, err error) {
	input := fmt.Sprintf("%d_%s_%s", d.counter, dataToBeSigned, d.lastSignature)
	signature, err := d.Signer.Sign([]byte(input))
	if err != nil {
		return
	}

	d.lastSignature = b64.StdEncoding.EncodeToString(signature)
	d.counter += 1

	response = &SignatureResponse{
		Signature:      d.lastSignature,
		DataToBeSigned: input,
	}
	return
}

// Verify takes the the user-data to be signed togehter with the signature and verifies the signature over the data
func (d *SignatureDevice) Verify(data, signature string) bool {
	rawSignature, err := b64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Println("Verification Error: not base64 decodable signature")
		return false
	}

	err = d.Verifier.VerifySignature([]byte(data), rawSignature)
	if err != nil {
		fmt.Println("Verification Error: ", err.Error())
	}
	return err == nil
}
