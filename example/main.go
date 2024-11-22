package main

import (
	"context"
	"time"

	"github.com/yuninks/loggerx"
)

func main() {
	ctx := context.Background()
	log := loggerx.NewLogger(ctx,
		// loggerx.SetPrintFile(false),
		loggerx.SetToConsole(),
		// loggerx.SetTimeZone(time.UTC),
		loggerx.SetTimeZone(time.FixedZone("CST", 8*3600)),
	)
	log.WriteAsync().Info(ctx, "哈哈哈2异步")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")

	for i := 0; i < 100; i++ {
		log.WriteAsync().Infof(ctx, "异步 %d", i)
	}

	time.Sleep(time.Second)
}
