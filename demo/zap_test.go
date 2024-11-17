package demo

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"testing"
)

func TestZap(t *testing.T) {
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	defaultLogLevel := zapcore.DebugLevel
	logFile, _ := os.OpenFile("./log-test-zap.json", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 06666)
	writer := zapcore.AddSync(logFile)
	logger := zap.New(zapcore.NewCore(fileEncoder, writer, defaultLogLevel), zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	defer logger.Sync()
	url := "https://www.test.com"
	logger.Info("write log to file", zap.String("url", url), zap.Int("atempt", 3))

}
