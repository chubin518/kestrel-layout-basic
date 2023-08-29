package graceful

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chubin518/kestrel-layout-basic/pkg/env"
	"github.com/chubin518/kestrel-layout-basic/pkg/logging"
	"github.com/chubin518/kestrel-layout-basic/pkg/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Graceful struct {
	engine        *gin.Engine
	options       *GracefulOptions
	lifecycle     *Lifecycle
	routesManager *RoutesManager
}

func (g *Graceful) RunWithContext(ctx context.Context) {
	g.engine = g.setupEngine()

	g.routesManager.ConfigRoutes(g.engine, ctx)

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", g.options.Port),
		Handler:        g.engine,
		ReadTimeout:    g.options.ReadTimeout,
		WriteTimeout:   g.options.WriteTimeout,
		MaxHeaderBytes: g.options.MaxHeaderBytes,
	}

	startCtx, cancel := context.WithTimeout(ctx, g.options.StartupTimeout)
	defer cancel()
	if err := withTimeout(startCtx, func(ctx context.Context) error {
		return g.lifecycle.Start(ctx)
	}); err != nil {
		logging.Errorf("Server startup failed: %v", err)
		return
	}

	logging.Infof("Starting server on ports %d ......", g.options.Port)

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logging.Errorf("Server listen failed: %v", err)
			return
		}
	}()

	quiet := make(chan os.Signal, 1)
	signal.Notify(quiet, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	sgl := <-quiet
	logging.Infof("Received shutdown signal %s ......", sgl.String())

	stopCtx, cancel := context.WithTimeout(ctx, g.options.ShutdownTimeout)
	defer cancel()
	if err := withTimeout(stopCtx, func(ctx context.Context) error {
		if err := srv.Shutdown(ctx); err != nil {
			return err
		}
		return g.lifecycle.Stop(ctx)
	}); err != nil {
		logging.Errorf("Server forced to shutdown: %v", err)
	}
	logging.Info("Server exiting")
}

func (g *Graceful) setupEngine() *gin.Engine {
	if env.IsDevelopment() {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	server := gin.New()

	server.RedirectTrailingSlash = true
	server.RedirectFixedPath = true
	server.HandleMethodNotAllowed = true
	server.ContextWithFallback = true

	if g.options.EnableCors {
		server.Use(cors.New(cors.Config{
			AllowAllOrigins: true,
			AllowHeaders:    []string{"*"},
			AllowMethods:    []string{"GET", "PUT", "POST", "PATCH", "DELETE", "OPTIONS"},
			MaxAge:          12 * time.Hour,
		}))
	}

	server.Use(middleware.RequestIdMiddleware())
	server.Use(middleware.RecoveryMiddleware())

	return server
}
