package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the SAAG version information",
	Run:   VersionLogic,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func VersionLogic(cmd *cobra.Command, args []string) {
	// FIXME: use a global variable
	const version = "0.0.0"

	fmt.Printf("SAAG - Silent As A Ghost\nVersion: %s\n", version)
}
