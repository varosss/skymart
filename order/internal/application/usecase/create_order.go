package usecase

import (
	aport "clirzy/order/internal/application/port"
	"clirzy/order/internal/domain"
	"clirzy/order/internal/domain/entity"
	"clirzy/order/internal/domain/port"
	"clirzy/order/internal/domain/valueobject"
	"context"
)

type CreateOrderItem struct {
	SellerID  valueobject.SellerID
	ProductID valueobject.ProductID
	Quantity  int64
}

type CreateOrderCommand struct {
	BuyerID valueobject.BuyerID
	Items   []CreateOrderItem
}

type CreateOrderUseCase struct {
	products aport.ProductGateway
	buyers   aport.BuyerGateway
	orders   port.OrderRepo
	bus      aport.EventBus
}

func (uc *CreateOrderUseCase) Execute(
	ctx context.Context,
	cmd CreateOrderCommand,
) (valueobject.OrderID, error) {

	buyer, err := uc.buyers.GetByID(ctx, cmd.BuyerID)
	if err != nil {
		return "", domain.ErrBuyerNotFound
	}

	if !buyer.IsActive {
		return "", domain.ErrInactiveBuyer
	}

	productIDs := make([]valueobject.ProductID, 0, len(cmd.Items))
	qtyByProduct := make(map[string]int64)

	for _, item := range cmd.Items {
		if item.Quantity <= 0 {
			return "", domain.ErrInvalidQuantity
		}

		if _, ok := qtyByProduct[item.ProductID.String()]; !ok {
			productIDs = append(productIDs, item.ProductID)
		}

		qtyByProduct[item.ProductID.String()] += item.Quantity
	}

	products, err := uc.products.GetProducts(ctx, productIDs)
	if err != nil {
		return "", err
	}

	if len(products) != len(productIDs) {
		return "", domain.ErrProductNotFound
	}

	orderItems := make([]entity.OrderItem, 0, len(productIDs))

	for _, p := range products {

		if !p.IsPublished {
			return "", domain.ErrProductNotAvailable
		}

		item := entity.NewOrderItem(
			p.ID,
			p.SellerID,
			p.Price,
			qtyByProduct[p.ID.String()],
		)

		orderItems = append(orderItems, item)
	}

	order, err := entity.NewOrder(
		cmd.BuyerID,
		orderItems,
	)
	if err != nil {
		return "", err
	}

	if err := uc.orders.Save(ctx, order); err != nil {
		return "", err
	}

	if err := uc.bus.Publish(ctx, order.PullEvents()); err != nil {
		return "", err
	}

	return order.ID(), nil
}
