package naveed

import "github.com/stretchr/testify/assert"
import "net/http"
import "net/http/httptest"
import "testing"

func TestNotification(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()

	var req *http.Request
	var res *httptest.ResponseRecorder
	router := Router()

	req, _ = http.NewRequest("GET", "/outbox", nil)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, 405, res.Code)

	req, _ = http.NewRequest("POST", "/outbox", nil)
	res = httptest.NewRecorder()
	router.ServeHTTP(res, req)

	assert.Equal(t, 202, res.Code)

	suite.Teardown()
}
