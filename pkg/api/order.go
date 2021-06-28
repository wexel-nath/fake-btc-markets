package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fake-btc-markets/pkg/log"
	"fake-btc-markets/pkg/order"
)

type orderRequest struct{
	MarketID string `json:"marketId"`
	Price    string `json:"price"`
	Amount   string `json:"amount"`
	Type     string `json:"type"`
	Side     string `json:"side"`
}

func placeOrder(_ http.ResponseWriter, r *http.Request) response {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	var request orderRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	o, err := order.NewOrder(
		request.MarketID,
		request.Price,
		request.Amount,
		request.Type,
		request.Side,
	)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	return newResponseWithStatusOK(o)
}
