package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Starting HTTP Server")
}

func main() {
	http.HandleFunc("/", handler)
	log.Info(http.ListenAndServer(":8080", nil))
}