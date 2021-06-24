package api

import (
	"net/http"

	"fake-btc-markets/pkg/config"
)

type healthResponse struct{
	Status  string `json:"status"`
	Version string `json:"version"`
}

func health(w http.ResponseWriter, _ *http.Request) {
	data := healthResponse{
		Status:  "ok",
		Version: config.Get().Version,
	}

	jsonResponse(w, http.StatusOK, data)
}
