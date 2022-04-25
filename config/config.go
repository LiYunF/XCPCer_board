package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	conf = &Conf{}
)

type MQ struct {
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type MySql struct {
	Msg MQ     `yaml:"DB_Message"`
	CF  string `yaml:"cfTable"`
	NK  string `yaml:"ncTable"`
	VJ  string `yaml:"vjTable"`
	LG  string `yaml:"lgTable"`
}

type Conf struct {
	Description string `yaml:"description"`
	Database    MySql  `yaml:"database"`
}

func init() {
	path := "config/config.yaml"
	if f, err := os.Open(path); err != nil {
		log.Errorf("Init config Error : %v", err)
	} else {
		yaml.NewDecoder(f).Decode(conf)
	}
}
func GetDBMsg() MySql {
	return conf.Database
}
