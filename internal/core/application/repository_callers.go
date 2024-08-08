package application

import "sub-hf-product-p5/internal/core/domain/entity/dto"

func (app application) GetProductByIDRepository(uuid string) (*dto.ProductDB, error) {
	return app.productRepo.GetProductByID(uuid)
}

func (app application) SaveProductRepository(product dto.ProductDB) (*dto.ProductDB, error) {
	return app.productRepo.SaveProduct(product)
}

func (app application) GetProductByCategoryRepository(category string) ([]dto.ProductDB, error) {
	return app.productRepo.GetProductByCategory(category)
}

func (app application) UpdateProductByIDRepository(uuid string, product dto.ProductDB) (*dto.ProductDB, error) {
	return app.productRepo.UpdateProductByID(uuid, product)
}

func (app application) DeleteProductByIDRepository(uuid string) error {
	return app.productRepo.DeleteProductByID(uuid)
}
