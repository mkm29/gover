package utils

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
)

func ReadFile(path string) ([]byte, error) {
	// read file
	// log.Printf("Reading file: %s\n", path)
	// check if file exists
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, fmt.Errorf("file does not exist: %s", path)
	}
	b, err := ioutil.ReadFile(path)
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

	return path.Join(dirname, p)
}
