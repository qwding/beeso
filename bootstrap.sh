#!/bin/sh

set -e

NS=reg.yunpro.cn/
VERSION=${VERSION:-latest}
REPO=dingqiwei/beeso
NAME=beeso
INSTANCE=1
PORTS="-p 8080:8080"

cd `dirname $0`
echo "image: ${NS}${REPO}:${VERSION}"
case "$1" in
    build)
	    docker build -t ${NS}${REPO}:${VERSION} .
        ;;
    build)
        go build beeso.go
        ;;
    push)
	    docker push ${NS}${REPO}:${VERSION}
        ;;
    exec)
	    docker exec -it ${NAME} /bin/sh
        ;;
    pull)
	    docker pull ${NS}${REPO}:${VERSION}
        ;;
    shell)
	    docker run --rm --name ${NAME}-${INSTANCE} -i -t ${PORTS} ${VOLUMES} ${ENV} ${NS}${REPO}:${VERSION} /bin/sh
        ;;
    run)
        echo docker run --rm --name ${NAME}-${INSTANCE} ${PORTS} ${VOLUMES} ${ENV} ${NS}${REPO}:${VERSION}
        docker run --rm --name ${NAME}-${INSTANCE} ${PORTS} ${VOLUMES} ${ENV} ${NS}${REPO}:${VERSION}
        ;;
    start)
        docker run -d --name ${NAME}-${INSTANCE} ${PORTS} ${VOLUMES} ${ENV} ${NS}${REPO}:${VERSION}
        ;;
    stop)
        docker stop ${NAME}-${INSTANCE}
        ;;
    rm)
        docker rm ${NAME}-${INSTANCE}
        ;;
    release)
	    docker build -t ${NS}${REPO}:${VERSION} .
	    docker push ${NS}${REPO}:${VERSION}
        ;;
    compose)
	    docker-compose up
        ;;
    remove-none-images)
        docker rmi $(docker images | grep "^<none>" | awk '{print $3}')
        ;;
    *)
        echo "Usage: sh bootstrap.sh {build|push|shell|run|start|stop|rm|release|remove-none-images}"
        exit 3
        ;;
esac
