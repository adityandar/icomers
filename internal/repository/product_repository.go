package repository

import (
	"icomers/internal/entity"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ProductRepository struct {
	Repository[entity.Product]
	Log *logrus.Logger
}

func NewProductRepository(log *logrus.Logger) *ProductRepository {
	return &ProductRepository{
		Log: log,
	}
}

func (r *ProductRepository) FindAll(db *gorm.DB) ([]entity.Product, error) {
	var products []entity.Product

	results := db.Find(&products)

	if results.Error != nil {
		return nil, results.Error
	}

	return products, nil

}
