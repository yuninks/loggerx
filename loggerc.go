package loggerx

import (
	"context"
	"fmt"
)

var loggerc *Logger

func init() {
	loggerc = NewLogger(context.Background())
}

func NewLoggerc(ctx context.Context, opts ...Option) {
	for _, apply := range opts {
		apply(&loggerc.option)
	}
}

func Channel(ch string) (r *Logger) {
	rr := *loggerc
	rr.channel = ch
	return &rr
}

func Info(ctx context.Context, v ...any) {
	loggerc.logger(ctx, "info", v...)
}

func Infof(ctx context.Context, format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	loggerc.logger(ctx, "info", s)
}

func Error(ctx context.Context, v ...any) {
	loggerc.logger(ctx, "error", v...)
}

func Errorf(ctx context.Context, format string, v ...any) {
	s := fmt.Sprintf(format, v...)
	loggerc.logger(ctx, "error", s)
}
