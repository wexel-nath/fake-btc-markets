package api

import (
	"net/http"

	"fake-btc-markets/pkg/log"
	"fake-btc-markets/pkg/market"

	"github.com/go-chi/chi"
)

const (
	movingAverageParam = "ma"
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
	queryParams := r.URL.Query()
	maNames := queryParams[movingAverageParam]

	timestamp := getTimestampFromRequest(r)
	ticker, err := market.GetTickerForTimestamp(marketID, timestamp, maNames)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusInternalServerError)
	}

	return newResponseWithStatusOK(ticker)
}
