package baskets

import (
	"context"
	"eda-in-golang/baskets/internal/application"
	"eda-in-golang/baskets/internal/grpc"
	"eda-in-golang/baskets/internal/handlers"
	"eda-in-golang/baskets/internal/logging"
	"eda-in-golang/baskets/internal/postgres"
	"eda-in-golang/baskets/internal/rest"
	"eda-in-golang/internal/ddd"
	"eda-in-golang/internal/monolith"
)

type Module struct{}

func (m *Module) Startup(ctx context.Context, mono monolith.Monolith) (err error) {
	domainDispatcher := ddd.NewEventDispatcher()
	basketRepository := postgres.NewBasketRepository("baskets.baskets", mono.DB())

	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	storeRepository := grpc.NewStoreRepository(conn)
	productRepository := grpc.NewProductRepository(conn)
	orderRepository := grpc.NewOrderRepository(conn)

	app := logging.LogApplicationAccess(
		application.New(
			basketRepository,
			storeRepository,
			productRepository,
			orderRepository,
			domainDispatcher,
		),
		mono.Logger(),
	)

	orderHandlers := logging.LogDomainEventHandlerAccess(
		application.NewOrderHandlers(orderRepository),
		mono.Logger(),
	)

	err = grpc.RegisterServer(app, mono.RPC())
	if err != nil {
		return err
	}

	err = rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	err = rest.RegisterSwagger(mono.Mux())
	if err != nil {
		return err
	}

	handlers.RegisterOrderHandlers(orderHandlers, domainDispatcher)
	return
}
