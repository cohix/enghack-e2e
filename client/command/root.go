package command

import (
	"fmt"
	"os"

	log "github.com/cohix/simplog"
	"github.com/spf13/cobra"
)

var authToken string

// Execute runs the tool
func Execute() {
	cmd := arrangeCommands()

	if err := cmd.Execute(); err != nil {
		log.LogError(err)
		os.Exit(1)
	}
}

func arrangeCommands() *cobra.Command {
	root := rootCmd()
	root.AddCommand(versionCmd())

	getCmd := getCmd()
	getCmd.AddCommand(getMessageCmd())
	root.AddCommand(getCmd)

	setCmd := setCmd()
	setCmd.AddCommand(setMessageCmd())
	root.AddCommand(setCmd)

	return root
}

func rootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "enghack",
		Short: "enghack is a demo of end-to-end encryption",
		Long:  `An example of building command-line clients in Go, and a primer for end-to-end encryption`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("%s\n\nuse `enghack --help` to view available commands, and `enghack version` to list version\n", cmd.Short)
		},
	}

	cmd.PersistentFlags().StringVar(&authToken, "token", "", "--token overrides the enghack_token env var")

	return cmd
}
