package main

import "fmt"
import "rsc.io/quote"
import "log"

func main() {
    fmt.Println("\nHello, World!")
	log.Println("\t",quote.Go(),"\n")
}