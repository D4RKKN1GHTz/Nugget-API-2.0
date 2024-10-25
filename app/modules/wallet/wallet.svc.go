package wallet

import (
	walletdto "app/app/modules/wallet/dto"
	walletent "app/app/modules/wallet/ent"

	"github.com/uptrace/bun"
)

type WalletService struct {
	db *bun.DB
}

func newService(db *bun.DB) *WalletService {
	return &WalletService{
		db: db,
	}
}

func (service *WalletService) Wallet() *walletdto.WalletDTOResponse {
	req := WalletEntityReturn()
	return &walletdto.WalletDTOResponse{
		Username: req.Username,
		Wallet:   req.Wallet,
	}
}

func WalletEntityReturn() *walletent.WalletEntity {
	return &walletent.WalletEntity{
		Username: "username1",
		Password: "password1",
		Wallet:   1,
	}
}
