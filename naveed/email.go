package naveed

import "os/exec"
import "errors"
import "strings"
import "io"

var Tokens string // XXX: only required for testing
var Mailx string  // XXX: only required for testing

func Sendmail(recipients []string, subject string, body string,
	token string) []string {
	app, err := checkToken(token)
	if err != nil {
		return nil // TODO: use proper error?
	}

	recipients = FilterRecipients(recipients, app)
	go dispatch(subject, resolveAddresses(recipients), body)
	return recipients
}

// mailx wrapper
func dispatch(subject string, recipients []string, body string) (output []byte) {
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

	output, err = proc.Output()
	if err == nil {
		return
	} else {
		ReportError(err, "sending e-mail")
		return nil
	}
}

func checkToken(token string) (app string, err error) { // TODO: cache to avoid file operations?
	if token == "" { // XXX: optimization; duplicates last line
		return "", errors.New("invalid token")
	}

	tokens := "tokens.cfg"
	if Tokens != "" {
		tokens = Tokens
	}

	appsByToken, err := ReadSettings(tokens, " #")
	if err != nil {
		return "", errors.New("could not read tokens")
	}

	for secret, app := range appsByToken {
		if token == secret {
			return app, nil
		}
	}
	return "", errors.New("invalid token")
}

// maps handles to e-mail addresses
// TODO: delegate to separate service (which might include validation)
func resolveAddresses(handles []string) (addresses []string) {
	for _, handle := range handles {
		addresses = append(addresses, handle+"@innoq.com")
	}
	return
}
