package naveed

import "fmt"

func ReportError(err error, ctx string) {
	if err != nil {
		fmt.Printf("ERROR %s: %s\n", ctx, err)
	}
}
