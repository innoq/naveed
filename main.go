package main

import "os"
import "strconv"
import "naveed"

func main() {
	port, _ := strconv.Atoi(os.Getenv("NAVEED_PORT"))
	naveed.Server(port)
}
