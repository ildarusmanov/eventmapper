package tests

import (
	"eventmapper/configs"

	"path/filepath"
)

func CreateConfig() *configs.Config {
	configFilePath, _ := filepath.Abs("/go/src/eventmapper/config_test.yml")
	config := configs.LoadConfigFile(configFilePath)

	return config
}
