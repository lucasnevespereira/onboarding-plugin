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

	bot model.Bot
}

// ServeHTTP demonstrates a plugin that handles HTTP requests by greeting the world.
func (p *Plugin) ServeHTTP(c *plugin.Context, w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, world!")
}

// OnActivate creates bot
func (p *Plugin) OnActivate() error {
	b := &model.Bot{
		Username:    "catherinebot",
		DisplayName: "catherinebot",
		Description: "Catherine increases your daily workflow",
	}

	pluginBot, err := p.API.CreateBot(b)
	if err != nil {
		return err
	}

	p.bot = *pluginBot

	p.LaunchBot()

	return nil
}

// See https://developers.mattermost.com/extend/plugins/server/reference/
