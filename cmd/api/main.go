package main

import (
	"fake-btc-markets/pkg/cron"
	"net/http"

	"fake-btc-markets/pkg/api"
	"fake-btc-markets/pkg/config"
	"fake-btc-markets/pkg/database"
	"fake-btc-markets/pkg/log"
)

func main() {
	err := config.Configure()
	if err != nil {
		log.Fatal(err)
	}

	_, err = database.Connect()
	if err != nil {
		log.Error(err)
	}

	// run once on startup (for now)
	cron.DoGetHistoricalData()

	startServer()
}

func startServer() {
	address := config.Get().GetListenAddress()
	log.Info("Listening on %s", address)
	log.Fatal(http.ListenAndServe(address, api.GetRouter()))
}
