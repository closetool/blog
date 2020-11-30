# BLOG

[![Build Status](https://www.travis-ci.org/closetool/blog.svg?branch=master)](https://www.travis-ci.org/closetool/blog) [![LICENSE](https://img.shields.io/github/license/closetool/blog)](https://github.com/closetool/blog)
---
>**此项目未测试**  
我本想写出一个golang微服务博客后端, 但是写完后发现我的服务器带不动, 学生党也没有太多的钱花在服务器上, 我对这个项目的兴趣减少了一半.  
>初次写完后回头看, 发现以前的代码简直不堪入目, 很罗嗦, 不够精炼, 所以回头将前期写的模块进行了完全的重构. 重构过程中我学到的东西越来越少了, 所以就此结题, 开始新的篇章.  
>这个项目从十月初开始写, 写了两个月, 算是把我学到的后端开发知识都大体用了一次, 比以前有了更深的认识.  
>很感谢大佬写的[微服务教程](https://segmentfault.com/a/1190000014894854), 我获益匪浅.  

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
* redis(储存token，token过时视为登录失效, 未完成)
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
> * /byteblogs/email/v1/send \> /auth/email/v1/send

## Category Service

> * /comments/v1/get \> /get/v1/comments
> * /category-tags/v1/list \> /list/v1/category-tags
> * /category/v1/list \> /list/v1/category
> * /tags/v1/list \> /list/v1/tags

### category

> * /category-tags/v1/list \> /list/v1/category-tags</br>
> * /category/v1/list \> /list/v1/category

### tags

> * /tags/v1/list \> /list/v1/tags

## Posts Service

### posts

> * /posts/v1/list \> /list/v1/posts

### comments

> * /comments/v1/get \> /get/v1/comments

## Menu Service

> * /menu/v1/list \> /list/v1/menu

## Link Service

> * /link/v1/list \> /list/v1/link

## FIXME

* docker swarm中不同node之间的同一overlay网络中的容器无法互相访问，nginx代理只能使用公网
* gin框架中不支持注册spring中的requestmapping，只能为路由多次注册不同方法
* gin框架中/info和/:id路由冲突