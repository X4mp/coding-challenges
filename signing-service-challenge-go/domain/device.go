package domain

import "github.com/google/uuid"

// TODO: signature device domain model ...

type SignatureDevice struct {
	label string
	id uuid.UUID
}

// CreateSignatureDeviceResponse data model
type CreateSignatureDeviceResponse struct {

}

// CreateSignatureDevice(id: string, algorithm: 'ECC' | 'RSA', [optional]: label: string): CreateSignatureDeviceResponse
func (d *SignatureDevice) Create(id string, algorithm string, label string) *CreateSignatureDeviceResponse {

	return &CreateSignatureDeviceResponse{}
}

// SignatureResponse data model
type SignatureResponse struct {

}

// SignTransaction(deviceId: string, data: string): SignatureResponse
func (d *SignatureDevice) Sign(deviceId string, data string) *SignatureResponse {

	return &SignatureResponse{}
}