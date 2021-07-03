package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fake-btc-markets/pkg/log"
	"fake-btc-markets/pkg/order"
)

func placeOrder(_ http.ResponseWriter, r *http.Request) response {
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	var request order.Order
	err = json.Unmarshal(body, &request)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	o, err := order.NewOrder(request)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	return newResponseWithStatusOK(o)
}
