package userindex

import "encoding/json"
import "io/ioutil"
import "errors"

// retrieves name and e-mail address for the given user handle
func ResolveUser(handle string) (name, email string, err error) {
	userData, err := ioutil.ReadFile("users.json") // XXX: hard-coded
	if err != nil {
		return "", "", errors.New("failed to read user data")
	}

	users := map[string]User{}
	err = json.Unmarshal(userData, &users)
	if err != nil {
		return "", "", errors.New("failed to decode JSON data")
	}

	user := users[handle]
	if user.Email == "" { // XXX: crude/insufficient?
		return "", "", errors.
			New("failed to retrieve e-mail address for user " + handle)
	}

	return user.Name, user.Email, nil
}
