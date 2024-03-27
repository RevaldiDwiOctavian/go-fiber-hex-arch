package service

import (
	"context"

	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/core/domain"
	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/core/port"
)

type ProductService struct {
	repo port.ProductRepository
}

func NewProductService(repo port.ProductRepository) *ProductService {
	return &ProductService{
		repo,
	}
}

func (ps *ProductService) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	product, err := ps.repo.CreateProduct(ctx, *product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	product, err := ps.repo.GetProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (ps *ProductService) GetProducts(ctx context.Context) ([]*domain.Product, error) {
	products, err := ps.repo.GetProducts(ctx)
	if err != nil {
		return nil, err
	}

	return products, nil
}

func (ps *ProductService) UpdateProduct(ctx context.Context, id string, product domain.Product) (*domain.Product, error) {
	productUpdate, err := ps.repo.UpdateProduct(ctx, id, product)
	if err != nil {
		return nil, err
	}

	return productUpdate, nil
}

func (ps *ProductService) DeleteProduct(ctx context.Context, id string) (*domain.Product, error) {
	productDelete, err := ps.repo.DeleteProduct(ctx, id)
	if err != nil {
		return nil, err
	}

	return productDelete, nil
}
