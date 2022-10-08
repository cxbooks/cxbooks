#!/usr/bin/env sh
#
# Created by tjuliuyou on 19/10/28.
#
set -e

# if command starts with an option, prepend cxbooks
if [ "${1:0:1}" = '-' ]; then
    set -- cxbooks "$@"
fi

if [ ! -f "/data/conf/conf.yml" ]; then
    mkdir -p /data/conf/
    mkdir -p /data/db
    cp /conf/conf.yml /data/conf/conf.yml
    cxbooks -conf /data/conf/conf.yml -verbose error -log stdout -init
fi

# cd workspace
# if command app only, add use default args
if [ "$1" = 'cxbooks' ] && [ "$#" -eq 1 ]; then
    exec cxbooks -conf ${CONFIG_FILE} -verbose ${VERBOSE} -log ${LOG_DIR} 
fi

exec "$@"
