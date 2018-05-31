package main

import (
	"Utils/Settings"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("entering...")
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func handler1(w http.ResponseWriter, r *http.Request) {
	log.Println("entering......")

	switch r.URL.Path[1:] {
	case "OK":
		//		m := Message{"HOT", "Hello", 1294706395881547000}
		Utils.BuildJsonResp(w, Message{"HOT", "Hello", 1294706395881547000})
		//		w.WriteHeader(200)
	case "OK2":
		m := Message{"HERE", "Hello", 1294706395881547000}
		//		w.WriteHeader(200)
		Utils.BuildJsonResp(w, m)
	case "3RD":
		m := Message{"THERE", "hi", 1294706395881547000}
		Utils.BuildJsonResp(w, m)
		w.WriteHeader(204)
	case "444":
		w.WriteHeader(204)
	case "NotFound":
		w.WriteHeader(404)
	case "InternalServerError":
		w.WriteHeader(500)
	default:
		log.Println("not such path")
		w.WriteHeader(500)
	}
}

func main() {
	http.HandleFunc("/", handler1)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
