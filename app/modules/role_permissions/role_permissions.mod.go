package role_permissions

type Role_permissionsModule struct {
	Ctl *Role_permissionsController
	Svc *Role_permissionsService
}

func New() *Role_permissionsModule {
	svc := newService()
	return &Role_permissionsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
