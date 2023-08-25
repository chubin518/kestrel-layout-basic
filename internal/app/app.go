package app

import (
	"github.com/chubin518/kestrel-layout-basic/pkg/graceful"
)

type App struct {
	Hooks  []*graceful.Hook
	Routes []*graceful.Routes
}

func NewApp(h *graceful.Hook, r *graceful.Routes) *App {
	return &App{
		Hooks:  []*graceful.Hook{h},
		Routes: []*graceful.Routes{r},
	}
}
