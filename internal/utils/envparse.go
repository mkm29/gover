package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"gover/internal/config"
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

func getBranch(b string, isMr bool) (string, error) {
	/*
		- if: "$CI_COMMIT_BRANCH == $CI_DEFAULT_BRANCH"
		- if: '$CI_COMMIT_BRANCH == "dev"'
		- if: '$CI_COMMIT_BRANCH == "develop"'
		- if: "$CI_COMMIT_BRANCH =~ /^RC/"
	*/
	// if branch is one of the above return
	// convert branch to lowercase
	if cfg.Debug {
		log.Printf("Branch: %+v\n", b)
	}
	branches := strings.Split(strings.ToLower(b), "/")
	switch branches[0] {
	case cfg.Variables.DefaultBranch, "dev", "develop":
		return branches[0], nil
	case "rc":
		return fmt.Sprintf("rc-%s", branches[1]), nil
	default:
		if isMr {
			return fmt.Sprintf("mr-%d", cfg.Variables.MergeRequestIid), nil
		}
		if cfg.Variables.CommitBranch == "" {
			return "", fmt.Errorf("commit branch is empty")
		}
		return cfg.Variables.CommitBranch, nil
	}
}

/*
	- if: "$CI_COMMIT_BRANCH =~ /^Release/"
	- if: "$CI_COMMIT_BRANCH =~ /^release/"
*/

func getTargetBranch() string {
	/*
			Need to distinguish between a branch build and a merge request build
			- if: $CI_PIPELINE_SOURCE == 'merge_request_event'
		    - if: $CI_PIPELINE_SOURCE == 'web'
		    - if: $CI_PIPELINE_SOURCE == 'trigger'
		    - if: $CI_PIPELINE_SOURCE == 'pipeline'
	*/
	var tb string
	if cfg.Debug {
		log.Printf("PipelineSource: %+v\n", cfg.Variables.PipelineSource)
	}
	var err error
	switch cfg.Variables.PipelineSource {
	case "merge_request_event":
		if cfg.Debug {
			log.Println("merge request build")
		}
		tb, err = getBranch(cfg.Variables.MergeRequestTargetBranchName, true)
		if err != nil {
			log.Fatalf("error getting target branch: %v", err)
		}
	case "web", "trigger", "pipeline":
		if cfg.Debug {
			log.Println("pipeline build")
		}
		tb, err = getBranch(cfg.Variables.CommitBranch, false)
		if err != nil {
			log.Fatalf("error getting target branch: %v", err)
		}
	case "push":
		if cfg.Debug {
			log.Printf("push build, commit branch: %s\n", cfg.Variables.CommitBranch)
		}
		tb, err = getBranch(cfg.Variables.CommitBranch, false)
		if err != nil {
			log.Fatalf("error getting target branch: %v", err)
		}
	default:
		if cfg.Debug {
			log.Println("branch build")
		}
		tb, err = getBranch(cfg.Variables.CommitBranch, false)
		if err != nil {
			log.Fatalf("error getting target branch: %v", err)
		}
	}
	// replace / with -
	tb = strings.ReplaceAll(tb, "/", "-")
	return tb
}

func isReleaseBranch() bool {
	// cfg.Variables.DefaultBranch || strings.Contains(strings.ToLower(tb), "release"
	return (strings.Contains(strings.ToLower(cfg.Variables.MergeRequestTargetBranchName), "release") ||
		strings.Contains(strings.ToLower(cfg.Variables.CommitBranch), "release"))
}

func (v *Version) String() string {
	// 1. construct base version string
	base := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)

	// 2. if on defalt or release branch return version string
	if isReleaseBranch() {
		return fmt.Sprintf("v%s", base)
	}
	// split MergeRequestTargetBranchName on /
	tb := getTargetBranch()
	if cfg.Debug {
		log.Printf("targetBranch: %s\n", tb)
	}

	// 3. if target branch is not default branch, append branch to base
	base = fmt.Sprintf("%s-%s", base, tb)
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
		if cfg.Debug {
			log.Println("GetVersion | No version file specified, using default version file")
		}
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
	if env["ADDOPTS"] != "" && !strings.Contains(env["ADDOPTS"], "#") {
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
