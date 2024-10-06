package service

import "github.com/product-manager/model"

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

func (s *service) ReadAllProducts() ([]model.Product, error) {
	results, err := s.ReadAllProducts()
	if err != nil {
		s.logger.Error("Service: failed read all products", "err", err)
		return nil, err
	}

	return results, nil
}
