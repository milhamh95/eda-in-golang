package grpc

import (
	"context"
	"eda-in-golang/baskets/internal/domain"
	"google.golang.org/grpc"
)

type ProductRepository struct {
}

var _ domain.ProductRepository = (*ProductRepository)(nil)

func NewProductRepository(conn *grpc.ClientConn) ProductRepository {
	return ProductRepository{}
}

func (r ProductRepository) Find(ctx context.Context, productID string) (*domain.Product, error) {
	return &domain.Product{
		ID:      "1",
		StoreID: "2",
		Name:    "abc",
		Price:   30,
	}, nil
}
