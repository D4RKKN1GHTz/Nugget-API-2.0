package permissions

type PermissionsModule struct {
	Ctl *PermissionsController
	Svc *PermissionsService
}

func New() *PermissionsModule {
	svc := newService()
	return &PermissionsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
