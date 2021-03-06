package naveed

import "errors"

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
	appsByToken, err = ReadSettings(Config.Tokens, " #")
	if err != nil {
		return nil, errors.New("could not read tokens")
	}
	return appsByToken, nil
}
