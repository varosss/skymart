package entity

import "clirzy/buyer/internal/domain/valueobject"

type Buyer struct {
	id     valueobject.BuyerID
	userID valueobject.UserID
}

func NewBuyer(userID valueobject.UserID) *Buyer {
	return &Buyer{
		id:     valueobject.NewBuyerID(),
		userID: userID,
	}
}

func (c *Buyer) ID() valueobject.BuyerID {
	return c.id
}

func (c *Buyer) UserID() valueobject.UserID {
	return c.userID
}
