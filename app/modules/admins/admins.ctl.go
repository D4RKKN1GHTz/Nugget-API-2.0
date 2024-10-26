package admins

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type AdminsController struct {
	adminsSvc *AdminsService
}

func newController(adminsService *AdminsService) *AdminsController {
	return &AdminsController{
		adminsSvc: adminsService,
	}
}

func (*AdminsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
