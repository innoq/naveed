package naveed

import "fmt"

func ReportError(err error) {
	if err != nil {
		fmt.Printf("ERROR: %s\n", err)
	}
}
