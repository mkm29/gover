#!/bin/sh

OS=linux # `uname -s | tr A-Z a-z`
ARCH=amd64 # `uname -m`
DOCKERFILE=Dockerfile
VERSION=$(./bin/gover version | sed 's/\+/-build./g')

echo Logging into Nexus
echo $NEXUS_PASS | docker login -u $NEXUS_USER --password-stdin $NEXUS_URL

IMG="$NEXUS_URL/ssf-tools/gover:$VERSION"
echo "Building $IMG"

if [ $ARCH == "amd64" ]; then
    docker build -t $IMG \
        --build-arg=TARGETOS=$OS --build-arg=TARGETARCH=$ARCH \
        -f $DOCKERFILE \
        --label GIT_COMMIT=$(git rev-parse HEAD) \
        --label GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD) \
        --label MAINTAINER="Mitchell Murphy<mitchell.k.murphy.ctr@socom.mil>" \
        --label VERSION=$VERSION \
        --label OS=$OS \
        --label ARCH=$ARCH \
        .
    docker push $NEXUS_URL/ssf-tools/gover:$VERSION
else
    docker buildx build --platform "${OS}/${ARCH}" \
        --build-arg=TARGETOS=$OS --build-arg=TARGETARCH=$ARCH \
        -f $DOCKERFILE --push \
        --label GIT_COMMIT=$(git rev-parse HEAD) \
        --label GIT_BRANCH=$(git rev-parse --abbrev-ref HEAD) \
        --label MAINTAINER="Mitchell Murphy<mitchell.k.murphy.ctr@socom.mil>" \
        --label VERSION=$VERSION \
        --label OS=$OS \
        --label ARCH=$ARCH \
        -t $IMG .
fi