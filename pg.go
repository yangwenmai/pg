package main

import (
	"flag"
	"fmt"

	"github.com/valyala/fasthttp"
)

var (
	addr = flag.String("addr", "localhost:8080", "TCP address to listen to")

	xGitlabToken = flag.String("x_gitlab_token", "WVDjAqMavneLkNFyrRbqPR8seghHLtuza", "x-gitlab-token")
	baseURL      = flag.String("base_url", "", "visit prd online URL")

	qiniuBucket = flag.String("qiniu_bucket", "", "qiniu bucket")
	accessKey   = flag.String("access_key", "", "qiniu access_key")
	secretKey   = flag.String("secret_key", "", "qiniu secret_key")

	qshellPath = flag.String("qshell_path", "", "qshell command path")
	jsonPath   = flag.String("json_path", "", "qshell json path")
	prdPath    = flag.String("prd_path", "", "git clone prd path")
)

func main() {
	flag.Parse()

	cfg := &Config{
		Addr:         *addr,
		XGitlabToken: *xGitlabToken,
		BaseURL:      *baseURL,
		QiniuBucket:  *qiniuBucket,
		AccessKey:    *accessKey,
		SecretKey:    *secretKey,
		QshellPath:   *qshellPath,
		JSONPath:     *jsonPath,
		PrdPath:      *prdPath,
	}
	InitConfig(cfg)

	requestHandler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/v1/webhook/process":
			WebhookHandler(ctx)
		default:
			fmt.Println("Unknown request URI...")
		}
	}

	// Start HTTP server.
	if len(*addr) > 0 {
		fmt.Printf("Starting HTTP server on %q\n", *addr)
		go func() {
			if err := fasthttp.ListenAndServe(*addr, requestHandler); err != nil {
				fmt.Printf("error in ListenAndServe: %s", err)
			}
		}()
		select {}
	} else {
		fmt.Printf("Start HTTP Server failed.\n")
	}
}
