package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadFile(file string, out interface{}) error {
	content, err := os.ReadFile(fmt.Sprintf("%s.yaml", file))
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(content, out)
	if err != nil {
		return err
	}
	return nil
}
