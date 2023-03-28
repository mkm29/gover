# Gover

```yaml
Author: Mitch Murphy
Date: 27 March 2023
```

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
0.1.0-development-alpha.1+55358
# merge request to rc/* branch
0.1.1-rc-8.2.0+55359
# merge request to release branch (official release)
v0.1.2
```

_Note_ that with how `MCS-COP` versions software (ie. `8.1.0`, `8.2.1`, etc.) this directly conflicts with [SemVer](https://semver.org). This needs to be addresses.