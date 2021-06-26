package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"

	"fake-btc-markets/pkg/api/request"
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
	r.Get("/health", request.Handle(health))

	r.Route("/v3", func(r chi.Router) {
		r.Route("/markets", func(r chi.Router) {
			r.Get("/", request.Handle(unsupported))
			r.Get("/tickers", request.Handle(unsupported))
			r.Get("/orderbooks", request.Handle(unsupported))

			r.Route("/{marketID}", func(r chi.Router) {
				r.Get("/ticker", request.Handle(unsupported))
				r.Get("/trades", request.Handle(unsupported))
				r.Get("/orderbook", request.Handle(unsupported))
				r.Get("/candles", request.Handle(unsupported))
			})
		})
	})
}
