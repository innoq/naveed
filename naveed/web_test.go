package naveed

import "github.com/stretchr/testify/assert"
import "net/http"
import "net/http/httptest"
import "testing"
import "htptest"
import "os"
import "fmt"

func TestNotification(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()

	req, _ := http.NewRequest("GET", "/outbox", nil)
	res := httptest.NewRecorder()
	Router().ServeHTTP(res, req)

	assert.Equal(t, 202, res.Code)

	suite.Teardown()
}

func TestDSL(t *testing.T) { // XXX: DEBUG
	pwd, _ := os.Getwd()
	fh, _ := os.Open(fmt.Sprintf("%s/../test/outbox.http", pwd))
	htptest.Process(fh)
}
