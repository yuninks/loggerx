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
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
	log.Info(ctx, "哈哈哈2")
}
