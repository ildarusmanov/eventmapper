package configs

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"errors"
)

var unkonwnHttpAuthParamError = errors.New("Unknown auth param")

type Config struct {
	ServerHost      string              `yaml:"server_host"`
	ServerTLSCrt    string              `yaml:"server_tls_crt"`
	ServerTLSKey    string              `yaml:"server_tls_key"`
	MqUrl           string              `yaml:"mq_url"`
	DisableHttp     bool                `yaml:"disable_http"`
	HttpAuthType    string              `yaml:"http_auth_type"`
	HttpAuthParams  map[string]string   `yaml:"http_auth_params"`
	DisableGrpc     bool                `yaml:"disable_grpc"`
	DisableHandlers bool                `yaml:"disable_handlers"`
	GrpcAddr        string              `yaml:"grpc_addr"`
	GrpcTls         bool                `yaml:"grpc_tls"`
	GrpcCertFile    string              `yaml:"grpc_cert"`
	GrpcKeyFile     string              `yaml:"grpc_key"`
	MqHandlers      []map[string]string `yaml:"mq_handlers"`
}

/**
 * Create new config object
 * @return *Config
 */
func CreateNewConfig() *Config {
	return &Config{}
}

/**
 * Load data from file
 * @param string
 * @return *Config
 */
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

/**
 * Get auth param by key
 * @param string
 * @return string
 */
func (c *Config) GetHttpAuthParamByKey(key string) (string, error) {
	if param, ok := c.HttpAuthParams[key]; ok {
		return param, nil
	}

	return "", unkonwnHttpAuthParamError
}
