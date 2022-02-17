package logger

import (
	"go.uber.org/zap"
	"log"
	"sync"
)

var doOnce sync.Once
var logger *zap.Logger

func GetInstance() *zap.Logger {
	doOnce.Do(func() {
		instance, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("Failed to initialize logger: %v", err)
		}
		log.Println("Initialize complete")
		logger = instance
	})

	return logger
}
