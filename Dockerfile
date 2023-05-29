ARG BUILDER_IMAGE=golang:buster
ARG DISTROLESS_IMAGE=gcr.io/distroless/base:nonroot
ARG TARGETOS=liunux
ARG TARGETARCH=amd64
ARG BASE_REGISTRY=docker.io
############################
# STEP 1 build executable binary
############################
FROM ${BUILDER_IMAGE} as builder

# Ensure ca-certficates are up to date
RUN update-ca-certificates

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum

# This only needs to be done when building on Gitlab CI
# RUN chmod +w -R /workspace

# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY . . 
# COPY cmd cmd
# COPY pkg pkg
# COPY main.go main.go
# COPY internal internal

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -a -o gover main.go

############################
# STEP 2 build a small image
############################
# using base nonroot image
# user:group is nobody:nobody, uid:gid = 65534:65534
FROM ${DISTROLESS_IMAGE}
# Refer to https://github.com/GoogleContainerTools/distroless for more details
WORKDIR /
COPY --from=builder /workspace/gover /bin/gover
# Gitlab CI needs a shell, mkdir and grep? Not very secure. Work on creating a better/more secure CI/CD systemn using just Golang
COPY --from=busybox:1.36.0-uclibc@sha256:58f16e69b626cfeed566288a6fe6d3950fb5601221bad4297474e7e93f90502b /bin/sh /bin/sh
COPY --from=busybox:1.36.0-uclibc@sha256:58f16e69b626cfeed566288a6fe6d3950fb5601221bad4297474e7e93f90502b /bin/grep /bin/grep
COPY --from=busybox:1.36.0-uclibc@sha256:58f16e69b626cfeed566288a6fe6d3950fb5601221bad4297474e7e93f90502b /bin/mkdir /bin/mkdir
USER 65532:65532

ENTRYPOINT ["/gover"]
