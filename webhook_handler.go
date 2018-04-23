package main

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/valyala/fasthttp"
)

// PushEvents 请求对应的实体
type PushEvents struct {
	After       string `json:"after"`
	Before      string `json:"before"`
	CheckoutSha string `json:"checkout_sha"`
	Commits     []struct {
		Added  []string `json:"added"`
		Author struct {
			Email string `json:"email"`
			Name  string `json:"name"`
		} `json:"author"`
		ID        string   `json:"id"`
		Message   string   `json:"message"`
		Modified  []string `json:"modified"`
		Removed   []string `json:"removed"`
		Timestamp string   `json:"timestamp"`
		URL       string   `json:"url"`
	} `json:"commits"`
	EventName  string `json:"event_name"`
	Message    string `json:"message"`
	ObjectKind string `json:"object_kind"`
	Project    struct {
		AvatarURL         string `json:"avatar_url"`
		CiConfigPath      string `json:"ci_config_path"`
		DefaultBranch     string `json:"default_branch"`
		Description       string `json:"description"`
		GitHTTPURL        string `json:"git_http_url"`
		GitSSHURL         string `json:"git_ssh_url"`
		Homepage          string `json:"homepage"`
		HTTPURL           string `json:"http_url"`
		Name              string `json:"name"`
		Namespace         string `json:"namespace"`
		PathWithNamespace string `json:"path_with_namespace"`
		SSHURL            string `json:"ssh_url"`
		URL               string `json:"url"`
		VisibilityLevel   int64  `json:"visibility_level"`
		WebURL            string `json:"web_url"`
	} `json:"project"`
	ProjectID  int64  `json:"project_id"`
	Ref        string `json:"ref"`
	Repository struct {
		Description     string `json:"description"`
		GitHTTPURL      string `json:"git_http_url"`
		GitSSHURL       string `json:"git_ssh_url"`
		Homepage        string `json:"homepage"`
		Name            string `json:"name"`
		URL             string `json:"url"`
		VisibilityLevel int64  `json:"visibility_level"`
	} `json:"repository"`
	TotalCommitsCount int64  `json:"total_commits_count"`
	UserAvatar        string `json:"user_avatar"`
	UserEmail         string `json:"user_email"`
	UserID            int64  `json:"user_id"`
	UserName          string `json:"user_name"`
	UserUsername      string `json:"user_username"`
}

func checkPermission(ctx *fasthttp.RequestCtx, xGitlabToken string) bool {
	havePermission := false

	if len(xGitlabToken) > 0 {
		ctx.Logger().Printf("%s", xGitlabToken)
		// 判断X-Gitlab-Token中的token是否与系统配置的一致
		if secretToken := strings.TrimSpace(xGitlabToken); secretToken == cfg.XGitlabToken {
			havePermission = true
		}
	}
	return havePermission
}

// WebhookHandler webhook to http response.
func WebhookHandler(ctx *fasthttp.RequestCtx) {
	ctx.Response.Reset()
	ctx.SetContentType("application/json; charset=utf-8")
	dingAccessToken := string(ctx.Request.URI().QueryArgs().Peek("dingAccessToken"))
	if len(dingAccessToken) > 0 {
		dingAccessToken = strings.TrimSpace(dingAccessToken)
	}
	ctx.Logger().Printf("dingAccessToken = %s", dingAccessToken)

	havePermission := checkPermission(ctx, string(ctx.Request.Header.Peek("X-Gitlab-Token")))
	if !havePermission {
		SendDinghook(dingAccessToken, "prd-generate", "Permission denied.")
		RenderError(ctx, "Permission denied.")
		return
	}

	prdMaps, projectSrc, err := parsePushEvents(ctx)
	if err != nil {
		SendDinghook(dingAccessToken, "prd-generate", err.Error())
		RenderError(ctx, "Permission denied.")
		return
	}
	// 后台线程运行上传，成功后通知钉钉
	go upload(prdMaps, projectSrc, dingAccessToken)

	RenderJSON(ctx, nil)
	return
}

func getFirstDirNodeName(path string) string {
	if strings.Contains(path, "/") {
		return strings.Split(path, "/")[0]
	}
	return path
}

func appendPRD(prdMap map[string]string, prdSlice []string) {
	for _, path := range prdSlice {
		prdName := getFirstDirNodeName(path)
		prdMap[prdName] = ""
	}
}

// parsePushEvents return prds, projectPath, error
func parsePushEvents(ctx *fasthttp.RequestCtx) (map[string]string, string, error) {
	var pushEvents PushEvents
	err := json.Unmarshal(ctx.Request.Body(), &pushEvents)
	if err != nil {
		return nil, "", errors.New("json unmarshal err" + err.Error())
	}

	prdMaps := map[string]string{}
	for _, c := range pushEvents.Commits {
		appendPRD(prdMaps, c.Added)
		appendPRD(prdMaps, c.Modified)
	}

	if len(prdMaps) == 0 {
		return nil, "", errors.New("没有需要更新的目录")
	}

	projectName := strings.Split(pushEvents.Project.PathWithNamespace, "/")[1]
	projectPath := cfg.PrdPath + "/" + projectName

	if PathIsExist(projectPath) {
		err := cmdRun(ctx, projectPath, "git", "pull")
		if err != nil {
			return prdMaps, "", errors.New("git pull 异常")
		}
	} else {
		err := cmdRun(ctx, cfg.PrdPath, "git", "clone", pushEvents.Repository.GitSSHURL)
		if err != nil {
			return prdMaps, "", errors.New("git clone 异常")
		}
	}
	return prdMaps, projectPath, nil
}
