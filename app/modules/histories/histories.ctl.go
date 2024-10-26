package histories

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type HistoriesController struct {
	historiesSvc *HistoriesService
}

func newController(historiesService *HistoriesService) *HistoriesController {
	return &HistoriesController{
		historiesSvc: historiesService,
	}
}

func (*HistoriesController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
