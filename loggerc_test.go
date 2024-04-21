package loggerx_test

import (
	"context"
	"testing"

	"github.com/yuninks/loggerx"
)

func TestLoggerc(t *testing.T) {
	loggerx.Info(context.Background(), "hhhhhh")
}
