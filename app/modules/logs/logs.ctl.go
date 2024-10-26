package logs

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type LogsController struct {
	logsSvc *LogsService
}

func newController(logsService *LogsService) *LogsController {
	return &LogsController{
		logsSvc: logsService,
	}
}

func (*LogsController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
