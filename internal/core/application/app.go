package application

import (
	"context"
	l "sub-hf-product-p5/external/logger"
	ps "sub-hf-product-p5/external/strings"
	"sub-hf-product-p5/internal/core/domain/broker"
	"sub-hf-product-p5/internal/core/domain/entity/dto"
	"sub-hf-product-p5/internal/core/domain/repository"
	"sub-hf-product-p5/internal/core/domain/useCase"
)

type Application interface {
	GetProductByID(msgID string, uuid string) (*dto.OutputProduct, error)
	SaveProduct(msgID string, product dto.RequestProduct) (*dto.OutputProduct, error)
	UpdateProductByID(msgID string, id string, product dto.RequestProduct) (*dto.OutputProduct, error)
	GetProductByCategory(msgID string, category string) ([]dto.OutputProduct, error)
	DeleteProductByID(msgID string, id string) error
}

type application struct {
	ctx           context.Context
	productBroker broker.ProductBroker
	productRepo   repository.ProductRepository
	productUC     useCase.ProductUseCase
}

func NewApplication(ctx context.Context, productBroker broker.ProductBroker, productRepo repository.ProductRepository, productUC useCase.ProductUseCase) Application {
	return application{
		ctx:           ctx,
		productBroker: productBroker,
		productRepo:   productRepo,
		productUC:     productUC,
	}
}

func (app application) GetProductByID(msgID string, uuid string) (*dto.OutputProduct, error) {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "GetProductByIDApp: ", " | ", uuid)

	inputBroker := dto.ProductBroker{
		UUID:      uuid,
		MessageID: msgID,
	}

	outBroker, err := app.productBroker.GetProductByID(inputBroker)

	if err != nil {
		l.Errorf(msgID, "GetProductByIDApp error: ", " | ", err)
		return nil, err
	}

	if outBroker == nil {
		return nil, nil
	}

	if err := app.GetProductByIDUseCase(outBroker.UUID); err != nil {
		l.Errorf(msgID, "GetProductByIDApp error: ", " | ", err)
		return nil, err
	}

	o, err := app.GetProductByIDRepository(outBroker.UUID)

	if err != nil {
		l.Errorf(msgID, "GetProductByIDApp error: ", " | ", err)
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	out := &dto.OutputProduct{
		UUID:          o.UUID,
		Name:          o.Name,
		Category:      o.Category,
		Image:         o.Image,
		Description:   o.Description,
		Price:         o.Price,
		CreatedAt:     o.CreatedAt,
		DeactivatedAt: o.DeactivatedAt,
	}

	l.Infof(msgID, "GetProductByIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) SaveProduct(msgID string, product dto.RequestProduct) (*dto.OutputProduct, error) {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "SaveProductApp: ", " | ", ps.MarshalString(product))

	inputBroker := dto.ProductBroker{
		UUID:          product.UUID,
		MessageID:     msgID,
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         product.Price,
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	oBroker, err := app.productBroker.SaveProduct(inputBroker)

	if err != nil {
		l.Errorf(msgID, "SaveProductApp error: ", " | ", err)
		return nil, err
	}

	if oBroker == nil {
		return nil, nil
	}

	pUC := dto.RequestProduct{
		UUID:          oBroker.UUID,
		Name:          oBroker.Name,
		Category:      oBroker.Category,
		Image:         oBroker.Image,
		Description:   oBroker.Description,
		Price:         oBroker.Price,
		CreatedAt:     oBroker.CreatedAt,
		DeactivatedAt: oBroker.DeactivatedAt,
	}

	if err := app.SaveProductUseCase(pUC); err != nil {
		l.Errorf(msgID, "SaveProductApp error: ", " | ", err)
		return nil, err
	}

	pDB := dto.ProductDB{
		UUID:          oBroker.UUID,
		Name:          oBroker.Name,
		Category:      oBroker.Category,
		Image:         oBroker.Image,
		Description:   oBroker.Description,
		Price:         oBroker.Price,
		CreatedAt:     oBroker.CreatedAt,
		DeactivatedAt: oBroker.DeactivatedAt,
	}

	o, err := app.SaveProductRepository(pDB)

	if err != nil {
		l.Errorf(msgID, "SaveProductApp error: ", " | ", err)
		return nil, err
	}

	out := &dto.OutputProduct{
		UUID:          o.UUID,
		Name:          o.Name,
		Category:      o.Category,
		Image:         o.Image,
		Description:   o.Description,
		Price:         o.Price,
		CreatedAt:     o.CreatedAt,
		DeactivatedAt: o.DeactivatedAt,
	}

	l.Infof(msgID, "SaveProductApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) GetProductByCategory(msgID string, category string) ([]dto.OutputProduct, error) {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "GetProductByCategoryApp: ", " | ", category)

	inputProker := dto.ProductBroker{
		MessageID: msgID,
		Category:  category,
	}

	outBroker, err := app.productBroker.GetProductByCategory(inputProker)

	if err != nil {
		l.Errorf(msgID, "GetProductByCategoryApp error: ", " | ", err)
		return nil, err
	}

	if outBroker == nil {
		return nil, nil
	}

	if err := app.GetProductByCategoryUseCase(outBroker.Category); err != nil {
		l.Errorf(msgID, "GetProductByIDApp error: ", " | ", err)
		return nil, err
	}

	oRepo, err := app.GetProductByCategoryRepository(outBroker.Category)

	if err != nil {
		l.Errorf(msgID, "GetProductByIDApp error: ", " | ", err)
		return nil, err
	}

	out := make([]dto.OutputProduct, 0)

	for _, o := range oRepo {
		op := &dto.OutputProduct{
			UUID:          o.UUID,
			Name:          o.Name,
			Category:      o.Category,
			Image:         o.Image,
			Description:   o.Description,
			Price:         o.Price,
			CreatedAt:     o.CreatedAt,
			DeactivatedAt: o.DeactivatedAt,
		}

		out = append(out, *op)

	}

	l.Infof(msgID, "GetProductByCategoryApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) UpdateProductByID(msgID string, id string, product dto.RequestProduct) (*dto.OutputProduct, error) {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "UpdateProductByIDApp: ", " | ", id, " | ", ps.MarshalString(product))

	inputBroker := dto.ProductBroker{
		UUID:          id,
		MessageID:     msgID,
		Name:          product.Name,
		Category:      product.Category,
		Image:         product.Image,
		Description:   product.Description,
		Price:         product.Price,
		CreatedAt:     product.CreatedAt,
		DeactivatedAt: product.DeactivatedAt,
	}

	oBroker, err := app.productBroker.UpdateProductByID(inputBroker)

	if err != nil {
		l.Errorf(msgID, "UpdateProductByIDApp error: ", " | ", err)
		return nil, err
	}

	if err != nil {
		l.Errorf(msgID, "UpdateProductByIDApp error: ", " | ", err)
		return nil, err
	}

	if oBroker == nil {
		return nil, nil
	}

	pUC := dto.RequestProduct{
		UUID:          oBroker.UUID,
		Name:          oBroker.Name,
		Category:      oBroker.Category,
		Image:         oBroker.Image,
		Description:   oBroker.Description,
		Price:         oBroker.Price,
		CreatedAt:     oBroker.CreatedAt,
		DeactivatedAt: oBroker.DeactivatedAt,
	}

	if err := app.UpdateProductByIDUseCase(pUC.UUID, pUC); err != nil {
		l.Errorf(msgID, "SaveProductApp error: ", " | ", err)
		return nil, err
	}

	pDB := dto.ProductDB{
		UUID:          oBroker.UUID,
		Name:          oBroker.Name,
		Category:      oBroker.Category,
		Image:         oBroker.Image,
		Description:   oBroker.Description,
		Price:         oBroker.Price,
		CreatedAt:     oBroker.CreatedAt,
		DeactivatedAt: oBroker.DeactivatedAt,
	}

	o, err := app.UpdateProductByIDRepository(pDB.UUID, pDB)

	if err != nil {
		l.Errorf(msgID, "UpdateProductByIDApp error: ", " | ", err)
		return nil, err
	}

	out := &dto.OutputProduct{
		UUID:          o.UUID,
		Name:          o.Name,
		Category:      o.Category,
		Image:         o.Image,
		Description:   o.Description,
		Price:         o.Price,
		CreatedAt:     o.CreatedAt,
		DeactivatedAt: o.DeactivatedAt,
	}

	l.Infof(msgID, "UpdateProductByIDApp output: ", " | ", ps.MarshalString(out))
	return out, nil
}

func (app application) DeleteProductByID(msgID string, id string) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "DeleteProductByIDApp: ", " | ", id)

	inputBroker := dto.ProductBroker{
		UUID:      id,
		MessageID: msgID,
	}

	oBroker, err := app.productBroker.DeleteProductByID(inputBroker)

	if err != nil {
		l.Errorf(msgID, "DeleteProductByIDApp error: ", " | ", err)
		return err
	}

	if err := app.DeleteProductByID(msgID, oBroker.UUID); err != nil {
		l.Errorf(msgID, "DeleteProductByIDApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "DeleteProductByIDApp output: ", " | ", "deleted with success!")
	return nil
}

func (app application) setMessageIDCtx(msgID string) {
	if app.ctx == nil {
		app.ctx = context.WithValue(context.Background(), l.MessageIDKey, msgID)
		return
	}
	app.ctx = context.WithValue(app.ctx, l.MessageIDKey, msgID)
}
