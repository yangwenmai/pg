# pg
[![Build Status](https://travis-ci.org/yangwenmai/pg.svg?branch=master)](https://travis-ci.org/yangwenmai/pg) [![Go Report Card](https://goreportcard.com/badge/github.com/yangwenmai/pg)](https://goreportcard.com/report/github.com/yangwenmai/pg)  [![Documentation](https://godoc.org/github.com/yangwenmai/pg?status.svg)](http://godoc.org/github.com/yangwenmai/pg) [![Coverage Status](https://coveralls.io/repos/github/yangwenmai/pg/badge.svg?branch=master)](https://coveralls.io/github/yangwenmai/pg?branch=master) [![GitHub issues](https://img.shields.io/github/issues/yangwenmai/pg.svg?label=Issue)](https://github.com/yangwenmai/pg/issues) [![license](https://img.shields.io/github/license/yangwenmai/pg.svg)](https://github.com/yangwenmai/pg/blob/master/LICENSE) [![Release](https://img.shields.io/github/release/yangwenmai/pg.svg?label=Release)](https://github.com/yangwenmai/pg/releases) [![star this repo](http://githubbadges.com/star.svg?user=yangwenmai&repo=pg)](http://github.com/yangwenmai/pg) [![fork this repo](http://githubbadges.com/fork.svg?user=yangwenmai&repo=pg)](http://github.com/yangwenmai/pg/fork)

## Stargazers over time

[![Stargazers over time](https://starcharts.herokuapp.com/yangwenmai/pg.svg)](https://starcharts.herokuapp.com/yangwenmai/pg)

一步一步教你搭建属于你或你们团队的在线文档。

>研发团队的工程效率的实践。

关键词：`Axure`, `Gitlab`, `Github`, `七牛`, `fasthttp`.

研发团队的工程效率实践，现在越来越多的人开始谈论这个话题，但是说实在的可操作性不大，或者说操作起来难度不小，毕竟每家公司都有自己的实际情况，很难可以直接套用。

本场 Chat 侧重于实践，不会有抽象的概念和理论知识。我会直接拟一个场景来实战《如何通过Gitlab+七牛云存储来构建在线产品需求文档》，一步一步的带领大家构建属于你自己或者你们团队的工程效率实践。

本场 Chat 您将学到如下内容：

- 了解 Axure 产品原型工具
- 了解 Gitlab/Github webhook
- 了解 七牛云存储 CDN 分发
- 基于 fasthttp 开发一个最简单的 HTTP Server
- 如何设计一个工程效率实践的产品
- 基于 Gitlab+七牛云存储 构建产品需求在线文档
- 给你的产品赋予生命力（整合钉钉/Slack等机器人）

### 目标 ###

1. 快速将 HTML 页面生成一个可预览的在线地址；
2. HTML 页面也能够被版本管理系统管理起来（Gitlab, Github, Coding, 码云等）；
  - 每一次改动都有迹可循，别以为你偷偷改产品文档我们就不知道了。
3. 能够及时同步新增和变更需求；

### 典型的应用场景 ###

产品经理的在线文档，简单的流程如下：

+ 产品经理撰写需求文档
+ 提交到 gitlab
+ 触发 Gitlab Webhook
+ 触发 pg 的处理接口(拉取项目需求文档到服务器-使用 qshell 上传需求文档到 qiniu bucket)
+ 通过钉钉机器人通知相关群组。

### 其他更多的应用场景 ###

最新公告，用户帮助，升级说明等。

**对于日常的产品需求文档变更，只需要做一件事情：**

  * 需求文档修改完成之后，Push 到 gitlab 即可。

## 安装

```shell
go build
```

## 运行


```shell
./pg -base_url "http://xxxxx.bkt.clouddn.com/" -qiniu_bucket "my-pg" -access_key "xxxxxx" -secret_key "xxxxxx" -qshell_path "/Users/yourname/xxx_tools/qshell-v2.1.7/qshell" -json_path "/Users/yourname/xxx_data/pg-test/json" -prd_path "/Users/yourname/xxx_data/pg-test/prd"
```

## 遇到的问题/坑 ##

1. Open account file error, open ~/.qshell/account.json: no such file or directory, please use `account` to set AccessKey and SecretKey first

# 参考资料

1. [一键生成 Github Go 项目 - gpt](https://github.com/yangwenmai/gpt)
