package api

import (
	"encoding/json"
	"net/http"

	"github.com/X4mp/coding-challenges/signing-service-challenge/persistence"
	"github.com/gorilla/mux"
)

// Response is the generic API response container.
type Response struct {
	Data interface{} `json:"data"`
}

// ErrorResponse is the generic error API response container.
type ErrorResponse struct {
	Errors []string `json:"errors"`
}

// Server manages HTTP requests and dispatches them to the appropriate services.
type Server struct {
	listenAddress string
	database *persistence.Database
}

// NewServer is a factory to instantiate a new Server.
func NewServer(listenAddress string) *Server {
	return &Server{
		listenAddress: listenAddress,
		database: persistence.NewDatabase(),
		// TODO: add services / further dependencies here ...
	}
}

// Run registers all HandlerFuncs for the existing HTTP routes and starts the Server.
func (s *Server) Run() error {
	httpMux := mux.NewRouter()

	httpMux.Handle("/api/v0/health", http.HandlerFunc(s.Health))
	httpMux.Handle("/api/v0/device", http.HandlerFunc(s.CreateSignatureDevice))
	httpMux.Handle("/api/v0/device/{deviceUUID}/sign", http.HandlerFunc(s.SignTransaction))
	// TODO: register further HandlerFuncs here ...

	return http.ListenAndServe(s.listenAddress, httpMux)
}

// WriteInternalError writes a default internal error message as an HTTP response.
func WriteInternalError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

// WriteErrorResponse takes an HTTP status code and a slice of errors
// and writes those as an HTTP error response in a structured format.
func WriteErrorResponse(w http.ResponseWriter, code int, errors []string) {
	w.WriteHeader(code)

	errorResponse := ErrorResponse{
		Errors: errors,
	}

	bytes, err := json.Marshal(errorResponse)
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}

// WriteAPIResponse takes an HTTP status code and a generic data struct
// and writes those as an HTTP response in a structured format.
func WriteAPIResponse(w http.ResponseWriter, code int, data interface{}) {
	w.WriteHeader(code)

	response := Response{
		Data: data,
	}

	bytes, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		WriteInternalError(w)
	}

	w.Write(bytes)
}
