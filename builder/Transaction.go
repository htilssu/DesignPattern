package builder

type Transaction struct {
	Amount float64
	From   string
	To     string
	Id     int32
}

func (t *Transaction) AddAmount() *Transaction {
	t.Amount = 100
	return t
}

func (t *Transaction) AddFrom(from string) *Transaction {
	t.From = from
	return t
}

func (t *Transaction) AddTo(to string) *Transaction {
	t.To = to
	return t
}
