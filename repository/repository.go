package repository

import (
	"golang.org/x/exp/slog"
	"gorm.io/gorm"
)

type IRepository interface {
	IProductRepository
}

type repo struct {
	*gorm.DB
	Logger *slog.Logger
}
