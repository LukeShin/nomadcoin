package cli

import (
	"flag"
	"fmt"
	"os"

	"github.com/LukeShin/nomadcoin/explorer"
	"github.com/LukeShin/nomadcoin/rest"
)

func ussage() {
	fmt.Printf("\nWecome to 노마드 코인\n")
	fmt.Printf("Please use the following flags:\n\n")
	fmt.Printf("-port=4000: Set the PORT of the server\n")
	fmt.Printf("-mode=rest: Choose mode from 'html', 'rest' and 'all'\n\n")
	os.Exit(0)
}

func Start() {
	if len(os.Args) == 1 {
		ussage()
	}

	port := flag.Int("port", 4000, "Set port of the server")
	mode := flag.String("mode", "rest", "Choose between 'html' and 'rest'")
	flag.Parse()

	switch *mode {
	case "rest":
		rest.Start(*port)
	case "html":
		explorer.Start(*port)
	case "all":
		go rest.Start(*port)
		explorer.Start(*port + 1)
	default:
		ussage()
	}
	fmt.Println(*port, *mode)
}
