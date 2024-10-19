package controller

import (
	"encoding/json"
	"github.com/product-manager/model"
	"github.com/product-manager/service"
	"golang.org/x/exp/slog"
	"net/http"
)

type IProductController interface {
	ReadAllProducts() ([]model.Product, error)
	ReadProductByID(id int) (*model.Product, error)
	ReadProductByName(name string) (*model.Product, error)
	ReadProductsByPromotion(promotionName string) ([]model.Product, error)
	ReadProductsByFilter(filter string) ([]model.Product, error)

	CreateProduct(*model.Product) (*model.Product, error)
	UpdateProduct(*model.Product) (*model.Product, error)

	DeleteProductByID(id int) error
}

type ProductController struct {
	ProductService service.IProductService
	Logger         *slog.Logger
}

func (c *ProductController) ReadAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	l := c.Logger.With("Controller", "Products")

	l.Info("Request on read all products",
		"method", r.Method,
		"url", r.URL.String(),
		"header", r.Header)

	products, err := c.ProductService.ReadAllProducts()
	if err != nil {
		l.Error("Failed on read all products", "error", err)
		http.Error(w, "Failed on read all products", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		l.Error("Failed on marshall all products", "error", err)
		http.Error(w, "Failed on marshall all products", http.StatusInternalServerError)
		return
	}
}
