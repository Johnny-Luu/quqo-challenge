package repository

import "quqo_challenge/domain/entity"

type ProductRepository interface {
	GetAllProducts() ([]entity.Product, error)
	GetProductById(id int64) (*entity.Product, error)
	CreateProducts([]entity.Product) ([]entity.Product, error)
	UpdateProduct(product entity.Product) (*entity.Product, error)
	DeleteProduct(id int64) error
	SearchProduct(key string) ([]entity.Product, error)
}
