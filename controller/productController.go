package controller

import (
	"encoding/json"
	"github.com/product-manager/model"
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

func (c *controller) ReadAllProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	c.logger.InfoContext(r.Context(), "Read")
	lggr := c.logger.With("Controller")

	lggr.Info("Request on read all products", "request", r)

	products, err := c.service.ReadAllProducts()
	if err != nil {
		lggr.Error("Failed on read all products", "error", err)
		http.Error(w, "Failed on read all products", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(products); err != nil {
		lggr.Error("Failed on marshall all products", "error", err)
		http.Error(w, "Failed on marshall all products", http.StatusInternalServerError)
		return
	}
}
