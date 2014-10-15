package naveed

import "github.com/stretchr/testify/assert"
import "testing"

func TestSendmail(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()
	defer suite.Teardown()

	recipients := Sendmail([]string{"st", "pg"}, "Hällo Wörld",
		"lörem ipsüm\ndolor sit ämet\n\n✓ … ✗", suite.token)
	assert.Equal(t, []string{"st", "pg"}, recipients)
}

func TestUserPreferences(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()
	defer suite.Teardown()

	recipients := Sendmail([]string{"st", "bn", "pg"}, "Hello World", "lipsum",
		suite.token)
	assert.Equal(t, []string{"st", "pg"}, recipients)
}
