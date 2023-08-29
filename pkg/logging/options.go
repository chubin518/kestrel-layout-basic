package logging

import "maps"

type LoggingOptions struct {
	Fields     map[string]any `mapstructure:"fields"`
	Level      string         `mapstructure:"level"`
	FileName   string         `mapstructure:"file_name"`
	MaxSize    int            `mapstructure:"max_size"`
	MaxAge     int            `mapstructure:"max_age"`
	MaxBackups int            `mapstructure:"max_backups"`
	LocalTime  bool           `mapstructure:"local_time"`
	Compress   bool           `mapstructure:"compress"`
}

type Option func(*LoggingOptions)

func NewLoggingOptions() *LoggingOptions {
	return &LoggingOptions{
		Fields:     make(map[string]any),
		Level:      "info",
		FileName:   "logs/app.log",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 14,
		LocalTime:  false,
		Compress:   true,
	}
}

// WithFields
func WithFields(fields map[string]any) Option {
	return func(args *LoggingOptions) {
		maps.Copy(fields, args.Fields)
	}
}

// WithLevel
func WithLevel(level string) Option {
	return func(args *LoggingOptions) {
		args.Level = level
	}
}

// WithFileName
func WithFileName(fileName string) Option {
	return func(args *LoggingOptions) {
		args.FileName = fileName
	}
}

// WithMaxSize
func WithMaxSize(maxSize int) Option {
	return func(args *LoggingOptions) {
		args.MaxSize = maxSize
	}
}

// WithMaxAge
func WithMaxAge(maxAge int) Option {
	return func(args *LoggingOptions) {
		args.MaxAge = maxAge
	}
}

// WithMaxBackups
func WithMaxBackups(maxBackups int) Option {
	return func(args *LoggingOptions) {
		args.MaxBackups = maxBackups
	}
}

// WithLocalTime
func WithLocalTime(localTime bool) Option {
	return func(args *LoggingOptions) {
		args.LocalTime = localTime
	}
}

// WithCompress
func WithCompress(compress bool) Option {
	return func(args *LoggingOptions) {
		args.Compress = compress
	}
}
