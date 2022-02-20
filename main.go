package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/pkw5240/blockchain/blockchain"
)

const port string = ":4000"

type homeData struct {
	PageTitle string
	Blocks    []*blockchain.Block
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/home.html"))
	data := homeData{"Home", blockchain.GetBlockchain().AllBlocks()}
	tmpl.Execute(rw, data)
	//fmt.Fprintf(rw, "hello from home!")
}

func main() {

	http.HandleFunc("/", home)
	fmt.Printf("listening on http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
