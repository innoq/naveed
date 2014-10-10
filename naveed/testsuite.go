package naveed

import "fmt"
import "os"

type TestSuite struct {
	path string
	stdout *os.File
}

func (suite *TestSuite) Setup() {
	pwd, _ := os.Getwd()
	Mailx = fmt.Sprintf("%s/../test/fixtures/bin/mailx", pwd)
}

func (suite *TestSuite) Teardown() {
	// NB: must not reset `Mailx` due to `Sendmail` being invoked as goroutine
}

func (suite *TestSuite) CaptureStdout() {
	suite.stdout = os.Stdout
	_, dummy, _ := os.Pipe()
	os.Stdout = dummy
}

func (suite *TestSuite) RestoreStdout() {
	os.Stdout.Close()
	os.Stdout = suite.stdout
}
