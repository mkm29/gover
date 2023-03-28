package cmd

import (
	"gover/internal/utils"
	"gover/pkg/config"
	"log"

	"github.com/spf13/cobra"
)

var cfg *config.Config

func NewRootCmd(c *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gover",
		Short: "gover is a tool to get project version",
		Long:  `gover gets the project vertsion using a combination of a VERSION file as well as GitLab CI/CD variables`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// load variables
			cfg = c
			if cfg.Debug {
				log.Println("Checking variables")
			}
			ok, mv := cfg.CheckVariables()
			if !ok {
				log.Fatalf("Missing variables: %v", mv)
			}
		},
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	cmd.AddCommand(newVersionCmd())
	return cmd
}

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gover",
		Long:  `All software has versions. This is gover's`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Println(utils.GetVersion(cfg))
		},
	}
	return cmd
}
