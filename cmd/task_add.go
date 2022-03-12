package cmd

import (
	"fmt"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
)

var taskAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Create a new task",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			prompt := promptui.Prompt{
				Label: "Task name",
			}

			result, err := prompt.Run()

			if err != nil {
				fmt.Printf("Prompt failed %v\n", err)
				return
			}

			fmt.Printf("Creating %q\n", result)
		}
	},
}

func init() {
	taskCmd.AddCommand(taskAddCmd)
}
