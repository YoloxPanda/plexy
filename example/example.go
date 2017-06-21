package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rbrick/plexy"
)

func main() {
	plexer := plexy.NewPlexy()

	plexer.HandleFunc("/sayhello/:from/:to", func(w http.ResponseWriter, r *http.Request, params *plexy.Params) {
		fmt.Fprintf(w, "%s says hello to %s", params.Get("from"), params.Get("to"))
	})

	log.Fatalln(http.ListenAndServe(":8080", plexer))
}
