package carts

type CartsModule struct {
	Ctl *CartsController
	Svc *CartsService
}

func New() *CartsModule {
	svc := newService()
	return &CartsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
