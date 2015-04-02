package userindex

import "encoding/json"
import "strings"
import "io/ioutil"
import "errors"

// retrieves name and e-mail address for the given user handle
func ResolveUser(handle, indexFile string) (name, email string, err error) { // XXX: passing `indexFile` should not be necessary
	userData, err := ioutil.ReadFile(indexFile)
	if err != nil {
		return "", "", errors.New("failed to read user data")
	}

	users := map[string]User{}
	err = json.Unmarshal(userData, &users)
	if err != nil {
		return "", "", errors.New("failed to decode JSON data")
	}

	handle = strings.ToLower(handle) // XXX: non-generic
	user := users[handle]
	if user.Email == "" { // XXX: crude/insufficient?
		return "", "", errors.
			New("failed to retrieve e-mail address for user " + handle)
	}

	return user.Name, user.Email, nil
}
