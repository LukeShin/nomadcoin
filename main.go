package main

import (
	"flag"
	"fmt"
	"os"
)

func ussage() {
	fmt.Printf("\nWecome to 노마드 코인\n")
	fmt.Printf("Please use the following commands:\n\n")
	fmt.Printf("explorer: Start the HTML Explorer\n")
	fmt.Printf("rest:     Start the REST API (recommended)\n\n")
	os.Exit(0)
}

func main() {

	if len(os.Args) < 2 {
		ussage()
	}

	rest := flag.NewFlagSet("rest", flag.ExitOnError)
	portFlag := rest.Int("port", 4000, "Sets the port of the server")

	switch os.Args[1] {
	case "explorer":
		fmt.Println("Start Explorer")
	case "rest":
		rest.Parse(os.Args[2:])
	default:
		ussage()
	}

	if rest.Parsed() {
		fmt.Println(portFlag)
		fmt.Println("Start Server")
	}

	fmt.Println(*portFlag)
}
