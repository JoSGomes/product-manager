package service

import (
	"github.com/product-manager/repository"
	"golang.org/x/exp/slog"
)

type IService interface {
	IProductService
}

type service struct {
	repository *repository.IRepository
	logger     *slog.Logger
}
