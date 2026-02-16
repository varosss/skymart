package event

import (
	"clirzy/product/internal/domain/valueobject"
	"time"
)

type ProductInfoUpdated struct {
	eventID     valueobject.EventID
	productID   valueobject.ProductID
	sellerID    valueobject.SellerID
	title       string
	description string
	occurredAt  time.Time
}

func NewProductInfoUpdated(
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	title string,
	description string,
) *ProductInfoUpdated {
	return &ProductInfoUpdated{
		eventID:     valueobject.NewEventID(),
		productID:   productID,
		sellerID:    sellerID,
		title:       title,
		description: description,
		occurredAt:  time.Now(),
	}
}

func ProductInfoUpdatedFromPrimitives(
	eventID valueobject.EventID,
	productID valueobject.ProductID,
	sellerID valueobject.SellerID,
	title string,
	description string,
	occurredAt time.Time,
) *ProductInfoUpdated {
	return &ProductInfoUpdated{
		eventID:     eventID,
		productID:   productID,
		sellerID:    sellerID,
		title:       title,
		description: description,
		occurredAt:  occurredAt,
	}
}

func (e *ProductInfoUpdated) ID() string {
	return e.eventID.String()
}

func (*ProductInfoUpdated) Type() string {
	return "product.info_updated"
}

func (e *ProductInfoUpdated) AggregateID() string {
	return e.productID.String()
}

func (*ProductInfoUpdated) AggregateType() string {
	return "product"
}

func (e *ProductInfoUpdated) OccurredAt() time.Time {
	return e.occurredAt
}

func (e *ProductInfoUpdated) ProductID() valueobject.ProductID {
	return e.productID
}

func (e *ProductInfoUpdated) SellerID() valueobject.SellerID {
	return e.sellerID
}

func (e *ProductInfoUpdated) Title() string {
	return e.title
}

func (e *ProductInfoUpdated) Description() string {
	return e.description
}
