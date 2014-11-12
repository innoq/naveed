package naveed

import "os"
import "os/exec"
import "log"
import "fmt"
import "strconv"
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
	go dispatch(subject, resolveAddresses(recipients), body, app)
	return recipients
}

// mailx wrapper
func dispatch(subject string, recipients []string, body string, app string) {
	cmd := "mailx"
	if Mailx != "" {
		cmd = Mailx
	}

	addresses := strings.Join(recipients, ", ")
	proc := exec.Command(cmd, "-s", subject, addresses)

	// FIXME: fugly workaround to avoid Unicode issues with mailx or sendmail
	body = strconv.QuoteToASCII(body)
	body = strings.Replace(body, "\\n", "\n", -1)

	stdin, err := proc.StdinPipe()
	ReportError(err, "accessing STDIN")
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
