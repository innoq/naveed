package naveed

import "github.com/stretchr/testify/assert"
import "net/http"
import "net/http/httptest"
import "testing"

func TestNotification(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()

	req, _ := http.NewRequest("GET", "/outbox", nil)
	res := httptest.NewRecorder()
	Router().ServeHTTP(res, req)

	assert.Equal(t, 202, res.Code)

	suite.Teardown()
}
