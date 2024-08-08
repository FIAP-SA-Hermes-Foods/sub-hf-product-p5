package repository

import "sub-hf-product-p5/internal/core/domain/entity/dto"

type ProductRepository interface {
	GetProductByID(uuid string) (*dto.ProductDB, error)
	SaveProduct(product dto.ProductDB) (*dto.ProductDB, error)
	UpdateProductByID(uuid string, product dto.ProductDB) (*dto.ProductDB, error)
	GetProductByCategory(category string) ([]dto.ProductDB, error)
	DeleteProductByID(uuid string) error
}
