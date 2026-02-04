package entity

import "clirzy/seller/internal/domain/valueobject"

type Seller struct {
	id     valueobject.SellerID
	userID valueobject.UserID
}

func NewSeller(userID valueobject.UserID) *Seller {
	return &Seller{
		id:     valueobject.NewSellerID(),
		userID: userID,
	}
}

func (c *Seller) ID() valueobject.SellerID {
	return c.id
}

func (c *Seller) UserID() valueobject.UserID {
	return c.userID
}
