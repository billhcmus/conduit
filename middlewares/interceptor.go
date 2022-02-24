package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Intercept(ctx *gin.Context) {
	zap.L().Info("Intercept", zap.String("full_path", ctx.FullPath()))
	ctx.Next()
}
