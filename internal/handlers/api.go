package handlers

import (
	"github.com/NucleonGodX/Distributed-Task-Queue-System/goapi/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r * chi.Mux) {
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router, chi.Router){
		r.Use(middleware.Authorization)
		r.Get("/coins", GetCoinBalance)
	})
}