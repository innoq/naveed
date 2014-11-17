package naveed

import "github.com/gorilla/mux"
import "github.com/gorilla/handlers"
import "html/template"
import "net/http"
import "os"
import "log"
import "fmt"
import "path"
import "sort"
import "strings"

type provider struct {
	Name  string
	Muted bool
}

func Server(host string, port int, pathPrefix string) {
	Router(pathPrefix)

	address := fmt.Sprintf("%s:%d", host, port)
	log.Printf("listening at http://%s", address)
	if Config.ExternalRoot != "" {
		log.Printf("... and also %s", Config.ExternalRoot)
	} else {
		log.Printf("WARN external URL not set")
	}
	err := http.ListenAndServe(address, nil)
	if err != nil {
		log.Printf("ERROR starting server")
	}
}

func Router(pathPrefix string) *mux.Router {
	router := mux.NewRouter()

	root := router
	if pathPrefix != "" {
		log.Printf("routing with path prefix %s\n", pathPrefix)
		root = router.PathPrefix(pathPrefix).Subrouter()
	}

	root.HandleFunc("/", FrontpageHandler)
	root.HandleFunc("/preferences/{user}", PreferencesHandler)
	root.HandleFunc("/outbox", NotificationHandler)

	http.Handle("/", handlers.LoggingHandler(os.Stdout, root))
	return root
}

func FrontpageHandler(res http.ResponseWriter, req *http.Request) {
	user := req.Header.Get("REMOTE_USER")
	if user == "" {
		res.WriteHeader(404) // FIXME: this is almost offensively wrong
		return
	}
	redirect("/preferences/"+user, req, res)
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
	redirect("/preferences/"+user, req, res)
}

func NotificationHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		res.WriteHeader(405)
		return
	}

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

	err := req.ParseForm()
	if err != nil {
		respond(res, 400, "invalid form data")
		return
	}

	sender := req.FormValue("sender")
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

	if SendMail(sender, recipients, subject, body, token) == nil {
		respond(res, 403, "invalid credentials")
		return
	}
	res.WriteHeader(202)
}

func render(res http.ResponseWriter, view string, data interface{}) {
	tmpl, _ := template.ParseFiles(path.Join(Config.Templates, view+".html"))
	tmpl.Execute(res, data)
}

func redirect(uri string, req *http.Request, res http.ResponseWriter) {
	uri = os.Getenv("NAVEED_PATH_PREFIX") + uri // XXX: breaks encapsulation -- TODO: use reverse routing
	http.Redirect(res, req, uri, http.StatusFound)
}

func respond(res http.ResponseWriter, status int, body string) {
	res.WriteHeader(status)
	res.Write([]byte(body + "\n"))
}
