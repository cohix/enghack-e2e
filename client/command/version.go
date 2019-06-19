package command

import (
	"fmt"

	"github.com/spf13/cobra"
)

const enghackVersion = 0.1

func versionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "list the version of enghack",
		Long:  `list the version of enghack`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(enghackVersion)
		},
	}

	return cmd
}
