package Utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

func Load_settings(conf interface{}, filename string) interface{} {
	log.SetFlags(log.Lshortfile)
	log.Printf("load setting...")

	path, _ := os.Getwd()
	setting, err := ioutil.ReadFile(path + "/" + filename)
	//	log.Print("\n" + string(setting[:]))
	if err != nil {
		log.Printf("read setting error: %v ", err)
	}
	err = yaml.Unmarshal(setting, conf)
	if err != nil {
		log.Printf("Unmarshal: %v", err)
	}
	//	log.Print(conf)
	return conf
}
