package naveed

import "github.com/BurntSushi/toml"

type settings struct {
	UserIndex string `toml:"userindex"`
	Sendmail string `toml:"sendmail"`
	Templates string `toml:"templates"` // XXX: only required for testing
}

var Config settings

func ReadConfig(filePath string) {
	_, err := toml.DecodeFile(filePath, &Config)
	if err != nil {
		panic("failed to read configuration")
	}
	// TODO: normalize paths (e.g. stripping trailing slashes)
}
