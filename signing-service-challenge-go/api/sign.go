package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	ErrNoUUID             error = fmt.Errorf("device-id is not a valid UUID")
	ErrInvalidSignRequest error = fmt.Errorf("ivalid sign-transaction request format")
)

type SignTransactionRequest struct {
	DeviceID string `json:"deviceId"`
	Data     string `json:"data"`
}

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	deviceIDString := vars["deviceUUID"]
	deviceUUID, err := uuid.Parse(deviceIDString)
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{ErrNoUUID.Error()})
		return
	}

	// Like a transaction
	unlockFunc := s.database.LockDevice(deviceUUID)
	defer unlockFunc()

	device, err := s.database.GetSignatureDevice(deviceUUID)
	if err != nil {
		WriteErrorResponse(response, http.StatusNotFound, []string{err.Error()})
		return
	}

	signTransactionRequest := SignTransactionRequest{}
	decoder := json.NewDecoder(request.Body)
	err = decoder.Decode(&signTransactionRequest)
	if err != nil {
		WriteErrorResponse(response, http.StatusBadRequest, []string{ErrInvalidSignRequest.Error()})
		return
	}

	signatureResponse, err := device.Sign(signTransactionRequest.Data)
	if err != nil {
		WriteInternalError(response)
		return
	}

	s.database.StoreSignatureDevice(device)
	WriteAPIResponse(response, 200, signatureResponse)
}
