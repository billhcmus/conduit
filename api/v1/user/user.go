package user

import (
	"github.com/billhcmus/conduit/logger"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

var log = logger.GetInstance()

func Logger(ctx *gin.Context) {
	t := time.Now()
	log.Info("Intercept", zap.String("full_path", ctx.FullPath()))
	ctx.Set("context", "my-value-context")
	ctx.Next()
	duration := time.Since(t)
	log.Info("Request time", zap.Int64("duration", duration.Microseconds()))
}

func RegisterUserRoutes(group *gin.RouterGroup) {
	group.Use(Logger)
	group.GET("/ping", Ping)
	group.GET("/user/:name", userRetrieve)
	group.GET("/user/:name/:action", userAction)
	group.GET("/user/groups", userGroup)
}

func userRetrieve(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(http.StatusOK, gin.H{"message": name})
}

func userAction(ctx *gin.Context) {
	name := ctx.Param("name")
	action := ctx.Param("action")
	ctx.JSON(http.StatusOK, gin.H{"message": name, "action": action})
}

func userGroup(ctx *gin.Context) {
	id := ctx.Query("id")
	label := ctx.Query("label")
	log.Info("Receive request", zap.String("id", id), zap.String("label", label))
	ctx.JSON(http.StatusOK, gin.H{"message": "xyz"})
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
}
