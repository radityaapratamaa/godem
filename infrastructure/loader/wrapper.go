package loader

import (
	"github.com/kodekoding/phastos/go/apps"
	"godem/domain/models"
	"godem/infrastructure/service/_internal/http/api"

	notifications2 "github.com/kodekoding/phastos/go/notifications"
)

func InitWrapper(cfg *models.Config) *api.WrapperHandler {
	notifService := notifications2.New(&cfg.Notifications)
	slackApp := apps.NewSlack(cfg.Apps.Slack.BotToken)

	return &api.WrapperHandler{Notif: notifService, Apps: slackApp}
}
