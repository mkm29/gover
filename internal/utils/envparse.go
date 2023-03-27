package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"gover/pkg/config"
	"log"
	"strconv"
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
	base := fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
	if v.Additional != "" {
		base = fmt.Sprintf("%s-%s", base, v.Additional)
	}
	return base
}

func GetVersion(variables *config.Variables) string {
	// version is in the format of vX.Y.Z
	// we want to return X.Y.Z (and optionally -ADDOPTS)
	env := make(map[string]string)
	vr := Version{
		Major: 0,
		Minor: 1,
		Patch: 0,
	}

	// read VERSION file
	r, err := ReadFile(fmt.Sprintf("%s/%s", GetProjectRoot("./"), "VERSION"))
	if err != nil {
		log.Fatalln(err)
		return vr.String()
	}

	scanner := bufio.NewScanner(bytes.NewReader(r))

	// Track line number
	i := 0

	// Main scan loop
	for scanner.Scan() {
		i++
		k, v, err := parseLine(scanner.Bytes())
		if err != nil {
			log.Println(parseError(i, err))
			return vr.String()
		}

		// Skip blank lines
		if len(k) > 0 {
			env[string(k)] = string(v)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Println(parseError(i, err))
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
	if env["ADDOPTS"] != "" {
		vr.Additional = env["ADDOPTS"]
	}
	return vr.String()
}

func parseLine(line []byte) ([]byte, []byte, error) {
	// Find the first equals sign
	i := bytes.IndexByte(line, '=')
	if i < 0 {
		return nil, nil, fmt.Errorf("no equals sign found")
	}

	// Split the line into two parts
	// split line by =
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
	return bytes.Equal(k, []byte("MAJOR")) || bytes.Equal(k, []byte("MINOR")) || bytes.Equal(k, []byte("PATCH")) || bytes.Equal(k, []byte("ADDOPTS"))
}
