package naveed

import "os"
import "os/exec"
import "fmt"
import "strings"
import "io"

var Mailx string // XXX: only required for testing

func Sendmail(recipients []string, subject string, body string,
	token string) []string {
	app, err := CheckAppToken(token)
	if err != nil {
		return nil // TODO: use proper error?
	}

	recipients = FilterRecipients(recipients, app)
	go dispatch(subject, resolveAddresses(recipients), body)
	return recipients
}

// mailx wrapper
func dispatch(subject string, recipients []string, body string) {
	cmd := "mailx"
	if Mailx != "" {
		cmd = Mailx
	}

	addresses := strings.Join(recipients, ", ")
	proc := exec.Command(cmd, "-s", subject, addresses)

	stdin, err := proc.StdinPipe()
	ReportError(err, "accessing STDIN")
	io.WriteString(stdin, body)
	sig := fmt.Sprintf("\n-- \nsent via Naveed - customize preferences:\n%s\n",
		os.Getenv("NAVEED_ROOT_URL")) // XXX: breaks encapsulation?
	io.WriteString(stdin, sig)
	stdin.Close()

	_, err = proc.Output()
	if err != nil {
		ReportError(err, "sending e-mail")
	}
}

// maps user handles to e-mail addresses
// TODO: delegate to separate service (which might include validation)
func resolveAddresses(users []string) (addresses []string) {
	for _, user := range users {
		addresses = append(addresses, user+"@innoq.com")
	}
	return
}
