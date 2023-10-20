package usecase

import (
	"example/grpc/domain/model"
)

type ProductUseCase struct {
	ProductRepository model.ProductRepositoryInterface
}

func (p *ProductUseCase) Register(name string, description string, price float32) (*model.Product, error) {
	validatedProduct, err := model.NewProduct(name, description, price)
	if err != nil {
		return nil, err
	}

	newProduct, err := p.ProductRepository.AddProduct(validatedProduct)
	if err != nil {
		return nil, err
	}

	return newProduct, nil
}

func (p *ProductUseCase) FindByKey(key string) (*model.Product, error) {
	product, err := p.ProductRepository.FindProduct(key)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (p *ProductUseCase) FindAll() ([]*model.Product, error) {
	products, err := p.ProductRepository.ShowAll()
	if err != nil {
		return nil, err
	}
	return products, nil
}
