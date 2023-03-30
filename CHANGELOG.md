All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.3] - 2023-03-28

## Added

- `sonar-project.properties` file.
- Created `WriteVersion` function in `utils` package. If the output flag is specified this will write the version to a file.
- Created `initializeCommand` to call `config.LoadConfig` as well as `bindFlags` to bind any Cobra flags to the Viper instance

## Changed

- Added `debug/d` flag to root command
- Added `output/o` flag to version subcommand
- Copy over the `mkdir` and `grep` binaries from the `busybox` image in Dockerfile (needed to run Gitlab CI). This should be addressed as nothing should need to be on the distroless image except the `gover` binary
- Created `initializeConfig` and `bindFlags` functions to 
- Moved all Viper initialization from `main` to `cmd.initializeCommand`
- Make sure to call the parents' `PersistentPreRunE` function in the version child command to initialize the configuration (Viper)

### Fixed

- Addressed error in parsing `VERSION` file when line is commented out (using a `#`)

## [0.1.2] - 2023-03-28

### fixed

- If `VERSION` file does not exists, return `0.0.0+<PipelineIid>`
- Added Docker instructions for building on Apple Silicon (enable multiplatform builds).

## [0.1.0] - 2023-03-27

### Added

- This CHANGELOG file to hopefully serve as an evolving example of a
  standardized open source project CHANGELOG.
- Dockerfile to build Golang binary and then final distroless layer with only compiled binary. 
- Makefile to automate all SDLC pertaining to project.
- `.gitlab-ci.yml` file to enable GitLab CI (basic SSF functionality).
- README contains basic information about project
- Setup all GitLab types for CI/CD variables.
- Setup basic CLI using Cobra.
- Setup basic configuration management using Viper.
- Added functionality to parse VERSION file.
- Created `Config` struct in config package to hold:
  - Debug (`bool`)
  - Variables (`config.Cariables{}`)
  - requiredVars (`map[string]string`)
- Set logic for constructing version string:
  - First, parse VERSION file to get `MAJOR`, `MINOR`, `PATCH` and (optionally) `ADDOPTS`
  - Split `MergeRequestTargetBranch` string by `/`. Eg. `rc/8.2.0` -> `["rc", "8.2.0"]` and `development` -> `["development"]`. Join by `-`, call this `tb`.
  - Final version string is: `<MAJOR>.<MINOR>.<PATCH>-<tb>(-<ADDOPTS>)+<PipelineIid>`