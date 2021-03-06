package common

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type CfgRd struct {
}

type CfgMg struct {
}

type CfgPg struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
	DbName   string `yaml:"DbName"`
	Schema   string `yaml:"Schema"`
}

type CfgDb struct {
	Postgre *CfgPg `yaml:"Postgre"`
	Redis   *CfgRd `yaml:"Redis"`
	Mongo   *CfgMg `yaml:"Mongo"`
}

type Service struct {
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	SSL        bool   `yaml:"SSL"`
	PrivateKey string `yaml:"PrivateKey"`
	PublicKey  string `yaml:"PublicKey"`
}

type Option struct {
	SmsUrl        string `yaml:"SmsUrl"`
	SmsApiToken   string `yaml:"SmsApiToken"`
	FireBaseUrl   string `yaml:"FireBaseUrl"`
	FireBaseToken string `yaml:"FireBaseToken"`
}

type Config struct {
	Database *CfgDb   `yaml:"Database"`
	Service  *Service `yaml:"Service"`
	Option   *Option  `yaml:"Option"`
}

func LoadConfig(cfgPath string) (*Config, error) {
	// Read config file as yaml
	yamlFile, err := ioutil.ReadFile(cfgPath)

	var _conf Config
	if err != nil {
		log.Printf("Cannot read the configuration file: #%v ", err)
		os.Exit(102)
	}
	err = yaml.Unmarshal(yamlFile, &_conf)
	if err != nil {
		log.Printf("Cannot pasre the configuration file: #%v ", err)
		os.Exit(103)
	}
	return &_conf, nil
}
