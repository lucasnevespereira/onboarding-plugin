package main

import (
	"github.com/jasonlvhit/gocron"
)

var interval = 1

// Sends a bot message every hour
func (p *Plugin) RunCronJob() {
	gocron.Every(uint64(interval)).Hour().Do(p.SendBotMsg)
	<-gocron.Start()
}
