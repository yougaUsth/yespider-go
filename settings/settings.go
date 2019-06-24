package settings

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

const confFilePath = "conf/app.ini"


type App struct {
}

type Server struct {
	RunMode string
	HttpPort  string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Database struct {

}

type Redis struct {
	Host     string
	PassWord string
}

var AppSettings = &App{}
var ServerSettings = &Server{}
var DatabaseSettings = &Database{}
var RedisSettings = &Redis{}

var conf *ini.File


func Setup() {
	var err error
	conf, err = ini.Load(confFilePath)
	if err != nil {
		log.Fatalf("Settings init error Fatil to load %v : %v", confFilePath, err)
	}

	mapTo("app", AppSettings)
	mapTo("server", ServerSettings)
	mapTo("database", DatabaseSettings)
	mapTo("redis", RedisSettings)


}

func mapTo(section string, v interface{}) {
	err := conf.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Conf Map to Settings obj error %v", section)
	}
}
