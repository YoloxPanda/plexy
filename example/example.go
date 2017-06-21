package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rbrick/plexy"
)

func main() {
	plexer := plexy.NewPlexy()

	plexer.HandleFunc("/hello/:world", func(w http.ResponseWriter, r *http.Request, params *plexy.Params) {
		// fmt.Fprintf(w, "Hello, %s", params.Get("world"))
		fmt.Println("Handled with Plexy")
	})

	log.Fatalln(http.ListenAndServe(":8080", plexer))
}
