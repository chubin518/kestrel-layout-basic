package logging

import (
	"context"
	"log"
	"os"
	"sync"
	"time"

	"github.com/chubin518/kestrel-layout-basic/pkg/env"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

var once sync.Once

// Init init zap log
func Init(opts ...Option) {
	InitWithOptions(NewLoggingOptions(), opts...)
}

func InitWithOptions(options *LoggingOptions, opts ...Option) {
	once.Do(func() {
		for _, apply := range opts {
			apply(options)
		}
		level := zap.InfoLevel
		if options.Level != "" {
			zapLevel, err := zapcore.ParseLevel(options.Level)
			if err != nil {
				log.Println("invalid level, defaulting to info:", err)
			}
			level = zapLevel
		}

		hook := &lumberjack.Logger{
			Filename:   options.FileName,
			MaxSize:    options.MaxSize,
			MaxAge:     options.MaxAge,
			MaxBackups: options.MaxBackups,
			LocalTime:  options.LocalTime,
			Compress:   options.Compress,
		}

		wss := make([]zapcore.WriteSyncer, 0)
		if env.IsDevelopment() {
			wss = append(wss, zapcore.AddSync(os.Stdout))
		}
		wss = append(wss, zapcore.AddSync(hook))

		encoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
			TimeKey:       "timestamp",
			LevelKey:      "level",
			NameKey:       "logger",
			CallerKey:     "caller",
			MessageKey:    "msg",
			FunctionKey:   zapcore.OmitKey,
			StacktraceKey: zapcore.OmitKey,
			LineEnding:    zapcore.DefaultLineEnding,
			EncodeLevel:   zapcore.LowercaseLevelEncoder,
			EncodeTime: func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
				pae.AppendString(t.Format("2006-01-02 15:04:05.999"))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		})

		zlog := zap.New(
			zapcore.NewCore(
				encoder,
				zapcore.NewMultiWriteSyncer(wss...),
				zap.NewAtomicLevelAt(level),
			),
			zap.AddCaller(),
			zap.AddCallerSkip(1),
			zap.AddStacktrace(zap.ErrorLevel),
		)
		zlog = zlog.With(mapToFields(options.Fields)...)

		zap.ReplaceGlobals(zlog)

		Log = zlog
	})
}

func Debug(msg string) {
	Log.Debug(msg)
}

func Debugf(msg string, args ...any) {
	Log.Sugar().Debugf(msg, args...)
}

func DebugContext(ctx context.Context, msg string) {
	WithContext(ctx).Debug(msg)
}

func DebugfContext(ctx context.Context, msg string, args ...any) {
	WithContext(ctx).Sugar().Debugf(msg, args...)
}

func Info(msg string) {
	Log.Info(msg)
}

func Infof(msg string, args ...any) {
	Log.Sugar().Infof(msg, args...)
}

func InfoContext(ctx context.Context, msg string) {
	WithContext(ctx).Info(msg)
}

func InfofContext(ctx context.Context, msg string, args ...any) {
	WithContext(ctx).Sugar().Infof(msg, args...)
}

func Warn(msg string) {
	Log.Warn(msg)
}

func Warnf(msg string, args ...any) {
	Log.Sugar().Warnf(msg, args...)
}

func WarnContext(ctx context.Context, msg string) {
	WithContext(ctx).Warn(msg)
}

func WarnfContext(ctx context.Context, msg string, args ...any) {
	WithContext(ctx).Sugar().Warnf(msg, args...)
}

func Error(msg string) {
	Log.Error(msg)
}

func Errorf(msg string, args ...any) {
	Log.Sugar().Errorf(msg, args...)
}

func ErrorContext(ctx context.Context, msg string) {
	WithContext(ctx).Error(msg)
}

func ErrorfContext(ctx context.Context, msg string, args ...any) {
	WithContext(ctx).Sugar().Errorf(msg, args...)
}

func Fatal(msg string) {
	Log.Fatal(msg)
}

func Fatalf(msg string, args ...any) {
	Log.Sugar().Fatalf(msg, args...)
}

func FatalContext(ctx context.Context, msg string) {
	WithContext(ctx).Fatal(msg)
}

func FatalfContext(ctx context.Context, msg string, args ...any) {
	WithContext(ctx).Sugar().Fatalf(msg, args...)
}
