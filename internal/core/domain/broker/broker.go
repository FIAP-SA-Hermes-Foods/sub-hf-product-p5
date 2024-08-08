package broker

import (
	"sub-hf-product-p5/internal/core/domain/entity/dto"
)

type ProductBroker interface {
	GetProductByID(input dto.ProductBroker) (*dto.ProductBroker, error)
	SaveProduct(input dto.ProductBroker) (*dto.ProductBroker, error)
	UpdateProductByID(input dto.ProductBroker) (*dto.ProductBroker, error)
	GetProductByCategory(input dto.ProductBroker) (*dto.ProductBroker, error)
	DeleteProductByID(input dto.ProductBroker) (*dto.ProductBroker, error)
}
