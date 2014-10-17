package naveed

import "github.com/gorilla/mux"
import "net/http"
import "net/http/httptest"
import "path"
import "os"
import "io"

type TestSuite struct {
	Router *mux.Router
	token  string
	path   string
	stdout *os.File
}

func (suite *TestSuite) Setup() {
	pwd, _ := os.Getwd()
	root := path.Join(pwd, "..", "test", "fixtures")
	Tokens = path.Join(root, "tokens.cfg")
	PreferencesDir = path.Join(root, "preferences")
	Mailx = path.Join(root, "bin", "mailx")
	suite.token = "9a790fc4-668b-4d19-aa9f-60c8a00d8621"

}

func (suite *TestSuite) Teardown() {
	// NB: must not reset `Mailx` due to it being invoked as goroutine
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

// TODO: improve signature to make `body` and `contentType` optional and allow
// for automatic form data conversion
func (suite *TestSuite) Request(method string, uri string, body io.Reader,
	headers map[string]string) *httptest.ResponseRecorder {
	if suite.Router == nil {
		panic("router unset")
	}
	req, _ := http.NewRequest(method, uri, body)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	res := httptest.NewRecorder()
	suite.Router.ServeHTTP(res, req)
	return res
}
