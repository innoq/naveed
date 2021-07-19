package main

import "time"
import "github.com/innoq/naveed/naveed"
import "github.com/innoq/naveed/userindex"

func main() {
	naveed.ReadConfig("naveed.ini")

	userindex.StartSync(naveed.Config.UserIndex, 3*time.Hour) // TODO: configurable interval

	cfg := naveed.Config
	naveed.Server(cfg.Host, cfg.Port, cfg.PathPrefix)
}
