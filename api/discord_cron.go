package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/disgoorg/disgo/webhook"
)

func DiscordCronHandler(w http.ResponseWriter, r *http.Request) {
	client, err := webhook.NewWithURL(os.Getenv("DISCORD_WEBHOOK_URL"))
	if err != nil {
		slog.Error("can't init webhook client", err)
		return
	}

	msg, err := client.CreateContent("hello world")
	if err != nil {
		slog.Error("can't init webhook client", err)
		return
	}
	slog.Info(fmt.Sprintf("sent msg id %s", msg.ID))
}
