package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/kelseyhightower/envconfig"
	"github.com/product-manager/repository"
	"github.com/product-manager/settings"
	"golang.org/x/exp/slog"
	_ "gorm.io/gorm/logger"
	"log"
	"net/http"
	"os"
)

func init() {
	handler := &slog.HandlerOptions{}
	jsonHandler := slog.NewJSONHandler(os.Stdout, handler)
	l := slog.New(jsonHandler)

	l.Info("Product server is initiating...", slog.Int("version", 1.0)) // <-

	var dbConfig settings.Database
	err := envconfig.Process("DATABASE", &dbConfig)
	if err != nil {
		l.Error("Error initializing environment database variables", "error", err)
	}

	err = repository.MustInit(l, dbConfig)
	if err != nil {
		l.Error("Error initializing database", "error", err)
	}
}

func main() {
	var sttngs settings.Settings
	handler := &slog.HandlerOptions{}
	jsonHandler := slog.NewJSONHandler(os.Stdout, handler)
	l := slog.New(jsonHandler)

	err := envconfig.Process("server", &sttngs)
	if err != nil {
		l.Error("Error initializing server settings", "error", err)
	}

	//Initialize repositories, services and controllers
	r := chi.NewRouter()

	r.Route(fmt.Sprintf("/%s", sttngs.Server.Context), func(r chi.Router) {
		r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			response, _ := json.Marshal("ok")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				l.Error("Failed to generate JSON response", "error", err)
				return
			}

			w.WriteHeader(http.StatusOK)
			w.Write(response)
		})
	})

	address := fmt.Sprintf("%s:%s", sttngs.Server.Host, sttngs.Server.Port)
	l.Info(fmt.Sprintf("Starting server on %s %s", address, sttngs.Server.Context))
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
