package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/disgoorg/disgo/webhook"
)

func DiscordCronHandler(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Authorization") != fmt.Sprintf("Bearer %s", os.Getenv("CRON_SECRET")) {
		slog.Error("unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
	}

	client, err := webhook.NewWithURL(os.Getenv("DISCORD_WEBHOOK_URL"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("can't init webhook client: %v", err)))
		return
	}

	msg, err := client.CreateContent("hello world")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(fmt.Sprintf("can't create content: %v", err)))
		return
	}
	slog.Debug(fmt.Sprintf("sent msg id %s", msg.ID))
	w.WriteHeader(http.StatusOK)
}
