package loggerx

import "context"

type Loggerx interface {
	Info(ctx context.Context, args ...any)
	Infof(ctx context.Context, format string, args ...any)
	Error(ctx context.Context, args ...any)
	Errorf(ctx context.Context, format string, args ...any)
}
