package api

import (
	"net/http"
	"time"

	"fake-btc-markets/pkg/log"
	"fake-btc-markets/pkg/market"

	"github.com/go-chi/chi"
)

func getMarketByID(_ http.ResponseWriter, r *http.Request) response {
	marketID := chi.URLParam(r, "marketID")

	m, err := market.GetMarketByID(marketID)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusInternalServerError)
	}

	return newResponseWithStatusOK(m)
}

func getMarkets(_ http.ResponseWriter, _ *http.Request) response {
	m, err := market.GetMarkets()
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusInternalServerError)
	}

	return newResponseWithStatusOK(m)
}

func getMarketTicker(_ http.ResponseWriter, r *http.Request) response {
	marketID := chi.URLParam(r, "marketID")

	timestamp, err := time.Parse(time.RFC3339, r.URL.Query().Get("timestamp"))
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusInternalServerError)
	}

	ticker, err := market.GetTickerForTimestamp(marketID, timestamp)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusInternalServerError)
	}

	return newResponseWithStatusOK(ticker)
}