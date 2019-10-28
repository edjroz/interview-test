package rpc

// "JSONResponse" is a metadata and data response in JSON format.
type JSONResponse struct {
	Data string `json:"data"`
}

// "JSONErrorResponse" is an error response in JSON format.
type JSONErrorResponse struct {
	Error *APIError `json:"error"`
}

// "APIError" is an error feedback structure containing a title and a status.
type APIError struct {
	Status int    `json:"code"`
	Title  string `json:"title"`
}
