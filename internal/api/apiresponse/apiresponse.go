package apiresponse

import (
	"encoding/json"
	"log"
	"myapp/internal/apperrors"
	"net/http"
)

type ResponseBody struct {
	Success bool              `json:"success"`
	Errors  map[string]string `json:"errors"`
	Data    interface{}       `json:"data"`
}

type ResponseBodyNew struct {
	Errors []apperrors.Error `json:"errors"`
}

// ErrorResponse attaches Response Status line and WRITES into response, so don't change anything after that
// just return
func ErrorResponse(errors []apperrors.Error, statusCode int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	jsonBytes, err := json.Marshal(ResponseBodyNew{
		Errors: errors,
	})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Printf("error on building error response: %v", err)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}

func SuccessResponse(data interface{}, statusCode int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")
	var errors []apperrors.Error
	jsonBytes, err := json.Marshal(ResponseBody{
		Success: true,
		Errors:  nil,
		Data:    data,
	})

	// w.WriteHeader(statusCode)
	if err != nil {
		log.Printf("error on building success response: %v", err)
		errors = append(errors, apperrors.Error{
			ErrorCode: apperrors.ErrInternalServerError,
			ErrorMsg:  "Internal error",
		})
		ErrorResponse(errors, http.StatusInternalServerError, w)
		return
	}

	w.WriteHeader(statusCode)
	w.Write(jsonBytes)
}
