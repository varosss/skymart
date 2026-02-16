package valueobject

import "errors"

type Money struct {
	amount   int64
	currency string
}

func NewMoney(amount int64, currency string) (Money, error) {
	if amount <= 0 {
		return Money{}, errors.New("amount must be positive")
	}

	if currency == "" {
		return Money{}, errors.New("currency is required")
	}

	return Money{
		amount:   amount,
		currency: currency,
	}, nil
}

func (m Money) Amount() int64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}
