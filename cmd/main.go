package main

import (
	l "log"

	"github.com/Arkosh744/chat-client/cmd/root"
	"github.com/Arkosh744/chat-client/internal/config"
	"github.com/Arkosh744/chat-client/internal/log"
)

func main() {
	inits := []func() error{
		config.Init,
		log.InitLogger,
	}

	for _, init := range inits {
		if err := init(); err != nil {
			l.Fatalf("failed to initialize logger: %v", err)
		}
	}

	root.Execute()
}
