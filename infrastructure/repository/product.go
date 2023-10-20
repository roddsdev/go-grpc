package repository

import (
	"fmt"

	"example/grpc/domain/model"

	"github.com/jinzhu/gorm"
)

type ProductRepositoryDb struct {
	Db *gorm.DB
}

func (r ProductRepositoryDb) AddProduct(product *model.Product) (*model.Product, error) {
	err := r.Db.Create(product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (r ProductRepositoryDb) FindProduct(id string) (*model.Product, error) {
	var product model.Product
	r.Db.First(&product, "id = ?", id)

	if product.ID == "" {
		return nil, fmt.Errorf("no product found")
	}
	return &product, nil
}

func (r ProductRepositoryDb) ShowAll() ([]*model.Product, error) {
	var products []*model.Product
	r.Db.Find(&products)

	if len(products) == 0 {
		return nil, fmt.Errorf("no product found")
	}
	return products, nil
}
