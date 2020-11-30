#!/bin/bash

set -e

#rabbitmq
docker build -t closetool/rabbitmq support/rabbitmq/

#configserver
cd support/config-server
./gradlew build
cd ../..
docker build -t closetool/configserver support/config-server/

#nginx
docker build -t closetool/nginx support/nginx/

#mysql
docker build -t closetool/mysql support/mysql/

cd scripts
bash build.sh `cat services.txt`