$ Gover

Very simple Golang project that simply parses a `VERSION` file and returns the full version string. 

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

This program is intended to run inside a CI/CD pipeline (in a container), but can also be run locally. Here are the commands:

```bash
./bin/gover version
```

Example output: 

```bash
0.1.0-development-alpha.1+55358
0.1.1-rc+55359
v0.1.2
```