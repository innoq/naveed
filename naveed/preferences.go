package naveed

import "os"
import "path"
import "strings"
import "bufio"

var PreferencesDir string // XXX: only required for testing

// discard recipients that have disabled notifications in their preferences
func FilterRecipients(recipients []string) []string {
	for i, recipient := range recipients {
		if isSuppressed(recipient, "default") {
			recipients = append(recipients[:i], recipients[i+1:]...)
		}
	}
	return recipients
}

func isSuppressed(handle string, app string) (suppressed bool) {
	filePath := path.Join(PreferencesDir, handle)
	settings := readSettings(filePath, ": ")
	if settings == nil {
		return false
	}
	return settings[app] == "suppressed"
}

// reads settings files where each line contains a key/value pair
func readSettings(filePath string, delimiter string) map[string]string {
	settings := map[string]string{}

	fh, err := os.Open(filePath)
	defer fh.Close()
	if err != nil {
		return nil // TODO: use proper error?
	}

	scanner := bufio.NewScanner(fh)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.SplitN(line, delimiter, 2)
		key := items[0]
		value := items[1]
		settings[key] = value
	}

	return settings
}
