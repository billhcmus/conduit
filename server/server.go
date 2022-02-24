package server

import (
	"context"
	"errors"
	"github.com/billhcmus/conduit/config"
	"github.com/billhcmus/conduit/middlewares"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type Http struct {
	router *gin.Engine
	server *http.Server
	config config.ServerConfig
	groups map[string]*gin.RouterGroup
	value1 int
}

func New(config config.ServerConfig, opts ...Option) *Http {
	router := gin.New()
	router.Use(middlewares.Intercept)
	server := &http.Server{
		Addr:    config.Addr,
		Handler: router,
	}

	h := &Http{
		router: router,
		server: server,
		config: config,
		groups: make(map[string]*gin.RouterGroup),
	}
	for _, opt := range opts {
		opt.apply(h)
	}
	return h
}

func (h *Http) Start() {
	zap.L().Info("Server ready to serve", zap.String("address", h.config.Addr))
	if err := h.server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		zap.L().Fatal("Failed to start server", zap.Error(err))
	}
}

func (h *Http) Group(path string) *gin.RouterGroup {
	if group, found := h.groups[path]; found {
		return group
	}
	h.groups[path] = h.router.Group(path)
	return h.groups[path]
}

func (h *Http) Stop() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.server.Shutdown(ctx); err != nil {
		zap.L().Error("Failed to shutdown server", zap.Error(err))
	}
	zap.L().Info("Server shutdown complete")
}
