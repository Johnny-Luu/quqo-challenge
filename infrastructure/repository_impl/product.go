package repository_impl

import (
	"quqo_challenge/domain/db"
	"quqo_challenge/domain/entity"
	"quqo_challenge/domain/repository"
)

type ProductRepositoryImpl struct {
	p *db.Persistence
}

func NewProductRepository(p *db.Persistence) repository.ProductRepository {
	return ProductRepositoryImpl{p: p}
}

func (repo ProductRepositoryImpl) GetAllProducts() ([]entity.Product, error) {
	var productList []entity.Product

	if err := repo.p.AppDb.Find(&productList).Error; err != nil {
		return nil, err
	}

	return productList, nil
}

func (repo ProductRepositoryImpl) GetProductById(id int64) (*entity.Product, error) {
	var product entity.Product

	if err := repo.p.AppDb.
		Where("id = ?", id).
		First(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo ProductRepositoryImpl) CreateProducts(productList []entity.Product) ([]entity.Product, error) {
	if err := repo.p.AppDb.
		CreateInBatches(productList, len(productList)).Error; err != nil {
		return nil, err
	}

	return productList, nil
}

func (repo ProductRepositoryImpl) UpdateProduct(product entity.Product) (*entity.Product, error) {
	if _, err := repo.GetProductById(product.ID); err != nil {
		return nil, err
	}

	if err := repo.p.AppDb.Save(&product).Error; err != nil {
		return nil, err
	}

	return &product, nil
}

func (repo ProductRepositoryImpl) DeleteProduct(id int64) error {
	if _, err := repo.GetProductById(id); err != nil {
		return err
	}

	if err := repo.p.AppDb.Delete(&entity.Product{}, id).Error; err != nil {
		return err
	}

	return nil
}

func (repo ProductRepositoryImpl) SearchProduct(key string) ([]entity.Product, error) {
	var productList []entity.Product

	if err := repo.p.AppDb.
		Where("name ilike ?", "%"+key+"%").
		Find(&productList).Error; err != nil {
		return nil, err
	}

	return productList, nil
}
