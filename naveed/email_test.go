package naveed

import "github.com/stretchr/testify/assert"
import "testing"

func TestSendmail(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()

	var res []byte

	res = Sendmail("fnd@innoq.com", "fnd@innoq.com", "Hello World", "lorem ipsum")
	assert.NotNil(t, res)

	res = Sendmail("INVALID", "INVALID", "INVALID", "INVALID")
	assert.Nil(t, res)

	suite.Teardown()
}
