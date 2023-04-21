package main

import (
	"gover/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
