package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/alireza12prom/SAAG/internal/config"
	"github.com/alireza12prom/SAAG/internal/helpers"
	"github.com/spf13/cobra"
)

var sweepCmd = &cobra.Command{
	Use:   "sweep",
	Short: "Gather user's home data.",
	Run:   VacuumLogic,
}

func init() {
	rootCmd.AddCommand(sweepCmd)
}

func VacuumLogic(cmd *cobra.Command, args []string) {
	err := helpers.RemoveDir(config.StagingDir)
	if err != nil {
		fmt.Println("Error removing staging directory:", err)
		os.Exit(0)
	}

	err = helpers.MakeDir(config.StagingDir)
	if err != nil {
		fmt.Println("Error creating staging directory:", err)
		os.Exit(0)
	}

	info, err := helpers.GetCurrentUserInfo()
	if err != nil {
		fmt.Println("Error getting user info:", err)
		os.Exit(0)
	}

	for _, white_list := range config.WhiteList {
		for _, dir := range white_list {
			src := path.Join(info.Home, dir)

			if !helpers.CheckDirExists(src) {
				fmt.Println("Skipping", src)
				continue
			}

			dest := helpers.GetDestinationPath(config.StagingDir, info.Username, dir)
			helpers.CopyDir(src, dest)
		}
	}
}
