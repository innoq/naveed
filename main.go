package main

import "time"
import "naveed"
import "userindex"

func main() {
	naveed.ReadConfig("naveed.ini")

	userindex.StartSync(naveed.Config.UserIndex, 3*time.Hour)

	cfg := naveed.Config
	naveed.Server(cfg.Host, cfg.Port, cfg.PathPrefix)
}
