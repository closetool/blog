#!/bin/bash

# RabbitMQ
docker service rm rabbitmq
docker build -t closetool/rabbitmq support/rabbitmq/
docker service create --name=rabbitmq --replicas=1 --network=my_network -p 1883:1883 -p 5672:5672 -p 15672:15672 closetool/rabbitmq

# Config Server
cd support/config-server
./gradlew build
cd ../..
docker build -t closetool/configserver support/config-server/
docker service rm configserver
docker service create --replicas 1 --name configserver -p 8888:8888 --network my_network --update-delay 10s --with-registry-auth  --update-parallelism 1 closetool/configserver
