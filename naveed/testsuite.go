package naveed

import "fmt"
import "os"

type TestSuite struct {
	path string
}

func (suite *TestSuite) Setup() {
	pwd, _ := os.Getwd()
	Mailx = fmt.Sprintf("%s/../test/fixtures/bin/mailx", pwd)
}

func (suite *TestSuite) Teardown() {
	// NB: must not reset `Mailx` due to `Sendmail` being invoked as goroutine
}
