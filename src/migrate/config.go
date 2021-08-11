package migrate

import (
	"errors"
	"twitter.go/src/utils/yaml"
)

type URL struct {
	Schema string `yaml:"schema"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Path   string `yaml:"path"`
}

type Configs struct {
	Server URL `yaml:"server"`
}

var ServerConfigs *Configs

func InitConfig() error {
	files := []string{"./config/default.yml"} // TODO different mode
	ServerConfigs = &Configs{}
	for _, file := range files {
		if len(file) > 0 {
			err := yaml.ReadYml(file, ServerConfigs)
			if err != nil {
				return err
			}
		}
	}
	if *ServerConfigs == (Configs{}) {
		return errors.New("Config is empty.\n")
	}
	// TODO change default config to local or prod
	// TODO in prod necessarily use Vault
	return nil
}
