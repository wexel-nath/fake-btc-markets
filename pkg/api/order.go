package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"fake-btc-markets/pkg/log"
	"fake-btc-markets/pkg/order"

	"github.com/go-chi/chi"
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

func getOrderByID(_ http.ResponseWriter, r *http.Request) response {
	orderID := chi.URLParam(r, "orderID")

	o, err := order.GetOrderByID(orderID)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	return newResponseWithStatusOK(o)
}

func getOrders(_ http.ResponseWriter, _ *http.Request) response {
	orders, err := order.GetOrders()
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	return newResponseWithStatusOK(orders)
}

func cancelOrder(_ http.ResponseWriter, r *http.Request) response {
	orderID := chi.URLParam(r, "orderID")

	cancelledOrder, err := order.CancelOrder(orderID)
	if err != nil {
		log.Error(err)
		meta := newMeta(err.Error())
		return newResponse(nil, meta, http.StatusBadRequest)
	}

	return newResponseWithStatusOK(cancelledOrder)
}
