package cmd

import (
	"fmt"
	"gover/internal/config"
	"gover/internal/utils"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	cfg     *config.Config
	debug   bool
	output  string
	version string
	rootCmd = &cobra.Command{
		Use:   "gover",
		Short: "gover is a tool to get project version",
		Long:  `gover gets the project vertsion using a combination of a VERSION file as well as GitLab CI/CD variables`,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cmd.Help(); err != nil {
				log.Fatal(err)
			}
		},
	}
	versionCmd = &cobra.Command{
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
					cfg.Output = "_version.txt"
				}
			} else {
				return fmt.Errorf("config is nil")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			// Working with OutOrStdout/OutOrStderr allows us to unit test our command easier
			out := cmd.OutOrStdout()
			if cfg.Debug {
				log.Println("Debug is enabled")
			}

			// Print the final resolved value from binding cobra flags and viper config
			// cfg.Output is not "" write to file
			if cfg.Output != "" {
				if err := utils.WriteVersion(cfg); err != nil {
					log.Fatal(err)
				}
				return
			}
			cfg.VersionFile = version
			if output != "" {
				fmt.Fprintln(out, utils.GetVersion(cfg))
			} else {
				// print to stdout
				fmt.Println(utils.GetVersion(cfg))
			}
		},
	}
) // Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug output")
	versionCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "Output file")
	rootCmd.AddCommand(versionCmd)
}

func initConfig() {
	v := viper.New()

	// Initialize config
	config.Init()
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Bind the current command's flags to viper
	bindFlags(rootCmd, v)
	bindFlags(versionCmd, v)
	c.Debug = debug
	c.Output = output
	cfg = c
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
