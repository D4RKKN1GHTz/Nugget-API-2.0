package admins

type AdminsModule struct {
	Ctl *AdminsController
	Svc *AdminsService
}

func New() *AdminsModule {
	svc := newService()
	return &AdminsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
