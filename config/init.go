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

func InitConfig(path string) error {

	if f, err := os.Open(path); err != nil {
		log.Errorf("Init config Error : %v", err)
		return err
	} else {
		yaml.NewDecoder(f).Decode(Config)
	}
	return nil
}
func InitAll() error {
	if err := InitConfig("config/config.yaml"); err != nil {
		return err
	}

	return nil
}
