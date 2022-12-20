package money

type Money struct {
	Amount   int
	Currency string
}

func New(a int, c string) Money {
	m := Money{
		Amount:   a,
		Currency: c,
	}
	return m
}
