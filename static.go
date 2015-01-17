package main

import (
	"flag"
	"log"
	"net/http"
)

var bind = flag.String("bind", ":18181", "bind address")
var root = flag.String("root", "/srv/http", "root directory")

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	http.Handle("/", loggingHandler(http.FileServer(http.Dir(*root))))
	log.Fatal(http.ListenAndServe(*bind, nil))
}
