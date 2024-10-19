package repository

import (
	"errors"
	"github.com/product-manager/model"
	"golang.org/x/exp/slog"
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

type ProductRepository struct {
	Repository *repo
	Logger     *slog.Logger
}

func (p ProductRepository) ReadAllProducts() ([]model.Product, error) {
	var products []model.Product

	result := p.Repository.Model(products).
		Preload("Promotion").
		Find(&products).
		Order("name ASC")

	if result.Error != nil {
		p.Logger.Error("Repository: failed to read all products on database.", "error", result.Error.Error())
		return nil, errors.New("Failed to read all products on database.")
	}

	return products, nil
}

func (p ProductRepository) ReadProductByID(id int) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) ReadProductByName(name string) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) ReadProductsByPromotion(promotionName string) ([]model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) ReadProductsByFilter(filter string) ([]model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) CreateProduct(product *model.Product) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) UpdateProduct(product *model.Product) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (p ProductRepository) DeleteProductByID(id int) error {
	//TODO implement me
	panic("implement me")
}
