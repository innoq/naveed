package naveed

import "os/exec"
import "io"

// returns `nil` if unsuccessful
func Sendmail(sender string, recipient string, subject string, body string) []byte {
	proc := exec.Command("mailx", "-s", subject, recipient, "--", "-f", sender)

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
