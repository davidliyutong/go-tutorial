package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rebirthmonkey/pkg/errors"
	"github.com/rebirthmonkey/pkg/gin/util"
	"github.com/rebirthmonkey/pkg/log"

	"github.com/rebirthmonkey/go/80_standards/20_coding/40_error/errcode"
)

func pingHandler2(ctx *gin.Context) {
	log.Info("pingHandler() called.")

	// Simulate an external error.
	err := errors.Errorf("auth error")
	if err != nil {
		log.Errorf("Error %+v", err)
		log.Panicf("Panic %+v", err)
		log.Info("xxxxxxxxxxxxxxxxxxxxxxx")
		util.WriteResponse(ctx, errors.WithCode(errcode.ErrInvalidAuthHeader, err.Error()), nil)
		return
	}
	util.WriteResponse(ctx, nil, "pong")
}

func main() {
	ginEngine := gin.Default()
	ginEngine.GET("/ping", pingHandler2)
	ginEngine.Run(":8080")
}
