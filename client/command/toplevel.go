package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

func checkCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "commands to check things",
		Long:  `check the status of enghack, such as auth`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Long)
		},
	}

	return cmd
}

func getCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "commands to get things",
		Long:  `get things, like the secret message`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Long)
		},
	}

	return cmd
}

func setCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "commands to set things",
		Long:  `set things, like the secret message`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(cmd.Long)
		},
	}

	return cmd
}
