package graceful

import "time"

type GracefulOptions struct {
	Port            int           `mapstructure:"port"`
	MaxHeaderBytes  int           `mapstructure:"max_header_bytes"`
	EnableCors      bool          `mapstructure:"enable_cors"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	StartupTimeout  time.Duration `mapstructure:"startup_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type Option func(*GracefulOptions)

// WithPort
func WithPort(port int) Option {
	return func(args *GracefulOptions) {
		args.Port = port
	}
}

// WithMaxHeaderBytes
func WithMaxHeaderBytes(maxHeaderBytes int) Option {
	return func(args *GracefulOptions) {
		args.MaxHeaderBytes = maxHeaderBytes
	}
}

// WithEnableCors
func WithEnableCors(enableCors bool) Option {
	return func(args *GracefulOptions) {
		args.EnableCors = enableCors
	}
}

// WithReadTimeout
func WithReadTimeout(readTimeout time.Duration) Option {
	return func(args *GracefulOptions) {
		args.ReadTimeout = readTimeout
	}
}

// WithWriteTimeout
func WithWriteTimeout(writeTimeout time.Duration) Option {
	return func(args *GracefulOptions) {
		args.WriteTimeout = writeTimeout
	}
}

// WithStartupTimeout
func WithStartupTimeout(startupTimeout time.Duration) Option {
	return func(args *GracefulOptions) {
		args.StartupTimeout = startupTimeout
	}
}

// WithShutdownTimeout
func WithShutdownTimeout(shutdownTimeout time.Duration) Option {
	return func(args *GracefulOptions) {
		args.ShutdownTimeout = shutdownTimeout
	}
}
