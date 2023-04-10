package main

import (
	"gover/cmd"
)

func main() {
	rootCmd := cmd.NewRootCommand()
	if e := rootCmd.Execute(); e != nil {
		panic(e)
	}
}
