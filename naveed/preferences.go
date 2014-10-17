package naveed

import "os"
import "errors"
import "path"
import "bufio"

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
func WritePreferences(handle string, preferences map[string]bool) (err error) {
	filePath := path.Join(PreferencesDir, handle)
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

func isMuted(handle string, app string) (muted bool) {
	filePath := path.Join(PreferencesDir, handle)
	settings, err := ReadSettings(filePath, ": ")
	if err != nil {
		return false
	}
	return settings[app] == "muted"
}
