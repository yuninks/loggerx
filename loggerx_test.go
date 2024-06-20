package loggerx_test

import (
	"bytes"
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/yuninks/loggerx"
)

// func TestMain(m *testing.M) {
// }

func TestLogger(t *testing.T) {

	b := bytes.NewBuffer(nil)

	l := loggerx.NewLogger(
		context.Background(),
		loggerx.SetErrorToInfo(),
		loggerx.SetToConsole(),
		loggerx.SetTimeZone(time.UTC),
		// loggerx.SetExtraDriver(b, Print{}),
		loggerx.SetFileSplit(loggerx.FileSplitTimeA),
	)

	l.Error(context.Background(), "test error")

	l.Channel("channel1").Error(context.Background(), "channel1 test error")
	l.Channel("channel2").Error(context.Background(), "channel2 test error")

	l.Info(context.Background(), "test info")

	fmt.Println(b.String())

	time.Sleep(time.Second * 5)
}

type Print struct {
}

func (pp Print) Write(p []byte) (n int, err error) {
	fmt.Print("ppppppppppppppp", string(p))
	return
}
