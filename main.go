package main

import (
	"gover/cmd"
	"gover/pkg/config"
	"log"
)

func main() {
	// load variables
	// extra vars (not needed when running in GitLab CI/CD)
	args := []string{".", ".env"}
	variables, err := config.Loadvariables(args...)
	if err != nil {
		log.Fatalf("Error loading variables: %s\n", err)
	}
	cmd := cmd.NewRootCmd(variables)
	cmd.Execute()
}
