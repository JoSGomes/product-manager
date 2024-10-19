package service

import (
	"github.com/product-manager/model"
	"github.com/product-manager/repository"
	"golang.org/x/exp/slog"
)

type IProductService interface {
	ReadAllProducts() ([]model.Product, error)
	ReadProductByID(id int) (*model.Product, error)
	ReadProductByName(name string) (*model.Product, error)
	ReadProductsByPromotion(promotionName string) ([]model.Product, error)
	ReadProductsByFilter(filter string) ([]model.Product, error)

	CreateProduct(*model.Product) (*model.Product, error)
	UpdateProduct(*model.Product) (*model.Product, error)

	DeleteProductByID(id int) error
}

type ProductService struct {
	ProductRepository repository.IProductRepository
	Logger            *slog.Logger
}

func (s *ProductService) ReadAllProducts() ([]model.Product, error) {
	results, err := s.ProductRepository.ReadAllProducts()
	if err != nil {
		s.Logger.Error("Service: failed read all products", "err", err)
		return nil, err
	}

	return results, nil
}

func (s *ProductService) ReadProductByID(id int) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductService) ReadProductByName(name string) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductService) ReadProductsByPromotion(promotionName string) ([]model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductService) ReadProductsByFilter(filter string) ([]model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductService) CreateProduct(product *model.Product) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductService) UpdateProduct(product *model.Product) (*model.Product, error) {
	//TODO implement me
	panic("implement me")
}

func (s *ProductService) DeleteProductByID(id int) error {
	//TODO implement me
	panic("implement me")
}
