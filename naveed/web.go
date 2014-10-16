package naveed

import "github.com/gorilla/mux"
import "net/http"
import "fmt"
import "path"
import "sort"

func Server(port int) {
	Router()

	address := fmt.Sprintf(":%d", port)
	fmt.Printf("→ http://localhost%s\n", address)
	err := http.ListenAndServe(address, nil)
	ReportError(err, "starting server")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", FrontpageHandler)
	router.HandleFunc("/preferences/{handle}", PreferencesHandler)
	router.HandleFunc("/outbox", NotificationHandler)
	http.Handle("/", router)
	return router
}

func FrontpageHandler(res http.ResponseWriter, req *http.Request) {
	handle := req.Header.Get("REMOTE_USER")
	if handle == "" {
		res.WriteHeader(404) // FIXME: this is almost offensively wrong
	}

	http.Redirect(res, req, "/preferences/"+handle, http.StatusFound)
}

func PreferencesHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	handle := params["handle"]

	appsByToken, err := ReadAppTokens()
	if err != nil {
		res.WriteHeader(500)
		res.Write([]byte("unexpected error\n"))
		return
	}

	filePath := path.Join(PreferencesDir, handle) // XXX: duplicates `isSuppressed`
	preferences, err := ReadSettings(filePath, ": ")
	if err != nil {
		preferences = map[string]string{}
	}

	var apps []string
	for _, app := range appsByToken {
		apps = append(apps, app)
	}
	sort.Strings(apps)

	for _, app := range apps {
		icon := "✓"
		if preferences[app] == "suppressed" { // XXX: duplicates `isSuppressed`
			icon = "✗"
		}
		res.Write([]byte(icon + " " + app + "\n"))
	}
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
	token := req.FormValue("token") // TODO: use `Authorization: Bearer ...` header
	if Sendmail(recipients, subject, body, token) == nil {
		res.WriteHeader(403)
		res.Write([]byte("token missing or invalid\n"))
		return
	}

	res.WriteHeader(202)
}
