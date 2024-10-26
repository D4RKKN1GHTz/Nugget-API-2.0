package logs

type LogsModule struct {
	Ctl *LogsController
	Svc *LogsService
}

func New() *LogsModule {
	svc := newService()
	return &LogsModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
