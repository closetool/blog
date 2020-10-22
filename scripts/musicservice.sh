#!/bin/bash

#server=http://39.108.114.242:8888
#pro=test
#bra=blog

cp ../config.yml ../services/musicservice/

cd ../utils/healthutils/
CGO_ENABLED=0 go build -o healthchecker-linux-amd64
cp ./healthchecker-linux-amd64 ../../services/musicservice
echo "built healthchecker-linux-amd64 in `pwd`"

cd ../../services/musicservice/
CGO_ENABLED=0 go build -o musicservice-linux-amd64
echo "built musicservice-linux-amd64 in `pwd`"


docker build -t closetool/musicservice ./

#echo "pushing images to aliyun"
#docker tag closetool/musicservice "$ALIYUN"closetool/musicservice
#docker push "$ALIYUN"closetool/musicservice

docker service rm musicservice
docker service create --network my_network --replicas 1 --name musicservice -p 2599:2599 "$ALIYUN"closetool/musicservice
rm config.yml
#echo "exec 'docker service create -e CONFIG_SERVER="$server" -e PROFILE="$pro" -e BRANCH="$bra" --replicas 1 --name musicservice -p 2599:2599 closetool/musicservice'"
#docker service create -e CONFIG_SERVER="$server" -e PROFILE="$pro" -e BRANCH="$bra" --replicas 1 --name musicservice -p 2599:2599 closetool/musicservice
#CONFIG_SERVER=http://39.108.114.242/8888 \
#	PROFILE=test BRANCH=blog \
#	docker run -e CONFIG_SERVER=`$CONFIG_SERVER` \
#	-e PROFILE=`$PROFILE` -e BRANCH=`$BRANCH` \
#	--name musicservice \
#	-p 2599:2599 \
#	closetool/musicservice
