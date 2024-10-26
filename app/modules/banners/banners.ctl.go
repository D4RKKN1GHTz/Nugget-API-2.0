package banners

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type BannersController struct {
	bannersSvc *BannersService
}

func newController(bannersService *BannersService) *BannersController {
	return &BannersController{
		bannersSvc: bannersService,
	}
}

func (*BannersController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
