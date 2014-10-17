package naveed

import "path"

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

func isMuted(handle string, app string) (muted bool) {
	filePath := path.Join(PreferencesDir, handle)
	settings, err := ReadSettings(filePath, ": ")
	if err != nil {
		return false
	}
	return settings[app] == "muted"
}
