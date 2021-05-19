package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func serveList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	PORT := os.Getenv("PORT")
	
	if PORT == "" {
		log.Fatal("No PORT env defined, aborting...")
	}

    router := httprouter.New()
    router.GET("/proxy-list", serveList)
	

    log.Fatal(http.ListenAndServe(os.Getenv("PORT"), router))
}