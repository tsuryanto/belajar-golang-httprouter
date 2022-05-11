package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var localhost string = "http://localhost:3000"
var localhostAddr string = "localhost:3000"

func main() {
	router := httprouter.New()
	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello httprouter")
	})

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	server.ListenAndServe()
}
