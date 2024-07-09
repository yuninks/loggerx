package main

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/yuninks/loggerx"
	"github.com/yuninks/loggerx/middleware"
)

// curl --location '127.0.0.1:8080/ping'

func main() {
	ctx := context.Background()
	log := loggerx.NewLogger(ctx, loggerx.SetToConsole())

	g := gin.Default()

	g.Use(middleware.SetGinTraceIdByLogger(log))
	g.Use(middleware.SetGinParams(log))

	g.GET("/ping", func(ctx *gin.Context) {
		log.Infof(ctx, "GET /ping")
		ctx.JSON(200, gin.H{"message": "pong"})
	})

	g.Run(":8080")

}
