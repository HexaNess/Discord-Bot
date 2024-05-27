package src

import (
	"Bot1/src/controllers/utils"

	"github.com/disgoorg/disgo/discord"
	"github.com/disgoorg/disgo/events"
)

var Listeners = &events.ListenerAdapter{
	OnReady:         onReady,
	OnMessageCreate: onMessageCreate,
}

func onReady(event *events.Ready) {
	channelID := utils.GetSnowflakeIDFromEnv("InfoRebootChannel")
	// channelID := snowflake.ID(1112034150875148422)
	// var channelID = snowflake.ID(os.Getenv("InfoRebootChannel"))
	var message string = "Ready"
	event.Client().Rest().CreateMessage(channelID, discord.NewMessageCreateBuilder().SetContent(message).Build())

}

func onMessageCreate(event *events.MessageCreate) {
	if event.Message.Author.Bot {
		return
	}
	if !utils.IsAdmin(*event.Message.Member) {
		return
	}
	var message string
	if event.Message.Content == "ping" {
		message = "pong"
	}
	if message != "" {
		event.Client().Rest().CreateMessage(event.ChannelID, discord.NewMessageCreateBuilder().SetContent(message).Build())
	}
}
