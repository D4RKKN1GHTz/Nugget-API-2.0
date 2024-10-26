package banners

type BannersModule struct {
	Ctl *BannersController
	Svc *BannersService
}

func New() *BannersModule {
	svc := newService()
	return &BannersModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
