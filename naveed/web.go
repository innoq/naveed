package naveed

import "github.com/gorilla/mux"
import "net/http"
import "fmt"

func Server(port int) {
	Router()

	address := fmt.Sprintf(":%d", port)
	fmt.Printf("→ http://localhost%s\n", address)
	err := http.ListenAndServe(address, nil)
	ReportError(err, "starting server")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/outbox", NotificationHandler)
	http.Handle("/", router)
	return router
}

func NotificationHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(405)
		return
	}

	err := req.ParseForm()
	if err != nil {
		res.WriteHeader(400)
		res.Write([]byte("invalid form data\n"))
		return
	}

	token := req.FormValue("token") // TODO: use header instead?
	if !checkToken(token) { // XXX: belongs into `Sendmail`
		res.WriteHeader(403)
		res.Write([]byte("token missing or invalid\n"))
		return
	}
	subject := req.FormValue("subject")
	if subject == "" {
		res.WriteHeader(400)
		res.Write([]byte("missing subject\n"))
		return
	}
	recipients := req.Form["recipient"]
	if recipients == nil {
		res.WriteHeader(400)
		res.Write([]byte("missing recipients\n"))
		return
	}
	body := req.FormValue("body") // TODO: rename?
	if body == "" {
		res.WriteHeader(400)
		res.Write([]byte("missing message body\n"))
		return
	}

	go Sendmail(recipients, subject, body)
	res.WriteHeader(202)
}
