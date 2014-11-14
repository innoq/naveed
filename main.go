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

	userindex.StartSync("users.json", 3*time.Hour)

	naveed.Server(host, port, pathPrefix)
}
