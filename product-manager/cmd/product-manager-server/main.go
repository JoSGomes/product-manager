package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/product-manager/settings"
	"golang.org/x/exp/slog"
	"os"
)

func init() {

}

func main() {
	handler := &slog.HandlerOptions{}
	jsonHandler := slog.NewJSONHandler(os.Stdout, handler)
	logger := slog.New(jsonHandler)

	logger.Info("Product server is initiating...", slog.Int("version", 1.0)) // <-

	r := chi.NewRouter()
	sttngs := settings.Settings{}

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", err)
	}

	r.Route(fmt.Sprintf("%s", sttngs.Server.Context), func(r chi.Router) {

	})
}
