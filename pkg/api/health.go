package api

import (
	"net/http"

	"fake-btc-markets/pkg/config"
)

type healthResponse struct{
	Status  string `json:"status"`
	Version string `json:"version"`
}

func health(_ http.ResponseWriter, _ *http.Request) response {
	data := healthResponse{
		Status:  "ok",
		Version: config.Get().Version,
	}

	return newResponseWithStatusOK(data)
}
