CUR_DIR=$(shell pwd)
BIN_DIR=${CUR_DIR}/bin
SRC_DIR=${CUR_DIR}/data
APP_NAME=product_analyzer
PACKAGE=${APP_NAME}/cmd
FILE?=
MOUNT=${SRC_DIR}/${FILE}:/app/${FILE}
ENV_FILE=FILE=${FILE}

run: docker_build
	docker run -p 8090:8090 -v ${MOUNT} -e ${ENV_FILE} --name ${APP_NAME} ${APP_NAME}

docker_build: build
	docker build -t ${APP_NAME} .

build: bin_dir
	go build -o ${BIN_DIR}/${APP_NAME} ${PACKAGE}

bin_dir:
	mkdir -p ${BIN_DIR}


