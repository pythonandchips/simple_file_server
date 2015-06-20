package main

import "net/http"
import "flag"
import "fmt"

var dir = flag.String("dir", ".", "Directory to serve files from")
var port = flag.String("port", "8080", "Port to run server on")

func main() {
	flag.Parse()
	fmt.Printf("Starting on port :%s\n", *port)
	fmt.Printf("Serving %s\n", *dir)
	panic(http.ListenAndServe(":"+*port, http.FileServer(http.Dir(*dir))))
}
