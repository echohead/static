package main

import (
	"flag"
	"log"
	"net/http"
)

var bind = flag.String("bind", ":8080", "bind address")
var root = flag.String("root", "", "root directory")

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	if *root == "" {
		log.Fatal("missing required flag --root")
	}
	http.Handle("/", loggingHandler(http.FileServer(http.Dir(*root))))
	log.Printf("serving %v on %v", *root, *bind)
	log.Fatal(http.ListenAndServe(*bind, nil))
}
