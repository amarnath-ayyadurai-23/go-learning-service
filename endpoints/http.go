package main

import (
   // "encoding/json"
    //"regexp"
    //"sync"
	
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"log"
)

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("\tgot / request")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	log.Println("\tgot /hello request")
	io.WriteString(w, "Hello, HTTP!\n")
}

func main() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	
	log.Println("\tStarted http Server @ port 3000")
	log.Println("\tPress Ctrl+C to shutdown... ")
	err := http.ListenAndServe(":3000", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}