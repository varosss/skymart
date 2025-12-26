package entity

import (
	"clirzy/product/internal/domain"
	"clirzy/product/internal/domain/event"
	"clirzy/product/internal/domain/valueobject"
)

type Product struct {
	id          valueobject.ProductID
	sellerID    valueobject.SellerID
	title       string
	description string
	price       valueobject.Money
	status      valueobject.Status

	events []event.Event
}

func NewProduct(
	id valueobject.ProductID,
	sellerID valueobject.SellerID,
	title string,
	description string,
	price valueobject.Money,
	status valueobject.Status,
) *Product {
	p := &Product{
		id:          id,
		sellerID:    sellerID,
		title:       title,
		description: description,
		price:       price,
		status:      status,
	}

	p.addEvent(event.NewProductCreated(p.id, p.sellerID))

	return p
}

func (p *Product) ID() valueobject.ProductID {
	return p.id
}

func (p *Product) SellerID() valueobject.SellerID {
	return p.sellerID
}

func (p *Product) Price() valueobject.Money {
	return p.price
}

func (p *Product) Publish() error {
	if p.status != valueobject.StatusDraft && p.status != valueobject.StatusUnpublished {
		return domain.ErrInvalidProductStatusTransition
	}

	p.status = valueobject.StatusPublished
	p.addEvent(event.NewProductPublished(p.id, p.sellerID))

	return nil
}

func (p *Product) Unpublish() error {
	if p.status != valueobject.StatusPublished {
		return domain.ErrInvalidProductStatusTransition
	}
	p.status = valueobject.StatusUnpublished
	return nil
}

func (p *Product) Archive() error {
	if p.status == valueobject.StatusArchived {
		return domain.ErrInvalidProductStatusTransition
	}
	p.status = valueobject.StatusArchived
	return nil
}

func (p *Product) UpdateInfo(title, description string) error {
	if p.status == valueobject.StatusArchived {
		return domain.ErrCannotChangeArchivedProduct
	}
	if title == "" {
		return domain.ErrInvalidTitle
	}

	p.title = title
	p.description = description
	return nil
}

func (p *Product) UpdatePrice(newPrice valueobject.Money) error {
	if p.status == valueobject.StatusArchived {
		return domain.ErrCannotChangeArchivedProduct
	}
	if newPrice.Amount < 0 {
		return domain.ErrInvalidTitle
	}

	p.price = newPrice
	return nil
}

func (p *Product) addEvent(e event.Event) {
	p.events = append(p.events, e)
}

func (p *Product) PullEvents() []event.Event {
	events := p.events
	p.events = nil
	return events
}
