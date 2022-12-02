package api

import "net/http"

func (s *Server) SignTransaction(response http.ResponseWriter, request *http.Request) {
	// STUB API handler, NEEDS IMPLEMENTATION
	WriteAPIResponse(response, 200, nil)
}
