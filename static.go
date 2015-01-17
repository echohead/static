package main

import (
	"flag"
	"log"
	"net/http"
)

var bind = flag.String("bind", ":8080", "bind address")
var root = flag.String("root", "/srv", "root directory")

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	http.Handle("/", loggingHandler(http.FileServer(http.Dir(*root))))
	log.Printf("starting on %v", *bind)
	log.Fatal(http.ListenAndServe(*bind, nil))
}
