package root

import (
	"github.com/Arkosh744/chat-client/internal/app"
	"github.com/Arkosh744/chat-client/internal/log"
	"github.com/Arkosh744/chat-client/internal/model"
	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Авторизует на сервере",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		username, err := cmd.Flags().GetString("user")
		if err != nil {
			log.Fatalf("failed to get username: %s\n", err.Error())
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get password: %s\n", err.Error())
		}

		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)

		err = handlerService.Login(ctx, &model.AuthInfo{
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Fatalf("failed to login: %s\n", err.Error())
		}

		log.Info("log-in successfully")
	},
}

func initLogin() {
	loginCmd.Flags().StringP("user", "u", "", "Имя пользователя")
	err := loginCmd.MarkFlagRequired("user")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}

	loginCmd.Flags().StringP("password", "p", "", "Пароль пользователя")
	err = loginCmd.MarkFlagRequired("password")
	if err != nil {
		log.Fatalf("failed to mark password flag as required: %s\n", err.Error())
	}
}
