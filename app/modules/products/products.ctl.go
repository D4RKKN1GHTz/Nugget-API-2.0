package products

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type ProductsController struct {
	productsSvc *ProductsService
}

func newController(productsService *ProductsService) *ProductsController {
	return &ProductsController{
		productsSvc: productsService,
	}
}

func (*ProductsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
