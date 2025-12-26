package entity

import "clirzy/buyer/internal/domain/valueobject"

type Buyer struct {
	id     valueobject.BuyerID
	userID valueobject.UserID
}

func NewBuyer(id valueobject.BuyerID, userID valueobject.UserID) *Buyer {
	return &Buyer{
		id:     id,
		userID: userID,
	}
}

func (c *Buyer) ID() valueobject.BuyerID {
	return c.id
}

func (c *Buyer) UserID() valueobject.UserID {
	return c.userID
}
