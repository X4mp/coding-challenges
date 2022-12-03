package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	ErrNoUUID error = fmt.Errorf("device-id is not a valid UUID")
)

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
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

	device.Sign(deviceIDString, "")
	WriteAPIResponse(response, 200, nil)
}
