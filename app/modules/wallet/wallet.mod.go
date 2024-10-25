package wallet

import "github.com/uptrace/bun"

type WalletModule struct {
	Ctl *WalletController
	Svc *WalletService
}

func New(db *bun.DB) *WalletModule {
	svc := newService(db)
	return &WalletModule{
		Ctl: newController(svc),
		Svc: svc,
	}
}
