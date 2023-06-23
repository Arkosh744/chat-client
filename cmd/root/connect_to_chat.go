package root

import (
	"github.com/Arkosh744/chat-client/internal/app"
	"github.com/Arkosh744/chat-client/internal/log"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Подключается к чат-серверу",
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

		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)

		err = handlerService.ConnectChat(ctx, chatID, user)
		if err != nil {
			log.Fatalf("failed to connect: %s\n", err.Error())
		}

		log.Info("chat finished")
	},
}

func initConnect() {

	connectCmd.Flags().StringP("chat-id", "c", "", "Chat id")
	err := connectCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag required: %s", err.Error())
	}

	connectCmd.Flags().StringP("user", "u", "", "provide username to check access")
	err = connectCmd.MarkFlagRequired("user")
	if err != nil {
		log.Fatalf("failed to mark user flag required: %s", err.Error())
	}

}
