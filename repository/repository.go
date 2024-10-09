package repository

import (
	"fmt"
	"github.com/product-manager/settings"
	"golang.org/x/exp/slog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type IRepository interface {
	IProductRepository
}

type repo struct {
	*gorm.DB
	Logger *slog.Logger
}

var Repo *repo

func MustInit(logger *slog.Logger, dbConfig settings.Database) error {
	Repo = &repo{
		Logger: logger,
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	var err error
	Repo.DB, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		Repo.Logger.Error("Failed to connect to database", "error", err)
		return err
	}

	return nil
}
