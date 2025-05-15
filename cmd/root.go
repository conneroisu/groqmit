// Package cmd is the main entry point for the groqmit command.
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RootCmd represents the base command when called without any subcommand.
var RootCmd = &cobra.Command{
	Use:   "groqmit",
	Short: "groqmit is a tool for generating git commits",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	home, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}
	configDir := filepath.Join(home, ".groqmit")
	viper.SetConfigName("groqmit")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir)

	return RootCmd.Execute()
}

func init() {
	RootCmd.PersistentFlags().StringP(
		"config",
		"c",
		"",
		`config file (default is $HOME/.config/groqmit/groqmit.yaml)
On Unix systems, it returns $XDG_CONFIG_HOME as specified by
https://specifications.freedesktop.org/basedir-spec/basedir-spec-latest.html if
non-empty, else $HOME/.config.
On Darwin, it returns $HOME/Library/Application Support.
On Windows, it returns %AppData%.
		`,
	)
}
