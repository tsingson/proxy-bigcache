package main

import (
	"github.com/valyala/fasthttp"
)

// /api/vodcodeauthenticate
// /api/livecodeauthenticate

func PlayAuth(ctx *fasthttp.RequestCtx) {
	var payload string
	payload = string(ctx.PostBody())


	ctx.Response.SetBodyString(payload)
	return
}
