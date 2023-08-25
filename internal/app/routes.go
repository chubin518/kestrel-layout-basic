package app

import (
	"io/fs"
	"net/http"

	"github.com/chubin518/kestrel-layout-basic/internal/handler"
	"github.com/chubin518/kestrel-layout-basic/pkg/graceful"
	"github.com/chubin518/kestrel-layout-basic/pkg/middleware"
	"github.com/chubin518/kestrel-layout-basic/pkg/result"
	"github.com/chubin518/kestrel-layout-basic/webui"
	"github.com/gin-gonic/gin"
)

func NewRoutes(userHandler *handler.UserHandler) *graceful.Routes {
	return &graceful.Routes{
		ConfigRoutes: func(routes *gin.Engine) error {
			// middlewares
			routes.Use(middleware.RequestLoggingMiddleware("^/$", "^/assets/.*"))

			// static files
			dist, _ := fs.Sub(webui.StaticFS, "dist/assets")
			routes.StaticFS("/assets", http.FS(dist))

			// home index
			routes.GET("/", func(ctx *gin.Context) {
				rawIndex, _ := webui.StaticFS.ReadFile("dist/index.html")
				ctx.Writer.Write(rawIndex)
			})

			// handlers
			user := routes.Group("user")
			{
				user.GET("/list", userHandler.List)
			}

			api := routes.Group("/api")
			{
				api.GET("/popular/list", func(ctx *gin.Context) {
					list := []struct {
						Key      string `json:"key"`
						Expand   string `json:"expand"`
						Disabled bool   `json:"disabled"`
						IsLeaf   bool   `json:"isLeaf"`
					}{
						{
							Key:      "1",
							Expand:   "1",
							Disabled: true,
							IsLeaf:   true,
						},
						{
							Key:      "1",
							Expand:   "2",
							Disabled: true,
							IsLeaf:   true,
						},
						{
							Key:      "1",
							Expand:   "3",
							Disabled: true,
							IsLeaf:   true,
						},
					}

					result.OK.WithData(list).JSON(ctx)
				})
			}
			return nil
		},
	}
}
