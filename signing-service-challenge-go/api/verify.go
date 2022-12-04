package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	ErrInvalidVerifyRequest error = fmt.Errorf("invalid signature-verification request format")
)

type VerifySignatureRequest struct {
	DeviceId  string `json:"deviceId"`
	Signature string `json:"signature"`
}

func (s *Server) VerifySignature(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	deviceIDString := vars["deviceUUID"]
	deviceUUID, err := uuid.Parse(deviceIDString)
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{ErrNoUUID.Error()})
		return
	}

	device, err := s.database.GetSignatureDevice(deviceUUID)
	if err != nil {
		WriteErrorResponse(response, http.StatusNotFound, []string{err.Error()})
		return
	}

	verifySignatureRequest := VerifySignatureRequest{}
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&verifySignatureRequest)
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{ErrInvalidVerifyRequest.Error()})
		return
	}
}
