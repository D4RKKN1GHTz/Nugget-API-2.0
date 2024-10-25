package modules

import (
	"app/internal/modules/config"
	"app/internal/modules/database"
	"app/internal/modules/log"
	"app/internal/modules/otel/collector"
	"sync"
)

type Modules struct {
	Conf *config.ConfigModule
	Log  *log.LogModule
	OTEL *collector.OTELCollectorModule
	DB   *database.DatabaseModule
}

func modulesInit() {
	conf := config.New()

	logMod := log.New(conf.Svc)

	otel := collector.New(conf.Svc.App())
	log.Info("otel module initialized")

	db := database.New(conf.Svc)
	log.Info("database module initialized")

	// rd := redis.New(conf.Svc)
	// log.Info("redis module initialized")

	mod = &Modules{
		Conf: conf,
		Log:  logMod,
		OTEL: otel,
		DB:   db,
	}
	log.Info("all modules initialized")
}

var (
	once sync.Once
	mod  *Modules
)

func Get() *Modules {
	once.Do(modulesInit)

	return mod
}
