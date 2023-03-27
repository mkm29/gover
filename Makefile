# Setup variables
APP := gover
REPO := code.sclzdev.net/ssf/ssf-tools/gover
# source VERSION File
include VERSION
VERSION := "v$(MAJOR).$(MINOR).$(PATCH)-$(ADDOPTS)"
APPVERSION := $(shell git describe --tags --abbrev=0)
GOVERSION := "1.20"
OS := $(shell uname -s | tr '[:upper:]' '[:lower:]')
ARCH := $(shell uname -m)

# Release info
RELEASE_NAME := "v$(VERSION)"
RELEASE_DESCRIPTION := "Support standalone mode (eg. using flags versus environment variables) and use samber/lo package for filtering"
NEXUS_REGISTRY := https://nexus.ssf.sclzdev.net/repository/ssf-tools/$(APP)/$(VERSION)
# need to eliminate the double quotes
ARTIFACT_URL := $(subst $\",,$(NEXUS_REGISTRY))

# HELP
# This will output the help for each task
.PHONY: help build compile run clean

# Tasks
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

bootstrap: ## Initialize Go module
	go mod init $(APP)
	# Create directories if they do not exist
	mkdir -p bin cmd internal pkg/config pkg/utils reports Security
	touch main.go cmd/root.go 
	echo Getting Go packages
	go get golang.org/x/tools/cmd/cover github.com/stretchr/testify \
	github.com/spf13/cobra github.com/spf13/viper


build: ## Build Go binary for current OS and Platform
	echo "Building for current OS and Platform"
	CGO_ENABLED=0 GOOS=${OS} GOARCH=$(ARCH) go build -a -o bin/"$(APP)-$(OS)-$(ARCH)" main.go

compile: ## Compile Go binary for every OS and Platform
	echo "Compiling for every OS and Platform"
	GOOS=darwin GOARCH=amd64 go build -o bin/$(APP)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o bin/$(APP)-darwin-arm64 main.go
	GOOS=linux GOARCH=amd64 go build -o bin/$(APP)-linux-amd64 main.go
	GOOS=linux GOARCH=arm64 go build -o bin/$(APP)-linux-arm64 main.go
	GOOS=windows GOARCH=amd64 go build -o bin/$(APP)-windows-amd64 main.go

test: ## Run tests and generate coverage report
	go test ./... -v -coverpkg=./... -coverprofile reports/cover.out
	go tool cover -html=reports/cover.out -o reports/cover.html

run: compile ## Run Go binary, need to pass in arguments
	./bin/"$(APP)-$(OS)-$(ARCH)" $(ACTION) --project-id ${PROJECT_ID} --path ${VARIABLES_PATH} --environment ${ENVIRONMENT} --keys-only ${KEYS_ONLY} --include-group-variables ${INCLUDE_GROUP}


clean: ## Clean up bin folder
	rm -f bin/*

upload_nexus: ## Upload binaries to Nexus
	for file in `ls bin`; do \
        curl -f -u "${NEXUS_USER}:${NEXUS_PASS}" --upload-file bin/$$file "https://nexus.ssf.sclzdev.net/repository/mcs-cop-raw/gover/v$(VERSION)/$$file" ; \
    done

git_tag: ## Create a new tag
	git tag -a $(VERSION) -m $(RELEASE_DESCRIPTION)
	git push origin $(VERSION)

git_release_tag: ## Create a new tag
	git tag -a v$(VERSION) -m $(RELEASE_DESCRIPTION)
	git push origin v$(VERSION)

release: ## Create a new release
	curl --header 'Content-Type: application/json' --header "PRIVATE-TOKEN: ${GITLAB_TOKEN}" \
     --data '{ "name": $(VERSION), "tag_name": $(VERSION), "ref": "main",\
     "description": $(RELEASE_DESCRIPTION),\
	 "assets": { "links": [\
	 { "name": "gover-darwin-amd64", "url": "${ARTIFACT_URL}/gover-darwin-amd64", "link_type":"other" },\
	 { "name": "gover-darwin-arm64", "url": "${ARTIFACT_URL}/gover-darwin-arm64", "link_type":"other" },\
	 { "name": "gover-linux-amd64", "url": "${ARTIFACT_URL}/gover-linux-amd64", "link_type":"other" },\
	 { "name": "gover-linux-arm64", "url": "${ARTIFACT_URL}/gover-linux-arm64", "link_type":"other" },\
	 { "name": "gover-windows-amd64", "url": "${ARTIFACT_URL}/gover-windows-amd64", "link_type":"other" }\
	 ] }\
	 }' \
     --request POST "https://code.sclzdev.net/api/v4/projects/$(PROJECT_ID)/releases"

appversion: ## Print the version number
	@echo $(VERSION)
