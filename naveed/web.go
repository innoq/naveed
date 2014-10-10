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
	msg := "Hello World\n\nlorem ipsum\ndolor sit amet\n\n-- \nNaveed"
	go Sendmail("fnd@innoq.com", "fnd@innoq.com", "Hello World", msg)

	// TODO: check auth token
	res.WriteHeader(202)
}
