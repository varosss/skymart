package event

type OrderItemSnapshot struct {
	ProductID string
	SellerID  string
	Price     int64
	Qty       int64
}
