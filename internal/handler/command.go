package handler

import (
	"os/exec"

	"github.com/chubin518/kestrel-layout-basic/pkg/result"
	"github.com/gin-gonic/gin"
)

type CmdHandler struct{}

func NewCmdHandler() *CmdHandler {
	return &CmdHandler{}
}

func (h *CmdHandler) Exec(ctx *gin.Context) {
	command, ok := ctx.GetPostForm("command")
	if !ok || command == "" {
		result.FAILED.JSON(ctx)
		return
	}
	cmd := exec.Command("bash", "-c", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		result.INTERNAL_SERVER_ERROR.WithMessage(err.Error()).JSON(ctx)
		return
	}
	result.OK.WithData(string(output)).JSON(ctx)
}
