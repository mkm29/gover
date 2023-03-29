# Build the manager binary
ARG BASE_REGISTRY=nexus.ssf.sclzdev.net/ironbank/google
# FROM nexus.ssf.sclzdev.net/dockerhub/golang:1.20.2 as builder
FROM nexus.ssf.sclzdev.net/ironbank/google/golang/golang-1.20:1.20.0 as builder
ARG TARGETOS
ARG TARGETARCH

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

# Build
# the GOARCH has not a default value to allow the binary be built according to the host where the command
# was called. For example, if we call make docker-build in a local env which has the Apple Silicon M1 SO
# the docker BUILDPLATFORM arg will be linux/arm64 when for Apple x86 it will be linux/amd64. Therefore,
# by leaving it empty we can ensure that the container and binary shipped on it will have the same platform.
RUN CGO_ENABLED=0 GOOS=${TARGETOS:-linux} GOARCH=${TARGETARCH:-amd64} go build -a -o gover main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM nexus.ssf.sclzdev.net/ironbank/google/distroless/base:nonroot
WORKDIR /
COPY --from=builder /workspace/gover /usr/local/bin/gover
# Gitlab CI needs a shell, super secure
COPY --from=nexus.ssf.sclzdev.net/dockerhub/busybox:1.36.0-uclibc /bin/sh /bin/sh
USER 65532:65532

# ENTRYPOINT ["/gover"]

# CMD ["version"]
