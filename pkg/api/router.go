package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

func GetRouter() chi.Router {
	router := chi.NewRouter()

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders: []string{"Authorization", "Content-Type"},
	})
	router.Use(c.Handler)

	addRoutes(router)
	return router
}

func addRoutes(r chi.Router) {
	r.Get("/health", handle(health))

	r.Route("/v3", func(r chi.Router) {
		r.Route("/markets", marketRoutes)
		r.Route("/orders", orderRoutes)
		r.Route("/trades", tradeRoutes)
	})
}

// marketRoutes base path: /v3/markets
func marketRoutes(r chi.Router) {
	r.Get("/", handle(getMarkets))
	r.Get("/tickers", handle(unsupported))
	r.Get("/orderbooks", handle(unsupported))

	r.Route("/{marketID}", func(r chi.Router) {
		r.Get("/", handle(getMarketByID))
		r.Get("/ticker", handle(getMarketTicker))
		r.Get("/trades", handle(unsupported))
		r.Get("/orderbook", handle(unsupported))
		r.Get("/candles", handle(unsupported))
	})
}

// orderRoutes base path: /v3/orders
func orderRoutes(r chi.Router) {
	r.Post("/", handle(placeOrder))
	r.Get("/", handle(getOrders))
	r.Delete("/", handle(unsupported))

	r.Route("/{orderID}", func(r chi.Router) {
		r.Get("/", handle(getOrderByID))
		r.Delete("/", handle(cancelOrder))
		r.Put("/", handle(unsupported))
	})
}

// tradeRoutes base path: /v3/trades
func tradeRoutes(r chi.Router) {
	r.Get("/", handle(getTrades))
	r.Get("/{tradeID}", handle(unsupported))
}
