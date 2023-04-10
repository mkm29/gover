package cmd

import (
	"fmt"
	"gover/internal/utils"
	"gover/pkg/config"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var cfg *config.Config

// Build the cobra command that handles our command line tool.
func NewRootCommand() *cobra.Command {
	// Store the result of binding cobra flags and viper config. In a
	// real application these would be data structures, most likely
	// custom structs per command. This is simplified for the demo app and is
	// not recommended that you use one-off variables. The point is that we
	// aren't retrieving the values directly from viper or flags, we read the values
	// from standard Go data structures.

	debug := false

	rootCmd := &cobra.Command{
		Use:   "gover",
		Short: "gover is a tool to get project version",
		Long:  `gover gets the project vertsion using a combination of a VERSION file as well as GitLab CI/CD variables`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			if cfg == nil {
				c, err := initializeConfig(cmd)
				if err != nil {
					return err
				}
				cfg = c
			}
			ok, mv := cfg.CheckVariables()
			if !ok {
				log.Fatalf("Missing variables: %v", mv)
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			if debug {
				log.Println("Debug is enabled")
			}
			cfg.Debug = debug
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
		},
	}

	// Define cobra flags, the default value has the lowest (least significant) precedence
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Enable debug output")
	rootCmd.AddCommand(NewVersionCommand())
	return rootCmd
}

// Build the cobra command that handles our command line tool.
func NewVersionCommand() *cobra.Command {
	// Store the result of binding cobra flags and viper config. In a
	// real application these would be data structures, most likely
	// custom structs per command. This is simplified for the demo app and is
	// not recommended that you use one-off variables. The point is that we
	// aren't retrieving the values directly from viper or flags, we read the values
	// from standard Go data structures.

	// Define our command
	output := ""
	version := ""
	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Print the version number of gover",
		Long:  `All software has versions. This is gover's`,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			// call parent PersistentPreRunE
			if err := cmd.Parent().PersistentPreRunE(cmd, args); err != nil {
				return err
			}
			// You can bind cobra and viper in a few locations, but PersistencePreRunE on the root command works well
			if cfg != nil {
				if output != "" {
					cfg.Output = output
				} else {
					cfg.Output = "VERSION"
				}
			} else {
				return fmt.Errorf("config is nil")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Working with OutOrStdout/OutOrStderr allows us to unit test our command easier
			out := cmd.OutOrStdout()

			// Print the final resolved value from binding cobra flags and viper config
			// cfg.Output is not "" write to file
			if cfg.Output != "" {
				if err := utils.WriteVersion(cfg); err != nil {
					log.Fatal(err)
				}
				return
			}
			cfg.VersionFile = version
			fmt.Fprintln(out, utils.GetVersion(cfg))
		},
	}

	// Define cobra flags, the default value has the lowest (least significant) precedence
	versionCmd.Flags().StringVarP(&output, "output", "o", "", "File to output version to")
	versionCmd.Flags().StringVarP(&version, "version", "v", "VERSION", "Version file to use")
	return versionCmd
}

func initializeConfig(cmd *cobra.Command) (*config.Config, error) {
	v := viper.New()

	// Initialize config
	config.Init()
	c, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Bind the current command's flags to viper
	bindFlags(cmd, v)

	return c, nil
}

// Bind each cobra flag to its associated viper configuration (config file and environment variable)
func bindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// Determine the naming convention of the flags when represented in the config file
		configName := f.Name

		// Apply the viper config value to the flag when the flag is not set and viper has a value
		if !f.Changed && v.IsSet(configName) {
			val := v.Get(configName)
			if err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val)); err != nil {
				log.Fatalf("unable to set flag '%s' from config: %v", f.Name, err)
			}
		}
	})
}
