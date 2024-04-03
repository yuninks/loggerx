package loggerx_test

import (
	"context"
	"testing"
	"time"

	"code.yun.ink/pkg/loggerx"
)

// func TestMain(m *testing.M) {
// }

func TestLogger(t *testing.T) {

	l := loggerx.NewLogger(loggerx.SetErrorToInfo())

	l.Error(context.Background(), "test error")

	l.Channel("channel1").Error(context.Background(), "channel1 test error")
	l.Channel("channel2").Error(context.Background(), "channel2 test error")

	l.Info(context.Background(), "test info")

	time.Sleep(time.Second * 5)
}
