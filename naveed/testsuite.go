package naveed

import "github.com/gorilla/mux"
import "net/http"
import "net/http/httptest"
import "fmt"
import "os"
import "io"

type TestSuite struct {
	Router *mux.Router
	path   string
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

func (suite *TestSuite) Request(method string, uri string, body io.Reader) *httptest.ResponseRecorder {
	if suite.Router == nil {
		panic("router unset")
	}
	req, _ := http.NewRequest(method, uri, body)
	res := httptest.NewRecorder()
	suite.Router.ServeHTTP(res, req)
	return res
}
