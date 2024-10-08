package repository

import (
	"fmt"
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

func (r *repo) MustInit() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/dbname?charset=utf8mb4&parseTime=True&loc=Local")

}
