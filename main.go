package main

import (
	"gover/cmd"
	"gover/pkg/config"
	"log"
)

func main() {
	// load variables
	variables, err := config.Loadvariables()
	if err != nil {
		log.Fatalf("Error loading variables: %s\n", err)
	}
	cmd := cmd.NewRootCmd(variables)
	cmd.Execute()
}
