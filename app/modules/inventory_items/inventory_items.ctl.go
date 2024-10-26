package inventory_items

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type Inventory_itemsController struct {
	inventory_itemsSvc *Inventory_itemsService
}

func newController(inventory_itemsService *Inventory_itemsService) *Inventory_itemsController {
	return &Inventory_itemsController{
		inventory_itemsSvc: inventory_itemsService,
	}
}

func (*Inventory_itemsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
