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

  产品经理撰写需求文档---->提交到 gitlab----->触发 Gitlab Webhook---->触发 pg 的处理接口---->拉取项目需求文档到服务器---->使用 qshell 上传需求文档到 qiniu bucket---->通过钉钉机器人通知相关群组。

### 其他更多的应用场景 ###

最新公告，用户帮助，升级说明等。

### 详细操作步骤 ###

**对于新增产品需求文档，只需要做以下几件事情：**

  1. 点击 [New Project](https://xxx.gitlab.com/projects/new) ，创建一个产品需求文档的分组和具体产品需求文档项目：填写 group_name 和 project_name
  2. 点击具体的项目 Settings->Integrations 进入该项目的 Integrations Settings 中配置 Webhook，在 URL 填写 http://ip:port/v1/webhook/process?dingAccessToken=xxxxxxx (此 token 为钉钉机器人的 token，如果不需要可以不附带此参数)，在 Secret Token 填写 xxx，Trigger 勾选 Push events，点击 Add webhook 完成添加。
  3. 在 xxx\_prd 项目下创建产品需求文档目录，比如 xxx\_v1.0。
  4. 将需求文档数据提交到 gitlab 的 xxx\_v1.0 即可。

**对于增加产品需求文档小版本时，只需要做两件事情：**

  1. 在 xxx\_prd 项目的根目录 /product/xxx\_prd 下创建需求文档目录，目录以 xxx\_v1.1、xxx\_v2.0 此类方式命名。
  2. 需求文档数据添加入 xxx\_v1.1、xxx\_2.0 之后，Push 到 gitlab 中即可。

**对于日常的产品需求文档变更，只需要做一件事情：**

  * 需求文档修改完成之后，Push 到 gitlab 即可。

### 访问在线资源 ###

  访问 http://xxx.bkt.clouddn.com/xxx\_v1.0/ 等于访问到 https://xxx.gitlab.com/group_name/xxx\_prd/xxx\_v1.0/ 。
  比如，需求文档 test\_prd\_v1 的首页 `index.html` 位于 `https://xxx.gitlab.com/group_name/test\_prd/test\_prd\_v1/index.html` ，那么只要访问 http://xxx.bkt.clouddn.com/test\_prd\_v1/index.html 就行了。

### 项目结构 ###

  ├── product  
  │----├── test_prd  
  │----│----├── test_v1.0  
  │----│----├── test_v1.1  
  │----│----└── test_v2.0  
  │----└── xxx_prd  

### 技术点 ###
  
  - Webhook 如何监听对应的变更（新增或者更新）；
  - 上传七牛 CDN，如何封装 qshell 的调用；
  - 钉钉机器人通知相关群组（推送的信息：发布之后的 URL 地址）
  - 七牛 CDN PRD URL 需要增加时间戳鉴权

### 其他 ###

  增加一个统一的产品需求展示页面（URL 链接需要增加时间戳鉴权处理，防止泄露）
  增加一个可以方便配置 Webhook 的管理控制台；

## 遇到的问题/坑 ##

1. Open account file error, open ~/.qshell/account.json: no such file or directory, please use `account` to set AccessKey and SecretKey first

# 参考资料

1. [https://github.com/yangwenmai/how-to-add-badge-in-github-readme](https://github.com/yangwenmai/how-to-add-badge-in-github-readme)
