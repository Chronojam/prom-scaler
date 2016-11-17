package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Scalars []RegistrableScalarConfig `yaml:"scalars"`
}

type RegistrableScalarConfig struct {
	Type    string
	Options map[string]interface{} `yaml:",inline"`
}

func Load(path string) (config *Config, err error) {
	var cfg Config

	f, err := os.Open(os.ExpandEnv(path))
	if err != nil {
		return
	}

	defer f.Close()

	d, err := ioutil.ReadAll(f)
	if err != nil {
		return
	}

	err = yaml.Unmarshal(d, &cfg)
	if err != nil {
		return
	}

	return &cfg, nil
}
