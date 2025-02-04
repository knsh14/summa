package summa

import "context"

type Transaction interface {
	ListCredit() []Credit
	ListDebt() []Debt
}

type Credit interface {
	implementCredit()
}

type Debt interface {
	implementDebt()
}

func CreateTransaction(ctx context.Context) (Transaction, error) {
	return nil, nil
}
