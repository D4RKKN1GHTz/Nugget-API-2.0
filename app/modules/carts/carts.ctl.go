package carts

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type CartsController struct {
	cartsSvc *CartsService
}

func newController(cartsService *CartsService) *CartsController {
	return &CartsController{
		cartsSvc: cartsService,
	}
}

func (*CartsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
