package histories

type HistoriesModule struct {
	Ctl *HistoriesController
	Svc *HistoriesService
}

func New() *HistoriesModule {
	svc := newService()
	return &HistoriesModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
