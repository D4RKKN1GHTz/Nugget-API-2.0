package qr_codes

type Qr_codesModule struct {
	Ctl *Qr_codesController
	Svc *Qr_codesService
}

func New() *Qr_codesModule {
	svc := newService()
	return &Qr_codesModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
