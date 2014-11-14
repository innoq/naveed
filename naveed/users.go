package naveed

import "encoding/json"
import "io/ioutil"
import "errors"
import "userindex"

// maps user handles to e-mail addresses
// TODO: delegate to userindex (which includes validation)
func ResolveUser(handle string) (name, email string, err error) {
	userData, err := ioutil.ReadFile("users.json") // XXX: hard-coded
	if err != nil {
		return "", "", errors.New("failed to read user data")
	}

	users := map[string]userindex.User{}
	err = json.Unmarshal(userData, &users)
	if err != nil {
		return "", "", errors.New("failed to decode JSON data")
	}

	user := users[handle]
	return user.Name, user.Email, nil
}
