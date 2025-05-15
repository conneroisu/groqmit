package cmd

import (
	"fmt"

	"github.com/conneroisu/groqmit/internal/git"
	"github.com/spf13/cobra"
)

// GenerateCmd represents the generate command.
var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate commit message",
	RunE: func(
		cmd *cobra.Command,
		args []string,
	) error {
		diff, err := git.Diff(cmd.Context())
		if err != nil {
			return err
		}
		fmt.Println(diff)

		return nil
	},
}

func init() {
	RootCmd.AddCommand(GenerateCmd)
}
