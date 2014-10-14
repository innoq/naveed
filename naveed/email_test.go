package naveed

import "github.com/stretchr/testify/assert"
import "testing"

func TestSendmail(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()
	defer suite.Teardown()

	var res []byte

	res = Sendmail([]string{"st", "pg"}, "Hällo Wörld",
		"lörem ipsüm\ndolor sit ämet\n\n✓ … ✗")
	expected := `Subject: Hällo Wörld
To: st@innoq.com, pg@innoq.com

lörem ipsüm
dolor sit ämet

✓ … ✗
`
	assert.Equal(t, expected, string(res))

	suite.CaptureStdout() // suppress "ERROR sending e-mail" to avoid confusion
	res = Sendmail([]string{"INVALID"}, "INVALID", "INVALID")
	assert.Nil(t, res)
	suite.RestoreStdout()
}
