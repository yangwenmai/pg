package main

// Config config
type Config struct {
	Addr         string
	XGitlabToken string `yaml:"x_gitlab_token"`
	BaseURL      string `yaml:"base_url"`
	QiniuBucket  string `yaml:"qiniu_bucket"`
	AccessKey    string `yaml:"access_key"`
	SecretKey    string `yaml:"secret_key"`
	QshellPath   string `yaml:"qshell_path"`
	JSONPath     string `yaml:"json_path"`
	PrdPath      string `yaml:"prd_path"`
}

var cfg *Config

// InitConfig init config
func InitConfig(config *Config) {
	cfg = config
}
