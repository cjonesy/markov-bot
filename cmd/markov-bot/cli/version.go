package cli

import (
	"fmt"

	"github.com/cjonesy/markov-bot/pkg/markov-bot/version"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version %s", version.Version)
	},
}
