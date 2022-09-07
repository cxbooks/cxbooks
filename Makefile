# Makefile.
#
# Create by: tjuliuyou At 2022-07-24
#
#

PROJECT_PATH=$(shell cd "$(dirname "$0" )" &&pwd)
PROJECT_NAME=$(shell basename "$(PWD)")
VERSION=$(shell git describe --tags | sed 's/\(.*\)-.*/\1/')
BUILD_DATE=$(shell date -u '+%Y-%m-%d_%I:%M:%S%p')
BUILD_HASH=$(shell git rev-parse HEAD)
LDFLAGS="-X github.com/cxbooks/cxbooks.buildstamp=${BUILD_DATE} -X github.com/cxbooks/cxbooks.githash=${BUILD_HASH} -X github.com/cxbooks/cxbooks.VERSION=${VERSION} -s -w"
DESTDIR=${PROJECT_PATH}/build

TARGETS = cxbooks


.PHONY: all

export

all : cxbooks

cxbooks: 
	@echo "创建 cxbooks-${VERSION}目录"
	@#debian上直接使用mkdir不会创建，需要额外调用 bash-c 
	@bash -c "mkdir -p ${DESTDIR}/cxbooks-${VERSION}/{ssl,conf,bin,i18n}"
	@echo "拷贝配置文件"
	@cp -f ${PROJECT_PATH}/cmd/conf.yml ${DESTDIR}/cxbooks-${VERSION}/conf/conf.yml

	@echo "编译 github.com/cxbooks/cxbooks/cmd"
	@env GOARCH=amd64 go build -ldflags ${LDFLAGS} -o ${DESTDIR}/cxbooks-${VERSION}/bin/cxbooks github.com/cxbooks/cxbooks/cmd

	@echo "打包文件 cxbooks-${VERSION}.tar.gz"
	@cd ${DESTDIR}; tar -czf cxbooks-${VERSION}.tar.gz cxbooks-${VERSION}

docker: cxbooks
	@cp -f ${PROJECT_PATH}/.dockerignore ${DESTDIR}/cxbooks-${VERSION}/
	@cp -f ${PROJECT_PATH}/docker-entrypoint.sh ${DESTDIR}/cxbooks-${VERSION}/
	@cp -f ${PROJECT_PATH}/Dockerfile ${DESTDIR}/cxbooks-${VERSION}/
	@chmod +x ${DESTDIR}/cxbooks-${VERSION}/docker-entrypoint.sh
	@cd ${DESTDIR}/cxbooks-${VERSION}/;docker build -t cxbooks:v0.0.1 ./


clean:
	rm -rf ${DESTDIR}
	docker rmi cxbooks:${VERSION}