package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ErrTypeError = "ERROR"
)

type ErrorModel struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Code    int    `json:"code"`
}

func CommonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func WriteErrModel(w http.ResponseWriter, statusCode int, errModel *ErrorModel) {
	jsonStr, err := json.Marshal(errModel)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	fmt.Fprint(w, string(jsonStr))
}

func NewErrorResponse(msg string) *ErrorModel {
	return &ErrorModel{
		Message: msg,
		Type:    ErrTypeError,
	}
}
