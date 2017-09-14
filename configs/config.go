package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	ServerHost      string              `yaml:"server_host"`
	MqUrl           string              `yaml:"mq_url"`
	AuthToken       string              `yaml:"auth_token"`
	DisableGrpc     bool                `yaml:"disable_grpc"`
	DisableHandlers bool                `yaml:"disable_handlers"`
	GrpcAddr        string              `yaml:"grpc_addr"`
	GrpcTls         bool                `yaml:"grpc_tls"`
	GrpcCertFile    string              `yaml:"grpc_cert"`
	GrpcKeyFile     string              `yaml:"grpc_key"`
	MqHandlers      []map[string]string `yaml:"mq_handlers"`
}

func CreateNewConfig() *Config {
	return &Config{}
}

func LoadConfigFile(configFilePath string) *Config {
	configData := CreateNewConfig()

	configFileData, err := ioutil.ReadFile(configFilePath)

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal([]byte(configFileData), configData)

	if err != nil {
		panic(err)
	}

	return configData
}
