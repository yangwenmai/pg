package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	"github.com/valyala/fasthttp"
)

const (
	// JSONTemplate 模板
	JSONTemplate = `{
  "src_dir" : "%s",
  "access_key": "%s",
  "secret_key": "%s",
  "bucket" : "%s",
  "key_prefix": "%s/",
  "skip_file_prefixes": ".git",
  "skip_path_prefixes": ".git",
  "skip_suffixes": ".DS_Store",
  "overwrite": true,
  "rescan_local": true
}`

	// MarkdownSuccessTemplate success 模板
	MarkdownSuccessTemplate = "产品文档已经更新，点击前往：[[%s]](%s)\n\n"
	// MarkdownFailureTemplate failure 模板
	MarkdownFailureTemplate = "产品文档更新失败：[%s][%s]\n\n"
)

// PathIsExist 判断目录/文件是否存在
func PathIsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// cmdRun exec command
func cmdRun(ctx *fasthttp.RequestCtx, path string, name string, args ...string) error {
	printCmd(ctx, path, name, args...)

	cmd := exec.Command(name, args...)
	cmd.Dir = path
	out, err := cmd.CombinedOutput()

	if err != nil {
		ctx.Logger().Printf("cmd.Run() failed with %s", err)
		return err
	}
	ctx.Logger().Printf("combined out:\n%s", string(out))
	return nil
}

// printCmd print command
func printCmd(ctx *fasthttp.RequestCtx, path string, name string, args ...string) {
	var buffer bytes.Buffer
	buffer.WriteString(name)
	for _, v := range args {
		buffer.WriteString(" ")
		buffer.WriteString(v)
	}
	ctx.Logger().Printf(buffer.String())
	ctx.Logger().Printf("cmd.Run() run in %s", path)
}

func upload(prdMaps map[string]string, projectSrc string, dingAccessToken string) {
	fmt.Println("start upload...")
	content := ""
	for prdName := range prdMaps {
		prdJSONPath := cfg.JSONPath + "/" + prdName + ".json"
		// JSON文件不存在则新建
		if !PathIsExist(prdJSONPath) {
			prdPath := projectSrc + "/" + prdName
			err := createJSON(prdName, prdJSONPath, prdPath)
			if err != nil {
				content += fmt.Sprintf(MarkdownFailureTemplate, prdName, "新建 JSON 失败\n"+err.Error())
				continue
			}
		}
		// 七牛上传
		err := quploadRun(prdJSONPath)
		if err == nil {
			url := cfg.BaseURL + prdName + "/index.html"
			content += fmt.Sprintf(MarkdownSuccessTemplate, prdName, url)
		} else {
			content += fmt.Sprintf(MarkdownFailureTemplate, prdName, "调用 Qupload 失败\n"+err.Error())
		}
	}

	// 钉钉机器人
	if dingAccessToken != "" {
		SendDinghook(dingAccessToken, "prd-generate", content)
	}
}

func createJSON(prdName string, prdJSONPath string, prdPath string) error {
	out, err := os.Create(prdJSONPath)
	defer out.Close()
	if err != nil {
		return err
	}
	out.WriteString(fmt.Sprintf(JSONTemplate, prdPath, cfg.AccessKey, cfg.SecretKey, cfg.QiniuBucket, prdName))
	return nil
}

// quploadRun 执行 qshell 的 qupload 命令，将产品文档的文件同步到七牛，并发上传的协程数量为 4。
func quploadRun(prdJSONPath string) error {
	cmd := exec.Command(cfg.QshellPath, "qupload", "4", prdJSONPath)
	_, err := cmd.CombinedOutput()
	return err
}
