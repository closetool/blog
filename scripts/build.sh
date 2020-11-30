#!/bin/bash

home=`pwd`

for i in $@
do
	servicename=${i%:*}
	port=${i#*:}

	if [ -z $servicename ] 
	then 
		echo 'please input a service name'
		exit 1
	fi

	cp ../config.yml ../services/"$servicename"/

	cd ../utils/healthutils/
	CGO_ENABLED=0 go build -o healthchecker-linux-amd64
	cp ./healthchecker-linux-amd64 ../../services/"$servicename"
	echo "built healthchecker-linux-amd64 in `pwd`"

	cd ../../services/"$servicename"/
	CGO_ENABLED=0 go build -o "$servicename"-linux-amd64
	echo "built $servicename-linux-amd64 in `pwd`"

	docker build -t closetool/"$servicename" ./

	rm config.yml
	cd $home
done
