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
	ProductID string
	Quantity  int64
}

type CreateOrderCommand struct {
	BuyerID string
	Items   []CreateOrderItem
}

type CreateOrderUseCase struct {
	products  aport.ProductQuery
	buyers    aport.BuyerQuery
	orders    port.OrdersRepo
	publisher aport.EventPublisher
}

func (uc *CreateOrderUseCase) Execute(
	ctx context.Context,
	cmd CreateOrderCommand,
) (valueobject.OrderID, error) {

	buyerID, err := valueobject.ParseBuyerID(cmd.BuyerID)
	if err != nil {
		return "", domain.ErrInvalidBuyerID
	}

	buyer, err := uc.buyers.GetByID(cmd.BuyerID)
	if err != nil {
		return "", domain.ErrNoBuyerFound
	}

	if !buyer.IsActive {
		return "", domain.ErrInactiveBuyer
	}

	productIDs := make([]string, 0, len(cmd.Items))
	qtyByProduct := make(map[string]int64)

	for _, i := range cmd.Items {
		if i.Quantity <= 0 {
			return "", domain.ErrInvalidQuantity
		}

		if _, ok := qtyByProduct[i.ProductID]; !ok {
			productIDs = append(productIDs, i.ProductID)
		}

		qtyByProduct[i.ProductID] += i.Quantity
	}

	products, err := uc.products.GetProducts(productIDs)
	if err != nil {
		return "", err
	}

	if len(products) != len(productIDs) {
		return "", domain.ErrProductNotFound
	}

	orderItems := make([]entity.OrderItem, 0, len(productIDs))

	var productID valueobject.ProductID
	var sellerID valueobject.SellerID
	var price valueobject.Money

	for _, p := range products {

		if !p.IsPublished {
			return "", domain.ErrProductNotAvailable
		}

		productID, err = valueobject.ParseProductID(p.ID)
		if err != nil {
			return "", domain.ErrInvalidProductID
		}

		sellerID, err = valueobject.ParseSellerID(p.SellerID)
		if err != nil {
			return "", domain.ErrInvalidSellerID
		}

		price = valueobject.NewMoney(p.Amount, p.Currency)

		item := entity.NewOrderItem(
			productID,
			sellerID,
			price,
			qtyByProduct[p.ID],
		)

		orderItems = append(orderItems, item)
	}

	order, err := entity.NewOrder(
		buyerID,
		orderItems,
	)
	if err != nil {
		return "", err
	}

	if err := uc.orders.Save(ctx, order); err != nil {
		return "", err
	}

	if err := uc.publisher.Publish(order.PullEvents()); err != nil {
		return "", err
	}

	return order.ID(), nil
}
