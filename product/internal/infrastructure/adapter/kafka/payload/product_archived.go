package payload

type ProductArchivedPayload struct {
	ProductID string `json:"product_id"`
	SellerID  string `json:"seller_id"`
}
