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
		res.Write([]byte("invalid form data"))
		return
	}

	subject := req.FormValue("subject")
	if subject == "" {
		res.WriteHeader(400)
		res.Write([]byte("missing subject"))
		return
	}
	recipients := req.FormValue("recipients")
	if recipients == "" {
		res.WriteHeader(400)
		res.Write([]byte("missing recipients"))
		return
	}
	body := req.FormValue("body") // TODO: rename?
	if body == "" {
		res.WriteHeader(400)
		res.Write([]byte("missing message body"))
		return
	}

	// TODO: check auth token
	go Sendmail("fnd@innoq.com", recipients, subject, body)

	res.WriteHeader(202)
}
