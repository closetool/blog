#!/bin/bash

docker build -t closetool/nginx support/nginx/
docker service rm nginx 
docker service create --name nginx -p 9000:9000 --replicas 1 --network my_network closetool/nginx

