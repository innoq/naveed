package naveed

import "github.com/stretchr/testify/assert"
import "net/http/httptest"
import "testing"
import "net/url"
import "fmt"
import "strings"

const formContentType = "application/x-www-form-urlencoded"

var suite *TestSuite

func TestPreferences(t *testing.T) {
	spawnSuite()
	defer suite.Teardown()

	var res *httptest.ResponseRecorder

	res = suite.Request("GET", "/", nil, nil)
	assert.Equal(t, 404, res.Code)

	res = suite.Request("GET", "/", nil, map[string]string{
		"REMOTE_USER": "johndoe",
	})
	assert.Equal(t, 302, res.Code)
	uri := fmt.Sprintf("%s/preferences/johndoe", Config.PathPrefix)
	assert.Equal(t, uri, res.Header()["Location"][0])

	res = suite.Request("GET", "/preferences/johndoe", nil, nil)
	assert.Equal(t, 200, res.Code)
	// TODO: parse HTML
	assert.Equal(t, true, strings.Contains(compact(res.Body.String()),
		`<legend>dummyapp</legend><label><input type="checkbox" name="dummyapp" value="muted">`))
	assert.Equal(t, true, strings.Contains(compact(res.Body.String()),
		`<legend>randomapp</legend><label><input type="checkbox" name="randomapp" value="muted">`))
	assert.Equal(t, true, strings.Contains(compact(res.Body.String()),
		`<legend>sampleapp</legend><label><input type="checkbox" name="sampleapp" value="muted">`))

	res = suite.Request("GET", "/preferences/bn", nil, nil)
	assert.Equal(t, 200, res.Code)
	// TODO: parse HTML
	assert.Equal(t, true, strings.Contains(compact(res.Body.String()),
		`<legend>dummyapp</legend><label><input type="checkbox" name="dummyapp" value="muted" checked>`))
	assert.Equal(t, true, strings.Contains(compact(res.Body.String()),
		`<legend>randomapp</legend><label><input type="checkbox" name="randomapp" value="muted">`))
	assert.Equal(t, true, strings.Contains(compact(res.Body.String()),
		`<legend>sampleapp</legend><label><input type="checkbox" name="sampleapp" value="muted">`))

	// TODO: test POST handler
}

func TestNotification(t *testing.T) {
	spawnSuite()
	defer suite.Teardown()

	uri := "/outbox"
	var res *httptest.ResponseRecorder

	res = suite.Request("GET", uri, nil, nil)
	assert.Equal(t, 405, res.Code)

	res = suite.Request("POST", uri, nil, nil)
	assert.Equal(t, 403, res.Code)
	assert.Equal(t, "missing credentials\n", res.Body.String())

	res = suite.Request("POST", uri, nil, map[string]string{
		"Authorization": "Bearer " + suite.token,
	})
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "invalid form data\n", res.Body.String())

	res = submitForm(uri, url.Values{
		"subject": {"Hello World"},
	}, suite)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "missing recipients\n", res.Body.String())

	res = submitForm(uri, url.Values{
		"recipient": {"fnd"},
	}, suite)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "missing subject\n", res.Body.String())

	res = submitForm(uri, url.Values{
		"subject":   {"Hello World"},
		"recipient": {"fnd"},
	}, suite)
	assert.Equal(t, 400, res.Code)
	assert.Equal(t, "missing message body\n", res.Body.String())

	res = submitForm(uri, url.Values{
		"subject":   {"Hello World"},
		"recipient": {"fnd", "st"},
		"body":      {"lorem ipsum\ndolor sit amet\n\n..."},
	}, suite)
	assert.Equal(t, 202, res.Code)
}

func submitForm(uri string, data url.Values, suite *TestSuite) *httptest.
	ResponseRecorder {
	body := strings.NewReader(data.Encode())
	res := suite.Request("POST", uri, body, map[string]string{
		"Authorization": "Bearer " + suite.token,
		"Content-Type":  formContentType,
	})
	return res
}

// NB: relies on global to avoid multiple registrations error in HTTP routing
func spawnSuite() {
	if suite == nil {
		suite = new(TestSuite)
		suite.Router = Router("")
	}
	suite.Setup()
}

func compact(str string) string {
	str = strings.Replace(str, "\n", "", -1)
	str = strings.Replace(str, "\t", "", -1)
	return str
}
