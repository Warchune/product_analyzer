CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
APP=product_analyzer
PACKAGE=${APP}/cmd
FILE?=

run: docker_build
	docker run -p 8090:8090 -v ${BINDIR}/${FILE}:/app/${FILE} -e FILE=${FILE} --name ${APP} ${APP}

docker_build: build
	docker build -t ${APP} .

build: bindir
	go build -o ${BINDIR}/${APP} ${PACKAGE}

bindir:
	mkdir -p ${BINDIR}


