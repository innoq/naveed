package naveed

import "os"
import "os/exec"
import "log"
import "fmt"
import "strings"
import "io"

var Sendmail string // XXX: only required for testing

func SendMail(sender string, recipients []string, subject, body,
	token string) []string {
	app, err := CheckAppToken(token)
	if err != nil {
		return nil // TODO: use proper error?
	}

	recipients = FilterRecipients(recipients, app)
	go dispatch(sender, resolveAddresses(recipients), subject, body, app)
	return recipients
}

// sendmail wrapper
func dispatch(sender string, recipients []string, subject, body, app string) {
	cmd := "/usr/sbin/sendmail"
	if Sendmail != "" {
		cmd = Sendmail
	}

	addresses := strings.Join(recipients, ",")
	proc := exec.Command(cmd, addresses)

	stdin, err := proc.StdinPipe()
	ReportError(err, "accessing STDIN")
	if sender != "" {
		sender = fmt.Sprintf("%s <%s>", sender, resolveAddress(sender))
	} else {
		sender = "Naveed <noreply@innoq.com>" // XXX: hard-coded
	}
	io.WriteString(stdin, fmt.Sprintf("From: %s\n", sender))
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

func resolveAddresses(users []string) (addresses []string) {
	for _, user := range users {
		addresses = append(addresses, resolveAddress(user))
	}
	return
}

// maps user handles to e-mail addresses
// TODO: delegate to separate service (which might include validation)
func resolveAddress(user string) (address string) {
	return user + "@innoq.com"
}
