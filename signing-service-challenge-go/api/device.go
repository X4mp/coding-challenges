package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/X4mp/coding-challenges/signing-service-challenge/domain"
)

var (
	ErrInvalidCreateRequest error = fmt.Errorf("invalid create-signature-device request format")
)

// CreateSignatureDeviceRequest data model for parsing the request body
type CreateSignatureDeviceRequest struct {
	Label     string `json:"label"`
	Algorithm string `json:"algorithm"`
}

type CreateSignatureDeviceResponse struct {
	Id        string `json:"id"`
	Label     string `json:"label"`
	Algorithm string `json:"algorithm"`
	Counter   uint64 `json:"signing_counter"`
}

func (s *Server) CreateSignatureDevice(response http.ResponseWriter, request *http.Request) {
	createSignatureDeviceRequest := CreateSignatureDeviceRequest{}

	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(&createSignatureDeviceRequest)
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{ErrInvalidCreateRequest.Error()})
		return
	}

	signatureDevice, err := domain.NewSignatureDevice(createSignatureDeviceRequest.Label, createSignatureDeviceRequest.Algorithm)
	if err != nil && err == domain.ErrInvalidAlgorithm {
		WriteErrorResponse(response, http.StatusBadRequest, []string{err.Error()})
		return
	} else if err != nil {
		WriteInternalError(response)
		return
	}

	s.database.StoreSignatureDevice(&signatureDevice)

	deviceResponse := &CreateSignatureDeviceResponse{
		Id:        signatureDevice.DeviceId.String(),
		Label:     signatureDevice.Label,
		Algorithm: signatureDevice.Algorithm,
		Counter:   0,
	}

	WriteAPIResponse(response, 200, deviceResponse)
}
