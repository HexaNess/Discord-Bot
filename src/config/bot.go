package config

import (
	"Bot1/src"
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/disgoorg/disgo"
	"github.com/disgoorg/disgo/bot"
	"github.com/disgoorg/disgo/cache"
	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/gateway"
	"github.com/joho/godotenv"
)

var (
	Test    = "truco"
	Client  bot.Client
	Intents = gateway.IntentGuilds | gateway.IntentGuildMessages | gateway.IntentDirectMessages | gateway.IntentGuildMembers | gateway.IntentMessageContent
)

func Bot() {
	godotenv.Load()
	var err error
	Client, err = disgo.New(os.Getenv("Token"), bot.WithGatewayConfigOpts(gateway.WithIntents(Intents), gateway.WithPresenceOpts(gateway.WithPlayingActivity("Regarde /help", gateway.WithActivityState("lol")), gateway.WithOnlineStatus(discord.OnlineStatusOnline))), bot.WithCacheConfigOpts(cache.WithCaches(cache.FlagRoles, cache.FlagMembers)), bot.WithEventListeners(src.Listeners))

	if err != nil {
		slog.Error("error while building disgo", slog.String("err", err.Error()))
		return
	}

	defer Client.Close(context.TODO())

	if err = Client.OpenGateway(context.TODO()); err != nil {
		slog.Error("errors while connecting to gateway", slog.String("err", err.Error()))
		return
	}

	slog.Info("Bot actif")
	slog.Info("disgo version : " + disgo.Version)

	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-s
}

func GetClient() bot.Client {
	return Client
}
