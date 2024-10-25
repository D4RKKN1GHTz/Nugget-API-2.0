package product

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService *ProductService
}

func newController(productService *ProductService) *ProductController {
	return &ProductController{
		ProductService: productService,
	}
}

func (o *ProductController) Get(ctx *gin.Context) {
	resp := o.ProductService.Product()
	base.Success(ctx, resp)
}
