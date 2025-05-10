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
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gita called")
	},
}

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
		"config file (default is $HOME/.config/groqmit/groqmit.yaml)",
	)
}
