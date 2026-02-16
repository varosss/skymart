package payload

type InvoiceCanceledPayload struct {
	InvoiceID string `json:"invoice_id"`
	BuyerID   string `json:"buyer_id"`
}
