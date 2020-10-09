#!/bin/bash

docker service rm mysql
docker build -t closetool/mysql support/mysql/
docker service create --name mysql -p 3306:3306 --replicas 1 --network my_network closetool/mysql
