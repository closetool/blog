#!/bin/bash

cp ../config.yml ../services/categoryservice/

cd ../utils/healthutils/
CGO_ENABLED=0 go build -o healthchecker-linux-amd64
cp ./healthchecker-linux-amd64 ../../services/categoryservice
echo "built healthchecker-linux-amd64 in `pwd`"

cd ../../services/categoryservice/
CGO_ENABLED=0 go build -o categoryservice-linux-amd64
echo "built categoryservice-linux-amd64 in `pwd`"


docker build -t closetool/categoryservice ./

#echo "pushing images to aliyun"
#docker tag closetool/categoryservice "$ALIYUN"closetool/categoryservice
#docker push "$ALIYUN"closetool/categoryservice

docker service rm categoryservice 
docker service create --network my_network --replicas 1 --name categoryservice -p 2600:2600 "$ALIYUN"closetool/categoryservice
rm config.yml
