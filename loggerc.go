package loggerx

import "context"

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
	loggerc.Info(ctx, v...)
}

func Infof(ctx context.Context, format string, v ...any) {
	loggerc.Infof(ctx, format, v...)
}

func Error(ctx context.Context, v ...any) {
	loggerc.Error(ctx, v...)
}

func Errorf(ctx context.Context, format string, v ...any) {
	loggerc.Errorf(ctx, format, v...)
}
