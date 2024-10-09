package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
	"github.com/product-manager/repository"
	"github.com/product-manager/settings"
	"golang.org/x/exp/slog"
	"os"
)

func init() {
	handler := &slog.HandlerOptions{}
	jsonHandler := slog.NewJSONHandler(os.Stdout, handler)
	logger := slog.New(jsonHandler)

	logger.Info("Product server is initiating...", slog.Int("version", 1.0)) // <-

	if err := godotenv.Load(); err != nil {
		logger.Error("Error loading .env file", err)
	}

	dbConfig := settings.Database{
		Username: os.Getenv("DATABASE_USERNAME"),
		Password: os.Getenv("DATABASE_PASSWORD"),
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Name:     os.Getenv("DATABASE_NAME"),
	}

	err := repository.MustInit(logger, dbConfig)
	if err != nil {
		logger.Error("Error initializing database", "error", err)
	}

}

func main() {

	//Initialize repositories, servicies and controllers
	r := chi.NewRouter()
	sttngs := settings.Settings{}

	r.Route(fmt.Sprintf("%s", sttngs.Server.Context), func(r chi.Router) {

	})
}
