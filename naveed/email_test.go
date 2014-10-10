package naveed

import "testing"

func TestSendmail(t *testing.T) {
	suite := new(TestSuite)
	suite.Setup()

	var res []byte

	res = Sendmail("fnd@innoq.com", "fnd@innoq.com", "Hello World", "lorem ipsum")
	if res == nil {
		t.Errorf("FAIL'd (1)")
	}

	res = Sendmail("INVALID", "INVALID", "INVALID", "INVALID")
	if res != nil {
		t.Errorf("FAIL'd (2)")
	}

	suite.Teardown()
}
