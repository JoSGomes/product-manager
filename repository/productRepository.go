package repository

import (
	"errors"
	"github.com/product-manager/model"
)

type IProductRepository interface {
	ReadAllProducts() ([]model.Product, error)
	ReadProductByID(id int) (*model.Product, error)
	ReadProductByName(name string) (*model.Product, error)
	ReadProductsByPromotion(promotionName string) ([]model.Product, error)
	ReadProductsByFilter(filter string) ([]model.Product, error)

	CreateProduct(*model.Product) (*model.Product, error)
	UpdateProduct(*model.Product) (*model.Product, error)

	DeleteProductByID(id int) error
}

func (r *repo) ReadAllProducts() ([]model.Product, error) {
	var products []model.Product

	result := r.Model(products).
		Find(&products).
		Order("name ASC")

	if result.Error != nil {
		r.Logger.Error("Repository: failed to read all products on database.", "error", result.Error.Error())
		return nil, errors.New("Failed to read all products on database.")
	}

	return products, nil
}
