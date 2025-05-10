package cmd

import (
	_ "embed"
	"fmt"

	"github.com/conneroisu/groqmit/internal/git"
	"github.com/spf13/cobra"
)

var GenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate commit messages",
	RunE: func(cmd *cobra.Command, args []string) error {
		diff, err := git.GetDiff()
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
