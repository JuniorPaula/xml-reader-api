package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// JsonResponse is a struct that represents a JSON response.
type JsonResponse struct {
	Message string      `json:"message"`
	Error   bool        `json:"error"`
	Data    interface{} `json:"data,omitempty"`
}

// ReadJSON reads a JSON request and decodes it into the provided data structure.
// It also checks if the request body is larger than 1MB.
func ReadJSON(w http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1048576
	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	err := dec.Decode(data)
	if err != nil {
		return err
	}
	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must only have a single JSON value")
	}
	return nil
}

// WriteJSON writes a JSON response with the specified status code.
// It also sets the Content-Type header to "application/json".
func WriteJSON(w http.ResponseWriter, status int, data interface{}, headers ...http.Header) error {
	out, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for k, v := range headers[0] {
			w.Header()[k] = v
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, err = w.Write(out)
	if err != nil {
		return err
	}
	return nil
}

// ErrorJSON writes an error message to the response writer
// in JSON format with the specified status code.
func ErrorJSON(w http.ResponseWriter, err error, status ...int) error {
	statusCode := http.StatusBadRequest
	if len(status) > 0 {
		statusCode = status[0]
	}

	var payload JsonResponse
	payload.Message = err.Error()
	payload.Error = true

	return WriteJSON(w, statusCode, payload)
}
