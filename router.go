package main

import (
	"github.com/tsingson/fastweb/middle"
	"github.com/tsingson/fastweb/zaplogger"
	"github.com/tsingson/phi"
	"go.uber.org/zap"
)

func MuxRouter(log *zap.Logger) *phi.Mux {
	//
	var router *phi.Mux
	router = phi.NewRouter()
	//
	httpLogger := zaplogger.InitZapLogger(log )
	router.Use(httpLogger.FastHttpZapLogHandler)

	//
	router.Use(middle.Recoverer)
	router.Get("/", helloHandler)
	//

	router.Post("/api/vodcodeauthenticate", PlayAuth)
	router.Post("/api/livecodeauthenticate", PlayAuth)
	//
	return router
}

// design and code by tsingson
