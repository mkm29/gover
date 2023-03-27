package cmd

import (
	"gover/internal/utils"
	"gover/pkg/config"

	"github.com/spf13/cobra"
)

var variables *config.Variables

func NewRootCmd(v *config.Variables) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "gover",
		Short: "gover is a tool to get project version",
		Long:  `gover gets the project vertsion using a combination of a VERSION file as well as GitLab CI/CD variables`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// load variables
			variables = v
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
			cmd.Println(utils.GetVersion(variables))
		},
	}
	return cmd
}
