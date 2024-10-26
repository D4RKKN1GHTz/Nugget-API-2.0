package inventory_items

type Inventory_itemsModule struct {
	Ctl *Inventory_itemsController
	Svc *Inventory_itemsService
}

func New() *Inventory_itemsModule {
	svc := newService()
	return &Inventory_itemsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
