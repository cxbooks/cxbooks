FROM debian:bullseye-slim

LABEL author="tjuliuyou@gmail.com" \
    description="cxbooks"
# 注意使用 .dockerignore 忽略其他文件

ADD ./ /

ENV CONFIG_FILE=/data/conf/conf.yml \
    VERBOSE=info \
    LOG_DIR=stdout
    
EXPOSE 80
VOLUME [ "/data" ]

ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["cxbooks"]