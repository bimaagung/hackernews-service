package response

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Status   	string  `json:"status"`
	Message 	string 	`json:"message"`
	Data    	any    	`json:"data,omitempty"`
}

func Success(w http.ResponseWriter, status int, data any, headers ...http.Header) error {

	dataResponse :=  &jsonResponse{
		Status: "ok",
		Message: "success",
		Data: data,
	}

	out, err := json.Marshal(dataResponse)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
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

func Error(w http.ResponseWriter, status int, message string, headers ...http.Header) error {

	dataResponse :=  &jsonResponse{
		Status: "failed",
		Message: message,
	}

	out, err := json.Marshal(dataResponse)
	if err != nil {
		return err
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			w.Header()[key] = value
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