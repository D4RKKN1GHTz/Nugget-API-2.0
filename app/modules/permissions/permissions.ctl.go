package permissions

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type PermissionsController struct {
	permissionsSvc *PermissionsService
}

func newController(permissionsService *PermissionsService) *PermissionsController {
	return &PermissionsController{
		permissionsSvc: permissionsService,
	}
}

func (*PermissionsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
