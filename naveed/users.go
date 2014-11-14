package naveed

// maps user handles to e-mail addresses
// TODO: delegate to userindex (which includes validation)
func ResolveUser(handle string) (name, email string, err error) {
	return handle, handle + "@innoq.com", nil
}
