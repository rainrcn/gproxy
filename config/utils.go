package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfig(path string, out any) error {
	file, err := os.ReadFile(path)

	if err != nil {
		return err
	}

	err = yaml.Unmarshal(file, out)

	if err != nil {
		return err
	}

	return nil
}
