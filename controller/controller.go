package controller

import (
	"github.com/product-manager/service"
	"golang.org/x/exp/slog"
)

type IController interface {
	IProductController
}

type controller struct {
	service *service.IService
	logger  *slog.Logger
}
