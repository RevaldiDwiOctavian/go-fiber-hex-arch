package port

import (
	"context"

	"github.com/revaldidwioctavian/go-fiber-hex-arch/internal/core/domain"
)

type ProductRepository interface {
	CreateProduct(ctx context.Context, product domain.Product) (*domain.Product, error)
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
	GetProducts(ctx context.Context) ([]*domain.Product, error)
	UpdateProduct(ctx context.Context, id string, product domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) (*domain.Product, error)
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error)
	GetProduct(ctx context.Context, id string) (*domain.Product, error)
	GetProducts(ctx context.Context) ([]*domain.Product, error)
	UpdateProduct(ctx context.Context, id string, product domain.Product) (*domain.Product, error)
	DeleteProduct(ctx context.Context, id string) (*domain.Product, error)
}
