package application

import (
	"context"
	"eda-in-golang/baskets/internal/domain"
	"eda-in-golang/internal/ddd"
	"github.com/stackus/errors"
)

type (
	StartBasket struct {
		ID         string
		CustomerID string
	}

	CancelBasket struct {
		ID string
	}

	CheckoutBasket struct {
		ID        string
		PaymentID string
	}

	AddItem struct {
		ID        string
		ProductID string
		Quantity  int
	}

	RemoveItem struct {
		ID        string
		ProductID string
		Quantity  int
	}

	GetBasket struct {
		ID string
	}

	App interface {
		StartBasket(ctx context.Context, start StartBasket) error
		CancelBasket(ctx context.Context, cancel CancelBasket) error
		CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error
		AddItem(ctx context.Context, add AddItem) error
		RemoveItem(ctx context.Context, remove RemoveItem) error
		GetBasket(ctx context.Context, get GetBasket) (*domain.Basket, error)
	}

	Application struct {
		basketRepository  domain.BasketRepository
		storeRepository   domain.StoreRepository
		productRepository domain.ProductRepository
		orderRepository   domain.OrderRepository
		domainPublisher   ddd.EventPublisher
	}
)

// interface checker
var _ App = (*Application)(nil)

func New(
	basketRepository domain.BasketRepository,
	storeRepository domain.StoreRepository,
	productRepository domain.ProductRepository,
	orderRepository domain.OrderRepository,
	domainPublisher ddd.EventPublisher,
) *Application {
	return &Application{
		basketRepository:  basketRepository,
		storeRepository:   storeRepository,
		productRepository: productRepository,
		orderRepository:   orderRepository,
		domainPublisher:   domainPublisher,
	}
}

func (a Application) StartBasket(ctx context.Context, start StartBasket) error {
	basket, err := domain.StartBasket(start.ID, start.CustomerID)
	if err != nil {
		return err
	}

	err = a.basketRepository.Save(ctx, basket)
	if err != nil {
		return err
	}

	err = a.domainPublisher.Publish(ctx, basket.GetEvents()...)
	if err != nil {
		return err
	}

	return nil
}

func (a Application) CancelBasket(ctx context.Context, cancel CancelBasket) error {
	basket, err := a.basketRepository.Find(ctx, cancel.ID)
	if err != nil {
		return err
	}

	err = basket.Cancel()
	if err != nil {
		return err
	}

	err = a.basketRepository.Update(ctx, basket)
	if err != nil {
		return err
	}

	err = a.domainPublisher.Publish(ctx, basket.GetEvents()...)
	if err != nil {
		return err
	}

	return nil
}

func (a Application) CheckoutBasket(ctx context.Context, checkout CheckoutBasket) error {
	basket, err := a.basketRepository.Find(ctx, checkout.ID)
	if err != nil {
		return err
	}

	err = basket.Checkout(checkout.PaymentID)
	if err != nil {
		return errors.Wrap(err, "baskets checkout")
	}

	err = a.basketRepository.Update(ctx, basket)
	if err != nil {
		return errors.Wrap(err, "basket checkout")
	}

	err = a.domainPublisher.Publish(ctx, basket.GetEvents()...)
	if err != nil {
		return err
	}

	return nil
}

func (a Application) AddItem(ctx context.Context, add AddItem) error {
	basket, err := a.basketRepository.Find(ctx, add.ID)
	if err != nil {
		return err
	}

	product, err := a.productRepository.Find(ctx, add.ProductID)
	if err != nil {
		return err
	}

	store, err := a.storeRepository.Find(ctx, product.StoreID)
	if err != nil {
		return err
	}

	err = basket.AddItem(store, product, add.Quantity)
	if err != nil {
		return err
	}

	err = a.basketRepository.Update(ctx, basket)
	if err != nil {
		return err
	}

	err = a.domainPublisher.Publish(ctx, basket.GetEvents()...)
	if err != nil {
		return err
	}

	return nil
}
