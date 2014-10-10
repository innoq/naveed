package naveed

import "fmt"
import "os"

type TestSuite struct {
	path string
}

func (suite *TestSuite) Setup() {
	// add mocks to PATH
	pwd, _ := os.Getwd()
	suite.path = fmt.Sprintf("%s", os.Getenv("PATH"))
	modified := fmt.Sprintf("%s/../test/fixtures/bin:%s", pwd, suite.path)
	os.Setenv("PATH", modified)
}

func (suite *TestSuite) Teardown() {
	os.Setenv("PATH", suite.path)
}
