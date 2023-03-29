#!/bin/sh

OS=linux # `uname -s | tr A-Z a-z`
ARCH=amd64 # `uname -m`
DOCKERFILE=Dockerfile
VERSION=0.1.2-development-alpha.3

if [ $ARCH == "amd64" ]; then
    docker build -t nexus.ssf.sclzdev.net/ssf-tools/gover:$VERSION \
        --build-arg=TARGETOS=$OS --build-arg=TARGETARCH=$ARCH \
        -f $DOCKERFILE .
    docker push nexus.ssf.sclzdev.net/ssf-tools/gover:$VERSION
else
    docker buildx build --platform "${OS}/${ARCH}" \
        --build-arg=TARGETOS=$OS --build-arg=TARGETARCH=$ARCH \
        -f $DOCKERFILE --push \
        -t "nexus.ssf.sclzdev.net/ssf-tools/gover:$VERSION" .
fi