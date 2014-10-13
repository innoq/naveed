package naveed

import "github.com/stretchr/testify/assert"
import "net/http/httptest"
import "testing"

func TestNotification(t *testing.T) {
	suite := new(TestSuite)
	suite.Router = Router()
	suite.Setup()
	defer suite.Teardown()

	var res *httptest.ResponseRecorder

	res = suite.Request("GET", "/outbox", nil)
	assert.Equal(t, 405, res.Code)

	res = suite.Request("POST", "/outbox", nil)
	assert.Equal(t, 202, res.Code)
}
