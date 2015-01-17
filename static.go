package main

import (
	"flag"
	"log"
	"net/http"
)

var bind = flag.String("bind", ":18181", "bind address")
var root = flag.String("root", "/srv/http", "root directory")

func main() {
	flag.Parse()
	http.Handle("/", http.FileServer(http.Dir(*root)))
	log.Fatal(http.ListenAndServe(*bind, nil))
}
