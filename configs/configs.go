package configs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"gopkg.in/yaml.v3"
)

type Redis struct {
	Uri    string `yaml:"uri"`
	Prefix string `yaml:"prefix"`
}

type Mysql struct {
	Uri  string `yaml:"uri"`
	Host string `yaml:"host"`
	User string `yaml:"user"`
	Pass string `yaml:"pass"`
}

//ServiceConfig - config
type ServiceConfig struct {
	Redis
	Mysql
}

var (
	Env        = os.Getenv("env")
	configPath = map[string]string{
		"dev": "./configs/",
	}
)

var sCfg ServiceConfig

var once sync.Once

//Get - get service config
func Get() ServiceConfig {

	once.Do(func() {

		configFile := "/tmp/config.dev.yaml"
		fmt.Println(configFile)
		configData, err := ioutil.ReadFile(configFile)
		if err != nil {
			log.Printf("Load config data error: %s", err)
		}

		err = yaml.Unmarshal([]byte(configData), &sCfg)
		if err != nil {
			log.Printf("Yaml unmarshal error: %s", err)
		}
	})

	return sCfg
}
