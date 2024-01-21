package service_impl

import (
	"quqo_challenge/domain/entity"
	"quqo_challenge/domain/repository"
	"quqo_challenge/domain/service"
	"quqo_challenge/infrastructure/constant"
	"strconv"
)

type ProductServiceImpl struct {
	ProductRepo repository.ProductRepository
	RedisRepo   repository.RedisRepository
}

func NewProductService(productRepo repository.ProductRepository, redisRepo repository.RedisRepository) service.ProductService {
	return ProductServiceImpl{ProductRepo: productRepo, RedisRepo: redisRepo}
}

func (s ProductServiceImpl) GetAllProductsService() ([]entity.Product, error) {
	return s.ProductRepo.GetAllProducts()
}

func (s ProductServiceImpl) GetProductByIdService(id int64) (*entity.Product, error) {
	var cachedProduct *entity.Product
	if err := s.RedisRepo.GetByKey(strconv.Itoa(int(id)), &cachedProduct); err != nil {
		return nil, err
	}

	if cachedProduct != nil {
		return cachedProduct, nil
	}

	cachedProduct, err := s.ProductRepo.GetProductById(id)
	if err != nil {
		return nil, err
	}

	if err := s.RedisRepo.SetByKey(strconv.Itoa(int(cachedProduct.ID)), *cachedProduct, constant.DefaultRedisExpirationTime); err != nil {
		// TODO: log caching error
		return cachedProduct, nil
	}

	return cachedProduct, nil
}

func (s ProductServiceImpl) CreateProductsService(productList []entity.Product) ([]entity.Product, error) {
	return s.ProductRepo.CreateProducts(productList)
}

func (s ProductServiceImpl) UpdateProductService(product entity.Product) (*entity.Product, error) {
	data, err := s.ProductRepo.UpdateProduct(product)
	if err != nil {
		return nil, err
	}

	var cachedProduct *entity.Product
	if err := s.RedisRepo.GetByKey(strconv.Itoa(int(data.ID)), &cachedProduct); err != nil {
		return data, nil
	}
	if cachedProduct != nil {
		if err := s.RedisRepo.SetByKey(strconv.Itoa(int(data.ID)), *data, constant.DefaultRedisExpirationTime); err != nil {
			// TODO: log caching error
			return data, nil
		}
	}

	return data, nil
}

func (s ProductServiceImpl) DeleteProductService(id int64) error {
	return s.ProductRepo.DeleteProduct(id)
}

func (s ProductServiceImpl) SearchProductService(key string) ([]entity.Product, error) {
	return s.ProductRepo.SearchProduct(key)
}
