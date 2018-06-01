package Mongo

import (
	"Utils/Settings"
	"gopkg.in/mgo.v2"
	"log"
)

type MongoConf struct {
	MongoServer string `yaml:"server"`
	MongoPort   string `yaml:"port"`
	MongoDB     string `yaml:"database"`
}

func Mongo_getSession() *mgo.Session {
	log.Println("getting mongo session...")

	var conf MongoConf
	Utils.Load_settings(&conf, "mongoSetting.yaml")

	session, err := mgo.Dial(conf.MongoServer + ":" + conf.MongoPort)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)

	return session
}
