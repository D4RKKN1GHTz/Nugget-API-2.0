package items

type ItemsModule struct {
	Ctl *ItemsController
	Svc *ItemsService
}

func New() *ItemsModule {
	svc := newService()
	return &ItemsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
