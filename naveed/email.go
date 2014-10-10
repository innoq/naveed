package naveed

import "os/exec"
import "io"

var Mailx string // XXX: only required for testing

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
