package main

import "fmt"
import "rsc.io/quote"
import "log"

func basic() string {
	res := fmt.Sprintf("\nHello, World!");
	return res
}

func main() {
    fmt.Print(basic(),"\n")
	log.Println("\t",quote.Go())
}