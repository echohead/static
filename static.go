package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "0.0.3"

var bind = flag.String("bind", ":8080", "bind address")
var root = flag.String("root", "", "root directory")
var showVersion = flag.Bool("version", false, "print version and exit")

func loggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		h.ServeHTTP(w, r)
	})
}

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		os.Exit(2)
	}
	if *root == "" {
		log.Fatal("missing required flag --root")
	}
	http.Handle("/", loggingHandler(http.FileServer(http.Dir(*root))))
	log.Printf("serving %v on %v", *root, *bind)
	log.Fatal(http.ListenAndServe(*bind, nil))
}
