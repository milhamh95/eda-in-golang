package grpc

import (
	"context"
	"eda-in-golang/baskets/internal/domain"
	"google.golang.org/grpc"
)

type StoreRepository struct {
}

var _ domain.StoreRepository = (*StoreRepository)(nil)

func NewStoreRepository(conn *grpc.ClientConn) StoreRepository {
	return StoreRepository{}
}

func (r StoreRepository) Find(ctx context.Context, storeID string) (*domain.Store, error) {
	return &domain.Store{
		ID:       "5",
		Name:     "toko baru",
		Location: "jakarta",
	}, nil
}
