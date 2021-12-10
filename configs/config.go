package configs

import (
	"github.com/spf13/viper"
	"log"
)

type Server struct {
	Addr  string `mapstructure:"addr"`
	Mode  string `mapstructure:"mode"`
	Pprof bool   `mapstructure:"pprof"`
}

type Database struct {
	Link string `mapstructure:"link"`
}

type Jwt struct {
	Secret string `mapstructure:"secret"`
	Ttl    int    `mapstructure:"ttl"`
}

type Config struct {
	Server   `mapstructure:"server"`
	Database `mapstructure:"database"`
	Jwt      `mapstructure:"jwt"`
}

var Conf = &Config{}

func ConfigInit() {
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}
	if err := viper.Unmarshal(Conf); err != nil {
		log.Fatal(err)
	}
}
