package routes

import (
	"app/app/modules"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WarpH(router *gin.RouterGroup, prefix string, handler http.Handler) {
	router.Any(fmt.Sprintf("%s/*w", prefix), gin.WrapH(http.StripPrefix(fmt.Sprintf("%s%s", router.BasePath(), prefix), handler)))
}

func api(r *gin.RouterGroup, mod *modules.Modules) {
	r.GET("/example/naming", mod.Example.Ctl.JSON)
	r.GET("/example/user", mod.Example.Ctl.User)
	r.GET("/user/show", mod.User.Ctl.Get)
	r.GET("/wallet/show", mod.Wallet.Ctl.Get)
	r.GET("/product/show", mod.Product.Ctl.Get)

	r.POST("/user/", mod.User.Ctl.UserCreate)
	r.GET("/user/", mod.User.Ctl.UserGetAll)
	r.GET("/user/:id", mod.User.Ctl.UserGetByID)
	r.PATCH("/user/:id", mod.User.Ctl.UserUpdate)
	r.PATCH("/user/resetpassword/:id", mod.User.Ctl.ResetPassword)
	r.DELETE("/user/:id", mod.User.Ctl.UserDelete)

	// router.POST("/create", ginIn.Handler(usr.Create))

}
