package api

import (
	"encoding/json"
	"net/http"

	"fake-btc-markets/pkg/log"
)

type response struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func jsonResponse(w http.ResponseWriter, status int, data interface{}) {
	r := response{
		Data: data,
		Meta: nil,
	}

	resp, err := json.Marshal(r)
	if err != nil {
		log.Error(err)

		status = http.StatusInternalServerError
		resp = []byte(`{}`)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_, _ = w.Write(resp)
}
