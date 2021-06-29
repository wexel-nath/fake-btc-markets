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
	ticker, err := doGetMarketTicker(r)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusInternalServerError)
	}

	return newResponseWithStatusOK(ticker)
}

func doGetMarketTicker(r *http.Request) (market.Ticker, error) {
	marketID := chi.URLParam(r, "marketID")
	timestampString := r.URL.Query().Get("timestamp")

	if timestampString == "" {
		return market.GetLatestTicker(marketID)
	}

	timestamp, err := time.Parse(time.RFC3339, timestampString)
	if err != nil {
		return market.Ticker{}, err
	}

	return market.GetTickerForTimestamp(marketID, timestamp)
}
