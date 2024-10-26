package users

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersSvc *UsersService
}

func newController(usersService *UsersService) *UsersController {
	return &UsersController{
		usersSvc: usersService,
	}
}

func (*UsersController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
