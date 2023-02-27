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
	"encoding/json"
)
type JSONResponse struct{
	message string 
	status_code int
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("\tgot / request")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	log.Println("\tgot /hello request")
	io.WriteString(w, "Hello, HTTP!\n")
}

func getJson(w http.ResponseWriter, r *http.Request) {
	jsonData := JSONResponse{"Working Json Output",201}
	fmt.Println(jsonData)
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
		"status": jsonData.message, 
		"statusCode": jsonData.status_code,
	})
	log.Println("\tgot /api/v1 request")
}

func runServer() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/api/v1",getJson)
	
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

func main(){
	runServer()
}