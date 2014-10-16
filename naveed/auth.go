package naveed

import "errors"

var Tokens string // XXX: only required for testing

func CheckAppToken(appToken string) (app string, err error) {
	if appToken == "" { // XXX: optimization; duplicates last line
		return "", errors.New("invalid token")
	}

	appsByToken, err := ReadAppTokens()
	if err != nil {
		return "", err
	}

	for token, app := range appsByToken {
		if appToken == token {
			return app, nil
		}
	}
	return "", errors.New("invalid token")
}

// TODO: cache to avoid file operations?
func ReadAppTokens() (appsByToken map[string]string, err error) {
	tokens := "tokens.cfg"
	if Tokens != "" {
		tokens = Tokens
	}

	appsByToken, err = ReadSettings(tokens, " #")
	if err != nil {
		return nil, errors.New("could not read tokens")
	}
	return appsByToken, nil
}
