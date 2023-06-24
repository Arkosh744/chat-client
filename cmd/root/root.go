package root

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Клиент лучшего в мире чата",
}

func Execute() {
	rootCmd.AddCommand(loginCmd)
	initLogin()

	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createChatCmd)
	initCreateChat()

	rootCmd.AddCommand(connectCmd)
	initConnect()

	rootCmd.AddCommand(addCmd)
	initAddToChat()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
