package main

import (
	"flag"
	"net/http"
)

func main() {
	domain := flag.String("domain", "", "Domain address to listen.")
	port := flag.String("port", ":9090", "Port address to listen.")

	http.HandleFunc("/numbers", processRequest)
	http.HandleFunc("/", invalidEndpoint)
	err := http.ListenAndServe(*domain+*port, nil) // Start server
	if err != nil {
		panic(err)
	}
}
