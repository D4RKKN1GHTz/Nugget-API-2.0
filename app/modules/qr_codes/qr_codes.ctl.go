package qr_codes

import (
	"app/app/modules/base"

	"github.com/gin-gonic/gin"
)

type Qr_codesController struct {
	qr_codesSvc *Qr_codesService
}

func newController(qr_codesService *Qr_codesService) *Qr_codesController {
	return &Qr_codesController{
		qr_codesSvc: qr_codesService,
	}
}

func (*Qr_codesController) Get(ctx *gin.Context) {
	base.Success(ctx, "ok")
}
