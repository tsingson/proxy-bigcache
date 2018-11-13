package main

import (
	"net"

	"github.com/tsingson/fastweb/zaplogger"
	"go.uber.org/zap"

	"github.com/tsingson/phi"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/reuseport"
)

func FasthttpServ(addr string, log *zap.Logger) {
	var (
		listener  net.Listener
		err, err1 error
		// 	fastLogger FastLogger
		router *phi.Mux
	)

	// fastLogger.Logger = log
	router = MuxRouter(log)

	// reuse port
	listener, err = reuseport.Listen("tcp4", addr)
	if err != nil {
		log.Info("working in Microsoft Windows" + addr)
		// for windows
		listener, err1 = net.Listen("tcp", addr)
		if err1 != nil {
			log.Fatal("tcp connect error", zap.Error(err1))
			panic("tcp connect error")
		}
	}

	// run fasthttp serv
	go func() {
		// fasthttp server setting here
		s := &fasthttp.Server{
			Handler:            router.ServeFastHTTP,
			Name:               ServerName,
			ReadBufferSize:     BufferSize,
			MaxConnsPerIP:      10,
			MaxRequestsPerConn: 10,
			MaxRequestBodySize: 1024 * 4, // MaxRequestBodySize: 100<<20, // 100MB
			Concurrency:        MaxFttpConnect,
			Logger:             zaplogger.InitZapLogger(log),
		}
		if err = s.Serve(listener); err != nil {
			log.Panic("fasthttp running error", zap.Error(err))
			panic("fasthttp running error")
		}
	}()
	log.Info("fasthttpgx server start success  " + addr)

}

// design and code by tsingson
