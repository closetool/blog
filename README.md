# BLOG

[![Build Status](https://www.travis-ci.org/closetool/blog.svg?branch=master)](https://www.travis-ci.org/closetool/blog)

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
