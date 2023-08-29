package graceful

import (
	"time"
)

type GracefulBuilder struct {
	options       *GracefulOptions
	lifecycle     *Lifecycle
	routesManager *RoutesManager
}

func NewBuilder(opts ...Option) *GracefulBuilder {
	options := &GracefulOptions{
		Port:            8080,
		ShutdownTimeout: 5 * time.Second,
		StartupTimeout:  5 * time.Second,
	}
	for _, apply := range opts {
		apply(options)
	}
	return NewBuilderWithOptions(options)
}

func NewBuilderWithOptions(options *GracefulOptions) *GracefulBuilder {
	return &GracefulBuilder{
		options:       options,
		lifecycle:     NewLifecycle(),
		routesManager: NewRoutesManager(),
	}
}

func (b *GracefulBuilder) UseHooks(hooks ...*Hook) *GracefulBuilder {
	b.lifecycle.Append(hooks...)
	return b
}

func (b *GracefulBuilder) UseRoutes(routes ...*Routes) *GracefulBuilder {
	b.routesManager.Append(routes...)
	return b
}

func (b *GracefulBuilder) Build() *Graceful {
	return &Graceful{
		options:       b.options,
		lifecycle:     b.lifecycle,
		routesManager: b.routesManager,
	}
}
