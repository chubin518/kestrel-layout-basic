package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// mapToFields
func mapToFields(args map[string]any) []zapcore.Field {
	list := make([]zapcore.Field, 0, len(args))
	for k, v := range args {
		list = append(list, zap.Any(k, v))
	}
	return list
}
