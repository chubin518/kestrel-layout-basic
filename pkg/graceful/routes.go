package graceful

import (
	"context"

	"github.com/gin-gonic/gin"
)

type RoutesManager struct {
	routes []*Routes
}

type Routes struct {
	ConfigRoutes func(*gin.Engine) error
}

func (r *RoutesManager) Append(routes ...*Routes) {
	r.routes = append(r.routes, routes...)
}

func (r *RoutesManager) ConfigRoutes(engine *gin.Engine, ctx context.Context) error {
	for _, route := range r.routes {
		if err := ctx.Err(); err != nil {
			return err
		}
		if route.ConfigRoutes == nil {
			continue
		}
		if err := route.ConfigRoutes(engine); err != nil {
			return err
		}
	}
	return nil
}
