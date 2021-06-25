package main

import (
	"fmt"
	"net/http"
	"sync"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

// Plugin implements the interface expected by the Mattermost server to communicate between the server and plugin processes.
type Plugin struct {
	plugin.MattermostPlugin

	// configurationLock synchronizes access to the configuration.
	configurationLock sync.RWMutex

	// configuration is the active plugin configuration. Consult getConfiguration and
	// setConfiguration for usage.
	configuration *configuration

	botId string

	userId string

	botChannel *model.Channel
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
	mmUserID := r.Header.Get("Mattermost-User-Id")
	p.userId = mmUserID
}

// OnActivate creates bot
func (p *Plugin) OnActivate() error {
	b := &model.Bot{
		Username:    "catherinebot",
		DisplayName: "catherinebot",
		Description: "Catherine increases your daily workflow",
	}

	id, err := p.Helpers.EnsureBot(b)
	if err != nil {
		return err
	}

	p.botId = id

	p.ConfigBot()

	go p.RunCronJob()

	return nil
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
