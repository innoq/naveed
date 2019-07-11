package naveed

import "os/exec"
import "log"
import "fmt"
import "strings"
import "io"
import "naveed/userindex"

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
	addresses := strings.Join(recipients, ",")
	proc := exec.Command(Config.Sendmail, addresses)

	stdin, err := proc.StdinPipe()
	if err != nil {
		log.Printf("ERROR accessing STDIN")
		return
	}
	if sender != "" {
		name, email, err := userindex.ResolveUser(sender, Config.UserIndex)
		if err == nil {
			sender = fmt.Sprintf("%s <%s>", name, email)
		} else {
			log.Printf("WARN failed to resolve user %s", sender)
			sender = ""
		}
	}
	if sender == "" {
		sender = Config.DefaultSender
	}
	io.WriteString(stdin, fmt.Sprintf("From: %s\n", sender))
	io.WriteString(stdin, fmt.Sprintf("Subject: %s\n", subject))
	io.WriteString(stdin, "Content-Type: text/plain; charset=utf-8\n")
	io.WriteString(stdin, body)
	sig := fmt.Sprintf("\n-- \nsent via Naveed - customize preferences:\n%s\n",
		Config.ExternalRoot)
	io.WriteString(stdin, sig)
	stdin.Close()

	_, err = proc.Output()
	if err == nil {
		log.Printf("[%s] %s: %s", app, addresses, subject)
	} else {
		log.Printf("ERROR invoking sendmail")
	}
}

func resolveAddresses(users []string) (addresses []string) {
	for _, handle := range users {
		_, email, err := userindex.ResolveUser(handle, Config.UserIndex)
		if err == nil {
			addresses = append(addresses, email)
		} else {
			log.Printf("WARN failed to resolve user %s", handle)
		}
	}
	return
}
