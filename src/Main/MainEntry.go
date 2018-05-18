package main

import (
	"encoding/json"
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
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		m := Message{"Alice", "Hello", 1294706395881547000}
		b, _ := json.Marshal(m)
		w.Write(b)
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
