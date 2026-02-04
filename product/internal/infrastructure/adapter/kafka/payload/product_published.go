package payload

type ProductPublishedPayload struct {
	ProductID string `json:"product_id"`
	SellerID  string `json:"seller_id"`
}
