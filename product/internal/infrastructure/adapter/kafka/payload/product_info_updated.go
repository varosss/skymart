package payload

type ProductInfoUpdatedPayload struct {
	ProductID   string `json:"product_id"`
	SellerID    string `json:"seller_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}
