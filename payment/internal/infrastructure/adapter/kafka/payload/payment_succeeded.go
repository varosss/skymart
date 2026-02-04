package payload

type PaymentSucceededPayload struct {
	PaymentID string `json:"payment_id"`
	InvoiceID string `json:"invoice_id"`
	Amount    int64  `json:"amount"`
	Currency  string `json:"currency"`
}
