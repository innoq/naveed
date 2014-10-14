package naveed

// discard recipients that have disabled notifications in their preferences
func FilterRecipients(recipients []string) []string {
	preferences := map[string]bool{ // XXX: hard-coded -- TODO: per-application settings
		"bn": true, // XXX: confusing; seemingly inverted
	}

	for i, recipient := range recipients {
		if preferences[recipient] {
			recipients = append(recipients[:i], recipients[i+1:]...)
		}
	}
	return recipients
}
