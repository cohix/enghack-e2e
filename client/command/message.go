package command

import (
	"fmt"
	"log"
	"strings"

	"github.com/cohix/enghack-e2e/client/action"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func setMessageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "message",
		Short: "set the secret message",
		Long:  `set the secret message`,
		Run: func(cmd *cobra.Command, args []string) {
			key, err := action.GetOrCreateKey()
			if err != nil {
				log.Fatal(errors.Wrap(err, "failed to GetOrCreateKey"))
			}

			message := strings.Join(args, " ")

			encMessage, err := key.Encrypt([]byte(message))
			if err != nil {
				log.Fatal(errors.Wrap(err, "failed to Encrypt message"))
			}

			if err := action.SetMessage(encMessage); err != nil {
				log.Fatal(errors.Wrap(err, "failed to SetMessage"))
			}

			fmt.Println("message set!")
		},
	}

	return cmd
}

func getMessageCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "message",
		Short: "get the secret message",
		Long:  `get the secret message`,
		Run: func(cmd *cobra.Command, args []string) {
			key, err := action.GetOrCreateKey()
			if err != nil {
				log.Fatal(errors.Wrap(err, "failed to GetOrCreateKey"))
			}

			encMessage, err := action.GetMessage()
			if err != nil {
				log.Fatal(errors.Wrap(err, "failed to SetMessage"))
			}

			message, err := key.Decrypt(encMessage)
			if err != nil {
				log.Fatal(errors.Wrap(err, "failed to Decrypt message"))
			}

			fmt.Println(string(message))
		},
	}

	return cmd
}
