package naveed

import "github.com/BurntSushi/toml"

type settings struct {
	UserIndex string `toml:"userindex"`
	Sendmail string `toml:"sendmail"`
}

var Config settings

func ReadConfig(filePath string) {
	_, err := toml.DecodeFile(filePath, &Config)
	if err != nil {
		panic("failed to read configuration")
	}
}
