package handler

import (
	"github.com/chubin518/kestrel-layout-basic/internal/service"
	"github.com/chubin518/kestrel-layout-basic/pkg/result"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) List(ctx *gin.Context) {
	list := h.userService.List()
	result.OK.WithData(list).JSON(ctx)
}
