package transactions

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type TransactionsController struct {
	transactionsSvc *TransactionsService
}

func newController(transactionsService *TransactionsService) *TransactionsController {
	return &TransactionsController{
		transactionsSvc: transactionsService,
	}
}

func (*TransactionsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
