#!/bin/bash


docker build -t closetool/mysql support/mysql/

docker tag closetool/mysql "$ALIYUN"closetool/mysql
docker push "$ALIYUN"closetool/mysql

docker service rm mysql
docker service create --name mysql -p 3306:3306 --replicas 1 --network my_network "$ALIYUN"closetool/mysql