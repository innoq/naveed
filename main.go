package main

import (
	"time"
	"naveed/naveed"
	"naveed/userindex"
)

func main() {
	naveed.ReadConfig("naveed.ini")

	userindex.StartSync(naveed.Config.UserIndex, 3*time.Hour) // TODO: configurable interval

	cfg := naveed.Config
	naveed.Server(cfg.Host, cfg.Port, cfg.PathPrefix)
	naveed.NatsSubscriber()
}
