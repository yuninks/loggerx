package main

import (
	"context"

	"github.com/yuninks/loggerx"
)

func main() {
	ctx := context.Background()
	log := loggerx.NewLogger(ctx,
		// loggerx.SetPrintFile(false),
		loggerx.SetToConsole(),
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
