package valueobject

type Money struct {
	amount   int64
	currency string
}

func NewMoney(amount int64, currency string) Money {
	return Money{
		amount:   amount,
		currency: currency,
	}
}

func (m Money) Amount() int64 {
	return m.amount
}

func (m Money) Currency() string {
	return m.currency
}
