package yaml

import (
	"bytes"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func ReadYml(file string, value interface{}) error {
	yamlFile, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	r := bytes.NewReader(yamlFile)
	dec := yaml.NewDecoder(r)
	lastData := value
	for dec.Decode(value) == nil {
		if value == nil {
			value = lastData
			return nil
		}
	}
	return nil
}
