package webscreen

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type SeleniumDriver = uint8

const (
	ChromeDriver SeleniumDriver = iota
	GeckoDriver
	HTMLUnitDriver
)

// contains the needed configuration for selenium
type SeleniumConfig struct {
	ServerJarPath string         `yaml:"serverjar"`
	DriverType    SeleniumDriver `yaml:"driver"`
	DriverPath    string         `yaml:"driver_path"`
}

func ConfigFromFile(path string) (*SeleniumConfig, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	conf := &SeleniumConfig{}
	return conf, yaml.Unmarshal(data, conf)
}
