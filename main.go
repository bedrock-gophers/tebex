package main

import (
	"github.com/bedrock-gophers/tebex/tebex"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"log/slog"
	"os"
	"time"
)

func main() {
	chat.Global.Subscribe(chat.StdoutSubscriber{})

	logger := slog.Default()
	conf, err := server.DefaultConfig().Config(logger)
	if err != nil {
		panic(err)
	}

	srv := conf.New()
	srv.CloseOnProgramEnd()

	store := loadStore(os.Getenv("TEBEX_KEY"), logger)
	srv.Listen()
	for p := range srv.Accept() {
		store.ExecuteCommands(p)
	}
}

// loadStore initializes the Tebex store connection.
func loadStore(key string, log *slog.Logger) *tebex.Client {
	store := tebex.NewClient(log, time.Second*5, key)
	name, domain, err := store.Information()
	if err != nil {
		log.Error("tebex: %v", err)
		os.Exit(1)
		return nil
	}
	log.Info("Connected to Tebex under %v (%v).", name, domain)
	return store
}
