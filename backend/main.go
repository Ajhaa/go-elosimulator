package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"ajhaa.fi/elosimulator/routes"
	"github.com/julienschmidt/httprouter"
)

func apiRoot(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "This is the api root")
}

func main() {
	seed := time.Now().Unix()
	rand.Seed(seed)

	router := httprouter.New()
	router.GET("/", routes.Index)
	router.GET("/simulate/:amount", routes.Simulate)
	router.GET("/api", apiRoot)

	fmt.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", router))
}
