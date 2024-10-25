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

	// router.POST("/create", ginIn.Handler(usr.Create))

}
