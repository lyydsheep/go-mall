package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	ctx     context.Context
	traceId string
	spanId  string
	pSpanId string
	_logger *zap.Logger
}

func NewLogger(ctx context.Context) *Logger {
	var l Logger
	if traceId := ctx.Value("traceId"); traceId != nil {
		l.traceId = traceId.(string)
	}
	if spanId := ctx.Value("spanId"); spanId != nil {
		l.spanId = spanId.(string)
	}
	if pSpanId := ctx.Value("pSpanId"); pSpanId != nil {
		l.pSpanId = pSpanId.(string)
	}
	l._logger = _logger
	_logger.Info("msg", zap.Any("key", "value"))
	return &l
}

func (l *Logger) log(level zapcore.Level, msg string, kv ...any) {
	// k-v成对
	if len(kv)%2 != 0 {
		kv = append(kv, "Unknown")
	}
	// 转化为zap.Field类型
	fields := make([]zap.Field, 0, len(kv)>>1)
	for i := 0; i < len(kv); i += 2 {
		// fmt.Sprintf("%v") ---> 字符串类型
		fields = append(fields, zap.Any(fmt.Sprintf("%v", kv[i]), kv[i+1]))
	}
	if ce := l._logger.Check(level, msg); ce != nil {
		ce.Write(fields...)
	}
}

func (l *Logger) Info(msg string, kv ...any) {
	l.log(zapcore.InfoLevel, msg, kv)
}

func (l *Logger) Debug(msg string, kv ...any) {
	l.log(zapcore.DebugLevel, msg, kv)
}

func (l *Logger) Warn(msg string, kv ...any) {
	l.log(zapcore.WarnLevel, msg, kv)
}

func (l *Logger) Error(msg string, kv ...any) {
	l.log(zapcore.ErrorLevel, msg, kv...)
}
