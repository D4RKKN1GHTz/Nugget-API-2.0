package modules

import (
	"app/app/modules/example"
	"app/app/modules/product"
	"app/app/modules/user"
	"app/app/modules/wallet"
	"app/internal/modules/config"
	"app/internal/modules/database"
	"app/internal/modules/log"
	"app/internal/modules/otel/collector"
	"sync"
)

type Modules struct {
	Conf    *config.ConfigModule
	Log     *log.LogModule
	OTEL    *collector.OTELCollectorModule
	DB      *database.DatabaseModule
	Example *example.ExampleModule
	User    *user.UserModule
	Wallet  *wallet.WalletModule
	Product *product.ProductModule
}

func modulesInit() {
	conf := config.New()

	logMod := log.New(conf.Svc)

	otel := collector.New(conf.Svc.App())
	log.Info("otel module initialized")

	db := database.New(conf.Svc)
	log.Info("database module initialized")

	wallet := wallet.New(db.Svc.DB())
	log.Info("wallet module initialized")

	product := product.New(db.Svc.DB())
	log.Info("product module initialized")

	// rd := redis.New(conf.Svc)
	// log.Info("redis module initialized")

	user := user.New(db.Svc.DB())
	log.Info("user module initialized")

	example := example.New(user.Service)
	log.Info("example module initialized")

	mod = &Modules{
		Conf:    conf,
		Log:     logMod,
		OTEL:    otel,
		DB:      db,
		User:    user,
		Example: example,
		Wallet:  wallet,
		Product: product,
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
