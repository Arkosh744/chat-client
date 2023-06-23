package root

import (
	"strings"

	"github.com/Arkosh744/chat-client/internal/app"
	"github.com/Arkosh744/chat-client/internal/log"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Вспомогательная команда для действий связанных с созданием",
}

var createChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Создает новый чат",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		user, err := cmd.Flags().GetString("user")
		if err != nil {
			log.Fatalf("failed to get username: %s\n", err.Error())
		}

		usernamesStr, err := cmd.Flags().GetString("usernames")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		withHistory, err := cmd.Flags().GetBool("history")
		if err != nil {
			log.Fatalf("failed to get usernames: %s\n", err.Error())
		}

		usernames := strings.Split(usernamesStr, ",")
		if len(usernames) == 0 {
			log.Fatalf("usernames must be not empty")
		}

		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)

		chatID, err := handlerService.CreateChat(ctx, user, usernames, withHistory)
		if err != nil {
			log.Fatalf("failed to create chat: %s\n", err.Error())
		}

		log.Infof("chat created with id: %s\n", chatID)
	},
}

func initCreateChat() {
	createChatCmd.Flags().StringP("usernames", "", "", "List of usernames for chat")
	err := createChatCmd.MarkFlagRequired("usernames")
	if err != nil {
		log.Fatalf("failed to mark usernames flag required: %s", err.Error())
	}

	createChatCmd.Flags().StringP("user", "u", "", "provide username to check access")
	err = createChatCmd.MarkFlagRequired("user")
	if err != nil {
		log.Fatalf("failed to mark user flag required: %s", err.Error())
	}

	// not required
	createChatCmd.Flags().Bool("history", false, "set true if you want save history of messages")
}
