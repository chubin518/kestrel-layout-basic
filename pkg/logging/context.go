package logging

import (
	"context"

	"go.uber.org/zap"
)

type loggerkey struct{}

// WithContext Returns a zap instance from the specified context
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return Log
	}
	if zlog, ok := ctx.Value(loggerkey{}).(*zap.Logger); ok {
		return zlog
	}
	return Log
}

// NewContext  Adds a field to the specified context
func NewContext(ctx context.Context, fields map[string]any) context.Context {
	return context.WithValue(ctx, loggerkey{}, WithContext(ctx).With(mapToFields(fields)...))
}
