package naveed

import "os"
import "os/exec"
import "strings"
import "bufio"
import "io"

var Tokens string // XXX: only required for testing
var Mailx string  // XXX: only required for testing

// returns `nil` if unsuccessful
// XXX: `sender` currently unused
func Sendmail(sender string, recipient string, subject string, body string) []byte {
	cmd := "mailx"
	if Mailx != "" {
		cmd = Mailx
	}

	proc := exec.Command(cmd, "-s", subject, recipient)

	stdin, err := proc.StdinPipe()
	ReportError(err, "accessing STDIN")
	io.WriteString(stdin, body)
	stdin.Close() // TODO: `defer`?

	out, err := proc.Output()
	if err == nil {
		return out
	} else {
		ReportError(err, "sending e-mail")
		return nil
	}
}

func checkToken(token string) bool { // TODO: cache to avoid file operations?
	if token == "" {
		return false
	}

	tokens := "tokens.cfg"
	if Tokens != "" {
		tokens = Tokens
	}

	fh, err := os.Open(tokens)
	if err != nil {
		panic("could not read tokens") // XXX: too crude?
	}
	defer fh.Close()

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		candidate := strings.SplitN(line, " #", 2)
		if candidate[0] == token {
			return true
		}
	}
	return false
}
