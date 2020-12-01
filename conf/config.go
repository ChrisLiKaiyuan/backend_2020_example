package conf

import (
	"github.com/BurntSushi/toml"
)

type config struct {
	Port int `toml:"port"`

	MySQL struct {
		User     string `toml:"user"`
		Password string `toml:"password"`
		Addr     string `toml:"addr"`
		Name     string `toml:"name"`
	} `toml:"mysql"`
}

var conf *config

func init() {
	conf = new(config)
	_, err := toml.DecodeFile("./config/example.toml", &conf)
	if err != nil {
		panic(err)
	}
}

func Get() *config {
	return conf
}
