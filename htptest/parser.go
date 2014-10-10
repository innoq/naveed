package htptest

import "fmt"

func parse(tokens chan string) {
	for {
		token := <-tokens
		fmt.Printf("RCV %s\n", token) // XXX: DEBUG
	}
}
