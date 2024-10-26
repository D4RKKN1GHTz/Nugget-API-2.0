package products

type ProductsModule struct {
	Ctl *ProductsController
	Svc *ProductsService
}

func New() *ProductsModule {
	svc := newService()
	return &ProductsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
