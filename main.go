package main

import (
	"gover/cmd"
)

func main() {
	cmd := cmd.NewRootCommand()
	cmd.Execute()
}
