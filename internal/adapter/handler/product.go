package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/core/domain"
	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/core/port"
)

type ProductHandler struct {
	svc port.ProductService
}

func NewProductHandler(svc port.ProductService) *ProductHandler {
	return &ProductHandler{
		svc,
	}
}

type createProductRequest struct {
	Product string `json:"product" validate:"required"`
	Stock   int    `json:"stock" validate:"required"`
}

func (ph *ProductHandler) CreateProduct(ctx *fiber.Ctx) error {
	var productRequest createProductRequest
	if err := ctx.BodyParser(&productRequest); err != nil {
		handleResponse(ctx, "Error", err)
		return err
	}

	product := domain.Product{
		Product: productRequest.Product,
		Stock:   productRequest.Stock,
	}

	result, err := ph.svc.CreateProduct(ctx.Context(), &product)
	if err != nil {
		handleResponse(ctx, "Error", err)
		return err
	}

	rsp := newProductResponse(result)

	handleResponse(ctx, "Success Creating Product", rsp)
	return nil
}

func (ph *ProductHandler) GetProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		handleResponse(ctx, "Error", "parameter id cannot be null")
	}

	result, err := ph.svc.GetProduct(ctx.Context(), id)
	if err != nil {
		return err
	}

	rsp := newProductResponse(result)

	handleResponse(ctx, "Success Getting Product", rsp)
	return nil
}

func (ph *ProductHandler) GetProducts(ctx *fiber.Ctx) error {
	result, err := ph.svc.GetProducts(ctx.Context())
	if err != nil {
		return err
	}

	handleResponse(ctx, "Success Getting Products", result)
	return nil
}

type updateProductRequest struct {
	Product string `json:"product" validate:"required"`
	Stock   int    `json:"stock" validate:"required"`
}

func (ph *ProductHandler) UpdateProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		handleResponse(ctx, "Error", "parameter id cannot be null")
	}

	var request updateProductRequest
	if err := ctx.BodyParser(&request); err != nil {
		handleResponse(ctx, "Error", err)
		return err
	}

	product := domain.Product{
		Product: request.Product,
		Stock:   request.Stock,
	}

	result, err := ph.svc.UpdateProduct(ctx.Context(), id, product)
	if err != nil {
		handleResponse(ctx, "Error", err)
		return err
	}

	rsp := newProductResponse(result)

	handleResponse(ctx, "Success Updating Product", rsp)
	return nil
}

func (ph *ProductHandler) DeleteProduct(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	if id == "" {
		handleResponse(ctx, "Error", "parameter id cannot be null")
	}

	result, err := ph.svc.DeleteProduct(ctx.Context(), id)
	if err != nil {
		return err
	}

	rsp := newProductResponse(result)

	handleResponse(ctx, "Success Deleting Product", rsp)
	return nil
}
