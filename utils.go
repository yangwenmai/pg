package main

import (
	"encoding/json"
	"fmt"

	"github.com/valyala/fasthttp"
	"github.com/yangwenmai/pg/dinghook"
)

// RenderError render error to json
func RenderError(ctx *fasthttp.RequestCtx, err interface{}) {
	o := err
	if err1, ok := err.(error); ok {
		o = err1.Error()
	}
	result := map[string]interface{}{
		"success": false,
		"error":   o,
	}
	bs, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	ctx.Logger().Printf("response: %s", string(bs))
	ctx.Write(bs)
}

// RenderJSON render to json
func RenderJSON(ctx *fasthttp.RequestCtx, data interface{}) {
	result := map[string]interface{}{
		"success": true,
		"data":    data,
	}
	bs, err := json.Marshal(result)
	if err != nil {
		panic(err)
	}
	ctx.Logger().Printf("response: %s", string(bs))
	ctx.Write(bs)
}

// SendDinghook 钉钉发送信息
func SendDinghook(dingAccessToken string, title string, content string) {
	fmt.Println("SendDinghook...")
	ding := dinghook.Ding{AccessToken: dingAccessToken}
	markdown := dinghook.Markdown{Title: title, Content: content}
	ding.Send(markdown)
}
