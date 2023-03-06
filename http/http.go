// Package main Golang Learning Service Microservice API.
//
// Documentation for Learning API
//
//	Schemes: http, https
//	Host: localhost
//	Description: This is a basic Learning APIs based on VCS service
//	BasePath: /
//	Version: 0.0.1
//	License: CC https://creativecommons.org/licenses/
//	Contact: Amar<amarnath.ayyadurai@pantheon.io>
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package main

import (
	// "encoding/json"
	//"regexp"
	//"sync"

	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type JSONResponse struct {
	message     string
	status_code int
}

// swagger:route GET / Basic basic
//
// This is the summary for getting a echo from base url
//
// This is the description for getting basic text like a ping
//
// responses:
//   200:
//		"$ref": "#/responses/rootResponse"
//   404: ErrorResponse
//   500: ErrorResponse

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Println("\tgot / request")
	// swagger:response rootResponse
	// in:body
	// Text
	//
	// example: This is my website

	io.WriteString(w, "This is my website!\n")
}

// swagger:route GET /hello Hello hello
//
// # This is the summary for getting a echo from base url
//
// # This is the description for getting basic text like a ping
//
// responses:
//
//	  200:
//		"$ref": "#/responses/helloResponse"
//	  404: ErrorResponse
//	  500: ErrorResponse
func getHello(w http.ResponseWriter, r *http.Request) {
	log.Println("\tgot /hello request")
	// swagger:response helloResponse
	// in:body
	// Text
	//
	// example: Hello, HTTP

	io.WriteString(w, "Hello, HTTP!\n")
}

// swagger:route GET /api/v1 JSONtest jsontest
//
// # This is the summary for getting a echo from base url
//
// # This is the description for getting basic text like a ping
//
// responses:
//
//	  200:
//			"$ref": "#/responses/jsonResponse"
//	  404: ErrorResponse
//	  500: ErrorResponse
func getJson(w http.ResponseWriter, r *http.Request) {
	jsonData := JSONResponse{"Working Json Output", 201}
	fmt.Println(jsonData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status":     jsonData.message,
		"statusCode": jsonData.status_code,
	})
	log.Println("\tgot /api/v1 request")
}

func runServer() {
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	http.HandleFunc("/api/v1", getJson)

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

func main() {
	runServer()
}
