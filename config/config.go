package config

import (
	"os"

	"github.com/spf13/viper"
)

var Conf *Config

type Config struct {
	Server *Server `yaml:"server"`
	N8n    *N8n    `yaml:"n8n"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}
type N8n struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Version string `yaml:"version"`
	Apikey  string `yaml:"apikey"`
	WebHook string `yaml:"webHook"`
}

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
}
