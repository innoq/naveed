package naveed

import "os"
import "errors"
import "path"
import "bufio"

const DefaultPreferencesDir = "preferences"

var PreferencesDir string // XXX: only required for testing

// discard recipients that have disabled notifications in their preferences
func FilterRecipients(recipients []string, app string) []string {
	for i, recipient := range recipients {
		if isMuted(recipient, app) {
			recipients = append(recipients[:i], recipients[i+1:]...)
		}
	}
	return recipients
}

// XXX: ambiguous contract; it's not obvious that booleans refer to muting
func WritePreferences(user string, preferences map[string]bool) (err error) {
	dir := prefPath()
	err = os.MkdirAll(dir, 0711)
	if err != nil {
		return errors.New("could not store preferences")
	}

	filePath := path.Join(dir, user)
	fh, err := os.Create(filePath)
	defer fh.Close()
	if err != nil {
		return errors.New("could not store preferences")
	}

	buffer := bufio.NewWriter(fh)
	defer buffer.Flush()
	for app, muted := range preferences {
		if muted {
			buffer.Write([]byte(app + ": muted\n"))
		}
	}

	return
}

func ReadPreferences(user string) (preferences map[string]string) {
	filePath := path.Join(prefPath(), user)
	preferences, err := ReadSettings(filePath, ": ")
	if err != nil {
		preferences = map[string]string{}
	}
	return
}

func isMuted(user string, app string) (muted bool) {
	filePath := path.Join(prefPath(), user)
	settings, err := ReadSettings(filePath, ": ")
	if err != nil {
		return false
	}
	return settings[app] == "muted"
}

func prefPath() string {
	if PreferencesDir == "" {
		PreferencesDir = DefaultPreferencesDir
	}
	return PreferencesDir
}
