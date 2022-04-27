package deus_cc

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// ResponseAPI represents the data structure needed to create a response
type ResponseAPI struct {
	Status      int            `json:"status"`
	Description string         `json:"description,omitempty"`
	Success     bool           `json:"success"`
	Data        map[string]int `json:"data,omitempty"`
}

// response sets the params to generate a JSON response
func response(w http.ResponseWriter, ra ResponseAPI) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(ra.Status)

	json.NewEncoder(w).Encode(ra)
}

// responseError prints the error using log and returns a response
func responseError(w http.ResponseWriter, description string) {
	log.Println(description)

	ra := ResponseAPI{
		Status:      http.StatusInternalServerError,
		Description: description,
		Success:     false,
	}
	response(w, ra)
}

// responseNotAllowed calls response function with proper data to generate a Not Allowed response
func responseNotAllowed(w http.ResponseWriter, method string) {
	ra := ResponseAPI{
		Status:      http.StatusMethodNotAllowed,
		Description: fmt.Sprintf("%s method not allowed", method),
		Success:     false,
	}

	response(w, ra)
}

// responseBadRequest calls response function with proper data to generate a Bad Request response
func responseBadRequest(w http.ResponseWriter) {
	ra := ResponseAPI{
		Status:  http.StatusBadRequest,
		Success: false,
	}

	response(w, ra)
}

// responseUnprocessable calls response function with proper data to generate a Unprocessable Entity response
func responseUnprocessable(w http.ResponseWriter, method string) {
	ra := ResponseAPI{
		Status:      http.StatusUnprocessableEntity,
		Description: fmt.Sprintf("%s input data not valid, expected {'url':'[valid url]','uuid':'[valid uuid]',}", method),
		Success:     false,
	}

	response(w, ra)
}

// responseOk calls response function with proper data to generate an OK response
func responseOk(w http.ResponseWriter, success bool, visitors int) {
	ra := ResponseAPI{
		Status:  http.StatusOK,
		Success: success,
	}
	response(w, ra)
}

// responseVisitors calls response function to show the number of visitors for an incoming request
func responseVisitors(w http.ResponseWriter, success bool, visitors int) {
	ra := ResponseAPI{
		Status:  http.StatusOK,
		Success: success,
		Data: map[string]int{
			"unique_visitors": visitors,
		},
	}
	response(w, ra)
}
