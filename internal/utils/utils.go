package utils

import (
	"fmt"
	"gover/pkg/config"
	"os"
	fp "path"
)

const (
	OsReadOnly  = 0400
	OsWrite     = 0200
	OsReadWrite = 0600
	OsAll       = 0777
)

// const (
// 	OS_READ_ONLY  = 0400
// 	OS_WRITE_ONLY = 0200
// 	OS_READ_WRITE = 0600
// )

func ReadFile(path string) ([]byte, error) {
	// read file
	// log.Printf("Reading file: %s\n", path)
	// check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", path)
	}
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func GetProjectRoot(p string) string {
	dirname, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fp.Join(dirname, p)
}

func WriteVersion(c *config.Config) error {
	// write version to file
	return os.WriteFile(c.Output, []byte(GetVersion(c)), os.FileMode(OsAll))
}

func GetEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v == "" {
		return fallback
	}
	return v
}
