package entity

import "clirzy/seller/internal/domain/valueobject"

type Seller struct {
	id     valueobject.SellerID
	userID valueobject.UserID
}

func NewSeller(id valueobject.SellerID, userID valueobject.UserID) *Seller {
	return &Seller{
		id:     id,
		userID: userID,
	}
}

func (c *Seller) ID() valueobject.SellerID {
	return c.id
}

func (c *Seller) UserID() valueobject.UserID {
	return c.userID
}
