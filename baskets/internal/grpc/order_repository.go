package grpc

import (
	"context"
	"eda-in-golang/baskets/internal/domain"
	"google.golang.org/grpc"
)

type OrderRepository struct {
}

var _ domain.OrderRepository = (*OrderRepository)(nil)

func NewOrderRepository(conn *grpc.ClientConn) OrderRepository {
	return OrderRepository{}
}

func (r OrderRepository) Save(ctx context.Context, basket *domain.Basket) (string, error) {
	return "d", nil
}
