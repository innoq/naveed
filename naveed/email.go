package naveed

import "os/exec"
import "fmt"
import "io"

func Sendmail(sender string, recipient string, subject string, body string) {
	proc := exec.Command("echo", // XXX: DEBUG
		"mailx", "-s", subject, recipient, "--", "-f", sender)

	stdin, err := proc.StdinPipe()
	ReportError(err)
	io.WriteString(stdin, body)
	stdin.Close() // TODO: `defer`?

	out, err := proc.Output()
	ReportError(err)
	fmt.Printf("%s", out)
}
