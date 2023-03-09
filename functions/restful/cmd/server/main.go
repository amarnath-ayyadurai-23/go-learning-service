// Package main for local testing
package main

import (
	"log"
	"os"

	_ "restful"

	"github.com/GoogleCloudPlatform/functions-framework-go/funcframework"
	// Blank-import the function package so the init() runs
	//_ "github.com/pantheon-systems/go-vcs-service/functions/restful"
)

func main() {

	// Use PORT environment variable, or default to 8080.
	port := "3000"
	if envPort := os.Getenv("PORT"); envPort != "" {
		port = envPort
	}
	if err := funcframework.Start(port); err != nil {
		log.Fatalf("funcframework.Start: %v\n", err)
	}
}
