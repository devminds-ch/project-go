package cmd

import (
	"fmt"
	"runtime/debug"

	"github.com/spf13/cobra"
)

var CommitHash = func() string {
	if info, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range info.Settings {
			if setting.Key == "vcs.revision" {
				return setting.Value
			}
		}
	}
	return "unknown"
}()

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of the Go Training Project",
	Long:  "All software has versions. This is the Go Training Project's",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go Training Project Version:", CommitHash)
	},
}
