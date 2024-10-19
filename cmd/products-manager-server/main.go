package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/kelseyhightower/envconfig"
	"github.com/product-manager/controller"
	"github.com/product-manager/middleware"
	"github.com/product-manager/repository"
	"github.com/product-manager/service"
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
		log.Fatal("Error initializing environment database variables ", "error", err)
	}

	err = repository.MustInit(l, dbConfig)
	if err != nil {
		log.Fatalf("Error initializing environment database variables ", "error", err)
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

	healthController := controller.HealthController{Logger: l}

	productRepository := repository.ProductRepository{
		Repository: repository.Repo,
		Logger:     l,
	}

	productService := service.ProductService{
		ProductRepository: productRepository,
		Logger:            l,
	}

	productController := controller.ProductController{
		ProductService: &productService,
		Logger:         l,
	}

	//Initialize repositories, services and controllers
	r := chi.NewRouter()
	r.Use(middleware.LoggingMiddleware)

	r.Route(fmt.Sprintf("/%s", sttngs.Server.Context), func(r chi.Router) {
		r.Get("/health", healthController.Health)
		r.Get("/products", productController.ReadAllProducts)
	})

	address := fmt.Sprintf("%s:%s", sttngs.Server.Host, sttngs.Server.Port)
	l.Info(fmt.Sprintf("Starting server on %s %s", address, sttngs.Server.Context))
	if err := http.ListenAndServe(address, r); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
