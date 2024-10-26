package roles

type RolesModule struct {
	Ctl *RolesController
	Svc *RolesService
}

func New() *RolesModule {
	svc := newService()
	return &RolesModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
