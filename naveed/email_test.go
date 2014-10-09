package naveed

import "testing"
import "fmt"
import "os"

func TestSendmail(t *testing.T) {
	// setup
	pwd, _ := os.Getwd()
	original := fmt.Sprintf(":%s", os.Getenv("PATH"))
	modified := fmt.Sprintf("%s/../test/fixtures/bin:%s", pwd, os.Getenv("PATH"))
	os.Setenv("PATH", modified)

	var res []byte

	res = Sendmail("fnd@innoq.com", "fnd@innoq.com", "Hello World", "lorem ipsum")
	if res == nil {
		t.Errorf("FAIL'd (1)")
	}

	res = Sendmail("INVALID", "...", "...", "...")
	if res != nil {
		t.Errorf("FAIL'd (2)")
	}

	// teardown
	os.Setenv("PATH", original)
}
