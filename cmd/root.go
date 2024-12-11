package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go_training_project",
	Short: "Go training project by devminds GmbH",
	Long: `Go training project by devminds GmbH
This Go project is used for DevOps CI/CD and Git trainings.
The project contains an application providing a CLI to calculate the sum of two numbers.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
