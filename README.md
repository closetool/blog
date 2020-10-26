# BLOG

[![Build Status](https://www.travis-ci.org/closetool/blog.svg?branch=master)](https://www.travis-ci.org/closetool/blog) [![LICENSE](https://img.shields.io/github/license/closetool/blog)](https://github.com/closetool/blog)
---
对[plumemo](https://github.com/byteblogs168/plumemo)项目后端的重构，使其成为了微服务后端，配合对应前端使用[theme-react-sakura](https://github.com/byteblogs168/theme-react-sakura/)/[plumemo-admin](https://github.com/byteblogs168/plumemo-admin)(部分api需要修改，见下方)

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

## 部署

* 安装jdk8和go1.13.x
* git clone --depth=1 https://github.com/closetool/blog
* 给根目录和scripts目录下的脚本和support/config-server/gradlew添加执行权限(chmod a+x *.sh scripts/*.sh support/config-server/gradlew)
* 先修改根目录下的config.yml，config.yml描述了configserver的位置和config branch/profile
* 运行buildall.sh脚本
* docker stack deploy --compose-file docker-compose.yml '集群名' 

## 配置中心

* 配置中心中必有service_port,service_name
* 配置中心使用log_file_path(default:logs)和log_file_name(default:appName_localTime.log)指定log文件的位置和名字
* 使用log_level修改日志输出级别(default:4 info)

## Music Service
* 需要music_playlist_id(网易云歌单id)
* 服务端口配置为2599

## User Service
> * /social/v1/list \> /list/v1/social</br> 
> * /social/v1/info \> /info/v1/social</br>
> * /social/v1/socials \> /socials/v1/social

## Category Service

### category
> * /category-tags/v1/list \> /list/v1/category-tags</br>
> * /category/v1/list \> /list/v1/category
### tags
> * /tags/v1/list \> /list/v1/tags

## AMQP

### user service
* VerifyToken body为vo.AuthUser的json串，验证token的正确性，如果正确回传相应的用户对象

## FIXME
* docker swarm中不同node之间的同一overlay网络中的容器无法互相访问，nginx代理只能使用公网
* gin框架中不支持注册spring中的requestmapping，只能为路由多次注册不同方法
* gin框架中/info和/:id路由冲突