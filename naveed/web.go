package naveed

import "github.com/gorilla/mux"
import "html/template"
import "net/http"
import "fmt"
import "path"
import "sort"
import "strings"

const DefaultTemplatesDir = "templates"

var TemplatesDir string // XXX: only required for testing

type provider struct {
	Name  string
	Muted bool
}

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
	router.HandleFunc("/preferences/{user}", PreferencesHandler)
	router.HandleFunc("/outbox", NotificationHandler)
	http.Handle("/", router)
	return router
}

func FrontpageHandler(res http.ResponseWriter, req *http.Request) {
	user := req.Header.Get("REMOTE_USER")
	if user == "" {
		res.WriteHeader(404) // FIXME: this is almost offensively wrong
		return
	}

	http.Redirect(res, req, "/preferences/"+user, http.StatusFound)
}

func PreferencesHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	user := params["user"]

	if req.Method == "POST" {
		updatePreferences(user, res, req)
		return
	}

	appsByToken, err := ReadAppTokens()
	if err != nil {
		respond(res, 500, "unexpected error")
		return
	}

	preferences := ReadPreferences(user)

	var apps []string
	for _, app := range appsByToken {
		apps = append(apps, app)
	}
	sort.Strings(apps)

	var providers []*provider
	for _, app := range apps {
		prov := new(provider)
		prov.Name = app
		prov.Muted = false
		if preferences[app] == "muted" { // XXX: duplicates `isMuted`
			prov.Muted = true
		}
		providers = append(providers, prov)
	}

	render(res, "preferences", providers)
}

func updatePreferences(user string, res http.ResponseWriter, req *http.Request) {
	var err error

	err = req.ParseForm()
	if err != nil {
		respond(res, 400, "invalid form data")
		return
	}

	preferences := map[string]bool{}
	for app, setting := range req.Form {
		preferences[app] = setting[0] == "muted"
	}

	err = WritePreferences(user, preferences)
	if err != nil {
		respond(res, 500, "unexpected error")
	}
	http.Redirect(res, req, "/preferences/"+user, http.StatusFound)
}

func NotificationHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(405)
		return
	}

	var scheme string
	var token string
	auth := req.Header.Get("Authorization")
	items := strings.SplitN(auth, " ", 2)
	if len(items) == 2 {
		scheme = items[0]
		token = items[1]
	}
	if scheme != "Bearer" {
		respond(res, 403, "invalid authorization scheme")
		return
	}

	err := req.ParseForm()
	if err != nil {
		respond(res, 400, "invalid form data")
		return
	}

	subject := req.FormValue("subject")
	if subject == "" {
		respond(res, 400, "missing subject")
		return
	}
	recipients := req.Form["recipient"]
	if recipients == nil {
		respond(res, 400, "missing recipients")
		return
	}
	body := req.FormValue("body") // TODO: rename?
	if body == "" {
		respond(res, 400, "missing message body")
		return
	}

	if Sendmail(recipients, subject, body, token) == nil {
		respond(res, 403, "invalid credentials")
		return
	}
	res.WriteHeader(202)
}

func render(res http.ResponseWriter, view string, data interface{}) {
	if TemplatesDir == "" {
		TemplatesDir = DefaultTemplatesDir
	}
	tmpl, _ := template.ParseFiles(path.Join(TemplatesDir, view+".html"))
	tmpl.Execute(res, data)
}

func respond(res http.ResponseWriter, status int, body string) {
	res.WriteHeader(status)
	res.Write([]byte(body + "\n"))
}
