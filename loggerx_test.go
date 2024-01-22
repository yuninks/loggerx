package loggerx_test

import (
	"context"
	"testing"

	"code.yun.ink/pkg/loggerx"
)

// func TestMain(m *testing.M) {
// }

func TestLogger(t *testing.T) {

	l := loggerx.NewLogger(loggerx.SetErrorToInfo())

	l.Error(context.Background(), "test error")

	l.Channel("test").Error(context.Background(), "test error")

	l.Info(context.Background(), "test info")

}
