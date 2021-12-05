package config

import (
	"log"

	"github.com/wd30gsrc/go-simple-todo-app/utils"
	"gopkg.in/go-ini/ini.v1"
)

type ConfigList struct {
	Port string
	SQLDriver string
	DbName string
	LogFile string
	Static string
}

var Config ConfigList

func init() {
	LoadConfig()
	utils.LoggingSettings(Config.LogFile)
}

func LoadConfig() {
	cfg, err := ini.Load("config/config.ini")
	if err != nil {
		log.Fatal(err)
	}
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustString("8080"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
		DbName: cfg.Section("db").Key("name").String(),
		LogFile: cfg.Section("web").Key("logfile").String(),
		Static: cfg.Section("web").Key("static").String(),
	}
}
