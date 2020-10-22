#!/bin/bash

cp ../config.yml ../services/userservice/

cd ../utils/healthutils/
CGO_ENABLED=0 go build -o healthchecker-linux-amd64
cp ./healthchecker-linux-amd64 ../../services/userservice
echo "built healthchecker-linux-amd64 in `pwd`"

cd ../../services/userservice/
CGO_ENABLED=0 go build -o userservice-linux-amd64
echo "built userservice-linux-amd64 in `pwd`"


docker build -t closetool/userservice ./

#echo "pushing images to aliyun"
#docker tag closetool/userservice "$ALIYUN"closetool/userservice
#docker push "$ALIYUN"closetool/userservice

docker service rm userservice 
docker service create --network my_network --replicas 1 --name userservice -p 2600:2600 "$ALIYUN"closetool/userservice
rm config.yml
