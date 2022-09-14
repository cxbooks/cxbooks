# stage 1: build vue frondend code
FROM node:16-alpine as nodebuilder

LABEL Author="tjuliuyou <tjuliuyou@gmail.com>"

ARG BUILD_COUNTRY=""

# 中国境内修改源，加速下载
RUN if [ "x${BUILD_COUNTRY}" = "xCN" ]; then \
    echo "using repo mirrors for ${BUILD_COUNTRY}"; \
    npm config set registry https://registry.npm.taobao.org; \
    fi

COPY /web /web/

WORKDIR /web
RUN npm install && npm run build-in-docker



# stage 2: build golang backend
FROM golang:alpine3.16 as gobuilder

ARG BUILD_COUNTRY=""
ARG GIT_VERSION=""
ARG BUILD_DATE=""
ARG BUILD_HASH=""

COPY ./ /cxbooks/
COPY --from=nodebuilder /web/dist /cxbooks/dist

WORKDIR /cxbooks

# 中国境内修改源，加速下载
RUN if [ "x${BUILD_COUNTRY}" = "xCN" ]; then \
    echo "using repo mirrors for ${BUILD_COUNTRY}"; \
    sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories; \
    go env -w GOPROXY=https://goproxy.cn,direct; \
    go env -w GO111MODULE=on;  \
    fi

RUN apk add build-base && \
    chmod +x /cxbooks/docker-entrypoint.sh && \
    mkdir -p /build/ssl && mkdir /build/conf && mkdir /build/bin && mkdir /build/i18n && \
    cp /cxbooks/cmd/conf.yml /build/conf/conf.yml && \
    go build -ldflags "-X github.com/cxbooks/cxbooks/server.buildstamp=${BUILD_DATE} -X github.com/cxbooks/cxbooks/server.githash=${BUILD_HASH} -X github.com/cxbooks/cxbooks/server.VERSION=${GIT_VERSION} -s -w" \
    -o /build/bin/cxbooks github.com/cxbooks/cxbooks/cmd


# stage 3: for production
FROM alpine:3.16

LABEL author="tjuliuyou@gmail.com" \
    description="cxbooks"
# RUN apk add --no-cache tzdata ca-certificates libc6-compat libgcc libstdc++

COPY --from=gobuilder /build/ ./
COPY --from=gobuilder /cxbooks/docker-entrypoint.sh ./

ENV CONFIG_FILE=/data/conf/conf.yml \
    VERBOSE=info \
    LOG_DIR=stdout
    
EXPOSE 80
VOLUME [ "/data" ]

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["cxbooks"]