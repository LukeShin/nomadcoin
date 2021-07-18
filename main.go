package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/LukeShin/nomadcoin/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	// tmpl, err := template.ParseFiles("templates/home.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
}
func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Listening on Http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
