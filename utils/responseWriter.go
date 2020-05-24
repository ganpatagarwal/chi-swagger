package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type ErrorResponseModel struct {
	Error string `json:"error"`
}

// ResponseType defines what kind of a response is sent for the request received
type ResponseType int

const (
	// ErrorResponse - Error response type Code
	ErrorResponse ResponseType = 1
	// SuccessResponse - Success response type Code
	SuccessResponse ResponseType = 2
)

// WriteResponse - writes the response to the http request
func WriteResponse(w http.ResponseWriter, responseType ResponseType, responseBody interface{}) {

	w.Header().Add("Content-Type", "application/json")
	switch responseType {
	case ErrorResponse:
		if err, ok := responseBody.(error); ok {
			var response ErrorResponseModel
			response.Error = err.Error()

			w.WriteHeader(http.StatusInternalServerError)

			responseBytes, err := json.Marshal(response)

			if err != nil {
				fmt.Printf("Failed to marshal error response structure - %+v\n", err)
				return
			}

			_, err = w.Write(responseBytes)
			if err != nil {
				log.Printf("Error in writing response: %s", err)
			}

			return
		}
	case SuccessResponse:
		w.WriteHeader(http.StatusOK)

		responseBytes, err := json.Marshal(responseBody)

		if err != nil {
			fmt.Printf("Failed to marshal response structure - %+v\n", err)
			return
		}

		_, err = w.Write(responseBytes)
		if err != nil {
			log.Printf("Error in writing response: %s", err)
		}
	}
}
