package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/pflag"
)

func main() {
	port := pflag.IntP("port", "p", 8080, "Port to serve on")
	directory := pflag.StringP("directory", "d", ".", "Directory to serve")
	pflag.Parse()

	fs := http.FileServer(http.Dir(*directory))
	http.Handle("/", fs)

	fmt.Printf("Serving directory %s on port %d\n", *directory, *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
