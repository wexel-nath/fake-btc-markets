package main

import (
	"net/http"

	"fake-btc-markets/pkg/api"
	"fake-btc-markets/pkg/config"
	"fake-btc-markets/pkg/log"
)

func main() {
	config.Configure()

	startServer()
}

func startServer() {
	address := config.Get().GetListenAddress()
	log.Info("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, api.GetRouter()))
}
