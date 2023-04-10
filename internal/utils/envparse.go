package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"gover/pkg/config"
	"log"
	"strconv"
	"strings"
)

// ParseError is returned whenever the Parse function encounters an error. It
// includes the line number and underlying error.
type ParseError struct {
	Line int
	Err  error
}

type Version struct {
	Major      int
	Minor      int
	Patch      int
	Additional string
}

var cfg *config.Config

func (e *ParseError) Error() string {
	if e.Line > 0 {
		return fmt.Sprintf("error on line %d: %v", e.Line, e.Err)
	}
	return fmt.Sprintf("error reading: %v", e.Err)
}

func parseError(line int, err error) error {
	return &ParseError{
		Line: line,
		Err:  err,
	}
}

func (v *Version) String() string {
	// 1. construct base version string
	base := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	// split MergeRequestTargetBranchName on /
	targetBranch := strings.Split(strings.ToLower(cfg.Variables.MergeRequestTargetBranchName), "/")
	// TODO: handle error when unable to split?
	var tb string
	if len(targetBranch) == 1 {
		tb = targetBranch[0]
	} else if len(targetBranch) == 2 {
		tb = fmt.Sprintf("%s-%s", targetBranch[0], targetBranch[1])
	}
	if cfg.Debug {
		log.Printf("targetBranch: %s\n", targetBranch)
	}
	// 2. if on defalt or release branch return version string
	if tb == cfg.Variables.DefaultBranch || strings.Contains(tb, "release") {
		return fmt.Sprintf("v%s", base)
	} else {
		// 3. if target branch is not default branch, append branch to base
		base = fmt.Sprintf("%s-%s", base, tb)
	}
	// 4. check if we need to append additional options
	if v.Additional != "" {
		base = fmt.Sprintf("%s-%s", base, v.Additional)
	}
	// 5. Add build number
	return fmt.Sprintf("%s+%d", base, cfg.Variables.PipelineIid)
}

func GetVersion(c *config.Config) string {
	// set variables
	cfg = c
	if cfg.Debug {
		log.Println("GetVersion | Setting variables")
	}
	// version is in the format of vX.Y.Z
	// we want to return X.Y.Z (and optionally -ADDOPTS)
	env := make(map[string]string)
	vr := Version{
		Major: 0,
		Minor: 0,
		Patch: 0,
	}

	// read VERSION file
	if cfg.VersionFile == "" {
		log.Println("GetVersion | No version file specified, using default version file")
		cfg.VersionFile = "VERSION"
	}
	r, err := ReadFile(cfg.VersionFile)
	if err != nil {
		log.Printf("File: %s | %s\n", cfg.VersionFile, err)
		return fmt.Sprintf("0.0.0+%d", cfg.Variables.PipelineIid)
	}

	scanner := bufio.NewScanner(bytes.NewReader(r))

	// Track line number
	i := 0

	// Main scan loop
	for scanner.Scan() {
		i++
		k, v, e := parseLine(scanner.Bytes())
		if e != nil {
			log.Println(parseError(i, e))
			return vr.String()
		}

		// Skip blank lines
		if len(k) > 0 {
			env[string(k)] = string(v)
		}
	}
	if e := scanner.Err(); e != nil {
		log.Println(parseError(i, e))
		return vr.String()
	}
	i, err = strconv.Atoi(env["MAJOR"])
	if err != nil {
		vr.Major = 0
	} else {
		vr.Major = i
	}
	i, err = strconv.Atoi(env["MINOR"])
	if err != nil {
		log.Println(err)
		vr.Minor = 0
	} else {
		vr.Minor = i
	}
	i, err = strconv.Atoi(env["PATCH"])
	if err != nil {
		vr.Patch = 0
	} else {
		vr.Patch = i
	}
	if env["ADDOPTS"] != "" && strings.Contains(env["ADDOPTS"], "#") {
		vr.Additional = env["ADDOPTS"]
	}
	return vr.String()
}

func parseLine(line []byte) (keys, values []byte, e error) {
	// if line contains a #, return (ignore line)
	if i := bytes.IndexByte(line, '#'); i >= 0 {
		return nil, nil, nil
	}
	// Find the first equals sign
	i := bytes.IndexByte(line, '=')
	if i < 0 {
		return nil, nil, fmt.Errorf("no equals sign found")
	}

	// Split the line into two parts by = sign
	kv := bytes.Split(line, []byte("="))
	k := bytes.TrimSpace(kv[0])
	v := bytes.TrimSpace(kv[1])

	// Check for empty key
	if len(k) == 0 {
		return nil, nil, fmt.Errorf("empty key")
	}

	// Check for empty value
	if len(v) == 0 {
		return nil, nil, fmt.Errorf("empty value")
	}

	// Check for invalid characters
	if !validKey(k) {
		return nil, nil, fmt.Errorf("invalid characters in key")
	}

	return k, v, nil
}

func validKey(k []byte) bool {
	// key must either be MAJOR, MINOR, PATCH or ADDOPTS
	return bytes.Equal(k,
		[]byte("MAJOR")) ||
		bytes.Equal(k, []byte("MINOR")) ||
		bytes.Equal(k, []byte("PATCH")) ||
		bytes.Equal(k, []byte("ADDOPTS"))
}
