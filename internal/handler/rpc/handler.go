package rpc

import (
	"context"
	l "sub-hf-product-p5/external/logger"
	"sub-hf-product-p5/internal/core/application"
	"sub-hf-product-p5/internal/core/domain/entity/dto"
	cp "sub-hf-product-p5/product_sub_proto"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedProductServer
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) GetProductByID(ctx context.Context, req *cp.GetProductByIDRequest) (*cp.GetProductByIDResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	outApp, err := h.app.GetProductByID(msgID, req.Uuid)
	if err != nil {
		return nil, err
	}

	out := &cp.GetProductByIDResponse{
		Uuid:          outApp.UUID,
		Name:          outApp.Name,
		Category:      outApp.Category,
		Image:         outApp.Image,
		Description:   outApp.Description,
		Price:         float32(outApp.Price),
		CreatedAt:     outApp.CreatedAt,
		DeactivatedAt: outApp.DeactivatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) CreateProduct(ctx context.Context, req *cp.CreateProductRequest) (*cp.CreateProductResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	input := dto.RequestProduct{
		Name:          req.Name,
		Category:      req.Category,
		Image:         req.Image,
		Description:   req.Description,
		Price:         float64(req.Price),
		CreatedAt:     req.CreatedAt,
		DeactivatedAt: req.DeactivatedAt,
	}

	outApp, err := h.app.SaveProduct(msgID, input)

	if err != nil {
		return nil, err
	}

	out := &cp.CreateProductResponse{
		Uuid:          outApp.UUID,
		Name:          outApp.Name,
		Category:      outApp.Category,
		Image:         outApp.Image,
		Description:   outApp.Description,
		Price:         float32(outApp.Price),
		CreatedAt:     outApp.CreatedAt,
		DeactivatedAt: outApp.DeactivatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) GetProductByCategory(ctx context.Context, req *cp.GetProductByCategoryRequest) (*cp.GetProductByCategoryResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	oApp, err := h.app.GetProductByCategory(msgID, req.Category)

	if err != nil {
		return nil, err
	}

	items := make([]*cp.GetProductByCategoryItem, 0)

	for _, outApp := range oApp {

		o := &cp.GetProductByCategoryItem{
			Uuid:          outApp.UUID,
			Name:          outApp.Name,
			Category:      outApp.Category,
			Image:         outApp.Image,
			Description:   outApp.Description,
			Price:         float32(outApp.Price),
			CreatedAt:     outApp.CreatedAt,
			DeactivatedAt: outApp.DeactivatedAt,
		}

		items = append(items, o)
	}

	out := cp.GetProductByCategoryResponse{
		Items: items,
	}

	return &out, nil
}

func (h *handlerGRPC) UpdateProduct(ctx context.Context, req *cp.UpdateProductRequest) (*cp.UpdateProductResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	input := dto.RequestProduct{
		Name:          req.Name,
		Category:      req.Category,
		Image:         req.Image,
		Description:   req.Description,
		Price:         float64(req.Price),
		CreatedAt:     req.CreatedAt,
		DeactivatedAt: req.DeactivatedAt,
	}

	outApp, err := h.app.UpdateProductByID(msgID, req.Uuid, input)

	if err != nil {
		return nil, err
	}

	out := &cp.UpdateProductResponse{
		Uuid:          outApp.UUID,
		Name:          outApp.Name,
		Category:      outApp.Category,
		Image:         outApp.Image,
		Description:   outApp.Description,
		Price:         float32(outApp.Price),
		CreatedAt:     outApp.CreatedAt,
		DeactivatedAt: outApp.DeactivatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) DeleteProductByID(ctx context.Context, req *cp.DeleteProductByIDRequest) (*cp.DeleteProductByIDResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	if err := h.app.DeleteProductByID(msgID, req.Uuid); err != nil {
		return nil, err
	}

	return nil, nil
}
