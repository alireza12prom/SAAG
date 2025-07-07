package cmd

import (
	"fmt"

	"github.com/alireza12prom/SAAG/internal/config"
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
	fmt.Printf("SAAG - Silent As A Ghost\nVersion: %s\n", config.AppVersion)
}
