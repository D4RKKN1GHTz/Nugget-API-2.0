package wallet

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type WalletController struct {
	WalletService *WalletService
}

func newController(walletService *WalletService) *WalletController {
	return &WalletController{
		WalletService: walletService,
	}
}

func (o *WalletController) Get(ctx *gin.Context) {
	data := o.WalletService.Wallet()
	base.Success(ctx, data)
}
