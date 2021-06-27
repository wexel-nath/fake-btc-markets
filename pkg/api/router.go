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
	r.Get("/health", Handle(health))

	r.Route("/v3", func(r chi.Router) {
		r.Route("/markets", marketRoutes)
	})
}

// marketRoutes base path: /v3/markets
func marketRoutes(r chi.Router) {
	r.Get("/", Handle(getMarkets))
	r.Get("/tickers", Handle(unsupported))
	r.Get("/orderbooks", Handle(unsupported))

	r.Route("/{marketID}", func(r chi.Router) {
		r.Get("/", Handle(getMarketByID))
		r.Get("/ticker", Handle(getMarketTicker))
		r.Get("/trades", Handle(unsupported))
		r.Get("/orderbook", Handle(unsupported))
		r.Get("/candles", Handle(unsupported))
	})
}
