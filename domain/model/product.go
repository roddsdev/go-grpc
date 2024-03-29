package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type ProductRepositoryInterface interface {
	AddProduct(product *Product) (*Product, error)
	FindProduct(id int32) (*Product, error)
	ShowAll() ([]*Product, error)
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Product struct {
	ID          int32     `json:"id" gorm:"primary_key;auto_increment;not null" valid:"-"`
	Name        string    `json:"name" gorm:"type:varchar(255);not null" valid:"notnull"`
	Description string    `json:"description" gorm:"type:varchar(255)" valid:"-"`
	Price       float32   `json:"number" gorm:"type:decimal(10,6);default:0" valid:"notnull"`
	CreatedAt   time.Time `json:"created_at" valid:"-"`
	UpdatedAt   time.Time `json:"updated_at" valid:"-"`
}

func (product *Product) isValid() error {
	_, err := govalidator.ValidateStruct(product)
	if err != nil {
		return err
	}
	return nil
}

func NewProduct(name string, description string, price float32) (*Product, error) {
	product := Product{
		Name:        name,
		Description: description,
		Price:       price,
	}

	// product.ID = uuid.NewV4().String()
	product.CreatedAt = time.Now()

	err := product.isValid()
	if err != nil {
		return nil, err
	}
	return &product, nil
}
