package configs

import (
    "io/ioutil"
    "gopkg.in/yaml.v2"
)

type Config struct {
	ServerHost string `yaml:"server_host"`
	MqUrl      string `yaml:"mq_url"`
	AuthToken  string `yaml:"auth_token"`
}

func CreateNewConfig() *Config {
	return &Config{}
}

func LoadConfigFile(configFilePath string) *Config {
	configData := CreateNewConfig()

	configFileData, err := ioutil.ReadFile(configFilePath)

	if (err != nil) {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(configFileData), configData)

	if (err != nil) {
		panic(err)
	}

	return configData
}