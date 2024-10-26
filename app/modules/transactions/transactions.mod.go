package transactions

type TransactionsModule struct {
	Ctl *TransactionsController
	Svc *TransactionsService
}

func New() *TransactionsModule {
	svc := newService()
	return &TransactionsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
