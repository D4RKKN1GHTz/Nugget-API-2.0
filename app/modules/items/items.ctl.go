package items

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type ItemsController struct {
	itemsSvc *ItemsService
}

func newController(itemsService *ItemsService) *ItemsController {
	return &ItemsController{
		itemsSvc: itemsService,
	}
}

func (*ItemsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
