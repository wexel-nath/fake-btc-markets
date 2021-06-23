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
	r.Get("/health", health)
}
