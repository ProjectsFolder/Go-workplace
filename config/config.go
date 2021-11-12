package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"sync"
)

type Configuration struct {
	DbDsn    string `yaml:"database_dsn"`
	GrpcHost string `yaml:"grpc_host"`
	GrpcPort string `yaml:"grpc_port"`
}

var (
	configuration *Configuration
	once          sync.Once
)

func GetConfig() *Configuration {
	once.Do(func() {
		file, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			log.Fatal("cannot read configuration:", err)
		}
		configuration = &Configuration{}
		if err = yaml.Unmarshal(file, configuration); err != nil {
			log.Fatal("cannot unmarshal config.yaml", err.Error())
		}
	})

	return configuration
}
