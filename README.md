# Gover

```yaml
Author: Mitch Murphy
Date: 28 March 2023
```

![gover gopher](media/gopher-gover.jpg)

_Generated using Midjourney_

## Description

Very simple Golang project that simply parses a `VERSION` file, incorporates any predefined Gitlab CI/CD variables and returns the full version string.  

## Structure

`VERSION`

```bash
MAJOR=<major_version>
MINOR=<minor_version>
PATCH=<patch_version>
# Optional
ADDOPTS=<additional_options>
```

## Usage

All relevant commands are listed and annotated in the [Makefile](Makefile).

This program is intended to run inside a CI/CD pipeline (in a container), but can also be run locally. Note that when running any pipeline in Gitlab, these [variables](https://docs.gitlab.com/ee/ci/variables/predefined_variables.html) are automatically added to any stage(kmage). Here are the commands:

```bash
./bin/gover version
```

Example output/versions: 

```bash
# merge request to development branch
0.1.0-development+55358
0.1.1-development-alpha.1+55359
# merge request to rc/* branch
0.1.2-rc-8.2.0+55360
# merge request to release branch (official release)
v0.1.3
```

_Note_ that with how `MCS-COP` versions software (ie. `8.1.0`, `8.2.1`, etc.) this directly conflicts with [SemVer](https://semver.org). This needs to be addresses.

## Docker

### Build

For non amd64 architectures (eg. Apple Silicon), Use buildkit.

```bash
OS=linux # `uname -s | tr A-Z a-z`
ARCH=amd64 # `uname -m`
DOCKERFILE=Dockerfile
VERSION=0.1.2-development-alpha.3

docker buildx build --platform "${OS}/${ARCH}" \
    --build-arg=TARGETOS=$OS --build-arg=TARGETARCH=$ARCH \
    -f $DOCKERFILE --push \
    -t "nexus.ssf.sclzdev.net/ssf-tools/gover:$VERSION" .
```

## Sonarqube

Running Sonarqube locally (using a Docker container) is easy. Just run the `sonarqube:lts-community` image and make sure that there is a `sonar-project.properties` file at the project root.

### Run Server

```bash

```

### Sonar-scanner

Either install the binary locally, or use the Docker scanning container.