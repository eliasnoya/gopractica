package config

import "github.com/tkanos/gonfig"

type Config struct {
	AppTitle, HttpPort string
}

func Get() Config {
	c := Config{}
	gonfig.GetConf("config/config.json", &c)
	return c
}
