package root

import (
	"os"
	"strings"

	"github.com/Arkosh744/chat-client/internal/app"
	"github.com/Arkosh744/chat-client/internal/log"
	"github.com/Arkosh744/chat-client/internal/model"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chat-client",
	Short: "Клиент лучшего в мире чата",
}

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Вспомогательная команда для действий связанных с созданием",
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Авторизует на сервере",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get username: %s\n", err.Error())
		}

		password, err := cmd.Flags().GetString("password")
		if err != nil {
			log.Fatalf("failed to get password: %s\n", err.Error())
		}

		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)

		refToken, err := handlerService.Login(ctx, &model.AuthInfo{
			Username: username,
			Password: password,
		})
		if err != nil {
			log.Fatalf("failed to login: %s\n", err.Error())
		}

		log.Info("log-in successfully")
		log.Infof("your refresh token: %s", refToken)
	},
}

var createChatCmd = &cobra.Command{
	Use:   "chat",
	Short: "Создает новый чат",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		refreshToken, err := cmd.Flags().GetString("refresh-token")
		if err != nil {
			log.Fatalf("failed to get refresh token: %s\n", err.Error())
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

		chatID, err := handlerService.CreateChat(ctx, usernames, refreshToken, withHistory)
		if err != nil {
			log.Fatalf("failed to create chat: %s\n", err.Error())
		}

		log.Infof("chat created with id: %s\n", chatID)
	},
}

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Подключается к чат-серверу",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		refreshToken, err := cmd.Flags().GetString("refresh-token")
		if err != nil {
			log.Fatalf("failed to get refresh token: %s\n", err.Error())
		}

		chatID, err := cmd.Flags().GetString("chat-id")
		if err != nil {
			log.Fatalf("failed to get chat id: %s\n", err.Error())
		}

		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)

		err = handlerService.ConnectChat(ctx, chatID, refreshToken)
		if err != nil {
			log.Fatalf("failed to connect: %s\n", err.Error())
		}

		log.Info("chat finished")
	},
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Добавить юзера в к существующему чату",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := cmd.Context()

		refreshToken, err := cmd.Flags().GetString("refresh-token")
		if err != nil {
			log.Fatalf("failed to get refresh token: %s\n", err.Error())
		}

		chatID, err := cmd.Flags().GetString("chat-id")
		if err != nil {
			log.Fatalf("failed to get chat id: %s\n", err.Error())
		}

		username, err := cmd.Flags().GetString("username")
		if err != nil {
			log.Fatalf("failed to get username: %s\n", err.Error())
		}

		serviceProvider := app.NewServiceProvider()
		handlerService := serviceProvider.GetHandlerService(ctx)

		if err = handlerService.AddToChat(ctx, chatID, username, refreshToken); err != nil {
			log.Fatalf("failed to add user to chat: %s\n", err.Error())
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	initLogin()

	initCreateChat()

	initAddToChat()

	initConnect()
}

func initCreateChat() {
	rootCmd.AddCommand(createCmd)
	createCmd.AddCommand(createChatCmd)
	createChatCmd.Flags().StringP("usernames", "u", "", "List of usernames for chat")

	err := createChatCmd.MarkFlagRequired("usernames")
	if err != nil {
		log.Fatalf("failed to mark usernames flag required: %s", err.Error())
	}

	createChatCmd.Flags().StringP("refresh-token", "r", "", "provide refresh token to check access")
	err = createChatCmd.MarkFlagRequired("refresh-token")
	if err != nil {
		log.Fatalf("failed to mark refresh-token flag required: %s", err.Error())
	}

	// not required
	createChatCmd.Flags().Bool("history", false, "set true if you want save history of messages")
}

func initAddToChat() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringP("refresh-token", "r", "", "provide refresh token to check access")

	err := addCmd.MarkFlagRequired("refresh-token")
	if err != nil {
		log.Fatalf("failed to mark refresh-token flag required: %s", err.Error())
	}

	addCmd.Flags().StringP("chat-id", "c", "", "Chat id")
	err = addCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag required: %s", err.Error())
	}

	addCmd.Flags().StringP("username", "u", "", "Username for chat")
	err = addCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark usernames flag required: %s", err.Error())
	}
}

func initConnect() {
	rootCmd.AddCommand(connectCmd)
	connectCmd.Flags().StringP("chat-id", "c", "", "Chat id")

	err := connectCmd.MarkFlagRequired("chat-id")
	if err != nil {
		log.Fatalf("failed to mark chat-id flag required: %s", err.Error())
	}

	connectCmd.Flags().StringP("refresh-token", "r", "", "provide refresh token to check access")
	err = connectCmd.MarkFlagRequired("refresh-token")
	if err != nil {
		log.Fatalf("failed to mark refresh-token flag required: %s", err.Error())
	}
}

func initLogin() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "Имя пользователя")
	err := loginCmd.MarkFlagRequired("username")
	if err != nil {
		log.Fatalf("failed to mark username flag as required: %s\n", err.Error())
	}

	loginCmd.Flags().StringP("password", "p", "", "Пароль пользователя")
	err = loginCmd.MarkFlagRequired("password")
	if err != nil {
		log.Fatalf("failed to mark password flag as required: %s\n", err.Error())
	}
}
