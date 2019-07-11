package naveed

import "sync"
import "log"
import nats "github.com/nats-io/nats.go"

// https://nats-io.github.io/docs/developer/receiving/structure.html?lang=go
// https://github.com/nats-io/go-nats-examples/blob/master/api-examples/publish_json/main.go
func NatsSubscriber() {

	// Connect to server
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	ec, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	defer ec.Close()

	// Define the object
	type notification struct {
		Sender string `json:"sender"`
		Recipients []string `json:"recipients"`
		Subject string `json:"subject"`
		Body string `json:"body"`
	}

	wg := sync.WaitGroup{}
	// TODO don't stop after one message
	wg.Add(1)

	// Subscribe
	if _, err := ec.Subscribe("notification.naveed", func(s notification) {
		log.Printf("sender: %s - recipient: %s - message: %s - %s", s.Sender, s.Recipients, s.Subject, s.Body)

		// TODO Authorization
		/* 		
        auth := req.Header.Get("Authorization")
		items := strings.SplitN(auth, " ", 2)
		if len(items) != 2 {
			respond(res, 403, "missing credentials")
			return
		}
		scheme := items[0]
		token := items[1]
		if scheme != "Bearer" {
			respond(res, 403, "invalid authorization scheme")
			return
		}
		*/

		token := "?"

		// TODO check if subject, recipient and body exists

		if SendMail(s.Sender, s.Recipients, s.Subject, s.Body, token) == nil {
			log.Fatal("Mail could not be send")
			return
		}

		wg.Done()

	}); err != nil {
		// TODO
		log.Fatal(err)
	}

	// Wait for a message to come in
	wg.Wait()

}
