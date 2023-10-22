package application

import (
	"context"
	"eda-in-golang/baskets/internal/domain"
	"eda-in-golang/internal/ddd"
)

type OrderHandlers struct {
	orderRepository domain.OrderRepository
	ignoreUnimplementedDomainEvents
}

var _ DomainEventHandlers = (*OrderHandlers)(nil)

func NewOrderHandlers(orderRepository domain.OrderRepository) OrderHandlers {
	return OrderHandlers{
		orderRepository: orderRepository,
	}
}

func (h OrderHandlers) OnBasketCheckedOut(ctx context.Context, event ddd.Event) error {
	checkedOut := event.(*domain.BasketCheckedOut)
	_, err := h.orderRepository.Save(ctx, checkedOut.Basket)
	return err
}
