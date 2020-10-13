# BLOG

[![Build Status](https://www.travis-ci.org/closetool/blog.svg?branch=master)](https://www.travis-ci.org/closetool/blog) [![LICENSE](https://img.shields.io/github/license/closetool/blog)](https://github.com/closetool/blog)
---
对[plumemo](https://github.com/byteblogs168/plumemo)项目后端的重构，使其成为了微服务后端，配合对应前端使用[theme-react-sakura](https://github.com/byteblogs168/theme-react-sakura/)/[plumemo-admin](https://github.com/byteblogs168/plumemo-admin)

## 运用的技术

* docker
* docker swarm(管理集群，创建网络，健康检查，服务发现)
* [httpmock](https://github.com/jarcoal/httpmock)
* [convey](https://github.com/smartystreets/goconvey)
* travis
* [jsoniter](https://github.com/json-iterator/go)
* [gin](https://github.com/gin-gonic/gin)
* mysql
* redis(储存token，token过时视为登录失效)
* nginx(反向代理，微服务网关)
* rabbitmq(消息总线)

## 全局

* 配置中心地址（http://localhost:8888）,profile,branch通过配置文件配置
* 配置文件位于./,$HOME/config,/$HOME/${servicename}/ or /etc/${servicename}目录下，名字为config.yml或${servicename}.yml
* 配置中心中必有service_port,service_name
* log_file_path默认为./,log_file_name默认为${servicename}_${time}.log
* log_level默认为4(info)

## Music Service
* 特有config属性music_playlist_id(网易云歌单id)
* 服务端口配置为2599
## User Service
* 因为[issue#338](https://github.com/gin-gonic/gin/issues/388)使`/social/v1/socials` `/social/v1/info` `/social/v1/list`与`social/v1/:id`无法兼容，将前三个路由v1改为v2
* 需要数据库支持，先进入根目录运行./mysql.sh脚本搭建mysql环境