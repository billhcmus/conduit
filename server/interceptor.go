package server

import (
	"github.com/billhcmus/conduit/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Intercept(ctx *gin.Context) {
	log := logger.GetInstance()
	log.Info("Intercept", zap.String("full_path", ctx.FullPath()))
	ctx.Next()
}
