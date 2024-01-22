package loggerx_test

import (
	"context"
	"testing"

	"code.yun.ink/pkg/loggerx"
)

// func TestMain(m *testing.M) {
// }

func TestLogger(t *testing.T) {

	l := loggerx.InitLogger("profix")

	l.Error(context.Background(), "test error")

}
