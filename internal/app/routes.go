package app

import (
	"io/fs"
	"net/http"

	"github.com/chubin518/kestrel-layout-basic/internal/handler"
	"github.com/chubin518/kestrel-layout-basic/pkg/graceful"
	"github.com/chubin518/kestrel-layout-basic/pkg/middleware"
	"github.com/chubin518/kestrel-layout-basic/webui"
	"github.com/gin-gonic/gin"
)

func homeHandler(ctx *gin.Context) {
	rawIndex, _ := webui.StaticFS.ReadFile("dist/index.html")
	ctx.Writer.Write(rawIndex)
}

func NewRoutes(userHandler *handler.UserHandler, cmdHandler *handler.CmdHandler) *graceful.Routes {
	return &graceful.Routes{
		ConfigRoutes: func(routes *gin.Engine) error {
			// middlewares
			routes.Use(middleware.RequestLoggingMiddleware("^/$", "^/assets/.*"))

			// static files
			dist, _ := fs.Sub(webui.StaticFS, "dist/assets")
			routes.StaticFS("/assets", http.FS(dist))

			// home index
			routes.GET("/", homeHandler)

			// handlers
			user := routes.Group("user")
			{
				user.GET("/list", userHandler.List)
			}
			cmd := routes.Group("cmd")
			{
				cmd.POST("exec", cmdHandler.Exec)
			}
			return nil
		},
	}
}
