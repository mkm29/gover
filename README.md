# Gover


<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="66" height="20" role="img" aria-label="Go: v1.20"><title>Go: v1.20</title><linearGradient id="s" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="r"><rect width="66" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#r)"><rect width="25" height="20" fill="#555"/><rect x="25" width="41" height="20" fill="#007ec6"/><rect width="66" height="20" fill="url(#s)"/></g><g fill="#fff" text-anchor="middle" font-family="Verdana,Geneva,DejaVu Sans,sans-serif" text-rendering="geometricPrecision" font-size="110"><text aria-hidden="true" x="135" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="150">Go</text><text x="135" y="140" transform="scale(.1)" fill="#fff" textLength="150">Go</text><text aria-hidden="true" x="445" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="310">v1.20</text><text x="445" y="140" transform="scale(.1)" fill="#fff" textLength="310">v1.20</text></g></svg> [![pipeline status](https://code.sclzdev.net/ssf/ssf-tools/gover/badges/main/pipeline.svg)](https://code.sclzdev.net/ssf/ssf-tools/gover/-/commits/main) 
[![coverage report](https://code.sclzdev.net/ssf/ssf-tools/gover/badges/main/coverage.svg)](https://code.sclzdev.net/ssf/ssf-tools/gover/-/commits/main) [![Latest Release](https://code.sclzdev.net/ssf/ssf-tools/gover/-/badges/release.svg)](https://code.sclzdev.net/ssf/ssf-tools/gover/-/releases)

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

Take a look at the [build](build.sh) script for more information on building an OCI image for platforms other than amd64.

## Sonarqube

Running Sonarqube locally (using a Docker container) is easy. Just run the `sonarqube:lts-community` image and make sure that there is a `sonar-project.properties` file at the project root.

### Run Server

You can run a simple local server using the Docker container:

```bash
docker run -d -p 9000:9000 sonarqube:lts-community
```

Once running, please visit `http://localhost:9000` and login with the default admin credetials of `admin:admin`. Then, create a project (gover) and generate a token. Add this token to the `sonar-project.properties` file (under `sonar.login`). 

### Sonar-scanner

Either install the binary locally, or use the Docker scanning container.

#### Mac OS X

Simply install with brew: `brew install sonar-scanner`. You can then scan by just calling `sonar-scanner`, and that will ingest the `sonar-project.properties` file, scan and upload the results to the local Sonarqube server.