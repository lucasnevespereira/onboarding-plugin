package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

func (p *Plugin) ConfigBot() {
	team, _ := p.API.GetTeamByName("groupe6")

	c, _ := p.API.GetChannelByName(team.Id, "Catherine", true)

	p.botChannel = c
}

func (p *Plugin) SendBotMsg() {
	post := &model.Post{
		UserId:    p.botId,
		ChannelId: p.botChannel.Id,
		Message:   "Welcome!",
	}

	p.API.SendEphemeralPost(p.userId, post)
}
