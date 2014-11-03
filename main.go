package main

import "os"
import "strconv"
import "naveed"

func main() {
	host := os.Getenv("NAVEED_HOST")
	port, _ := strconv.Atoi(os.Getenv("NAVEED_PORT"))
	pathPrefix := os.Getenv("NAVEED_PATH_PREFIX")
	naveed.Server(host, port, pathPrefix)
}
