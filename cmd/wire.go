//go:build wireinject
// +build wireinject

package main

import (
	"github.com/chubin518/kestrel-layout-basic/internal/app"
	"github.com/chubin518/kestrel-layout-basic/internal/handler"
	"github.com/chubin518/kestrel-layout-basic/internal/service"
	"github.com/google/wire"
)

var ServiceSet = wire.NewSet(
	service.NewUserService,
)

var HandlerSet = wire.NewSet(
	handler.NewUserHandler,
	handler.NewCmdHandler,
)

var RoutesSet = wire.NewSet(
	app.NewRoutes,
)

var HooksSet = wire.NewSet(
	app.NewAppHook,
)

var AppSet = wire.NewSet(
	app.NewApp,
)

func newApp() (*app.App, error) {
	panic(wire.Build(ServiceSet, HandlerSet, RoutesSet, HooksSet, AppSet))
}
