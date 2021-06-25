package main

import (
	"github.com/jasonlvhit/gocron"
)

var hourInterval = 1

func (p *Plugin) RunCronJob() {
	gocron.Every(uint64(hourInterval)).Hour().Do(p.SendBotMsg)
	<-gocron.Start()
}
