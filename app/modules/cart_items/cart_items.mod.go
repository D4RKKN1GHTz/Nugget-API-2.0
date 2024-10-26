package cart_items

type Cart_itemsModule struct {
	Ctl *Cart_itemsController
	Svc *Cart_itemsService
}

func New() *Cart_itemsModule {
	svc := newService()
	return &Cart_itemsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
