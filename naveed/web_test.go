package naveed

import "github.com/stretchr/testify/assert"
import "net/http/httptest"
import "testing"
import "net/url"
import "strings"

const formContentType = "application/x-www-form-urlencoded"

func TestNotification(t *testing.T) {
	suite := new(TestSuite)
	suite.Router = Router()
	suite.Setup()
	defer suite.Teardown()

	uri := "/outbox"
	var res *httptest.ResponseRecorder

	res = suite.Request("GET", uri, nil, "")
	assert.Equal(t, 405, res.Code)

	res = suite.Request("POST", uri, nil, "")
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "invalid form data\n", res.Body.String())

	res = submitForm(uri, url.Values{
		"token":   {suite.token},
		"subject": {"Hello World"},
	}, suite)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "missing recipients\n", res.Body.String())

	res = submitForm(uri, url.Values{
		"token":     {suite.token},
		"recipient": {"fnd"},
	}, suite)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "missing subject\n", res.Body.String())

	res = submitForm(uri, url.Values{
		"token":     {suite.token},
		"subject":   {"Hello World"},
		"recipient": {"fnd"},
	}, suite)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "missing message body\n", res.Body.String())

	res = submitForm(uri, url.Values{
		"token":     {suite.token},
		"subject":   {"Hello World"},
		"recipient": {"fnd", "st"},
		"body":      {"lorem ipsum\ndolor sit amet\n\n..."},
	}, suite)
	assert.Equal(t, 202, res.Code)
}

func submitForm(uri string, data url.Values, suite *TestSuite) *httptest.
	ResponseRecorder {
	body := strings.NewReader(data.Encode())
	res := suite.Request("POST", uri, body, formContentType)
	return res
}
