package api

import (
	"net/http"

	"fake-btc-markets/pkg/log"
	"fake-btc-markets/pkg/order"
)

func getTrades(_ http.ResponseWriter, r *http.Request) response {
	orderID := r.URL.Query().Get("orderId")

	trades, err := order.GetTradesByOrderID(orderID)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	return newResponseWithStatusOK(trades)
}
