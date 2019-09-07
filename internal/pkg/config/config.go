package config

import (
	"errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)


type Config struct {
	DebugMode	bool		`yaml:"debug"`
	Token		string		`yaml:"token"`
	Timeout     int         `yaml:"timeout"`
}

func (c *Config) UpdateFromBytes(content []byte) error {
	if nil == c {
		return errors.New("not initialized yet")
	}

	err := yaml.Unmarshal(content, c)
	if nil != err {
		return err
	}

	return nil
}

func (c *Config) UpdateFromFile(file *os.File) error {
	content, err := ioutil.ReadAll(file)
	if nil != err {
		return err
	}

	return c.UpdateFromBytes(content)
}


var config = &Config{
	DebugMode: false,
	Timeout: 60,
}


func GetConfig() *Config {
	return config
}
