package role_permissions

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type Role_permissionsController struct {
	role_permissionsSvc *Role_permissionsService
}

func newController(role_permissionsService *Role_permissionsService) *Role_permissionsController {
	return &Role_permissionsController{
		role_permissionsSvc: role_permissionsService,
	}
}

func (*Role_permissionsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
