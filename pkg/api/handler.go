package api

import (
	"encoding/json"
	"net/http"

	"fake-btc-markets/pkg/log"
)

type meta struct{
	Message string `json:"message"`
}

func newMeta(message string) *meta {
	return &meta{
		Message: message,
	}
}

type response struct{
	Data       interface{} `json:"data"`
	Meta       *meta       `json:"meta"`
	StatusCode int         `json:"-"`
}

func newResponse(data interface{}, meta *meta, statusCode int) response {
	return response{
		Data:       data,
		Meta:       meta,
		StatusCode: statusCode,
	}
}

func newResponseWithStatusOK(data interface{}) response {
	return newResponse(data, nil, http.StatusOK)
}

type Handler func(w http.ResponseWriter, r *http.Request) response

func handle(handler Handler) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		response := handler(w, r)

		body, err := json.Marshal(response)
		if err != nil {
			log.Error(err)

			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(`{}`))
			return
		}

		w.WriteHeader(response.StatusCode)
		_, err = w.Write(body)
		if err != nil {
			log.Error(err)
		}
	}
}
