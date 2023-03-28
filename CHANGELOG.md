All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]


## [0.1.0] - 2023-03-28

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