package app

import (
	"context"

	"github.com/chubin518/kestrel-layout-basic/pkg/graceful"
)

func NewAppHook() *graceful.Hook {
	hook := &graceful.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return nil
		},
	}
	return hook
}
