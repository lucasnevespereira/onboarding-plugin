package main

import (
	"fmt"
	"log"

	"github.com/mattermost/mattermost-server/v5/model"
)

var (
	client *model.Client4
	//webSocketClient *model.WebSocketClient
	botChannel *model.Channel
)

const (
	mmSiteUrl      = "http://localhost:8065"
	mmTeamName     = "groupe6"
	botChannelName = "Catherine"
)

func (p *Plugin) LaunchBot() {
	fmt.Println(p.bot.Username)

	client = model.NewAPIv4Client(mmSiteUrl)
	team, resp := client.GetTeamByName(mmTeamName, "")
	if resp.Error != nil {
		log.Println(resp.Error)
	}

	createBotChannel(team)

	sendBotMsg("this is a test")

}

func createBotChannel(team *model.Team) {
	// Check for existing channel
	if existingChan, resp := client.GetChannelByName(botChannelName, team.Id, ""); resp.Error != nil {
		log.Println(resp.Error)
	} else {
		botChannel = existingChan
		return
	}

	// Create channel
	c := &model.Channel{}
	c.Name = botChannelName
	c.DisplayName = "Catherine Bot Channel"
	c.Type = model.CHANNEL_OPEN
	c.TeamId = team.Id

	if channel, resp := client.CreateChannel(c); resp.Error != nil {
		log.Println(resp.Error)
	} else {
		botChannel = channel
	}
}

func sendBotMsg(msg string) {
	post := &model.Post{}
	post.ChannelId = botChannel.Id
	post.Message = msg

	_, resp := client.CreatePost(post)
	if resp.Error != nil {
		log.Println(resp.Error)
	}
}
