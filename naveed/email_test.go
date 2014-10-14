package naveed

import "github.com/stretchr/testify/assert"
import "testing"

func TestSendmail(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()
	defer suite.Teardown()

	var msg []byte

	msg = Sendmail([]string{"st", "pg"}, "Hällo Wörld",
		"lörem ipsüm\ndolor sit ämet\n\n✓ … ✗")
	expected := `Subject: Hällo Wörld
To: st@innoq.com, pg@innoq.com

lörem ipsüm
dolor sit ämet

✓ … ✗
`
	assert.Equal(t, expected, string(msg))

	suite.CaptureStdout() // suppress "ERROR sending e-mail" to avoid confusion
	msg = Sendmail([]string{"INVALID"}, "INVALID", "INVALID")
	assert.Nil(t, msg)
	suite.RestoreStdout()
}

func TestUserPreferences(t *testing.T) {
	msg := Sendmail([]string{"st", "bn", "pg"}, "Hello World", "lipsum")
	expected := `Subject: Hello World
To: st@innoq.com, pg@innoq.com

lipsum
`
	assert.Equal(t, expected, string(msg))
}
