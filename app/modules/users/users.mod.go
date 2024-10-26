package users

type UsersModule struct {
	Ctl *UsersController
	Svc *UsersService
}

func New() *UsersModule {
	svc := newService()
	return &UsersModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
