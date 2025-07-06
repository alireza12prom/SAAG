package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "saag",
	Short: "SAAG - Silent As A Ghost",
	Long:  "Be like a ghost, silent and unseen.",
	Run:   RootLogic,
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	DisableFlagsInUseLine: true,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func RootLogic(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Help()
	} else {
		fmt.Printf("Unknown command: %s\n", args[0])
		fmt.Println("Use 'saag --help' to see available commands.")
	}
}
