package cart_items

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type Cart_itemsController struct {
	cart_itemsSvc *Cart_itemsService
}

func newController(cart_itemsService *Cart_itemsService) *Cart_itemsController {
	return &Cart_itemsController{
		cart_itemsSvc: cart_itemsService,
	}
}

func (*Cart_itemsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
