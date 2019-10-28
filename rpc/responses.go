package rpc

import (
	"encoding/json"
	"net/http"
)

// ignore rpc error handling
// "WriteJSONResponse" writes a JSON response.
func WriteJSONResponse(w http.ResponseWriter, b []byte) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(b)
}

// ignore rpc error handling
// "WriteErrorResponse" writes an error JSON response.
func WriteErrorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)
	_ = json.NewEncoder(w).Encode(&JSONErrorResponse{Error: &APIError{Status: errorCode, Title: errorMsg}})
}
