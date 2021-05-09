package config

import "github.com/tkanos/gonfig"

type Config struct {
	AppTitle, HttpPort, DbEngine, DbAddress, DbPort, DbUsr, DbPsw, Dbname string
}

func Get() Config {
	c := Config{}
	gonfig.GetConf("config/config.json", &c)
	return c
}
