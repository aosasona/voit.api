package main

import (
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"voit.api/handler"
)

func main() {
	c := chi.NewRouter()

	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Use(cors.AllowAll().Handler)

	c.Get("/ping", handler.Ping)

	// API routes
	c.Route("/api", func(r chi.Router) {
		c.Get("/ping", handler.Ping)
		c.Route("/whisper", func(r chi.Router) {
			c.Get("/", handler.GetAllWhisperModels)
		})
	})

	port := "8080"
	if p, exists := os.LookupEnv("PORT"); exists {
		port = p
	}

	log.Infof("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, c))
}
