package request

import (
	"encoding/json"
	"net/http"

	"fake-btc-markets/pkg/log"
)

type Meta struct {
	Message string `json:"message"`
}

type Response struct {
	Data       interface{} `json:"data"`
	Meta       *Meta       `json:"meta"`
	StatusCode int
}

type Handler func(w http.ResponseWriter, r *http.Request) Response

func Handle(handler Handler) func(http.ResponseWriter, *http.Request) {
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
		_, _ = w.Write(body)
	}
}
