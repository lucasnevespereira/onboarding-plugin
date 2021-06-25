package main

import (
	"github.com/mattermost/mattermost-server/v5/model"
)

var (
//webSocketClient *model.WebSocketClient
//client          *model.Client4
)

func (p *Plugin) LaunchBot() {
	team, _ := p.API.GetTeamByName("groupe6")

	c, _ := p.API.GetChannelByName(team.Id, "Catherine", true)

	post := &model.Post{
		UserId:    p.botId,
		ChannelId: c.Id,
		Message:   "Welcome!",
	}

	_ = p.API.SendEphemeralPost(p.userId, post)

	// TODO: Create a Scheduler

	// websocket?

	//client = model.NewAPIv4Client("http://localhost:8065")

	// webSocketClient, _ := model.NewWebSocketClient4("ws://localhost:8065", client.AuthToken)

	// webSocketClient.Listen()

	// go func() {
	// 	for resp := range webSocketClient.EventChannel {
	//		p.API.PublishWebSocketEvent(resp)
	// 		fmt.Println(resp)
	// 		c := cron.New()
	// 		c.AddFunc("@every 1s", func() {
	// 			_ = p.API.SendEphemeralPost(p.userId, post)
	// 		})
	// 		c.Start()
	// 	}
	// }()

	// To loop
	//select {}

}
