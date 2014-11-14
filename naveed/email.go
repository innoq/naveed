package naveed

import "os"
import "os/exec"
import "log"
import "fmt"
import "strings"
import "io"

var Sendmail string // XXX: only required for testing

func SendMail(recipients []string, subject string, body string,
	token string) []string {
	app, err := CheckAppToken(token)
	if err != nil {
		return nil // TODO: use proper error?
	}

	recipients = FilterRecipients(recipients, app)
	go dispatch(subject, resolveAddresses(recipients), body, app)
	return recipients
}

// sendmail wrapper
func dispatch(subject string, recipients []string, body string, app string) {
	cmd := "/usr/sbin/sendmail"
	if Sendmail != "" {
		cmd = Sendmail
	}

	addresses := strings.Join(recipients, ",")
	proc := exec.Command(cmd, addresses)

	stdin, err := proc.StdinPipe()
	ReportError(err, "accessing STDIN")
	io.WriteString(stdin, "From: Naveed <noreply@innoq.com>\n") // XXX: hard-coded
	io.WriteString(stdin, fmt.Sprintf("Subject: %s\n", subject))
	io.WriteString(stdin, "Content-Type: text/plain; charset=utf-8\n")
	io.WriteString(stdin, body)
	sig := fmt.Sprintf("\n-- \nsent via Naveed - customize preferences:\n%s\n",
		os.Getenv("NAVEED_ROOT_URL")) // XXX: breaks encapsulation?
	io.WriteString(stdin, sig)
	stdin.Close()

	_, err = proc.Output()
	if err == nil {
		log.Printf("[%s] <%s>: %s", app, strings.Join(recipients, ">, <"),
			subject)
	} else {
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
