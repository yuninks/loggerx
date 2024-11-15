package loggerx

import "context"

// 简单使用的日志接口
type LoggerInterface interface {
	Info(ctx context.Context, args ...any)
	Infof(ctx context.Context, format string, args ...any)
	Error(ctx context.Context, args ...any)
	Errorf(ctx context.Context, format string, args ...any)
}
