package main

import "time"
import "os"
import "strconv"
import "naveed"
import "userindex"

func main() {
	host := os.Getenv("NAVEED_HOST")
	port, _ := strconv.Atoi(os.Getenv("NAVEED_PORT"))
	pathPrefix := os.Getenv("NAVEED_PATH_PREFIX")
	naveed.ReadConfig("naveed.ini")

	userindex.StartSync(naveed.Config.UserIndex, 3*time.Hour)

	naveed.Server(host, port, pathPrefix)
}
