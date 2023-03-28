package main

import (
	"gover/cmd"
	"gover/pkg/config"
	"log"
)

func main() {
	// load variables
	// extra vars (not needed when running in GitLab CI/CD)
	var args []string
	// args := []string{".", "config.env"}
	cfg, err := config.LoadConfig(args...)
	if err != nil {
		log.Fatalf("Error loading variables: %s\n", err)
	}
	// is cfg nil?
	if cfg == nil {
		log.Fatalf("Config is nil")
	}
	cmd := cmd.NewRootCmd(cfg)
	cmd.Execute()
}
