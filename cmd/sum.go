package cmd

import (
	"fmt"
	"strconv"

	"github.com/devminds-ch/project-go/calculate"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(sumCmd)
}

var sumCmd = &cobra.Command{
	Use:   "sum <number 1> <number 2>",
	Short: "Show the sum of two numbers on the console",
	Long:  "This command will calculate the sum of two numbers and show the result on the console",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.ExactArgs(2)(cmd, args); err != nil {
			return err
		}
		if _, err := strconv.ParseFloat(args[0], 64); err != nil {
			return fmt.Errorf("invalid number specified: %s", args[0])
		}
		if _, err := strconv.ParseFloat(args[1], 64); err != nil {
			return fmt.Errorf("invalid number specified: %s", args[1])
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		a, _ := strconv.ParseFloat(args[0], 64)
		b, _ := strconv.ParseFloat(args[1], 64)
		fmt.Printf("The sum of %f and %f is %f\n", a, b, calculate.Sum(a, b))
	},
}
