package api

import (
	"net/http"

	"fake-btc-markets/pkg/config"
)

type healthResponse struct{
	Status string `json:"status"`
	Image  string `json:"image"`
}

func health(w http.ResponseWriter, _ *http.Request) {
	data := healthResponse{
		Status: "ok",
		Image:  config.Get().ImageTag,
	}

	jsonResponse(w, http.StatusOK, data)
}
