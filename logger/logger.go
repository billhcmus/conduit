package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"sync"
)

var (
	_doOnce sync.Once
)

// InitLogger return the zap logger instance
func InitLogger(level zapcore.Level) {
	_doOnce.Do(func() {
		ws, err := getWriteSyncer()
		if err != nil {
			log.Fatalf("Failed to init logger: %v", err)
		}
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, ws, level)
		zap.ReplaceGlobals(zap.New(core, zap.AddCaller()))
	})
}

// getWriteSyncer create the file use for the logger
func getWriteSyncer() (zapcore.WriteSyncer, error) {
	file, err := os.Create("conduit.log")
	if err != nil {
		return nil, err
	}

	return zap.CombineWriteSyncers(os.Stdout, file), nil
}

// getEncoder return the log encoder
func getEncoder() zapcore.Encoder {
	encoderConf := zap.NewProductionEncoderConfig()
	encoderConf.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConf.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewJSONEncoder(encoderConf)
}
