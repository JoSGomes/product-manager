package repository

import (
	"fmt"
	"github.com/product-manager/settings"
	"golang.org/x/exp/slog"
	"gorm.io/driver/postgres"
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

	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Name,
	)

	var err error
	Repo.DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}))
	if err != nil {
		Repo.Logger.Error("Failed to connect to database", "error", err)
		return err
	}

	return nil
}

func GetRepository() *gorm.DB {
	return Repo.DB
}
