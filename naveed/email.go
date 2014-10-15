package naveed

import "os"
import "os/exec"
import "strings"
import "bufio"
import "io"

var Tokens string // XXX: only required for testing
var Mailx string  // XXX: only required for testing

// returns `nil` if unsuccessful
func Sendmail(recipients []string, subject string, body string) []byte {
	recipients = FilterRecipients(recipients)
	return dispatch(subject, resolveAddresses(recipients), body)
}

// mailx wrapper
func dispatch(subject string, recipients []string, body string) []byte {
	cmd := "mailx"
	if Mailx != "" {
		cmd = Mailx
	}

	addresses := strings.Join(recipients, ", ")
	proc := exec.Command(cmd, "-s", subject, addresses)

	stdin, err := proc.StdinPipe()
	ReportError(err, "accessing STDIN")
	io.WriteString(stdin, body)
	stdin.Close()

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
	defer fh.Close()
	if err != nil {
		panic("could not read tokens") // XXX: too crude?
	}

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

// maps handles to e-mail addresses
// TODO: delegate to separate service (which might include validation)
func resolveAddresses(handles []string) []string {
	var addresses []string
	for _, handle := range handles {
		addresses = append(addresses, handle+"@innoq.com")
	}
	return addresses
}
