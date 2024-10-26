package roles

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type RolesController struct {
	rolesSvc *RolesService
}

func newController(rolesService *RolesService) *RolesController {
	return &RolesController{
		rolesSvc: rolesService,
	}
}

func (*RolesController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
