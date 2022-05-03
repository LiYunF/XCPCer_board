package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	Config = &Conf{}
)

type MQ struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type DB struct {
	CF  string `yaml:"cfTable"`
	NK  string `yaml:"nkTable"`
	VJ  string `yaml:"vjTable"`
	LG  string `yaml:"lgTable"`
	Msg MQ     `yaml:"Message"`
}

type Conf struct {
	Description string `yaml:"description"`
	Database    DB     `yaml:"database"`
}

func InitConfig(path string) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		log.Errorf("Init config Error : %v", err)
	} else {
		yaml.NewDecoder(f).Decode(Config)
	}
}
func init() {
	InitConfig("config/config.yml")
}
