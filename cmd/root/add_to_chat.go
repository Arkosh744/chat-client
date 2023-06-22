package root

import (
	"github.com/Arkosh744/chat-client/internal/app"
	"github.com/Arkosh744/chat-client/internal/log"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить юзера в к существующему чату",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			log.Fatalf("failed to get refresh token: %s\n", err.Error())
		}

		chatID, err := cmd.Flags().GetString("chat-id")
		if err != nil {
			log.Fatalf("failed to get chat id: %s\n", err.Error())
		}

		addUser, err := cmd.Flags().GetString("add")
		if err != nil {
			log.Fatalf("failed to get username: %s\n", err.Error())
		}

		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)

		if err = handlerService.AddToChat(ctx, chatID, user, addUser); err != nil {
			log.Fatalf("failed to add user to chat: %s\n", err.Error())
		}
	},
}

func initAddToChat() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("user", "u", "", "provide username to check access")
	err := addCmd.MarkFlagRequired("user")
	if err != nil {
		log.Fatalf("failed to mark user flag required: %s", err.Error())
	}

	addCmd.Flags().StringP("chat-id", "c", "", "Chat id")
	err = addCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag required: %s", err.Error())
	}

	addCmd.Flags().StringP("add", "a", "", "add user to chat")
	err = addCmd.MarkFlagRequired("add")
	if err != nil {
		log.Fatalf("failed to mark usernames flag required: %s", err.Error())
	}
}
