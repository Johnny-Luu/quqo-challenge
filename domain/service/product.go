package service

import "quqo_challenge/domain/entity"

type ProductService interface {
	GetAllProductsService() ([]entity.Product, error)
	GetProductByIdService(id int64) (*entity.Product, error)
	CreateProductsService([]entity.Product) ([]entity.Product, error)
	UpdateProductService(entity.Product) (*entity.Product, error)
	DeleteProductService(id int64) error
	SearchProductService(key string) ([]entity.Product, error)
}
