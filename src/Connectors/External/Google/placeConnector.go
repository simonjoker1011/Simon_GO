package Connectors

import (
	"Constant"
	"Utils/Settings"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	//	"os"
)

type conf struct {
	Key    string `yaml:"api_key"`
	Server string `yaml:"api_server"`
}

func GetPlaceDetail(placeid string) []byte {
	log.Printf("start getting place detail for id: %v", placeid)

	var placeConf conf
	Utils.Load_settings(&placeConf, "placeSetting.yaml")

	req := make(map[string]string)
	req["placeid"] = placeid
	req["key"] = placeConf.Key

	var buffer bytes.Buffer
	buffer.WriteString(placeConf.Server)
	buffer.WriteString(Constant.Api_Method_Detail)
	buffer.WriteString("json?")
	buffer.WriteString(To_Map(req))

	if u, err := url.Parse(buffer.String()); err == nil {
		if resp, httpErr := http.Get(u.String()); httpErr == nil {
			if b, readErr := ioutil.ReadAll(resp.Body); readErr == nil {
				//TODO: to mongo
				return b
			}
		} else {
			log.Printf("http error: %v", httpErr)
		}
	} else {
		log.Printf("parse url %v failed, err: %v", buffer.String(), err)
	}
	return nil
}

func GetNearby(geoX string, geoY string, radius string, placetype string) []byte {
	log.Printf("start search for { %v, %v } nearby %v meters with type: %v", geoX, geoY, radius, placetype)

	var placeConf conf
	Utils.Load_settings(&placeConf, "placeSetting.yaml")

	req := make(map[string]string)
	req["location"] = geoX + "," + geoY
	req["radius"] = radius
	req["type"] = placetype
	req["key"] = placeConf.Key

	var buffer bytes.Buffer
	buffer.WriteString(placeConf.Server)
	buffer.WriteString(Constant.Api_Method_NearbySearch)
	buffer.WriteString("json?")
	buffer.WriteString(To_Map(req))

	if u, err := url.Parse(buffer.String()); err == nil {
		if resp, httpErr := http.Get(u.String()); httpErr == nil {
			if b, readErr := ioutil.ReadAll(resp.Body); readErr == nil {
				//TODO: to mongo
				return b
			}
		} else {
			log.Printf("http error: %v", httpErr)
		}
	} else {
		log.Printf("parse url %v failed, err: %v", buffer.String(), err)
	}
	return nil
}
