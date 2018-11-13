package main

import (
	"fmt"

	"github.com/valyala/fasthttp"
)

func helloHandler(ctx *fasthttp.RequestCtx) {
	//	fmt.Fprintf(ctx, "Hello, world!\n\n")
	/**
	fmt.Fprintf(ctx, "Request method is %q\n", ctx.Method())
	fmt.Fprintf(ctx, "RequestURI is %q\n", ctx.RequestURI())
	fmt.Fprintf(ctx, "Requested path is %q\n", ctx.Path())
	fmt.Fprintf(ctx, "Host is %q\n", ctx.Host())
	fmt.Fprintf(ctx, "Query string is %q\n", ctx.QueryArgs())
	*/
	ctx.SetContentType("text/plain; charset=utf8")

	// fmt.Fprintf(ctx, "User-Agent is %q\n", ctx.UserAgent())
	// fmt.Fprintf(ctx, "Connection has been established at %s\n", ctx.ConnTime())
	// fmt.Fprintf(ctx, "Request has been started at %s\n", ctx.Time())
	// fmt.Fprintf(ctx, "Serial request number for the current connection is %d\n", ctx.ConnRequestNum())
	fmt.Fprintf(ctx, "Your ip is %q\n\n", ctx.RemoteIP())

	// Set arbitrary headers
	//	ctx.Response.Header.Set("X-My-Header", "my-header-value")
	// Set cookies
	/**
	var c fasthttp.Cookie
	c.SetKey("cookie-name")
	c.SetValue("cookie-value")
	ctx.Response.Header.SetCookie(&c)
	*/
	return
}

// design and code by tsingson
