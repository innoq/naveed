package naveed

import "github.com/stretchr/testify/assert"
import "testing"

func TestSendmail(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()
	defer suite.Teardown()

	var res []byte

	res = Sendmail("fnd@innoq.com", "st@innoq.com, pg@innoq.com",
		"Hello World", "lorem ipsum\ndolor sit amet\n\n...")
	expected := `Subject: Hello World
To: st@innoq.com, pg@innoq.com

lorem ipsum
dolor sit amet

...
`
	assert.Equal(t, expected, string(res))

	suite.CaptureStdout() // suppress "ERROR sending e-mail" to avoid confusion
	res = Sendmail("INVALID", "INVALID", "INVALID", "INVALID")
	assert.Nil(t, res)
	suite.RestoreStdout()
}
