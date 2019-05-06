package cmd

import (
	"github.com/spf13/cobra"
)

// debugCmd represents the debug command
var debugCmd = &cobra.Command{
	Use:   "debug",
	Short: "debug all subcommand flags",
	Run: func(cmd *cobra.Command, args []string) {
		for _, c := range rootCmd.Commands() {
			c.DebugFlags()
		}
	},
}

func init() {
	rootCmd.AddCommand(debugCmd)
}
